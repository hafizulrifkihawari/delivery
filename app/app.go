package app

import (
	"delivery/app/config"
	"delivery/migration"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp is used to Start the Application
func StartApp() {
	setupConfig()
	initEnv()
	registerRoutes()

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err)
	}
}

// setupConfig is used to load environment variable
func setupConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	envVariable := godotenv.Load(environmentPath)
	if envVariable != nil {
		log.Fatal("Error loading .env file")
	}

	initConfig()
}

func initConfig() {
	var (
		dbHost     string = os.Getenv("DB_HOST")
		dbUser     string = os.Getenv("DB_USER")
		dbPassword string = os.Getenv("DB_PASSWORD")
		dbPort     string = os.Getenv("DB_PORT")
		dbName     string = os.Getenv("DB_NAME")
		err        error
	)

	// setup connection DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	config.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	config.DBORM, err = gorm.Open(postgres.New(postgres.Config{
		Conn: config.DB,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}

	allowCors()
	migration.InitMigration()
}

func initEnv() {
}

//endregion functions
