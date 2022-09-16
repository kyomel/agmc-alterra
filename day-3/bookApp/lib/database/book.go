package database

import (
	"bookApp/models"

	"gorm.io/gorm"
)

type LibBook struct {
	DB *gorm.DB
}

type BookRepo interface {
	CreateBook(book models.Book) (*models.Book, error)
	GetAllBooks() (*[]models.Book, error)
	GetBookById(id int) (models.Book, error)
	UpdateBook(id int, book *models.Book) error
	DeleteBook(id int) error
}

func InitBook(DB *gorm.DB) BookRepo {
	return &LibBook{DB}
}

func (b *LibBook) CreateBook(book models.Book) (*models.Book, error) {
	if err := b.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *LibBook) GetAllBooks() (*[]models.Book, error) {
	var (
		books *[]models.Book
		err   error
	)

	if err = b.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (b *LibBook) GetBookById(id int) (models.Book, error) {
	var book models.Book

	res := b.DB.Where("id = ?", id).First(&book)
	return book, res.Error
}

func (b *LibBook) UpdateBook(id int, book *models.Book) error {
	if err := b.DB.Model(book).Where("id = ?", id).Updates(book).Error; err != nil {
		return err
	}

	return nil
}

func (b *LibBook) DeleteBook(id int) error {
	var book *models.Book

	if err := b.DB.Where("id = ?", id).Delete(&book, id).Error; err != nil {
		return err
	}

	return nil
}
