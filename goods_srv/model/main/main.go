package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	// "strings"

	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/anaskhan96/go-password-encoder"
	"mxshop_srvs/goods_srv/model"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString((Md5.Sum(nil)))
}

func main() {
	dsn := "root:@tcp(localhost:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodePwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodePwd)
	fmt.Println(newPassword)

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("boy%d", i),
			Mobile:   fmt.Sprintf("1321111111%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}

	// _ = db.AutoMigrate(&model.User{})

	// options := &password.Options{16, 100, 32, sha512.New}
	// salt, encodePwd := password.Encode("generic password", options)
	// newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodePwd)
	// fmt.Println(newPassword)

	// passwordInfo := strings.Split(newPassword, "$")
	// fmt.Println(passwordInfo)
	// check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	// fmt.Println(check)
}
