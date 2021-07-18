package main

import (
	"chat/model"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
)
import "github.com/gin-gonic/gin"

func main() {
	stdout := os.Stdout
	logfile, err := os.OpenFile("log/log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.MultiWriter(stdout, logfile))

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		} else {
			logrus.Info("user", user)
			c.JSON(http.StatusOK, gin.H{"user": user, "message": "success register"})
		}
	})

	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		} else {
			logrus.Info("user", user)
			c.JSON(http.StatusOK, gin.H{"user": user, "message": "success login"})
		}
	})

	r.Run(":5030")

}
