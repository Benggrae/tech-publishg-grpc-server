package main

import (
	"context"
	"log"
	"net"

	"github.com/kbh0581/techPublish-grpc/scrapperService"
	pb "github.com/kbh0581/techPublish-grpc/test/sampleProto"
	"github.com/kbh0581/techPublish-grpc/test/sampleService"
)

const serverPort = ":9000"

type testServer struct {
	pb.TestServer
}

// 구현
func (s *testServer) GetSample(ctx context.Context, req *pb.ReqList) (*pb.ResponsList, error) {
	testmesage := make([]*pb.Response, len(sampleService.TestData))
	for i, u := range sampleService.TestData {
		print("data:" + "i")
		print("data u:" + "u")
		testmesage[i] = u
	}

	return &pb.ResponsList{
		Res: testmesage,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", serverPort)

	log.Print(lis.Addr().String())
	//nil 없음
	// go mod init "패키지 매니저"
	// go get -u  최신 버전 가져옴

	//GetHtml("https://woowabros.github.io/")
	//scrapper.WoowaScrapper()
	//mongo.GetConnetion()
	scrapperService.ScrapperService()

	if err != nil {
		log.Fatal(err)
	}

	// log.Print("grpcServeron")
	// grpcServer := grpc.NewServer()
	// pb.RegisterTestServer(grpcServer, &testServer{})

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("fail serve : %s", err)
	// }

}
