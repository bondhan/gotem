package manager

import (
	"sync"

	"github.com/bondhan/gotem/config"
	"github.com/bondhan/gotem/internal/driver"
	"github.com/bondhan/gotem/persistence/repository"
	dbservice "github.com/bondhan/gotem/persistence/service"
	"go.uber.org/dig"
)

// Manager ...
type Manager struct {
	container *dig.Container
}

var (
	singleton *Manager
	once      sync.Once
)

// BuildContainer ...
func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(driver.NewLogDriver)
	container.Provide(config.NewDbConfig)

	container.Provide(repository.NewProductsRepository)
	container.Provide(repository.NewCountryRepository)
	container.Provide(repository.NewProvinceRepository)
	container.Provide(repository.NewCityRepository)
	container.Provide(repository.NewZipCodeRepository)
	container.Provide(repository.NewUserRepository)
	container.Provide(repository.NewBuyerRepository)
	container.Provide(repository.NewSellerRepository)
	container.Provide(repository.NewVariantsRepository)
	container.Provide(repository.NewProductsRepository)

	container.Provide(dbservice.NewInventoryService)

	return container
}

// GetContainer ...
func GetContainer() *dig.Container {
	once.Do(func() {
		singleton = &Manager{buildContainer()}
	})

	return singleton.container
}
