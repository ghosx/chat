package db

import (
	"chat/model"
	"github.com/sirupsen/logrus"
)

func Create(user model.User) {
	// 插入数据
	stmt, err := DB.Prepare("INSERT INTO userinfo(username, password) values(?,?)")
	if err != nil {
		logrus.Error("insert error")
	}

	res, err := stmt.Exec(user.Username, user.Password)

	if err != nil {
		logrus.Error("insert error")
	}

	id, err := res.LastInsertId()

	if err != nil {
		logrus.Error("insert error")
	}
	logrus.Println("id", id)

}
