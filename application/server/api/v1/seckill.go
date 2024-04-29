package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"application/pkg/app"
)

type SeckillGoods struct {
	GoodID string `json:"good_id"` //操作人ID
	UserID string `json:"user_id"` //所有者(业主)(业主AccountId)
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

	var data []map[string]interface{}

	appG.Response(http.StatusOK, "成功", data)
}
