package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/Faroukhamadi/Banketeer/usermgmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

type UserManagementServer struct {
	conn *pgx.Conn
	pb.UnimplementedUserManagementServer
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
	return s.Serve(lis)
}

func (server *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())

	createSql := `CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY
		name TEXT
		age int
	);`

	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}

	createdUser := &pb.User{Name: in.GetName(), Age: in.GetAge()}
	transaction, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}

	_, err = transaction.Exec(context.Background(), "INSERT INTO users(name, age) VALUES ($1, $2)", createdUser.Name, createdUser.Age)
	if err != nil {
		log.Fatalf("transaction execution failed: %v", err)
	}

	transaction.Commit(context.Background())

	return createdUser, nil
}

func (server *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	var userList *pb.UserList = &pb.UserList{}
	rows, err := server.conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := pb.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		userList.Users = append(userList.Users, &user)
	}

	return userList, nil
}

func main() {
	databaseUrl := "postgres://postgres:16%2F04%2F2002@localhost:5432/postgres"
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())
	var userManagementServer *UserManagementServer = NewUserManagementServer()
	userManagementServer.conn = conn

	if err := userManagementServer.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
