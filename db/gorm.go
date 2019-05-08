package db

import (
	"fmt"

	"../config"
	"../model"
	"../util/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ORM ...
var ORM, _ = GormInit()

// GormInit ...
func GormInit() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.MysqlDSL())
	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()
	db.AutoMigrate(
		&model.User{},
		&model.Company{},
		&model.Team{},
		&model.Role{},
		&model.Admin{},
		&model.ChatRoom{},
		&model.Reminder{},
	)
	fmt.Println("------Migration OK!")
	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Disable table name's pluralization
	// db.SingularTable(true)
	if config.Environment == "DEVELOPMENT" {
		// db.LogMode(true)
		// db.DropTable(&model.User{}, "UserFollower")
		// db.AutoMigrate(&model.UserFollower{})
		// db.AutoMigrate(&model.User{}, &model.Role{}, &model.Connection{}, &model.Language{}, &model.Article{}, &model.Location{}, &model.Comment{}, &model.File{})
		// db.Model(&model.User{}).AddIndex("idx_user_token", "token")

	}
	log.CheckError(err)

	return db, err
}
