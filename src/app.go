package main

import (
	"log"
	"github.com/gin-gonic/gin"
	Endless "github.com/fvbock/endless"
	KafkaService "xmservice.com/kafka"
	Api "xmservice.com/api"
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
	
	if err := Endless.ListenAndServe("localhost:8090", router); err != nil {
		log.Fatal("Failed run app: ", err)
	}

}
















































