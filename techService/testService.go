package techService

import (
	"fmt"
	"time"

	"github.com/kbh0581/techPublish-grpc/mongoUtill"
	"github.com/kbh0581/techPublish-grpc/scrapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ScrapperSoruce string

const (
	wooha = ScrapperSoruce("wooha")
)

type TechDoc struct {
	Types  ScrapperSoruce
	Author string
	Time   time.Time
	Title  string
	Link   string
	Detail string
}

func (s *TechDoc) wohaTechDocToTechdoc(data scrapper.WohaTechDoc) {
	s.Author = data.Author
	s.Time = data.Time
	s.Title = data.Title
	s.Link = data.Link
	s.Detail = data.Detail
}

func deleteAll() {
	client, ctx, cancel := mongoUtill.GetConnetion()
	postCollection := mongoUtill.GetCollection(client, "postInfo")
	postCollection.DeleteMany(ctx, bson.D{{}})
	defer client.Disconnect(ctx)
	defer cancel()
}

func insertWoowa(ch chan *mongo.SingleResult, v scrapper.WohaTechDoc) {

	client, ctx, cancel := mongoUtill.GetConnetion()
	postCollection := mongoUtill.GetCollection(client, "postInfo")
	defer client.Disconnect(ctx)
	defer cancel()

	doc := TechDoc{}
	doc.Types = wooha
	doc.wohaTechDocToTechdoc(v)
	go func() {
		ch <- postCollection.FindOne(ctx, bson.D{{"link", doc.Link}}, options.FindOne().SetProjection(bson.D{{"_id", 1}}))
	}()

	res := <-ch
	var result bson.D
	res.Decode(&result)
	fmt.Println(result)
	if result == nil {
		insertResult, err := postCollection.InsertOne(ctx, doc)
		fmt.Println(insertResult)
		fmt.Println(err)
	}
}

func ScrapperService() {

	scrapper.WoowaScrapper(func(a []scrapper.WohaTechDoc) {
		dataChan := make(chan *mongo.SingleResult, 4)
		done := make(chan bool, 1)

		deleteAll()
		fmt.Println(len(a))

		func() {
			for index, v := range a {
				fmt.Println(index)
				go insertWoowa(dataChan, v)
			}
			done <- true
		}()
		<-done

	})

	//var datas []bson.M

	//리드 데이터	//res, err := mongo.GetCollection(client, "postInfo").Find(ctx, bson.D{{"data", "zz"}})
	//if err != nil {
	//fmt.Println(err)
	//}

	//resultone, _ := mongo.GetCollection(client, "postInfo").InsertOne(ctx, bson.D{{"data", "xx"}})

	// fmt.Println(resultone)

	// // 루프로 읽어야함
	// for res.Next(context.TODO()) {
	// 	var result bson.D
	// 	err = res.Decode(&result)
	// 	fmt.Println(result)

	// }

}
