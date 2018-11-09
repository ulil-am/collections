package grpc

import (
	ctrl "collections/controllers/grpc"
	"collections/helper/constant"
	pb "collections/proto"
	"collections/structs"
)

var (
	prefix = "/" + constant.GOAPP + "/" + constant.VERSION
)

type fnRouteRPC func(
	*pb.DoReq,
	*structs.TypeGRPCError,
	*[]byte,
)

var routeMap map[string]fnRouteRPC

func init() {
	Router()
}

func Router() {
	routeMap = map[string]fnRouteRPC{
		/*:STARTGRPC*/
		prefix + "/collections": ctrl.RPCtrlCollections,
		/*:ENDGRPC*/
	}
}
