package techService

import (
	"context"
	"fmt"

	"github.com/kbh0581/techPublish-grpc/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func ScrapperService() {
	client, ctx, cancel := mongo.GetConnetion()

	defer client.Disconnect(ctx)
	defer cancel()

	//var datas []bson.M

	res, err := mongo.GetCollection(client, "postInfo").Find(ctx, bson.D{{"data", "zz"}})
	if err != nil {
		fmt.Println(err)
	}

	// 루프로 읽어야함
	for res.Next(context.TODO()) {
		var result bson.D
		err = res.Decode(&result)
		fmt.Println(result)

	}

}
