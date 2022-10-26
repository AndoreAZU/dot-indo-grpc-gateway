package repository

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DAO interface {
	NewLoginQuery() LoginQuery
	NewRegisterQuery() RegisterQuery
	NewGetUserQuery() GetUserQuery
	NewUpdateQuery() UpdateQuery
	NewFollowQuery() FollowQuery
}

type dao struct {
}

var DB *gorm.DB

func NewDAO(db *gorm.DB) DAO {
	DB = db
	return &dao{}
}

func NewDBGorm() *gorm.DB {
	logrus.Error("initializing connection db...")

	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s port=%s",
		os.Getenv("DBHOST"), os.Getenv("DBNAME"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("SSLMODE"), os.Getenv("DBPORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("error while connection db: ", err)
		return nil
	}

	logrus.Error("finish initialize connection db")
	return db
}

func (x *dao) NewLoginQuery() LoginQuery {
	return &loginQuery{}
}

func (x *dao) NewRegisterQuery() RegisterQuery {
	return &registerQuery{}
}

func (x *dao) NewGetUserQuery() GetUserQuery {
	return &getUserQuery{}
}

func (x *dao) NewUpdateQuery() UpdateQuery {
	return &updateQuery{}
}

func (x *dao) NewFollowQuery() FollowQuery {
	return &followQuery{}
}
