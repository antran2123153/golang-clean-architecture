# Go Clean Architecture

- Clean architecture: [go-clean-architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- HTTP web framework: [gin](https://github.com/gin-gonic/gin)
- ORM: [gorm](https://github.com/go-gorm/gorm)
- Configuration: [viper](https://github.com/spf13/viper)
- API documentation: [swagger](https://github.com/swaggo/swag)
- Mocking framework: [gomock](https://github.com/golang/mock)

<img src="https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg" width="500"/>


# Command
```
go fmt ./...

go run cmd/api/main.go

go run cmd/migrate/main.go

swag init -g cmd/api/main.go

mockgen -source="internal/user/repository.go" -destination="internal/user/mock/repository.go" -package=mock

mockgen -source="internal/user/usecase.go" -destination="internal/user/mock/usecase.go" -package=mock

```