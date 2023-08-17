package interceptor

import (
	"fmt"
	"context"

	"google.golang.org/grpc"
	"github.com/Daaaai0809/grpc-auth-practice/data"
	"github.com/Daaaai0809/grpc-auth-practice/auth"
	proto "github.com/Daaaai0809/grpc-auth-practice/proto"
)

type MethodList map[string][]float64

var methodList = MethodList {
	proto.AuthSampleService_LoginMethod_FullMethodName: {data.Guest},
	proto.AuthSampleService_AdminMethod_FullMethodName: {data.Admin},
	proto.AuthSampleService_RequiredAuthMethod_FullMethodName: {data.Admin, data.User},
	proto.AuthSampleService_NotRequiredAuthMethod_FullMethodName: {data.User, data.Guest},
}

func Contains(list []float64, role float64) bool {
	for _, v := range list {
		if v == role {
			return true
		}
	}
	return false
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	roles, ok := methodList[info.FullMethod]
	if !ok {
		return nil, fmt.Errorf("method not found")
	}

	if (!Contains(roles, data.Guest)) {
		if (!Contains(roles, data.User)) {
			if (!Contains(roles, data.Admin)) {
				return nil, fmt.Errorf("method not found")
			}
			_req := req.(*proto.AdminRequest)
			if err := auth.CheckAdminToken(_req); err != nil {
				return nil, fmt.Errorf("permission denied")
			}
			return handler(ctx, _req)
		}
		_req := req.(*proto.RequiredAuthRequest)
		if err := auth.CheckUserToken(_req); err != nil {
			return nil, fmt.Errorf("permission denied")
		}

		return handler(ctx, _req)
	}

	return handler(ctx, req)
}