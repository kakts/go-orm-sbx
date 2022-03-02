package repository

import (
	"fmt"
	"strconv"

	"github.com/kakts/go-orm-sbx/gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

var (
	users models.Users
)

const (
	DBUser     = "user"
	DBPass     = "pass"
	DBProtocol = "tcp(127.0.0.1:3306)"
	DBName     = "test_database"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Port       int
	Connection *gorm.DB
}

func InitDB() {
	fmt.Println("NewDB called")
	c := NewConfig()
	newDB(&DB{
		Host:     c.DB.Local.Host,
		Port:     c.DB.Local.Port,
		Username: c.DB.Local.Username,
		Password: c.DB.Local.Password,
		DBName:   c.DB.Local.DBName,
	})
}

// makeDsn returns dsn
func makeDsn(d *DB) string {
	return d.Username + ":" + d.Password + "@tcp(" + d.Host + ":" + strconv.Itoa(d.Port) + ")/" + d.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func newDB(d *DB) *DB {
	// mysql

	dsn := makeDsn(d)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	autoMigrate()
	return d
}

// Begin begins a transaction
func (d *DB) Begin() *gorm.DB {
	return d.Connection.Begin()
}

func (d *DB) Connect() *gorm.DB {
	return d.Connection
}

// autoMigrate migrates DB tables from models
func autoMigrate() {
	fmt.Println("[autoMigrate] migrate DB tables from models.")
	db.AutoMigrate(&users)
}

// GetDB gets db instance.
func GetDB() *gorm.DB {
	return db
}
