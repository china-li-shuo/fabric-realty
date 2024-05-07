package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	bc "application/blockchain"
	"application/pkg/app"
)

// 公安局上链的参数
type BadRecordBody struct {
	Name   string `json:"name" binding:"required" form:"name"`
	IdCard string `json:"id_card" binding:"required" form:"id_card"`
	IsLock string `json:"is_lock" binding:"required" form:"is_lock"`
}

type QueryBadRecordBody struct {
	Name   string `json:"name" form:"name"`
	IdCard string `json:"id_card" binding:"required" form:"id_card"`
}

// BadRecordAdd 不良信息记录添加
func BadRecordAdd(c *gin.Context) {
	//gin框架的上下文
	appG := app.Gin{C: c}
	body := new(BadRecordBody)

	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(body.IdCard))
	bodyBytes = append(bodyBytes, []byte(body.IsLock))
	//调用智能合约
	resp, err := bc.ChannelExecute("badRecordAdd", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "添加公安局不良记录失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "数据转化失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "添加公安局不良记录成功", data)

}

// QueryBadRecordLatestByIdCard  公安局-根据身份证号查询用户最新的不良记录
func QueryBadRecordLatestByIdCard(c *gin.Context) {
	//gin框架的上下文
	appG := app.Gin{C: c}
	body := new(QueryBadRecordBody)

	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(body.IdCard))

	//调用智能合约
	//resp 多个账户信息（区块链）
	resp, err := bc.ChannelQuery("queryBadRecord", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	if len(data) == 0 {
		appG.Response(http.StatusBadRequest, "该用户不再本系统", nil)
		return
	}
	appG.Response(http.StatusOK, "成功", data[0])
}

// QueryBadRecordListByIdCard 公安局-根据身份证号查询历史的不良记录信息
func QueryBadRecordListByIdCard(c *gin.Context) {
	//gin框架的上下文
	appG := app.Gin{C: c}
	body := new(QueryBadRecordBody)

	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(body.IdCard))

	//调用智能合约
	//resp 多个账户信息（区块链）
	resp, err := bc.ChannelQuery("queryBadRecord", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
