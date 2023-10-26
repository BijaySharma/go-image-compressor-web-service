package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"web-service/routes"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func MakeRequest(method string, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.ServeHTTP(writer, request)
	return writer
}

func TestCreateProduct(t *testing.T) {
	t.Run("CreateProduct", func(t *testing.T) {
		writer := MakeRequest("POST", "/products", map[string]interface{}{
			"user_id":             1,
			"product_name":        "Test Product",
			"product_description": "abc",
			"product_price":       23.4,
			"product_images":      []string{"abc", "xyz"},
		})

		if writer.Code != http.StatusCreated {
			t.Errorf("Expected status %d but got %d", http.StatusCreated, writer.Code)
		}

		var response map[string]interface{}
		json.Unmarshal(writer.Body.Bytes(), &response)

		// Check if the response is correct
		if response["status"] != "success" {
			t.Errorf("Expected status to be success but got %s", response["status"])
		}
	})
}
