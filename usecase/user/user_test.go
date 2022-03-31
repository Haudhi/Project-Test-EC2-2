package user

import (
	_entities "be7/layered/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, "haudhi", data[0].Name)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetUserById(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.GetUserById(1)
		assert.Nil(t, err)
		assert.Equal(t, "odi", data.Name)
	})

	t.Run("TestGetUserByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.GetUserById(1)
		assert.NotNil(t, err)
		assert.Equal(t, _entities.User{}, data)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("TestCreateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreateUser(_entities.User{})
		assert.Nil(t, err)
		assert.Equal(t, "haudhi", data.Name)
	})

	t.Run("TestCreateUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.CreateUser(_entities.User{})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.User{}, data)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.UpdateUser(1, _entities.User{})
		assert.Nil(t, err)
		assert.Equal(t, "kirana", data.Name)
	})

	t.Run("TestUpdateUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.UpdateUser(1, _entities.User{})
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("TestDeleteUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		err := userUseCase.DeleteUser(1)
		assert.Nil(t, err)
		
	})

	t.Run("TestDeleteUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		err := userUseCase.DeleteUser(1)
		assert.NotNil(t, err)
		
	})
}

// === mock success ===
type mockUserRepository struct{}

func (m mockUserRepository) GetAll() ([]_entities.User, error) {
	return []_entities.User{
		{Name: "haudhi", Email: "haudhi@mail.com", Password: "lalala"},
	}, nil
}

func (m mockUserRepository) GetUserById(id int) (_entities.User, error) {
	return _entities.User{
		Name: "odi", Email: "odi@mail.com", Password: "lalala",
	}, nil
}

func (m mockUserRepository) CreateUser(request _entities.User) (_entities.User, error) {
	return _entities.User{
		Name: "haudhi", Email: "haudhi@mail.com", Password: "lalala",
	}, nil
}

func (m mockUserRepository) UpdateUser(id int, request _entities.User) (_entities.User, error) {
	return _entities.User{
		Name: "kirana", Email: "kirana@mail.com", Password: "lalala",
	}, nil
}

func (m mockUserRepository) DeleteUser(id int) error {
	return nil
}


// === mock error ===

type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) GetAll() ([]_entities.User, error) {
	return nil, fmt.Errorf("error")
}

func (m mockUserRepositoryError) GetUserById(id int) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error get data user")
}

func (m mockUserRepositoryError) CreateUser(request _entities.User) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error create data user")
}

func (m mockUserRepositoryError) UpdateUser(id int, request _entities.User) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error update data user")
}

func (m mockUserRepositoryError) DeleteUser(id int) error {
	return fmt.Errorf("error update data user")
}