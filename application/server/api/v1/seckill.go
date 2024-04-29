package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"application/model"
	"application/pkg/app"
	"application/service"
)

type SeckillGoods struct {
	GoodID     int `json:"good_id" form:"good_id"`         //商品ID
	UserID     int `json:"user_id" form:"user_id"`         //用户ID
	ActivityID int `json:"activity_id" form:"activity_id"` //活动ID
}

// SecKill 秒杀
func SecKill(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SeckillGoods)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	//去查询活动的信息
	var activityData model.Activity
	tx := model.DB.Model(&model.Activity{}).First(&activityData, body.ActivityID)
	if tx.RowsAffected == 0 {
		appG.Response(http.StatusNotFound, "没有该活动", nil)
		return
	}
	//判断活动是否开启
	//获取当前时间戳
	timestamp := time.Now().Unix()

	if int(timestamp) < activityData.StartTime {
		appG.Response(http.StatusBadRequest, "活动未开启", nil)
		return
	}

	//判断用户是否已经抢购到这件商品了 set_buy_scekill_good_1
	SeckillGoodsSet := "set_buy_scekill_good_" + strconv.Itoa(body.GoodID)

	//链接redis
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       7,  // use default DB
	})

	//查找用户id有没有在redis的set集合里
	if rdb.SIsMember(ctx, SeckillGoodsSet, body.UserID).Val() {
		appG.Response(http.StatusBadRequest, "你已经参与过秒杀了！", nil)
		return
	}

	//判断商品是否有库存(去查询redis的队列长度) list_scekill_good_1
	SeckillGoodsList := "list_scekill_good_" + strconv.Itoa(body.GoodID)

	//判断redis的key是否存在，如果队列不存在，第一个人先去数据库查库存，然后同步到redis的队列中（最好是提前预热）
	var goodsData model.Goods

	//判断队列的长度
	listLen := rdb.LLen(ctx, SeckillGoodsList).Val()
	if listLen <= 0 {
		appG.Response(http.StatusBadRequest, "库存不足！", nil)
		return
	}

	//扣减本地库存到redis
	rdb.LPop(ctx, SeckillGoodsList)

	//把秒杀成功的uid写入到set集合里
	rdb.SAdd(ctx, SeckillGoodsSet, body.UserID)

	//生成订单
	snowflake, err := service.NewSnowflake(1)
	if err != nil {
		return
	}

	tx = model.DB.Model(&model.Goods{}).First(&goodsData, body.GoodID)
	if tx.RowsAffected == 0 {
		appG.Response(http.StatusNotFound, "没有该商品", nil)
		return
	}

	orderData := model.Order{
		OrderSn:    "itshujia_" + strconv.Itoa(int(snowflake.GetID())),
		UserID:     body.UserID,
		GoodsID:    body.GoodID,
		ActivityID: body.ActivityID,
		Price:      goodsData.Price,
		Stock:      1,
	}
	tx = model.DB.Create(&orderData)
	if tx.RowsAffected == 0 {
		appG.Response(http.StatusNotFound, "订单生成失败", nil)
		return
	}
	//生成延迟队列，用来做库存归还

	appG.Response(http.StatusOK, "秒杀成功", nil)
}

// 库存预热
func StockHot(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SeckillGoods)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	//链接redis
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       7,  // use default DB
	})
	//判断redis的key是否存在，如果队列不存在，第一个人先去数据库查库存，然后同步到redis的队列中（最好是提前预热）
	var goodsData model.Goods

	SeckillGoodsList := "list_scekill_good_" + strconv.Itoa(body.GoodID)

	if rdb.Exists(ctx, SeckillGoodsList).Val() == 0 {
		tx := model.DB.Model(&model.Goods{}).First(&goodsData, body.GoodID)
		if tx.RowsAffected == 0 {
			appG.Response(http.StatusNotFound, "没有该商品", nil)
			return
		}
		//把商品的库存写入到队列中
		for i := 0; i < goodsData.Stock; i++ {
			rdb.RPush(ctx, SeckillGoodsList, 1)
		}
	}

	appG.Response(http.StatusOK, "预热成功", nil)
}
