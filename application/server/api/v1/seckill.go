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

	//判断用户是否已经抢购到这件商品了 buy_scekill_good_1
	IsBuyMember := "buy_scekill_good_" + strconv.Itoa(body.GoodID)

	//链接redis
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       7,  // use default DB
	})

	//查找用户id有没有在redis的set集合里
	fmt.Println(rdb.SIsMember(ctx, IsBuyMember, body.UserID).Val())
	if rdb.SIsMember(ctx, IsBuyMember, body.UserID).Val() {
		appG.Response(http.StatusBadRequest, "你已经参与过秒杀了！", nil)
		return
	}

	//var data []map[string]interface{}

	appG.Response(http.StatusOK, "成功", body)
}
