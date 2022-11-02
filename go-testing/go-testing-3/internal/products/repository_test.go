package products

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStore struct {
	callReadMethod bool
	datamock       []Product
	errRead        string
	errWrite       string
}

func (mockdb *mockStore) Read(data interface{}) error {
	if mockdb.errRead != "" {
		return fmt.Errorf(mockdb.errRead)
	}
	datainf := data.(*[]Product)
	*datainf = mockdb.datamock
	mockdb.callReadMethod = true
	return nil
}

func (mockdb *mockStore) Write(data interface{}) error {
	if mockdb.errWrite != "" {
		return fmt.Errorf(mockdb.errWrite)
	}
	datainf := data.([]Product)
	mockdb.datamock = datainf
	return nil
}

func TestServiceIntegrationUpdate(t *testing.T) {
	// Arrange
	database := []Product{
		{
			Id:        1,
			Name:      "Before Update",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
	}
	mockdb := mockStore{
		datamock:       database,
		callReadMethod: false,
		errRead:        "",
		errWrite:       "",
	}
	// Act
	repository := NewRepository(&mockdb)
	service := NewService(repository)
	resultado, err := service.Update(1, "After Update", database[0].Color,
		database[0].Code, database[0].Price, database[0].Stock, database[0].Published)
	database[0].Name = "After Update"
	// Assert
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(resultado, database[0]))
	assert.True(t, mockdb.callReadMethod)
}

func TestServiceIntegrationDelete(t *testing.T) {
	// Arrange
	database := []Product{
		{
			Id:        1,
			Name:      "Before Update",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
		{
			Id:        2,
			Name:      "Iphone",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
		{
			Id:        3,
			Name:      "Iphone",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
		{
			Id:        10,
			Name:      "Iphone",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
	}
	mockdb := mockStore{
		datamock:       database,
		callReadMethod: false,
		errRead:        "",
		errWrite:       "",
	}
	// Act
	repo := NewRepository(&mockdb)
	service := NewService(repo)
	err := service.Delete(1)
	err2 := service.Delete(10)
	fmt.Printf("%+v\n", mockdb.datamock)
	result, err3 := service.GetById(4)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 0, result.Id)
	assert.Nil(t, err2)
	assert.NotNil(t, err3)
}
func TestServiceIntegrationGetAll(t *testing.T) {
	// Arrange
	database := []Product{
		{
			Id:        1,
			Name:      "Before Update",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
		{
			Id:        2,
			Name:      "Iphone 13",
			Color:     "White",
			Price:     15000,
			Stock:     10,
			Code:      "abc321",
			Published: true,
		},
	}
	mockStore := mockStore{
		datamock: database,
		errRead:  "",
		errWrite: "",
	}
	// Act
	repo := NewRepository(&mockStore)
	service := NewService(repo)
	result, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, mockStore.datamock, result)
}
