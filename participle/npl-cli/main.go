package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/caoyuewen/participle/npl-server/proto"
)


func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("qa.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewNplService("go.micro.srv.npl", service.Client())

	// Call the greeter
	rsp, err := greeter.GetParticiple(context.TODO(), &proto.SentenceRequest{Sentence: "自动化机制"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	ps:=rsp.Participle
	for i:=0;i<len(ps) ;i++  {
		fmt.Println(ps[i].Word,ps[i].Wordtype)
	}

}
