package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	structs "main/src/Structs"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"syreclabs.com/go/faker"
)

func makeStudent() structs.Student {
	return structs.Student{
		Name:           fmt.Sprintf("%s - %s", faker.Name().Name(), uuid.New()),
		Age:            faker.Number().NumberInt32(2),
		CurrentPayment: float64(faker.Number().NumberInt(3)),
		BalancePayment: float64(faker.Number().NumberInt(4)),
	}
}

func comparedUsers(t *testing.T, student1 map[string]interface{}, student2 structs.Student) {
	assert.True(t, student1["Name"].(string) == student2.Name, "Expected the names are equals")
	assert.True(t, student1["Age"].(float64) == float64(student2.Age), "Expected the age are equals")
	assert.True(t, student1["CurrentPayment"].(float64) == student2.CurrentPayment, "Expected the current payment are equals")
	assert.True(t, student1["BalancePayment"].(float64) == student2.BalancePayment, "Expected the balance payment are equals")
}

type StudentHttpRequest struct {
	user     structs.User
	student  structs.Student
	url      string
	urlAuth  string
	method   string
	callBack func(ctx echo.Context) error
}

func httpRequest(studentHttpRequest StudentHttpRequest) (apiResult structs.ApiResult, err error) {
	apiResult, err = Auth(studentHttpRequest.user, studentHttpRequest.urlAuth)
	if err != nil {
		return
	}

	e := echo.New()
	inputJSON, err := json.Marshal(studentHttpRequest.student)
	if err != nil {
		return
	}

	req := httptest.NewRequest(studentHttpRequest.method, studentHttpRequest.url, bytes.NewBuffer(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("token", apiResult.Data.(string))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err != nil {
		return
	}

	studentHttpRequest.callBack(c)

	if err = json.Unmarshal(rec.Body.Bytes(), &apiResult); err != nil {
		return
	}

	return
}

var Students = []structs.Student{makeStudent(), makeStudent(), makeStudent()}

func TestCreateStudent(t *testing.T) {
	studentsController := StudentsController{}
	apiResult, err := httpRequest(StudentHttpRequest{
		user:     Users[1],
		student:  Students[0],
		urlAuth:  "/register",
		url:      "/user/create",
		method:   http.MethodPost,
		callBack: studentsController.CreateStudent,
	})

	if err != nil {
		t.Fatal(err)
	}

	student := apiResult.Data.(map[string]interface{})
	comparedUsers(t, student, Students[0])
	Students[0].ID = uint(student["ID"].(float64))
}

func TestFindStudent(t *testing.T) {
	studentsController := StudentsController{}
	apiResult, err := httpRequest(StudentHttpRequest{
		user:     Users[1],
		student:  Students[0],
		urlAuth:  "/login",
		url:      "/user/get",
		method:   http.MethodGet,
		callBack: studentsController.CreateStudent,
	})
	if err != nil {
		t.Fatal(err)
	}

	student := apiResult.Data.(map[string]interface{})
	comparedUsers(t, student, Students[0])
}

func TestUpdateStudent(t *testing.T) {
	studentsController := StudentsController{}
	apiResult, err := httpRequest(StudentHttpRequest{
		user:     Users[1],
		student:  Students[1],
		urlAuth:  "/login",
		url:      fmt.Sprintf("/user/update?studentId=%s", strconv.FormatUint(uint64(Students[0].ID), 10)),
		method:   http.MethodPut,
		callBack: studentsController.UpdateStudents,
	})
	if err != nil {
		t.Fatal(err)
	}
	if apiResult.Error != "" {
		t.Fatal(apiResult.Error)
	}
	student := apiResult.Data.(map[string]interface{})
	comparedUsers(t, student, Students[1])
}

func TestDeleteStudent(t *testing.T) {
	studentsController := StudentsController{}
	apiResult, err := httpRequest(StudentHttpRequest{
		user:     Users[1],
		urlAuth:  "/login",
		url:      fmt.Sprintf("/user/delete?studentId=%s", strconv.FormatUint(uint64(Students[0].ID), 10)),
		method:   http.MethodDelete,
		callBack: studentsController.DeleteStudent,
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	if apiResult.Message != "Student deleted successfully" {
		t.Fatal(fmt.Printf("User is not deleted successfully message result  %s\n", apiResult.Message))
	}
}
