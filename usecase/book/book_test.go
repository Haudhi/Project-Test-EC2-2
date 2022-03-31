package book

import (
	_entities "be7/layered/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestGetAll(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, err := bookUseCase.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, "odi", data[0].Title)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, err := bookUseCase.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetBookById(t *testing.T) {
	t.Run("TestGetBookByIdSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, err := bookUseCase.GetBookById(1)
		assert.Nil(t, err)
		assert.Equal(t, "haudhi", data.Title)
	})

	t.Run("TestGetBookByIdError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, err := bookUseCase.GetBookById(1)
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Book{}, data)
	})
}

func TestCreateBook(t *testing.T) {
	t.Run("TestCreateBookSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, err := bookUseCase.CreateBook(_entities.Book{})
		assert.Nil(t, err)
		assert.Equal(t, "odi", data.Title)
	})

	t.Run("TestCreateBookError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, err := bookUseCase.CreateBook(_entities.Book{})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Book{}, data)
	})
}

func TestUpdateBook(t *testing.T) {
	t.Run("TestUpdateBookSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, err := bookUseCase.UpdateBook(1, _entities.Book{})
		assert.Nil(t, err)
		assert.Equal(t, "kirana", data.Title)
	})

	t.Run("TestUpdateBookError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, err := bookUseCase.UpdateBook(1, _entities.Book{})
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}

func TestDeleteBook(t *testing.T) {
	t.Run("TestDeleteBookSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		err := bookUseCase.DeleteBook(1)
		assert.Nil(t, err)
		
	})

	t.Run("TestDeleteBookError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		err := bookUseCase.DeleteBook(1)
		assert.NotNil(t, err)
		
	})
}

// === mock success ===
type mockBookRepository struct{}

func (m mockBookRepository) GetAll() ([]_entities.Book, error) {
	return []_entities.Book{
		{Title: "odi", Author: "odi@mail.com", Publisher: "odi"},
	}, nil
}

func (m mockBookRepository) GetBookById(id int) (_entities.Book, error) {
	return _entities.Book{
		Title: "haudhi", Author: "haudhi@mail.com", Publisher: "haudhi",
	}, nil
}

func (m mockBookRepository) CreateBook(request _entities.Book) (_entities.Book, error) {
	return _entities.Book{
		Title: "odi", Author: "odi@mail.com", Publisher: "odi",
	}, nil
}

func (m mockBookRepository) UpdateBook(id int,request _entities.Book) (_entities.Book, error) {
	return _entities.Book{
		Title: "kirana", Author: "kirana@mail.com", Publisher: "kirana",
	}, nil
}

func (m mockBookRepository) DeleteBook(id int) error {
	return nil
}



// === mock error ===

type mockBookRepositoryError struct{}

func (m mockBookRepositoryError) GetAll() ([]_entities.Book, error) {
	return nil, fmt.Errorf("error")
}

func (m mockBookRepositoryError) GetBookById(id int) (_entities.Book, error) {
	return _entities.Book{}, fmt.Errorf("error get data book")
}

func (m mockBookRepositoryError) CreateBook(request _entities.Book) (_entities.Book, error) {
	return _entities.Book{}, fmt.Errorf("error create data book")
}

func (m mockBookRepositoryError) UpdateBook(id int, request _entities.Book) (_entities.Book, error) {
	return _entities.Book{}, fmt.Errorf("error update data book")
}

func (m mockBookRepositoryError) DeleteBook(id int) error {
	return fmt.Errorf("error update data book")
}