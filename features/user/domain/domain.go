package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Nama     string
	HP       string
	Password string
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	Update(updateData Core) (Core, error)
	Delete(ID uint) error
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
	Login(newUser Core) (Core, error)
}

type Service interface {
	AddUser(newUser Core) (Core, error)
	UpdateProfile(updateData Core) (Core, error)
	Profile(ID uint) (Core, error)
	ShowAllUser() ([]Core, error)
	// LoginUser(newUser Core) (Core, error)
	// GenerateToken(id uint) string
	Delete(id uint) error
}

type Handler interface {
	AddUser() echo.HandlerFunc
	updateUser() echo.HandlerFunc
	ShowAllUser() echo.HandlerFunc
	ShowUserByID() echo.HandlerFunc
	DeleteByID() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
}
