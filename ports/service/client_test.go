package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock_domain "github.com/vinigracindo/tdd-golang/ports/domain/mocks"
	"github.com/vinigracindo/tdd-golang/ports/service"
)

func TestClientService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_domain.NewMockClientInterface(ctrl)
	persistence := mock_domain.NewMockClientRepository(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(client, nil).AnyTimes()

	service := service.ClientService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, client, result)
}
