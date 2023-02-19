package main

import (
	"log"
	"github.com/gin-gonic/gin"
	Endless "github.com/fvbock/endless"
	KafkaService "xmservice.com/kafka"
	Api "xmservice.com/api"
	"github.com/joho/godotenv"
	"os"
	// Migration "xmservice.com/migration"
)

func main(){

    router := gin.Default();
	router.MaxMultipartMemory = 100 << 20  // 50 MiB

	// migrateDatabase :=  func(){
	// 	var migration Migration.Migration = &Migration.MigrationService{}
	// 	migration.MigrateTables();
	// } 

	// migrateDatabase();

	router.MaxMultipartMemory = 32 << 20;

	//KAFKA SERVICE;
	var kafkaService KafkaService.KafkaService;
	kafkaService.CreateKafKaTopic();

	var kafkaRecieverService KafkaService.KafkaRecieverService;
	go kafkaRecieverService.ConsumeKafkaTopic();

	setRoutes := func(){
		var companyApi Api.CompanyApi;
		var authApi Api.AuthApi;
		companyApi.Router(router);
		authApi.Router(router)
	}

    setRoutes();

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var host string = "localhost:" + os.Getenv("LISTENING_PORT");
	if err := Endless.ListenAndServe(host, router); err != nil {
		log.Fatal("Failed run app: ", err)
	}

}
















































