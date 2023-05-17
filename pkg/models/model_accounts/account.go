package models

import (
	"fmt"
	"project_demo/pkg/config"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type TimeWrapper struct {
	time.Time
}

type Accounts struct {
	gorm.Model
	Id_account    uuid.UUID `gorm:"type:varchar(36);primary_key" json:"id_account"`
	AdminName     string    `gorm:"" json:"adminName"`
	PasswordAdmin string    `gorm:"" json:"passwordAdmin"`
}

func (p *Accounts) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id_account = uuid.New()
	return nil
}
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Accounts{})
}

func (b *Accounts) CreateAccount() *Accounts {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllAccounts() []Accounts {
	var Accounts []Accounts
	db.Find(&Accounts)
	return Accounts
}

func GetAccountById(Id string) (*Accounts, *gorm.DB) {
	var getAccount Accounts
	// db := db.Where("ID=?", Id).Find(&getProuct)
	id, err := uuid.Parse(Id)
	if err != nil {
		fmt.Println("error while parsing UUID")
	}
	db := db.Where("Id_account=?", id).Find(&getAccount)
	return &getAccount, db
}
func DeleteAccount(Id string) Accounts {
	var getAccount Accounts
	// db.Where("ID=?", Id).Delete(getAccount)
	db.Where("Id_account=?", Id).Delete(getAccount)
	return getAccount
}
