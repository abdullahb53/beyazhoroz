package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/abdullahb53/beyazhoroz/api"
	"google.golang.org/grpc"
)

func GRPC_serve() {

	var grpc_port string = "localhost:443"
	lis, err := net.Listen("tcp", grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	api.RegisterFurkanServiceServer(srv, &myFurkanService{})
	fmt.Println("starting gRPC server...", grpc_port)
	go func() {
		panic(srv.Serve(lis))
	}()

}

type myFurkanService struct {
	api.UnimplementedFurkanServiceServer
}

func (c *myFurkanService) CreateUser(ctx context.Context, req *api.UserCreateReq) (*api.UserCreateRes, error) {

	//- MONGO DB CREATE USER -

	//________________________

	return &api.UserCreateRes{
		CreatedUserItems: &api.ResponseState{
			Status: 200,
			Mesaj:  "Success-Create-User",
			Data: &api.UsersBook{
				People: []*api.User{
					&api.User{Uid: req.PersonItem.Uid, Name: req.PersonItem.Name, Email: req.PersonItem.Email, City: req.PersonItem.City, Country: req.PersonItem.Country, Explain: req.PersonItem.Explain, Phone: req.PersonItem.Phone},
				},
			},
		},
	}, nil
}

func (c *myFurkanService) ListUsers(ctx context.Context, req *api.UserListReq) (*api.UserListRes, error) {

	//- MONGO DB LIST USERS -

	//_______________________

	return &api.UserListRes{
		ListedUserItems: &api.ResponseState{
			Status: 200,
			Mesaj:  "Sucess-List-Users",
			Data: &api.UsersBook{
				People: []*api.User{
					&api.User{Uid: "1", Name: "Abdullah BIYIK", Email: "abdullahb5355@gmail.com", City: "TRABZON", Country: "ORTAHISAR", Explain: "site ADMIN", Phone: "5301215386"},
					&api.User{Uid: "2", Name: "Invoker BIYIK", Email: "abdullahb5355@gmail.com", City: "TRABZON", Country: "ORTAHISAR", Explain: "site ADMIN", Phone: "5301215386"},
					&api.User{Uid: "3", Name: "Phantom Assasin BIYIK", Email: "abdullahb5355@gmail.com", City: "TRABZON", Country: "ORTAHISAR", Explain: "site ADMIN", Phone: "5301215386"},
				},
			},
		},
	}, nil
}

func (c *myFurkanService) FindUser(ctx context.Context, req *api.UserFindReq) (*api.UserFindRes, error) {

	//- MONGO DB FIND USER -

	//______________________

	return &api.UserFindRes{
		FindedUserItem: &api.ResponseState{
			Status: 200,
			Mesaj:  "Success-Find-User",
			Data: &api.UsersBook{
				People: []*api.User{
					&api.User{Uid: "1", Name: "Abdullah BIYIK", Email: "abdullahb5355@gmail.com", City: "TRABZON", Country: "ORTAHISAR", Explain: "site ADMIN", Phone: "5301215386"},
				},
			},
		},
	}, nil
}

func (c *myFurkanService) DeleteUser(ctx context.Context, req *api.UserDeleteReq) (*api.UserDeleteRes, error) {

	return &api.UserDeleteRes{
		UserDeletedUserItem: &api.ResponseState{
			Status: 200,
			Mesaj:  "Success-Deleted-User",
			Data: &api.UsersBook{
				People: []*api.User{
					&api.User{Uid: "1", Name: "Abdullah BIYIK", Email: "abdullahb5355@gmail.com", City: "TRABZON", Country: "ORTAHISAR", Explain: "site ADMIN", Phone: "5301215386"},
				},
			},
		},
	}, nil
}

func (c *myFurkanService) UpdateUser(ctx context.Context, req *api.UserUpdateReq) (*api.UserUpdateRes, error) {

	return &api.UserUpdateRes{
		UserUpdatedItems: &api.ResponseState{
			Status: 200,
			Mesaj:  "Success-Update-User",
			Data: &api.UsersBook{
				People: []*api.User{
					&api.User{Uid: req.PersonItem.Uid, Name: req.PersonItem.Name, Email: req.PersonItem.Email, City: req.PersonItem.City, Country: req.PersonItem.Country, Explain: req.PersonItem.Explain, Phone: req.PersonItem.Phone},
				},
			},
		},
	}, nil
}
