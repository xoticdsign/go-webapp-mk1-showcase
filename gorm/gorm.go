package gorm

import (
	"go-webapp-mark1-showcase/gobcrypt"

	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GLOBAL VARIABLES

type User struct {
	ID       uint   `gorm:"type:BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	Role     string `gorm:"type:VARCHAR(50) NOT NULL"`
	Username string `gorm:"type:VARCHAR(50) NOT NULL UNIQUE"`
	Nickname string `gorm:"type:VARCHAR(50)"`
	Bio      string `gorm:"type:VARCHAR(50)"`
	Bday     time.Time
	Password string `gorm:"type:VARCHAR(150) NOT NULL"`
	Email    string `gorm:"type:VARCHAR(100) NOT NULL UNIQUE"`
}

var users = []User{
	{Role: "admin", Username: "admin", Nickname: "bigbrother", Bday: time.Date(2004, 02, 04, 00, 00, 00, 00, time.UTC), Password: "admin", Email: "admin@localhost.com"},
	{Role: "user", Username: "john", Nickname: "johny1", Bio: "Your fav", Password: "john15", Email: "john15@gmail.com"},
	{Role: "user", Username: "elvira", Nickname: "scarface", Bio: "+++", Bday: time.Date(1996, 10, 31, 00, 00, 00, 00, time.UTC), Password: "ElviraSanders1", Email: "elvira_sanders@yandex.ru"},
	{Role: "user", Username: "greg", Bday: time.Date(1999, 02, 20, 00, 00, 00, 00, time.UTC), Password: "gregIsCool", Email: "gregIsCool@outlook.com"},
	{Role: "user", Username: "will", Password: "will_from_boston", Email: "will_from_boston@mail.ru"},
}

var db *gorm.DB

// GORM

func ConfigPostgreSQL() error {
	var (
		host     string
		user     string
		password string
		dbname   string
		sslmode  string
	)

	var err error

	host = os.Getenv("POSTGRESQL_HOST")
	user = os.Getenv("POSTGRESQL_USER")
	password = os.Getenv("POSTGRESQL_PASSWORD")
	dbname = os.Getenv("POSTGRESQL_DBNAME")
	sslmode = os.Getenv("POSTGRESQL_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbname, sslmode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		ok := tx.Migrator().HasTable(&User{})
		if !ok {
			err := tx.Migrator().CreateTable(&User{})
			if err != nil {
				return err
			}
			for _, val := range users {
				passwordCrypt, err := gobcrypt.Encrypt(val.Password)
				if err != nil {
					return err
				}
				user := User{
					Role:     val.Role,
					Username: val.Username,
					Nickname: val.Nickname,
					Bio:      val.Bio,
					Bday:     val.Bday,
					Password: passwordCrypt,
					Email:    val.Email,
				}

				tx := tx.Table("users").Create(&user)
				if tx.Error != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func SelectUser(username string, password string) (User, error) {
	var userCreds User

	tx := db.Table("users").Where("username=?", username).Find(&userCreds)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	ok := gobcrypt.Decrypt(userCreds.Password, password)
	if !ok {
		return User{}, bcrypt.ErrMismatchedHashAndPassword
	}
	return userCreds, nil
}

func SelectProfile(username string) (User, error) {
	var userDets User

	tx := db.Table("users").Where("username=?", username).Find(&userDets)
	if tx.Error != nil {
		return User{}, tx.Error
	}
	return userDets, nil
}

func CreateUser(username string, nickname string, bio string, password string, email string) (User, error) {
	var userCreds User

	err := db.Transaction(func(tx *gorm.DB) error {
		passwordCrypt, err := gobcrypt.Encrypt(password)
		if err != nil {
			return err
		}

		tx = tx.Table("users").Create(&User{Role: "user", Username: username, Nickname: nickname, Bio: bio, Password: passwordCrypt, Email: email})
		if tx.Error != nil {
			return tx.Error
		}

		tx = tx.Table("users").Where("username=?", username).Find(&userCreds)
		if tx.Error != nil {
			return tx.Error
		}

		ok := gobcrypt.Decrypt(userCreds.Password, password)
		if !ok {
			return bcrypt.ErrMismatchedHashAndPassword
		}
		return nil
	})
	if err != nil {
		return User{}, err
	}
	return userCreds, nil
}

func UpdateUser(username string, nickname string, bio string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Table("users").Where("username=?", username).Updates(User{Nickname: nickname, Bio: bio})
		if tx.Error != nil {
			return tx.Error
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
