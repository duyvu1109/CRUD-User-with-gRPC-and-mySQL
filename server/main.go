package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/vund5/usermanager/ent"
	"gitlab.com/vund5/usermanager/ent/user"
	pb "gitlab.com/vund5/usermanager/pkg/api"
	"google.golang.org/grpc"
)

var (
	port  = flag.Int("port", 8080, "The server port")
	users []*pb.User
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUserManagerServer
	db *ent.Client
}

// CreateUser implements api.UserManagerServer
func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	u, err := s.db.User.
		Create().
		SetName(in.GetName()).
		SetEmail(in.GetEmail()).
		SetPassword(in.GetPassword()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return &pb.CreateUserReply{
		Message: "Hello " + u.Name,
	}, nil
}

// GetUser implements api.UserManagerServer
func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u, err := s.db.User.
		Query().
		Where(user.ID(int(in.GetId()))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return &pb.GetUserReply{
		Message: "Name: " + u.Name + " your ID is: " + strconv.FormatInt(int64(u.ID), 10),
	}, nil
}

func (s *server) GetIdByEmail(ctx context.Context, in string) (id int) {
	u, _ := s.db.User.
		Query().
		Where(user.Email(in)).
		Only(ctx)
	log.Println("id: ", u)
	return u.ID
}

// UpdateUser implements api.UserManagerServer
func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	id := 0
	id = s.GetIdByEmail(ctx, in.GetUser().GetEmail())
	u, err := s.db.User.
		UpdateOneID(id).
		SetName(in.GetUser().GetName()).
		SetEmail(in.GetUser().GetEmail()).
		SetPassword(in.GetUser().GetPassword()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was updated: ", u)
	return &pb.UpdateUserReply{
		Message: "Hello " + u.Name + " your ID is: " + strconv.FormatInt(int64(id), 10),
	}, nil
}

// DeleteUser implements api.UserManagerServer
func (s *server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	err := s.db.User.
		DeleteOneID(int(in.GetId())).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed delete user: %w", err)
	}
	log.Println("user's id returned: ", strconv.Itoa(int(in.GetId())))
	return &pb.DeleteUserReply{
		Message: "Deleted: " + strconv.Itoa(int(in.GetId())),
	}, nil

}

// ListUsers implements api.UserManagerServer
func (s *server) GetAllUsers(ctx context.Context, in *pb.GetAllUsersRequest) (*pb.GetAllUsersReply, error) {
	rep := &pb.GetAllUsersReply{}
	u, _ := s.db.User.
		Query().
		All(ctx)
	for _, element := range u {
		var temp = &pb.User{}
		temp.Name = element.Name
		temp.Email = element.Email
		temp.Password = element.Password
		rep.UserList = append(rep.UserList, temp)
	}
	return rep, nil
}

func main() {
	flag.Parse()
	client, err := ent.Open("mysql", "root:duyvu1109@tcp(172.17.0.3:3306)/userManager?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagerServer(s, &server{pb.UnimplementedUserManagerServer{}, client})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
