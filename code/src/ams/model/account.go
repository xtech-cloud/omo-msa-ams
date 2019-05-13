package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Account struct {
	Embedded gorm.Model `gorm:"embedded"`
	UUID     string     `gorm:"column:uuid;not null;unique"`
	Username string     `gorm:"column:username;not null;unique"`
	Password string     `gorm:"column:password;not null"`
	Profile  string     `gorm:"column:profile"`
}

func (Account) TableName() string {
	return "Account"
}

type AccountDAO struct {
}

func NewAccountDAO() *AccountDAO {
	return &AccountDAO{}
}

func (AccountDAO) Insert(_account Account) error {
	db, err := openDB()
	if nil != err {
		return err
	}
	defer closeDB(db)

	isBlank := db.NewRecord(_account)
	if !isBlank {
		return errors.New("account is exists")
	}
	return db.Create(&_account).Error
}

/*
func (AccountDAO) Upsert(_account Account) error {
	var account Account
	err := db.Where(Account{UUID: _account.UUID}).FirstOrCreate(&account).Error
	if nil != err {
		return err
	}
	err = db.Model(&account).Updates(_account).Error
	return err
}

func (AccountDAO) List() ([]Account, error) {
	var accounts []Account
	err := db.Find(&accounts).Error
	return accounts, err
}
*/

func (AccountDAO) Find(_uuid string) (Account, error) {
	var account Account
	db, err := openDB()
	if nil != err {
		return account, err
	}
	defer closeDB(db)

	err = db.Where("uuid = ?", _uuid).First(&account).Error
	return account, err
}

func (AccountDAO) WhereUsername(_username string) (Account, error) {
	var account Account
	db, err := openDB()
	if nil != err {
		return account, err
	}
	defer closeDB(db)

	res := db.Where("username= ?", _username).First(&account)
	if res.RecordNotFound() {
		return Account{}, nil
	}
	return account, res.Error
}
