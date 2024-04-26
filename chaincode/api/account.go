package api

import (
	"encoding/json"
	"fmt"

	"chaincode/model"
	"chaincode/pkg/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// QueryAccountList 查询账户列表
func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//多个区块链的账户信息
	var accountList []model.Account
	//通过部分组合键获取状态
	//model.AccountKey 是个常量，常量值是 account-key
	//args   {"accountId":"5feceb66ffc8"}
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//进行循环处理账户信息
	for _, v := range results {
		//如果区块链里查到了有交易的账号信息
		if v != nil {
			//定义结构体获取单个账户信息
			var account model.Account
			//把json数据转化为结构体
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	//又序列号之后把数据返回给调用者（beego,gin ）
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}
