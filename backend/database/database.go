package database

import (
	"backend/gen/models"
	"github.com/jinzhu/gorm"
)

type Database struct {
	db *gorm.DB
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		panic(err)
	}
}

func NewDatabase() *Database {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.StoredBook{})
	return &Database{db: db}
}

func (d *Database) Create(book *models.Book) {
	d.db.Create(book)
}

func (d *Database) FetchAll() *[]models.StoredBook {
	var books []models.StoredBook
	d.db.Find(&books)
	return &books
}

func (d *Database) Delete(id int64) {
	m := models.StoredBook{}
	m.ID = &id
	d.db.Delete(m)
}
