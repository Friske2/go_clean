package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_todoHttpDelivery "go_clean/todolist/delivery/http"
	_todoRepo "go_clean/todolist/repository/mysql"
	_todoUseCase "go_clean/todolist/usecase"
	"log"
	"net/http"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}
func main() {
	port := viper.GetString("server.address")
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.Get(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	todoRepo := _todoRepo.NewMysqlAuthorRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	td := _todoUseCase.NewTodoListUseCase(todoRepo, timeoutContext)
	_todoHttpDelivery.NewTodolistHandler(r, td)
	log.Fatal(r.Run(port))
}
