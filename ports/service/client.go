package service

import "github.com/vinigracindo/tdd-golang/ports/domain"

type ClientServiceInterface interface {
	GetAll() ([]domain.Client, error)
}

type ClientService struct {
	Persistence domain.ClientRepository
}

func (c *ClientService) GetAll() ([]domain.ClientInterface, error) {
	client, err := c.Persistence.GetAll()
	if err != nil {
		return nil, err
	}
	return client, nil
}
