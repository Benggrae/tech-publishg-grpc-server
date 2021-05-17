package scrapperService

import (
	"context"
	"fmt"
	"sync"
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

// 모든 작업 삭제
func deleteAll() {
	client, ctx, cancel := mongoUtill.GetConnetion()
	postCollection := mongoUtill.GetCollection(client, "postInfo")
	postCollection.DeleteMany(ctx, bson.D{{}})
	defer client.Disconnect(ctx)
	defer cancel()
}

type MongoPostConnet struct {
	Client         *mongo.Client
	Ctx            context.Context
	Cancle         context.CancelFunc
	PostCollection *mongo.Collection
}

func getMongoPostConnet() MongoPostConnet {
	client, ctx, cancel := mongoUtill.GetConnetion()
	postCollection := mongoUtill.GetCollection(client, "postInfo")

	mongoPostConnet := MongoPostConnet{Client: client, Ctx: ctx, Cancle: cancel, PostCollection: postCollection}

	return mongoPostConnet

}

// 존재하는 index 만 리턴
func existWooha(arr []scrapper.WohaTechDoc) []interface{} {
	connet := getMongoPostConnet()

	//wait 그룹 생성
	var wg sync.WaitGroup
	wg.Add(len(arr))
	var jobs []interface{}
	var matsetindex int

	defer connet.Client.Disconnect(connet.Ctx)
	defer connet.Cancle()
	defer fmt.Println(matsetindex, " jobs existWooha end")

	for index, v := range arr {
		go func(i int, value scrapper.WohaTechDoc) {
			doc := TechDoc{}
			doc.Types = wooha
			doc.wohaTechDocToTechdoc(value)
			res := connet.PostCollection.FindOne(connet.Ctx, bson.D{{"link", doc.Link}}, options.FindOne().SetProjection(bson.D{{"_id", 1}}))

			var result bson.D
			res.Decode(&result)

			if result == nil {

				fmt.Println(res)
				jobs = append(jobs, value)
				//jobs = append(jobs, v)
			}
			matsetindex++
			wg.Done()
		}(index, v)
	}

	wg.Wait()
	fmt.Println(matsetindex)
	return jobs

}

func insertWooha(jobs []interface{}) {
	connet := getMongoPostConnet()
	defer connet.Client.Disconnect(connet.Ctx)
	defer connet.Cancle()

	fmt.Println(jobs)

	_, error := connet.PostCollection.InsertMany(connet.Ctx, jobs)
	if error != nil {
		fmt.Println(error)
	}
}

func ScrapperService() {
	scrapper.WoowaScrapper(func(a []scrapper.WohaTechDoc) {
		//deleteAll()
		jobs := existWooha(a)
		if len(a) != 0 {
			insertWooha(jobs)
		}

	})

}
