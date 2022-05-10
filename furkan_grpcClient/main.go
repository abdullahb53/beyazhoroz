package grpcclient

import (
	"context"
	"fmt"

	"github.com/abdullahb53/beyazhoroz/api"
	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:443"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewFurkanServiceClient(conn)

	ctx := context.Background()
	resp, err := client.ListUsers(ctx, &api.UserListReq{})
	if err != nil {
		panic(err)
	}

	fmt.Println("users:")
	for _, user := range resp.ListedUserItems.Data.People {
		fmt.Println("\ts: %s", user.GetName())
	}

}
