package domain

import "github.com/asaskevich/govalidator"

type ClientInterface interface {
	isValid() (bool, error)
}

type ClientRepository interface {
	GetAll() ([]ClientInterface, error)
}

type Client struct {
	Name  string `valid:"required"`
	Email string `valid:"required,email"`
	Phone string `valid:"required,numeric"`
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return false, err
	}
	return true, nil
}
