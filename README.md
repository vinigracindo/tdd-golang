go install github.com/golang/mock/mockgen@v1.6.0
go get github.com/golang/mock/gomock

mockgen --destination=ports/domain/mocks/client.go --source=ports/domain/client.go domain
mockgen --destination=ports/service/mocks/client.go --source=ports/service/client.go service