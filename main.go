package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kbh0581/techPublish-grpc/test/sampleProto"
	"github.com/kbh0581/techPublish-grpc/test/sampleService"
	"google.golang.org/grpc"
)

const serverPort = ":9000"

type testServer struct {
	sampleProto.TestServer
}

func (s *testServer) GetList(ctx context.Context, req *sampleProto.ReqList) *sampleProto.ResponsList {
	testmesage := make([]*sampleProto.Response, len(sampleService.TestData))
	for i, u := range sampleService.TestData {
		print("data:" + "i")
		print("data u:" + "u")
		testmesage[i] = u
	}

	return &sampleProto.ResponsList{
		Res: testmesage,
	}
}

func main() {

	lis, err := net.Listen("tcp", serverPort)

	log.Print(lis.Addr().String())
	//nil 없음
	// go mod init "패키지 매니저"
	// go get -u  최신 버전 가져옴
	if err != nil {
		log.Fatal(err)
	}

	log.Print("grpcServeron")
	grpcServer := grpc.NewServer()
	sampleProto.RegisterTestServer(grpcServer, &testServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail serve : %s", err)
	}

	fmt.Print("hellow")
}
