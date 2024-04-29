package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"application/model"
	"application/pkg/app"
)

type SeckillGoods struct {
	GoodID     string `json:"good_id" form:"good_id"`         //商品ID
	UserID     string `json:"user_id" form:"user_id"`         //用户ID
	ActivityID string `json:"activity_id" form:"activity_id"` //活动ID
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

	//var data []map[string]interface{}

	appG.Response(http.StatusOK, "成功", body)
}
