package test

import (
	"chat/db"
	"chat/model"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestUserCreate(t *testing.T) {
	u := model.User{
		Username: "jack",
		Password: "123456",
	}
	logrus.Println(u)
	db.Create(u)
}
