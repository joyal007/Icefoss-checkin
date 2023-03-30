package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoUrl string

func getenv(){
	if err:=godotenv.Load(".env");err != nil{
		log.Fatal("")
	}
	MongoUrl = os.Getenv("MONGODB_URL")
}

func Connect(){
	getenv()
	err := mgm.SetDefaultConfig(nil, "hackfit2023", options.Client().ApplyURI(MongoUrl))
	if err != nil{
		log.Fatal("Error while connecting to DB")
	}else{
		log.Println("Successfully Connected to DB")
	}
}