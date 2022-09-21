package controllers

import (
	"bookApp/internal/models"
	"bookApp/internal/repository"
	"bookApp/pkg/utils"
)

type BookControllers struct {
	Libbook repository.BookRepo
}

type BookControllersInterface interface {
	CreateBook(book models.Book) (*utils.Response, error)
	GetAllBooks() (*utils.Response, error)
	GetBookById(id int) (*utils.Response, error)
	UpdateBook(id int, book *models.Book) (*utils.Response, error)
	DeleteBook(id int) (*utils.Response, error)
}

func InitBook(Libbook repository.BookRepo) BookControllersInterface {
	return &BookControllers{Libbook}
}

func (bc BookControllers) CreateBook(book models.Book) (*utils.Response, error) {
	dataBook, err := bc.Libbook.CreateBook(book)
	if err != nil {
		return nil, err
	}

	data := &utils.Response{
		Code:    201,
		Status:  "Success",
		Message: "Success to create new book",
		Result:  dataBook,
	}

	return data, err
}

func (bc BookControllers) GetAllBooks() (*utils.Response, error) {
	var arrBooks []utils.BookData

	book, err := bc.Libbook.GetAllBooks()
	if err != nil {
		return nil, err
	}

	for _, v := range *book {
		dataBook := utils.BookData{
			ID:     int(v.ID),
			Title:  v.Title,
			ISBN:   v.ISBN,
			Writer: v.Writer,
		}

		arrBooks = append(arrBooks, dataBook)
	}

	data := &utils.Response{
		Code:    200,
		Status:  "Success",
		Message: "Success to get all book",
		Result:  arrBooks,
	}

	return data, nil
}

func (bc BookControllers) GetBookById(id int) (*utils.Response, error) {
	book, err := bc.Libbook.GetBookById(id)
	if err != nil {
		return nil, err
	}

	dataBook := &utils.BookData{
		ID:        int(book.ID),
		Title:     book.Title,
		ISBN:      book.ISBN,
		Writer:    book.Writer,
		CreatedAt: book.CreatedAt,
	}

	data := &utils.Response{
		Code:    200,
		Status:  "Success",
		Message: "Success get data user by ID",
		Result:  dataBook,
	}

	return data, nil
}

func (bc BookControllers) UpdateBook(id int, book *models.Book) (*utils.Response, error) {
	bookData := &models.Book{
		Title:  book.Title,
		ISBN:   book.ISBN,
		Writer: book.Writer,
	}

	err := bc.Libbook.UpdateBook(id, bookData)
	if err != nil {
		return nil, err
	}

	data := &utils.Response{
		Code:    201,
		Status:  "Success",
		Message: "Success to update book",
		Result:  err,
	}

	return data, nil
}

func (bc BookControllers) DeleteBook(id int) (*utils.Response, error) {
	err := bc.Libbook.DeleteBook(id)
	if err != nil {
		return nil, err
	}

	data := &utils.Response{
		Code:    200,
		Status:  "Success",
		Message: "Success data book delete",
		Result:  err,
	}

	return data, nil
}
