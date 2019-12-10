package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.HandlerFunc, method, path string, body io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, body)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	return rec, err
}

func TestLoginUserHandler(t *testing.T) {
	//reqBody := gin.H{
	//	"login": "migmatore",
	//	"password": "12345678",
	//}
	//
	//_router := router.SetupRouter()
	//
	//req, err := performRequest(_router, "GET", "/", url.Values{"login": {"migmatore"}, "password": {"12345678"}})
	//if err != nil {
	//	fmt.Errorf("Error: %v", err.Error())
	//}
}
