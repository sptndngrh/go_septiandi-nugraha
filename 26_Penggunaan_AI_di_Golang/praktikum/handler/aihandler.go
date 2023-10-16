package handler

import (
	"fmt"
	"net/http"
	"os"
	"praktikum_26/usecase"

	"github.com/labstack/echo/v4"
)

type AILaptopHandler struct {
	AILaptopUsecase usecase.AILaptopUsecase // Inject the use case
}

type AIresponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func NewAILaptopHandler(usecase usecase.AILaptopUsecase) *AILaptopHandler {
	return &AILaptopHandler{
		AILaptopUsecase: usecase,
	}
}

func (h *AILaptopHandler) AIRecommended(c echo.Context) error {
	var requestData map[string]interface{}
	err := c.Bind(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	money, moneyFix := requestData["money"].(float64)
	purpose, purposeFix := requestData["purpose"].(string)
	brand, brandFix := requestData["brand"].(string)
	ram, ramFix := requestData["ram"].(string)
	cpu, cpuFix := requestData["cpu"].(string)
	sizeScreen, sizeScreenFix := requestData["size_screen"].(string)

	if !moneyFix || !purposeFix || !brandFix || !ramFix || !cpuFix || !sizeScreenFix {
		return c.JSON(http.StatusBadRequest, "Invalid request format")
	}

	userInput := fmt.Sprintf("Rekomendasi Laptop merk %s untuk %s dengan spesifikasi %s ram %s cpu %s ukuran layar dan dengan uang sebesar Rp %.0f.", brand, purpose, ram, cpu, sizeScreen, money)

	answer, err := h.AILaptopUsecase.AIRecommended(userInput, brand, ram, cpu, sizeScreen, os.Getenv("alta_key_tugas26"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error dalam mmebuat rekomendasi")
	}

	responseData := AIresponse{
		Status: "success",
		Data:   answer,
	}

	return c.JSON(http.StatusOK, responseData)
}
