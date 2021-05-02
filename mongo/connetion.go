package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 앞에 소문자 private..s
type MongoAuth struct {
	Username string
	Password string
	Hostname string
	Clustor  string
	DbName   string
}

func getAuth() MongoAuth {
	data, err := os.Open("./mongo_auth.json")

	if err != nil {
		fmt.Println(err)
	}

	var authData MongoAuth
	byteValue, err := ioutil.ReadAll(data)

	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(byteValue, &authData)

	if err != nil {
		fmt.Println(err)
	}

	return authData
}

func GetConnetion() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	//타임아웃 설정
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	auth := getAuth()
	connectionUrl := fmt.Sprintf("mongodb+srv://%s:%s%s/%s?retryWrites=true&w=majority", auth.Username, url.QueryEscape(auth.Password), auth.Hostname, auth.Clustor)

	// 커넥션 옵션
	clinetOption := options.Client().ApplyURI(connectionUrl)
	//커넥션을 맺는다
	client, err := mongo.Connect(ctx, clinetOption)

	if err != nil {
		fmt.Println("connection Err")
		fmt.Println(err)
	}

	return client, ctx, cancel
}

//컬랙션 가져오기
func GetCollection(client *mongo.Client, colName string) *mongo.Collection {
	return client.Database(getAuth().DbName).Collection(colName)
}
