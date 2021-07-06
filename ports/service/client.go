package service

import "github.com/vinigracindo/tdd-golang/ports/domain"

type ClientServiceInterface interface {
	Get(name string) (domain.Client, error)
}

type ClientService struct {
	Persistence domain.ClientRepository
}

func (c *ClientService) Get(name string) (domain.ClientInterface, error) {
	client, err := c.Persistence.Get(name)
	if err != nil {
		return nil, err
	}
	return client, nil
}
