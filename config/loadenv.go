package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)





func LoadVariables()   {

	err := godotenv.Load()

	if err != nil{
		log.Fatalf("Failed To Load Environmental Variables")
	}
	
}



func GetDatabaseDSN() string  {
	return os.Getenv("POSTGRES_DSN")
}


func GetJWTSecret() string  {
	return os.Getenv("SECRET_KEY")
}