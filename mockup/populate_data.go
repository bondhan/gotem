package mockup

import (
	"github.com/bondhan/gotem/manager"
	"github.com/bondhan/gotem/persistence/repository"
)

func insertCountry() {
	manager.GetContainer().Invoke(func(c repository.CountryRepository) {
		c.InsertCountry("62", "indonesia")
		c.InsertCountry("966", "saudi arabia")
	})
}

func insertProvince() {
	manager.GetContainer().Invoke(func(c repository.ProvinceRepository) {
		c.InsertProvince("01", "DKI Jakarta", "62")
		c.InsertProvince("02", "Jawa Barat", "62")
		c.InsertProvince("03", "Jawa Tengah", "62")

		c.InsertProvince("01", "Eastern", "966")
		c.InsertProvince("02", "Central", "966")
		c.InsertProvince("03", "West", "966")

	})
}

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
	insertCountry()
	insertProvince()
	insertBuyer()
}
