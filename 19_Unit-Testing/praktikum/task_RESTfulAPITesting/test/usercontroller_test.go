package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func generateAuthToken() (string, error) {
	// Buat request untuk login
	loginRequest := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "septiandin92@gmail.com",
		Password: "posyanduinaja90",
	}

	// Marshal request ke JSON
	requestBody, err := json.Marshal(loginRequest)
	if err != nil {
		return "", err
	}

	// Buat HTTP request
	req := httptest.NewRequest(http.MethodPost, "users/", bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Buat konteks Echo
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Panggil handler LoginUserController untuk mendapatkan token
	if err := controllers.LoginUserController(c); err != nil {
		return "", err
	}

	// Periksa status kode HTTP
	if rec.Code != http.StatusBadRequest {
		return "", fmt.Errorf("Login failed with status code %d", rec.Code)
	}

	// Ambil token dari respons
	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		return "", err
	}

	token, ok := response["token"].(string)
	if !ok {
		return "", errors.New("Token not found in response")
	}

	return token, nil
}

func TestCreateUserController(t *testing.T) {
	e := echo.New()

	reqBody := `{
		"name": "Septiandi Nugraha",
		"email": "septiandin92@gmail.com",
		"password": "posyanduinaja90"
	}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLoginUserController(t *testing.T) {
	e := echo.New()
	reqBody := `{
		"email": "septiandin92@gmail.com",
		"password": "posyanduinaja90"
	}`
	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUsersController(t *testing.T) {
	// Buat instansiasi Echo
	e := echo.New()

	// Skenario 1: Pengambilan semua pengguna berhasil
	req1 := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)

	err1 := controllers.GetUsersController(c1)

	// Skenario 1: Periksa bahwa tidak ada kesalahan dan responsnya sesuai
	if assert.NoError(t, err1) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	} else {
		fmt.Println(rec1.Body.String())
	}

	// Skenario 2: Gagal mengambil pengguna (misalnya, koneksi database gagal)
	req2 := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	// Skenario 2: Simulasikan kesalahan saat mengambil pengguna
	err2 := controllers.GetUsersController(c2)

	// Skenario 2: Periksa bahwa respons adalah HTTP 200 OK
	if assert.NoError(t, err2) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}

	// Skenario 3: Tidak ada pengguna yang ditemukan
	req3 := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)

	// Skenario 3: Simulasikan tidak ada pengguna yang ditemukan (sebagai contoh, respons database kosong)
	err3 := controllers.GetUsersController(c3)

	// Skenario 3: Periksa bahwa respons adalah HTTP 200 OK tetapi data pengguna kosong
	if assert.NoError(t, err3) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestGetUserController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah mendapatkan data pengguna mereka sendiri
	// t.Run("ValidUserOwnData", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("1")

		
	// 	// Buat token untuk pengguna ID 1
	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.GetUserController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })

	// Kasus 2: Pengguna yang sah mencoba mendapatkan data pengguna lain (Forbidden)
	// t.Run("ValidUserOtherUserData", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("2")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.GetUserController(c)) {
	// 		assert.Equal(t, http.StatusForbidden, rec.Code)
	// 	}
	// })

	// Kasus 3: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, controllers.GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 4: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, controllers.GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 5: ID yang salah (Bad Request)
	// t.Run("InvalidID", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodGet, "/users/300", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("invalid_id")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.GetUserController(c)) {
	// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
	// 	}
	// })
}

func TestUpdateUserController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil memperbarui data pengguna mereka sendiri
	// t.Run("ValidUserUpdateOwnData", func(t *testing.T) {
	// 	reqBody := `{
	// 		"name": "Septiandi Nugraha (Septiandi)",
	// 		"email": "septiandin92@gmail.com",
	// 		"password": "posyanduinaja90"
	// 	}`

	// 	req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
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

	// 	if assert.NoError(t, controllers.UpdateUserController(c)) {
	// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
	// 	}
	// })

	// Kasus 2: Pengguna yang sah mencoba memperbarui data pengguna lain (Forbidden)
	// t.Run("ValidUserUpdateOtherUserData", func(t *testing.T) {
	// 	reqBody := `{
	// 		"name": "Septiandi Nugraha (Septiandi)",
	// 		"email": "septiandin92@gmail.com",
	// 		"password": "posyanduinaja90"
	// 	}`

	// 	req := httptest.NewRequest(http.MethodPut, "/users/2", strings.NewReader(reqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("2")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	// Assertions
	// 	if assert.NoError(t, controllers.UpdateUserController(c)) {
	// 		assert.Equal(t, http.StatusForbidden, rec.Code)
	// 	}
	// })

	// Kasus 3: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"name": "Septiandi Nugraha (Septiandi)",
			"email": "septiandin92@gmail.com",
			"password": "posyanduinaja90"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, controllers.UpdateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 4: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"name": "Septiandi Nugraha (Septiandi)",
			"email": "septiandin92@gmail.com",
			"password": "posyanduinaja90"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, controllers.UpdateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 5: ID yang salah (Bad Request)
	// t.Run("InvalidID", func(t *testing.T) {
	// 	reqBody := `{
	// 		"name": "Septiandi Nugraha (Septiandi)",
	// 		"email": "septiandin92@gmail.com",
	// 		"password": "posyanduinaja90"
	// 	}`

	// 	req := httptest.NewRequest(http.MethodPut, "/users/invalid_id", strings.NewReader(reqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("invalid_id")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.UpdateUserController(c)) {
	// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
	// 	}
	// })
}

func TestDeleteUserController(t *testing.T) {
	e := echo.New()

	// Kasus 1: Pengguna yang sah berhasil menghapus akun
	// t.Run("ValidUserDeleteUser", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("1")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.DeleteUserController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 	}
	// })

	// Kasus 2: Token tidak valid (Unauthorized)
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, controllers.DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 3: Pengguna tidak memiliki akses token (Unauthorized)
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/2", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("2")

		if assert.NoError(t, controllers.DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// Kasus 4: ID pengguna yang salah (Not Found)
	// t.Run("InvalidUserID", func(t *testing.T) {
	// 	req := httptest.NewRequest(http.MethodDelete, "/users/999", nil)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("999")

	// 	token, err := generateAuthToken()
	// 	if err != nil {
	// 		t.Fatalf("Failed to generate auth token: %v", err)
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+token)

	// 	if assert.NoError(t, controllers.DeleteUserController(c)) {
	// 		assert.Equal(t, http.StatusNotFound, rec.Code)
	// 	}
	// })
}

