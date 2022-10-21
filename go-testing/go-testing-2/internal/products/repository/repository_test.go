package products

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

func (sdb *stubStore) Read(data interface{}) error {
	dataInf := data.(*[]Product)
	*dataInf = []Product{
		{
			Id:        1,
			Name:      "Cellphone",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
		{
			Id:        2,
			Name:      "Mac Pro",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
	}
	return nil
}

func (sdb *stubStore) Write(data interface{}) error {
	return nil
}

type mockStore struct {
	CallUpdateName bool
	Datamock       []Product
}

func (mockdb *mockStore) Read(data interface{}) error {
	mockdb.CallUpdateName = true
	datainf := data.(*[]Product)
	*datainf = mockdb.Datamock
	return nil
}

func (mockdb *mockStore) Write(data interface{}) error {
	datainf := data.([]Product)
	mockdb.Datamock = append(mockdb.Datamock, datainf[len(datainf)-1])
	return nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	sdb := stubStore{}
	repository := NewRepository(&sdb)
	esperado := []Product{
		{
			Id:        1,
			Name:      "Cellphone",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
		{
			Id:        2,
			Name:      "Mac Pro",
			Color:     "Blue",
			Price:     15000,
			Stock:     10,
			Code:      "abc123",
			Published: false,
		},
	}
	// Act
	resultado, err := repository.GetAll()
	// Assert
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(esperado, resultado))
}

func TestUpdateName(t *testing.T) {
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
		Datamock: database,
	}
	repository := NewRepository(&mockdb)

	// Act
	resultado, err := repository.UpdateName(1, "After Update")
	database[0].Name = "After Update"
	// Assert
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(resultado, database[0]))

}
