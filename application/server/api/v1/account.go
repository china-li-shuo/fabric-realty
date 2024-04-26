package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bc "application/blockchain"
	"application/pkg/app"

	"github.com/gin-gonic/gin"
)

type AccountIdBody struct {
	AccountId string `json:"accountId"`
}

type AccountRequestBody struct {
	Args []AccountIdBody `json:"args"`
}

// QueryAccountList 获取账户列表
func QueryAccountList(c *gin.Context) {
	//gin框架的上下文
	appG := app.Gin{C: c}
	body := new(AccountRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	//追加登录账户id，多个账户
	var bodyBytes [][]byte
	for _, val := range body.Args {
		bodyBytes = append(bodyBytes, []byte(val.AccountId))
	}
	//调用智能合约
	//resp 多个账户信息（区块链）
	resp, err := bc.ChannelQuery("queryAccountList", bodyBytes)
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
