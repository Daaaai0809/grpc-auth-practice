package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/Daaaai0809/grpc-auth-practice/interceptor"
	"github.com/Daaaai0809/grpc-auth-practice/auth"
	proto "github.com/Daaaai0809/grpc-auth-practice/proto"
)

type Server struct{
	proto.UnimplementedAuthSampleServiceServer
}

var port = ":8080"

func (s *Server) LoginMethod(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	userName := req.GetUserName()
	if userName == "" {
		return nil, fmt.Errorf("user name is empty")
	}
	password := req.GetPassword()
	if password == "" {
		return nil, fmt.Errorf("password is empty")
	}

	role, err := auth.CheckPassword(userName, password)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(userName, role)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Token: token,
	}, nil
}

func (s *Server) AdminMethod(ctx context.Context, req *proto.AdminRequest) (*proto.AdminResponse, error) {
	message := req.GetMessage()
	if message == "" {
		return nil, fmt.Errorf("message is empty")
	}
	
	msg := fmt.Sprintf("[AdminMethod] %s", message)

	return &proto.AdminResponse{
		Message: msg,
	}, nil
}

func (s *Server) RequiredAuthMethod(ctx context.Context, req *proto.RequiredAuthRequest) (*proto.RequiredAuthResponse, error) {
	message := req.GetMessage()
	if message == "" {
		return nil, fmt.Errorf("message is empty")
	}
	
	msg := fmt.Sprintf("[RequiredAuthMethod] %s", message)

	return &proto.RequiredAuthResponse{
		Message: msg,
	}, nil
}

func (s *Server) NotRequiredAuthMethod(ctx context.Context, req *proto.NotRequiredAuthRequest) (*proto.NotRequiredAuthResponse, error) {
	name := req.GetName()
	if name == "" {
		return nil, fmt.Errorf("message is empty")
	}
	
	msg := fmt.Sprintf("[NotRequiredAuthMethod] Hello, %s!!", name)

	return &proto.NotRequiredAuthResponse{
		Message: msg,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor),
	)

	proto.RegisterAuthSampleServiceServer(s, &Server{})
	reflection.Register(s)

	go func() {
		log.Printf("start server on port%s", port)
		s.Serve(lis)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("stopping server...")
	s.GracefulStop()
}