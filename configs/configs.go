package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type setting struct {
	Host     string
	User     string
	Password string
	Port     string
	DBNAME   string
}

func ImportPasskey() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return os.Getenv("passkey")
}
func ImportserverKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return os.Getenv("midtranskey")
}

func ImportSetting() setting {
	var result setting
	err := godotenv.Load(".env")
	if err != nil {
		return setting{}
	}
	result.Host = os.Getenv("poshost")
	result.User = os.Getenv("posuser")
	result.Password = os.Getenv("pospw")
	result.Port = os.Getenv("posport")
	result.DBNAME = os.Getenv("dbname")
	return result
}

func ConnectDB(s setting) (*gorm.DB, error) {
	var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBNAME)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	return db, nil
}
