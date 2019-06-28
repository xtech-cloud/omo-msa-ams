package model

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	uuid "github.com/satori/go.uuid"
)

var base64Coder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

var dialectName string
var dialectArgs string
var saltSuffix string

func SetupEnv() {
	saltSuffix = os.Getenv("AMS_SALT_SUFFIX")
}

func AutoMigrateDatabase() {
	dialectName = os.Getenv("AMS_DATABASE_DRIVER")
	if "" == dialectName {
		panic("AMS_DATABASE_DRIVER not found")
	}
	mysql_addr := os.Getenv("AMS_MYSQL_ADDR")
	mysql_user := os.Getenv("AMS_MYSQL_USER")
	mysql_passwd := os.Getenv("AMS_MYSQL_PASSWORD")
	mysql_database := os.Getenv("AMS_MYSQL_DATABASE")
	sqlite_filepath := os.Getenv("AMS_SQLITE_FILEPATH")

	if "mysql" == dialectName {
		if "" == mysql_user {
			panic("AMS_MYSQL_USER not found")
		}
		if "" == mysql_passwd {
			panic("AMS_MYSQL_PASSWORD not found")
		}
		if "" == mysql_database {
			panic("AMS_MYSQL_DATABASE not found")
		}
		if "" == mysql_addr {
			panic("AMS_MYSQL_ADDR not found")
		}
	} else if "sqlite" == dialectName {
		dialectName = "sqlite3"
		if "" == sqlite_filepath {
			panic("AMS_SQLITE_FILEPATH not found")
		}
	}

	if "mysql" == dialectName {
		dialectArgs = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True", mysql_user, mysql_passwd, mysql_addr, mysql_database)
	} else if "sqlite3" == dialectName {
		dialectArgs = sqlite_filepath
	}
	db, err := openDB()
	if nil != err {
		panic(err)
	}
	defer closeDB(db)

	db.AutoMigrate(&Account{})
}

func openDB() (*gorm.DB, error) {
	return gorm.Open(dialectName, dialectArgs)
}

func closeDB(_db *gorm.DB) {
	_db.Close()
}

func NewUUID() string {
	guid, _ := uuid.NewV4()
	h := md5.New()
	h.Write(guid.Bytes())
	return hex.EncodeToString(h.Sum(nil))
}

func ToUUID(_content string) string {
	h := md5.New()
	h.Write([]byte(_content))
	return hex.EncodeToString(h.Sum(nil))
}

func ToBase64(_content []byte) string {
	return base64Coder.EncodeToString(_content)
}
