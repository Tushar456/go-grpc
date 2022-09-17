package auth

import (
	"fmt"
	"github.com/Tushar456/go-grpc/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/Tushar456/go-grpc/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/Tushar456/go-grpc/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)

}
func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
