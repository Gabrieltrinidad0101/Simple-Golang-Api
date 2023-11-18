package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	structs "main/src/Structs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"syreclabs.com/go/faker"
)

func makeUser() structs.User {
	return structs.User{
		Name:     fmt.Sprintf("%s - %s", faker.Name().Name(), uuid.New()),
		Password: faker.Internet().Password(4, 5),
	}
}

var Users = []structs.User{makeUser(), makeUser(), makeUser()}

var jwtHeader = "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"

func Auth(user structs.User, url string) (apiResult structs.ApiResult, err error) {
	e := echo.New()
	inputJSON, err := json.Marshal(&user)

	if err != nil {
		return
	}

	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err != nil {
		return
	}

	userController := UserController{}
	if url == "/login" {
		userController.Login(c)
	} else if url == "/register" {
		userController.Register(c)
	}

	if err = json.Unmarshal(rec.Body.Bytes(), &apiResult); err != nil {
		return
	}

	if err = verifyJwt(apiResult); err != nil {
		return
	}

	return
}

func verifyJwt(apiResult structs.ApiResult) error {
	if apiResult.Data == nil {
		return fmt.Errorf("jwt is null")
	}

	tokenParts := strings.Split(apiResult.Data.(string), ".")
	decodedBytes, err := base64.StdEncoding.DecodeString(tokenParts[0])
	if err != nil || len(tokenParts) != 3 || string(decodedBytes) != jwtHeader {
		return fmt.Errorf("Error format in the jwt %s", apiResult.Data)
	}

	return nil
}

func TestLogin(t *testing.T) {
	_, err := Auth(Users[0], "/register")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoginWithoutAccount(t *testing.T) {
	apiResult, _ := Auth(Users[1], "/login")
	if apiResult.Error != "The name or password is incorrect" {
		t.Fatal("User is login without account")
	}
}
