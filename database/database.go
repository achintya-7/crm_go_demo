package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// referencing to the SQL ORM
var (
	DBConn *gorm.DB
)

