package database

import (
	"backend/common"
	"context"
	_ "fmt"
	"log"

	utils "github.com/ItsMeSamey/go_utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database
var UserDB Collection[User]
var ClientDB Collection[Client]
var SentMailDB Collection[Client]

func init() {
	
	client, err := mongo.Connect(options.Client().ApplyURI(common.Cfg.MongoURI))

	if err != nil {
		log.Fatalln(utils.WithStack(err))
	}

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalln(utils.WithStack(err))
		panic(err)
	}

	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	DB = client.Database(common.Cfg.DBName)
	UserDB = Collection[User]{DB.Collection("users")}
	ClientDB = Collection[Client]{DB.Collection("clients")}
	SentMailDB = Collection[Client]{DB.Collection("sent-mails")}
	log.Println(UserDB.Collection.Name())
}
