package controllers

import (
	"context"
	"fmt"
	"healthcare/constanta"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

func Chatbot(c echo.Context) error {
	var request web.ChatbotRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidBody))
	}

	if err := helper.ValidateStruct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	filePath := "utils/helper/prompt/prompt.txt"
	promptContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	response, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("%s %s", string(promptContent), request.Request),
			},
		},
	})

	// fmt.Println(string(promptContent), request.Request) // for debugging

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet+"recommendation"))
	}

	if response.Choices == nil || len(response.Choices) == 0 {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet+"recommendation"))
	}

	aiResponse := response.Choices[0].Message.Content
	if strings.Contains(strings.ToLower(aiResponse), "maaf") {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(aiResponse))
	} else {
		return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"recommendation", aiResponse))
	}
}

func CustomerService(c echo.Context) error {
	var request web.ChatbotRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidBody))
	}

	if err := helper.ValidateStruct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if strings.Contains(strings.ToLower(request.Request), "pembayaran obat") {
		filePath := "utils/helper/prompt/content1.txt"
		content1, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"recommendation", string(content1)))
	} else if strings.Contains(strings.ToLower(request.Request), "rating dokter") {
		filePath := "utils/helper/prompt/content2.txt"
		content2, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"recommendation", string(content2)))
	}

	filePath := "utils/helper/prompt/content3.txt"
	content3, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"recommendation", string(content3)))

}
