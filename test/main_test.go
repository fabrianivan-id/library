package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8080"

func TestLogin(t *testing.T) {
	payload := map[string]string{
		"username": "admin",
		"password": "admin123",
	}
	body, _ := json.Marshal(payload)

	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(data), "token")
}

func TestGetBooks(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"/books", nil)
	req.Header.Add("Authorization", "Bearer YOUR_TEST_TOKEN")

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(data), "title")
}

func TestAddBook(t *testing.T) {
	client := &http.Client{}
	payload := map[string]interface{}{
		"title":          "Test Book",
		"author":         "John Doe",
		"published_year": 2022,
		"stock":          10,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", baseURL+"/books", bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer YOUR_ADMIN_TOKEN")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
