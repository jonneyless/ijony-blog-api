package common

import (
    "fmt"
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type DatabaseParams struct {
    Host       string
    UserName   string
    Password   string
    Database   string
    AuthSource string
}

type Database struct {
    dsn string
    db  *gorm.DB
}

var database *Database

func InitDatabase(c *DatabaseParams) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.UserName, c.Password, c.Host, c.Database)
    database = &Database{dsn: dsn}
}

func GetDatabase() *Database {
    return database
}

func DB() *gorm.DB {
    return database.db
}

func (d *Database) Connect() *gorm.DB {
    var err error

    d.db, err = gorm.Open(mysql.Open(d.dsn), &gorm.Config{})
    if err != nil {
        log.Panicln(err)
    }

    return d.db
}
