package main

import (
	proto "github.com/caoyuewen/participle/npl-server/proto"
	"github.com/micro/go-micro"
	"context"
	"fmt"
	"runtime"
	"github.com/caoyuewen/participle/npl-server/thulac"
)

type Server struct {
}

func (g *Server) GetParticiple(ctx context.Context, req *proto.SentenceRequest, rsp *proto.NplResponse) error {

	cut := thulac.Cut(req.Sentence, 0)
	items := cut.Seqs

	for i := 0; i < len(items); i++ {
		item := items[i]
		var p proto.Participle
		p.Word = item.Word
		p.Wordtype = int32(item.Typ)
		rsp.Participle = append(rsp.Participle, &p)
	}

	return nil
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	thulac.Init(8)
	microService()
	defer thulac.Destory()
}

func microService() {
	service := micro.NewService(
		micro.Name("go.micro.srv.npl"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterNplServiceHandler(service.Server(), new(Server))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
