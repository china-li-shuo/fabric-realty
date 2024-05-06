package api

import (
	"encoding/json"
	"fmt"

	"chaincode/model"
	"chaincode/pkg/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// 添加不良记录
func BadRecordAdd(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	name := args[0]
	idCard := args[1]
	isLock := args[2]
	if name == "" || idCard == "" || isLock == "" {
		return shim.Error("参数存在空值")
	}

	badRecord := &model.BadRecord{
		Name:   name,
		IdCard: idCard,
		IsLock: isLock,
	}

	// 写入账本
	if err := utils.WriteLedger(badRecord, stub, model.BadRecordKey, []string{badRecord.Name, badRecord.IdCard, badRecord.IsLock}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	//将成功创建的信息返回
	badRecordByte, err := json.Marshal(badRecord)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(badRecordByte)
}
