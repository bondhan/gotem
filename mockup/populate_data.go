package mockup

import (
	"github.com/bondhan/gotem/manager"
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/bondhan/gotem/persistence/repository"
	log "github.com/sirupsen/logrus"
)

func insertCountry() {
	manager.GetContainer().Invoke(func(c repository.CountryRepository) {
		c.InsertCountry("62", "indonesia")
		c.InsertCountry("966", "saudi arabia")
	})
}

func insertProvince() {
	manager.GetContainer().Invoke(func(c repository.ProvinceRepository) {
		c.InsertProvince("621", "DKI Jakarta", "62")
		c.InsertProvince("622", "Jawa Barat", "62")
		c.InsertProvince("623", "Jawa Tengah", "62")

		c.InsertProvince("9661", "Eastern", "966")
		c.InsertProvince("9662", "Central", "966")
		c.InsertProvince("9663", "West", "966")
	})
}

func insertCity() {
	manager.GetContainer().Invoke(func(c repository.CityRepository) {
		c.InsertCity("22", "Jakarta Selatan", "621")
		c.InsertCity("23", "Jakarta Barat", "621")
		c.InsertCity("24", "Jakarta Timur", "621")
		c.InsertCity("25", "Jakarta Utara", "621")

		c.InsertCity("111", "Dammam", "9661")
		c.InsertCity("112", "Jakarta Barat", "9661")

	})
}

func insertZipCode() {
	manager.GetContainer().Invoke(func(z repository.ZipCodeRepository) {
		z.InsertZipCode("12345", "22")
		z.InsertZipCode("12346", "23")
		z.InsertZipCode("12347", "24")
	})
}

// insertBuyer ...
func insertBuyer() {
	manager.GetContainer().Invoke(func(b repository.BuyerRepository) {
		buyer := domain.Buyer{
			Name:      "Bondhan Novandy",
			Mobile:    "1234567890",
			Email:     "bondhan@github.com",
			Address:   "RT 5 RW 3 No 10 Jalan Cinta",
			ZipCodeID: 1,
		}
		b.InsertABuyer(buyer)
	})
}

// insertSeller ...
func insertSeller() {
	manager.GetContainer().Invoke(func(s repository.SellerRepository) {
		seller := domain.Seller{
			Name:      "Muhammad Hisyam",
			Mobile:    "32947384732",
			Email:     "hisyam@github.com",
			Address:   "RT 11 RW 13 No 110 Jalan Keadilan",
			ZipCodeID: 2,
		}
		s.InsertASeller(seller)
	})
}

// insertVariants ...
func insertVariants() {
	manager.GetContainer().Invoke(func(v repository.VariantsRepository) {
		v.InsertVariant("S")
		v.InsertVariant("M")
		v.InsertVariant("L")
		v.InsertVariant("XL")
		v.InsertVariant("XXL")
	})
}

// insertProducts ...
func insertProducts() {
	manager.GetContainer().Invoke(func(productsRepo repository.ProductsRepository, variantRepo repository.VariantsRepository) {
		variants, err := variantRepo.GetAVariantFromSize("XL")
		if err != nil {
			log.Error(err)
		}

		products := domain.Products{
			Name:        "kemeja",
			Category:    "shirt",
			VariantID:   variants.ID,
			Price:       15000,
			Description: "original shirt made from Indonesia",
			Quantity:    100,
		}

		err = productsRepo.AddAproduct(products)
		if err != nil {
			log.Error(err)
		}
	})
}

// DoPopulateData ...
func DoPopulateData() {
	insertCountry()
	insertProvince()
	insertCity()
	insertZipCode()
	insertBuyer()
	insertSeller()
	insertVariants()
	insertProducts()
}
