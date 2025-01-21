package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func CreateDbConfig(host string, port string, user string, password string, name string) DbConfig {
	return DbConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}
}

type Db struct {
	Instance *sql.DB
	Config   DbConfig
}

func NewDb(sql *sql.DB, host string, port string, user string, password string, name string) (*Db, error) {
	return &Db{
		Instance: sql,
		Config:   CreateDbConfig(os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("NAME")),
	}, nil
}

// const (
// 	host     = "postgres-db"
// 	port     = "5432"
// 	user     = "testuser"
// 	password = "1234"
// 	dbbame   = "postgres"
// )

func (db *Db) ConnectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", db.Config.Host, db.Config.Port, db.Config.User, db.Config.Password, db.Config.Name)

	fmt.Println(psqlInfo)

	database, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to " + db.Config.Name)

	return database, nil
}
