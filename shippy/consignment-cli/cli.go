package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	Address         = "localhost:50051"
	DefaultInfoFile = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	//fmt.Println("consignment.json\n",string(bytes))

	consignment := new(pb.Consignment)
	//var consignment *pb.Consignment
	err = json.Unmarshal(bytes, consignment)
	if err != nil {
		return nil, err
	}

	return consignment, nil

}

func main() {

	microClient()
}

func microClient() {
	// create a new service
	service := micro.NewService()

	service.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", service.Client())

	consignment, err := parseFile(DefaultInfoFile)
	if err != nil {
		log.Fatalf("--------------get consignment fail:%v\n", err)
	}

	createConsignment(client, consignment)

	getConsignments(client)

}

/*测试调用微服务接口 createConsignment*/
func createConsignment(client pb.ShippingServiceClient, consignment *pb.Consignment) {
	response, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("--------------conn fail:%v\n", err)
	}

	fmt.Println("response->created", response.Created)
	fmt.Println("response->consignment", response.Consignment)
	fmt.Println("response->consignments", response.Consignments)
	fmt.Println("createConsignment func end!-----------")
	fmt.Println("")
}

func getConsignments(client pb.ShippingServiceClient) {
	response, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("--------------conn fail:%v\n", err)
	}
	fmt.Println("response->created", response.Created)
	fmt.Println("response->consignment", response.Consignment)
	fmt.Println("response->consignments", response.Consignments)
	fmt.Println("getConsignments func end!-----------")
	fmt.Println("")

}

//func gRPCClient()  {
//	// 连接到gRPC服务器
//	conn, err := grpc.Dial(Address, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("did not connection:%v", err)
//	}
//	defer conn.Close()
//
//	// 初始化gRPC客户端
//	client := pb.NewShippingServiceClient(conn)
//
//	// 在命令行中指定新的货物信息 json 文件
//	infoFile := DefaultInfoFile
//	if len(os.Args) > 1 {
//		infoFile = os.Args[1]
//	}
//
//	// 解析货物信息
//	consignment, err := parseFile(infoFile)
//	if err != nil {
//		log.Fatalf("parse info file error: %v", err)
//	}
//
//	// 调用 RPC
//	// 将货物存储到我们自己的仓库里
//	response, err := client.CreateConsignment(context.Background(), consignment)
//	if err != nil {
//		log.Fatalf("create consignment err:%v", err)
//	}
//
//	//新货物是否托运成功
//	log.Printf("Created :%t", response.Created)
//
//	//获取货物信息
//	consignments, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
//	if err != nil {
//		log.Fatalf("create consignment err:%v", err)
//	}
//	for _, v := range consignments.Consignments {
//		fmt.Println(v.Id)
//		fmt.Println(v.Containers)
//		fmt.Println(v.Description)
//		fmt.Println(v.Weight)
//		fmt.Println(v.VesselId)
//		fmt.Println()
//	}
//
//}
