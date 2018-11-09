package grpc

import (
	"encoding/json"
	"collections/helper"
	pb "collections/proto"
	"collections/structs"
	rpcStructs "collections/structs/api/grpc"
)

// RPCtrlCollections - RPCtrlCollections Controllers
func RPCtrlCollections(
	in *pb.DoReq,
	errRPCCode *structs.TypeGRPCError,
	body *[]byte,
) {

	var (
		req rpcStructs.ReqTest
		res rpcStructs.ResTest
	)

	err := json.Unmarshal(in.GetBody(), &req)
	if err != nil {
		helper.CheckErr("failed unmarshal @RPCtrlCollections", err)
		structs.ErrorCode.UnexpectedError.String(&errRPCCode.Error)
		return
	}

	res.ID = req.ID
	res.Res = "response"
	resBy, err := json.Marshal(res)
	if err != nil {
		helper.CheckErr("failed marshal &RPCtrlCollections", err)
		structs.ErrorCode.UnexpectedError.String(&errRPCCode.Error)
		return
	}

	*body = resBy
}
