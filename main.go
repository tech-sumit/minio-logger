package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
)

func main(){
	host := os.Getenv("HOST")
	logfile := os.Getenv("LOG_FILE")

	router := gin.Default()
	router.POST("/log", func(c *gin.Context) {
		var data string
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		str := buf.String()
		str=str+","
		log.Print(data)
		f, err := os.OpenFile(logfile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		if _, err := f.WriteString(str); err != nil {
			log.Println(err)
		}
		c.Status(http.StatusOK)
	})
	err := router.Run(host)
	if err!=nil{
		log.Fatal(err)
	}
}