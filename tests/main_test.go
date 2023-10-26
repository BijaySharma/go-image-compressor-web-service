package tests

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"web-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func loadConfig() {
	configFilePath := flag.String("config", "../conf/", "Path to config file")
	flag.Parse()
	fmt.Println("Config file path:", *configFilePath)

	viper.SetConfigName("app")
	viper.AddConfigPath(*configFilePath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	logrus.Info("Config file loaded successfully")
}

func TestMain(m *testing.M) {
	loadConfig()
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

func TestGetProductImages(t *testing.T) {
	t.Run("GetProductImages", func(t *testing.T) {
		writer := MakeRequest("GET", "/products/images/1", nil)

		if writer.Code != http.StatusOK {
			// print the response body
			fmt.Println("Response Body", writer.Body.String())
			t.Errorf("Expected status %d but got %d", http.StatusOK, writer.Code)
		}

		var response []string
		json.Unmarshal(writer.Body.Bytes(), &response)
	})
}

func TestAddCompressedImages(t *testing.T) {
	t.Run("AddCompressedImages", func(t *testing.T) {
		writer := MakeRequest("POST", "/products/compressed-images/1", []string{"abc", "xyz"})
		if writer.Code != http.StatusCreated {
			t.Errorf("Expected status %d but got %d", http.StatusCreated, writer.Code)
		}

		var response []string
		json.Unmarshal(writer.Body.Bytes(), &response)

		// Check if the response is correct
		if len(response) != 2 {
			t.Errorf("Expected 2 images but got %d", len(response))
		}
	})
}
