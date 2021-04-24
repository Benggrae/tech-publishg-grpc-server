package sampleService

import "github.com/kbh0581/techPublish-grpc/test/sampleProto"

var TestData = []*sampleProto.Response{
	{
		Test: "a",
		A:    0,
	},
	{
		Test: "2",
		A:    3,
	},
	{
		Test: "4",
		A:    5,
	},
}
