package migration

import (
	Entity "xmservice.com/entity"
	"time"
    "fmt"
	"log"
	"os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"github.com/joho/godotenv"
)

type Migration interface {
	MigrateTables() *gorm.DB
	ConnectToDb() *gorm.DB
}

type MigrationService struct {
   
}

func (migration *MigrationService) MigrateTables() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var connectionString string = os.Getenv("POSTGRES_HOST") + "/" + os.Getenv("POSTGRES_DB")
	// dsn := "host=db-postgresql-sfo3-21964-nov-29-backup-do-user-9772821-0.b.db.ondigitalocean.com user=mh-production password=AVNS_RB0gik8akCPKDtOVoPB dbname=mh-production-db port=25060 sslmode=require TimeZone=UTC";
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
    }), &gorm.Config{})

    if err != nil {
        log.Fatal("Cannot connect to DB at this time, please try again");
    }

    db.AutoMigrate(&Entity.Company{});
	db.AutoMigrate(&Entity.KafkaEvent{});

    postgresDB, err1 := db.DB();
    if err1 == nil {
        postgresDB.SetConnMaxLifetime(time.Minute * 10);
        fmt.Println("Database connection timeout has been set to 10mins")
    }
    
	return db;

}

func (migration *MigrationService) ConnectToDb() *gorm.DB {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var connectionString string = os.Getenv("POSTGRES_HOST") + "/" + os.Getenv("POSTGRES_DB")
	// + "user=" + os.Getenv("POSTGRES_USER") + " " + "password=" + os.Getenv("POSTGRES_PASSWORD") + " " + "dbname=" + os.Getenv("POSTGRES_DB")
	// dsn := "host=db-postgresql-sfo3-21964-nov-29-backup-do-user-9772821-0.b.db.ondigitalocean.com user=mh-production password=AVNS_RB0gik8akCPKDtOVoPB dbname=mh-production-db port=25060 sslmode=require TimeZone=UTC";
    db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
    }), &gorm.Config{});

	if err != nil {
        log.Fatal("Cannot connect to DB at this time, please try again");
    }

    postgresDB, err1 := db.DB();
    if err1 == nil {
        postgresDB.SetConnMaxLifetime(time.Minute * 10)
        fmt.Println("Database connection timeout has been set to 10mins")
    }

    return db;
}








