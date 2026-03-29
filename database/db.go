package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PW")
    name := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
    ssl := os.Getenv("DB_SSL")

    log.Printf("DB_HOST: '%s'", host)
    log.Printf("DB_USER: '%s'", user)
    log.Printf("DB_NAME: '%s'", name)
    log.Printf("DB_PORT: '%s'", port)
    log.Printf("DB_SSL: '%s'", ssl)

    fmt.Println("DB_HOST: '%s'", host)
    fmt.Println("DB_USER: '%s'", user)
    fmt.Println("DB_NAME: '%s'", name)
    fmt.Println("DB_PORT: '%s'", port)
    fmt.Println("DB_SSL: '%s'", ssl)

	if ssl == nil || ssl == "" {
		ssl = "require"
	}

    if host == "" || user == "" || name == "" || port == "" || ssl == "" {
     	fmt.Println("Variáveis de ambiente não carregadas corretamente")
        log.Panic("Variáveis de ambiente não carregadas corretamente")
    }

    stringDeConexao := "host="+ host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=" + ssl
// 	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Printf("Erro ao conectar com banco de dados: %v", err)
// 		os.Exit(1)
	}

	DB.AutoMigrate(&models.Aluno{})
}
