package config

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Alg_jwt, Secret_jwt, Service_biodata, Link_api_services, My_token, List_services_files, Token_service string

func init() {

	Alg_jwt = getEnv("ALG_JWT", "HS256")
	Secret_jwt = getEnv("SECRET_KEY_JWT", "verysecretWOW")
	Link_api_services = getEnv("API_SERVICE_URL", "http://localhost:8081")
	Service_biodata = getEnv("LINK_BIODATA", "http://siacc-employe.rumahlogic.id")
	My_token = getEnv("MY_TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYmY1NWYwMjEtOTI3YS00MWIwLWE3NDYtMGVlZGM3ZDI1ZGZiIn0.LbeVNIrCR4bQ783jZB91u83A6NRTMy2PH4g1uH3-Wbs")
	List_services_files = getEnv("LIST_SERVICES_FILES", "list_router.json")
	Token_service = getEnv("TOKEN_SERVICE", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYjUzNTgyYTUtOWQzOS00MzFkLTlkN2ItYjI0MTEyNmE5N2I0In0.tklEGXo6P5KMWClcFhvtHuS8CwVEeZNariJr2QFTREo")

}

func DataBase() (*gorm.DB, *sql.DB) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
		getEnv("DATABASE_HOST", "localhost"),
		getEnv("DATABASE_PORT", "5432"),
		getEnv("DATABASE_USER", "postgres"),
		getEnv("DATABASE_NAME", "auth"),
		getEnv("DATABASE_PASSWORD", "postgres"),
		getEnv("SSL_MODE", "disable"),
		getEnv("TIME_ZONE", "Asia/Jakarta"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error database connection :( ")
		panic(err)
	}

	sql, err := DB.DB()

	return DB, sql
}
