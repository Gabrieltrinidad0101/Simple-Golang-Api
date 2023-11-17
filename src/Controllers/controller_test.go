package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	structs "main/src/Structs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

var user = structs.User{
	Name:     "jose",
	Password: "123456",
}

var jwtHeader = "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"

func TestLogin(t *testing.T) {
	e := echo.New()
	inputJSON, err := json.Marshal(&user)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)

	var apiResult structs.ApiResult
	if err := json.Unmarshal(rec.Body.Bytes(), &apiResult); err != nil {
		t.Fatal(err)
	}

	tokenParts := strings.Split(apiResult.Data.(string), ".")
	decodedBytes, err := base64.StdEncoding.DecodeString(tokenParts[0])
	if len(tokenParts) != 3 || err != nil || string(decodedBytes) != jwtHeader {
		t.Fatal("Error format in the jwt")
	}
}
