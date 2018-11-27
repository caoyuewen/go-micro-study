package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	PORT = ":50051"
)

//仓库接口
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) //存放新货物
	GetAll() []*pb.Consignment
}

// 我们存放多批货物的仓库，实现了 IRepository 接口
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)

	//---test---
	fmt.Printf("-----total consignments:%d\n", len(repo.consignments))
	for i := 0; i < len(repo.consignments); i++ {
		fmt.Printf("consignments[%d]:%s\n", i, repo.consignments[i].VesselId)
	}

	return consignment, nil

}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//定义为服务
type service struct {
	repo Repository
}

//
// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
//
// 托运新的货物
//func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {   //--gRPC
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error { //micro
	//接受承运的货物
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	fmt.Println()
	//resp = &pb.Response{Created: true, Consignment: consignment}
	resp.Created = true
	resp.Consignment = consignment
	return nil
}

// 获取目前所有托运的货物
//func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {    //gRPC
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error { //micro
	allConsignments := s.repo.GetAll()
	//resp = &pb.Response{Consignments: allConsignments}
	resp.Consignments = allConsignments

	return nil
}

func main() {
	microService()
}

func microService() {
	mserver := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// 解析命令行参数
	mserver.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(mserver.Server(), &service{repo})

	err := mserver.Run()
	if err != nil {
		log.Fatalf("failed to serve :%v", err)
	}

}

func gRpcService() {
	//listener, err := net.Listen("tcp", PORT)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//log.Printf("listen on: %s\n", PORT)
	//
	//server := grpc.NewServer()
	//repo := Repository{}
	//
	//// 向 rRPC 服务器注册微服务
	//// 此时会把我们自己实现的微服务 service 与协议中的 ShippingServiceServer 绑定
	//pb.RegisterShippingServiceServer(server, &service{repo})
	//
	//err = server.Serve(listener)
	//if err != nil {
	//	log.Fatalf("failed to serve:%v",err)
	//}
}
