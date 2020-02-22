package mockup

import (
	"github.com/bondhan/gotem/manager"
	"github.com/bondhan/gotem/persistence/repository"
)

// insertBuyer ...
func insertBuyer() {
	manager.GetContainer().Invoke(func(b repository.BuyerRepository) {
		b.GetABuyer(2)
	})
}

// insertSeller ...
func insertSeller() {

}

// insertProducts ...
func insertProducts() {

}

// DoPopulateData ...
func DoPopulateData() {
	insertBuyer()
}
