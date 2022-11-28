package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/config"
	"main.go/features/user/domain"
	"main.go/mocks"
)

func TestAddUserr(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Nama: "same", HP: "0822"}, nil).Once()
		srv := New(repo)
		input := domain.Core{Nama: "same", HP: "0822", Password: "same"}
		res, err := srv.AddUser(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID, "Seharusnya ada ID yang berhasil dibuat")
		assert.NotEqual(t, res.Password, input.Password, "Password tidak terenkripsi")
		assert.Equal(t, input.Nama, res.Nama, "Nama user harus sesuai")
		repo.AssertExpectations(t)
	})
	t.Run("Duplicated", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
		srv := New(repo)
		input := domain.Core{Nama: "same", HP: "0822", Password: "same"}
		res, err := srv.AddUser(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DUPLICATED_DATA, "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
	t.Run("Gaga; Hash", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("cannot encript password")).Once()
		srv := New(repo)
		input := domain.Core{Nama: "same", HP: "0822", Password: ""}
		res, err := srv.AddUser(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "cannot encript password", "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
}

func TestShowAllUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get All User", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.Core{{ID: uint(1), Nama: "same", HP: "0822"}}, nil).Once()
		srv := New(repo)
		res, err := srv.ShowAllUser()
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res, 1, "Data kembalian harus lebih dari 0")
		assert.Equal(t, uint(1), res[0].ID, "Data seharusnya berisi 1")
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get All User", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.Core{}, nil).Once()
		srv := New(repo)
		res, err := srv.ShowAllUser()
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "no data", "Pesan error kurang sesuai")
		repo.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{ID: uint(1), Nama: "same", HP: "0822"}, nil).Once()
		srv := New(repo)
		input := uint(1)
		res, err := srv.Profile(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		// assert.Len(t, res, 1, "Data kembalian harus lebih dari 0")
		assert.Equal(t, uint(1), res.ID, "Data seharusnya berisi 1")
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		input := uint(7)
		res, err := srv.Profile(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "no data", "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
}

func TestDeleteByID(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()
		srv := New(repo)
		input := uint(1)
		err := srv.Delete(input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New("no data")).Once()
		srv := New(repo)
		input := uint(7)
		err := srv.Delete(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "no data", "Pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Update", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{ID: uint(1), Nama: "same", HP: "0822"}, nil).Once()
		srv := New(repo)
		input := domain.Core{ID: 1, Nama: "same", HP: "0822", Password: "same"}
		res, err := srv.UpdateProfile(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID, "Seharusnya ada ID user yang diupdate")
		assert.NotEqual(t, res.Password, input.Password, "Password tidak terenkripsi")
		assert.Equal(t, input.Nama, res.Nama, "Nama user harus sesuai")
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Update", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{}, errors.New("rejected from database")).Once()
		srv := New(repo)
		input := domain.Core{Nama: "same", HP: "0822", Password: "same"}
		res, err := srv.UpdateProfile(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "rejected from database", "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
}
