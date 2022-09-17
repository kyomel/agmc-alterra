package helpers_test

import (
	"bookApp/config"
	"bookApp/controllers"
	"bookApp/helpers"
	"bookApp/lib/database"
	"bookApp/middleware"
	"bookApp/mock"
	"bookApp/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	t.Parallel()

	var (
		echoMock   = mock.EchoMock{E: echo.New()}
		databases  = database.InitBook(config.InitDB())
		controller = controllers.InitBook(databases)
		h          = helpers.InitBook(controller)
	)

	t.Run("Success", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/v1/books")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetAllBooks(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
		c.SetPath("/v1/books")
		c.SetParamNames("keywords")
		c.SetParamValues("test")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetAllBooks(c)) {
			asserts.Equal(404, rec.Code)
		}
	})
}

func TestCreateBook(t *testing.T) {
	t.Parallel()

	var (
		echoMock    = mock.EchoMock{E: echo.New()}
		databases   = database.InitBook(config.InitDB())
		controller  = controllers.InitBook(databases)
		h           = helpers.InitBook(controller)
		payloadBook = &models.Book{
			Title:  "Mencari Cinta",
			ISBN:   "123-456",
			Writer: "Francesca Vania",
		}
		userId = int(1)
	)

	t.Run("Success", func(t *testing.T) {
		payload, err := json.Marshal(payloadBook)
		if err != nil {
			t.Fatal(err)
		}

		c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
		c.SetPath("/jwt/v1/books/")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.CreateBook(c)) {
			asserts.Equal(201, rec.Code)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		payload, err := json.Marshal(payloadBook)
		if err != nil {
			t.Fatal(err)
		}

		c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
		c.SetPath("/jwt/v1/books/")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.CreateBook(c)) {
			asserts.Equal(401, rec.Code)
		}
	})
}

func TestGetBookById(t *testing.T) {
	t.Parallel()

	var (
		databases  = database.InitBook(config.InitDB())
		controller = controllers.InitBook(databases)
		h          = helpers.InitBook(controller)
	)

	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/v1/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetBookById(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("Falied", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/v1/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("100")

		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetBookById(c)) {
			asserts.Equal(404, rec.Code)
		}
	})
}

func TestUpdateBook(t *testing.T) {
	t.Parallel()

	var (
		echoMock    = mock.EchoMock{E: echo.New()}
		databases   = database.InitBook(config.InitDB())
		controller  = controllers.InitBook(databases)
		h           = helpers.InitBook(controller)
		payloadBook = &models.Book{
			Title:  "Mencari Cinta Update",
			ISBN:   "123-456-7",
			Writer: "Francesca Vania",
		}
		userId = int(1)
	)

	t.Run("Success", func(t *testing.T) {
		payload, err := json.Marshal(payloadBook)
		if err != nil {
			t.Fatal(err)
		}

		c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
		c.SetPath("/v1/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.UpdateBook(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		payload, err := json.Marshal(payloadBook)
		if err != nil {
			t.Fatal(err)
		}

		c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
		c.SetPath("/v1/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("100")
		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.UpdateBook(c)) {
			asserts.Equal(404, rec.Code)
		}
	})
}
func TestDeleteBook(t *testing.T) {
	t.Parallel()

	var (
		databases  = database.InitBook(config.InitDB())
		controller = controllers.InitBook(databases)
		h          = helpers.InitBook(controller)
		echoMock   = mock.EchoMock{E: echo.New()}
		userId     = int(1)
	)

	t.Run("Success", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
		c.SetPath("/v1/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetBookById(c)) {
			asserts.Equal(200, rec.Code)
		}
	})

	t.Run("Failed", func(t *testing.T) {
		c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
		c.SetPath("/v1/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("100")

		token, err := middleware.CreateToken(userId)
		if err != nil {
			t.Fatal(err)
		}
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		//! asserts
		asserts := assert.New(t)
		if asserts.NoError(h.GetBookById(c)) {
			asserts.Equal(404, rec.Code)
		}
	})
}
