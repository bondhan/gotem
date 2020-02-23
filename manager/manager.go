package manager

import (
	"sync"

	"github.com/bondhan/gotem/config"
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
	container.Provide(config.NewLogConfig)
	container.Provide(config.NewDbConfig)

	container.Provide(repository.NewBuyerRepository)
	container.Provide(repository.NewProductsRepository)
	container.Provide(repository.NewCountryRepository)
	container.Provide(repository.NewProvinceRepository)

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