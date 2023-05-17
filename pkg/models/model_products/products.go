package models

import (
	"fmt"
	"project_demo/pkg/config"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Products struct {
	gorm.Model
	Id_product uuid.UUID `gorm:"type:varchar(36);primary_key" json:"id_product"`
	Name       string    `gorm:""json:"name"`
	Image      string    `json:"image"`
	Price      int64     `json:"price"`
	Amount     int64     `json:"amount"`
	Sale       int64     `json:"sale"`
	Set        string    `json:"set"`
	// CreatedAt   time.Time `gorm:"type:datetime"`
	// UpdatedAt time.Time  `gorm:"type:datetime"`
	//DeletedAt *time.Time `gorm:"type:datetime"`
}

// ID     uuid.UUID `gorm:"type:varchar(36);primary_key" json:"id"`
func (p *Products) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id_product = uuid.New()
	return nil
}
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Products{})
}

func (b *Products) Createproduct() *Products {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllProducts() []Products {
	var Products []Products
	db.Find(&Products)
	return Products
}

func GetProductById(Id string) (*Products, *gorm.DB) {
	var getProuct Products
	id, err := uuid.Parse(Id)
	if err != nil {
		fmt.Println("error while parsing UUID")
	}
	db := db.Where("Id_product=?", id).Find(&getProuct)
	return &getProuct, db
}
func DeleteProduct(Id string) Products {
	var product Products
	db.Where("id_product=?", Id).Delete(product)
	return product
}
