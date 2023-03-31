package main

import (
	"context"
	"database/sql"
	"errors"
	"log"

	cust "example.com/customer"

	_ "github.com/mattn/go-sqlite3"

	pb "github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/customer"
)

type CustomerServer struct {
	pb.UnimplementedCustomerServiceServer
	DB *sql.DB
}

func newServer() *CustomerServer {
	db, err := sql.Open("sqlite3", "./sqlcustomer.db")
	if err != nil {
		log.Fatal(err)
	}

	s := &CustomerServer{
		DB: db,
	}

	return s
}

func (cs *CustomerServer) Sigin(ctx context.Context, request *pb.SigninRequest) (*pb.SigninResponse, error) {
	log.Println("gRPC CustomerServer Signup")

	customer := request.GetCustomer()

	var c cust.Customer 
	c.ID = int(customer.GetId())
	c.Username = customer.GetUsername()
	c.Passwd = customer.GetPassword()
	c.Email = customer.GetEmail()

	if cust.ExistingUser(cs.DB, &c) {
		return nil, errors.New("User already exists")
	}

	err := cust.Signup(cs.DB, &c)
	if err != nil {
		return nil, err
	}

	return &pb.SigninResponse{
		Header: request.GetHeader(),
		Customer: &pb.Customer{
			Id: int32(c.ID),
			Username: c.Username,
			Password: c.Passwd,
			Email: c.Email,
		},
	}, nil
}

func (cs *CustomerServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("gRPC CustomerServer Login")

	customer := request.GetCustomer()

	var c cust.Customer 
	c.ID = int(customer.GetId())
	c.Username = customer.GetUsername()
	c.Passwd = customer.GetPassword()
	c.Email = customer.GetEmail()

	_, err := cust.Login(cs.DB, &c)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Header: request.GetHeader(),
		Customer: &pb.Customer{
			Id: int32(c.ID),
			Username: c.Username,
			Password: c.Passwd,
			Email: c.Email,
		},
	}, nil
}