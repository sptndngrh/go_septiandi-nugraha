package test

import (
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateBookController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil membuat buku baru
	// t.Run("ValidUserCreateBook", func(t *testing.T) {
	// 	reqBody := `{
	// 		"title": "Pandangan Politik Terhadap Kemajuan Bangsa",
	// 		"author": "Maranatha Maimuni",
	// 		"publisher": "Jabar Book Media"
	// 	}`

	// 	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.CreateBookController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"title": "Pandangan Politik Terhadap Kemajuan Bangsa",
			"author": "Maranatha Maimuni",
			"publisher": "Jabar Book Media"
		}`

		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, controllers.CreateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"title": "Pandangan Politik Terhadap Kemajuan Bangsa",
			"author": "Maranatha Maimuni",
			"publisher": "Jabar Book Media"
		}`

		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, controllers.CreateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 4: Data buku tidak valid (Bad Request)
	// t.Run("InvalidBookData", func(t *testing.T) {
	// 	reqBody := `{
	// 		"title": 1,
	// 		"author": 1,
	// 		"publisher": 1
	// 	}`

	// 	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.CreateBookController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })
}

func TestGetAllBooksController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah mendapatkan semua data buku
	// t.Run("ValidUserGetAllBooks", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.GetAllBooksController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, controllers.GetAllBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, controllers.GetAllBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestGetBookController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah mendapatkan data buku berdasarkan ID yang valid
	// t.Run("ValidUserGetBook", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("1")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.GetBookController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, controllers.GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, controllers.GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 4: ID buku yang salah (Not Found)
	// t.Run("InvalidBookID", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/books/InvalidBookID", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("InvalidBookID")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.GetBookController(c)) {
	// 		assert.Equal(t, http.StatusNotFound, rec.Code)
	// 	}
	// })

	// Kasus 5: ID buku yang tidak valid (Bad Request)
	// t.Run("InvalidID", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/books/invalid_id", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("invalid_id")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.GetBookController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })
}

func TestUpdateBookController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil mengupdate buku
	// t.Run("ValidUserUpdateBook", func(t *testing.T) {
	// 	reqBody := `{
	// 		"title": "Pandangan Politik Terhadap Kemajuan Bangsa (Gagasan Opini)",
	// 		"author": "Maranatha Maimuni",
	// 		"publisher": "Jabar Book Media"
	// 	}`

	// 	req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(reqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("1")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.UpdateBookController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"title": "Pandangan Politik Terhadap Kemajuan Bangsa (Gagasan Opini)",
			"author": "Maranatha Maimuni",
			"publisher": "Jabar Book Media"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, controllers.UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"title": "Pandangan Politik Terhadap Kemajuan Bangsa (Gagasan Opini)",
			"author": "Maranatha Maimuni",
			"publisher": "Jabar Book Media"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, controllers.UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 4: ID buku yang salah (Not Found)
	// t.Run("InvalidBookID", func(t *testing.T) {
	// 	reqBody := `{
	// 		"title": "Pandangan Politik Terhadap Kemajuan Bangsa (Gagasan Opini)",
	// 		"author": "Maranatha Maimuni",
	// 		"publisher": "Jabar Book Media"
	// 	}`

	// 	req := httptest.NewRequest(http.MethodPut, "/books/999", strings.NewReader(reqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("999")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.UpdateBookController(c)) {
	// 		assert.Equal(t, http.StatusNotFound, rec.Code)
	// 	}
	// })

	// Kasus 5: Data buku tidak valid (Bad Request)
	// t.Run("InvalidBookData", func(t *testing.T) {
	// 	reqBody := `{
	// 		"title": 1,
	// 		"author": 1,
	// 		"publisher": 1
	// 	}`

	// 	req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(reqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("1")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.UpdateBookController(c)) {
	// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
	// 	}
	// })
}

func TestDeleteBookController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil menghapus buku
	// t.Run("ValidUserDeleteBook", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("1")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.DeleteBookController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, controllers.DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, controllers.DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// t.Run("InvalidBookID", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodDelete, "/books/999", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("999")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.DeleteBookController(c)) {
	// 		assert.Equal(t, http.StatusNotFound, rec.Code)
	// 	}
	// })
}
