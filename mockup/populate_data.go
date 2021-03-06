package mockup

import (
	"github.com/bondhan/gotem/internal/errorhandler"
	"github.com/bondhan/gotem/manager"
	"github.com/bondhan/gotem/persistence/domain"
	"github.com/bondhan/gotem/persistence/repository"
	"github.com/jinzhu/gorm"
)

func insertCountry() {
	manager.GetContainer().Invoke(func(c repository.CountryRepository) {
		c.InsertCountry("62", "indonesia")
		c.InsertCountry("966", "saudi arabia")
	})
}

func insertProvince() {
	manager.GetContainer().Invoke(func(p repository.ProvinceRepository, c repository.CountryRepository) {
		countryIna, err := c.GetACountryByCountryCode("62")
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				errorhandler.DBError.Newf("Country 62 not found, err:", err)
			}
			return
		}

		p.InsertProvince(domain.Province{ProvinceCode: "621", ProvinceName: "DKI Jakarta", CountryID: countryIna.ID})
		p.InsertProvince(domain.Province{ProvinceCode: "622", ProvinceName: "Jawa Barat", CountryID: countryIna.ID})
		p.InsertProvince(domain.Province{ProvinceCode: "623", ProvinceName: "Jawa Tengah", CountryID: countryIna.ID})

		countryKsa, err := c.GetACountryByCountryCode("966")
		if err != nil {
			// log.Error(err)
		}
		p.InsertProvince(domain.Province{ProvinceCode: "9661", ProvinceName: "Eastern", CountryID: countryKsa.ID})
		p.InsertProvince(domain.Province{ProvinceCode: "9662", ProvinceName: "Central", CountryID: countryKsa.ID})
		p.InsertProvince(domain.Province{ProvinceCode: "9663", ProvinceName: "West", CountryID: countryKsa.ID})

	})
}

func insertCity() {
	manager.GetContainer().Invoke(func(c repository.CityRepository, p repository.ProvinceRepository) {
		provJkt, err := p.GetAProvinceByProvinceCode("621")
		if err != nil {
			// log.Error(err)
		}
		c.InsertCity(domain.City{CityCode: "22", CityName: "Jakarta Selatan", ProvinceID: provJkt.ID})
		c.InsertCity(domain.City{CityCode: "23", CityName: "Jakarta Barat", ProvinceID: provJkt.ID})
		c.InsertCity(domain.City{CityCode: "24", CityName: "Jakarta Timur", ProvinceID: provJkt.ID})
		c.InsertCity(domain.City{CityCode: "25", CityName: "Jakarta Utara", ProvinceID: provJkt.ID})

		provSaudi, err := p.GetAProvinceByProvinceCode("9661")
		if err != nil {
			// log.Error(err)
		}

		c.InsertCity(domain.City{CityCode: "111", CityName: "Dammam", ProvinceID: provSaudi.ID})
		c.InsertCity(domain.City{CityCode: "112", CityName: "Khobar", ProvinceID: provSaudi.ID})
	})
}

func insertZipCode() {
	manager.GetContainer().Invoke(func(z repository.ZipCodeRepository, c repository.CityRepository) {
		c22, err := c.GetACityByCityCode("22")
		if err != nil {
			// log.Error(err)
		}
		z.InsertZipCode(domain.ZipCode{ZipCode: "12345", CityID: c22.ID})

		c23, err := c.GetACityByCityCode("23")
		if err != nil {
			// log.Error(err)
		}
		z.InsertZipCode(domain.ZipCode{ZipCode: "12346", CityID: c23.ID})

		c24, err := c.GetACityByCityCode("24")
		if err != nil {
			// log.Error(err)
		}
		z.InsertZipCode(domain.ZipCode{ZipCode: "12347", CityID: c24.ID})
	})
}

// insertBuyer ...
func insertBuyer() {
	manager.GetContainer().Invoke(func(b repository.BuyerRepository, u repository.UserRepository, z repository.ZipCodeRepository) {
		zc1, err := z.GetZipCode("12345")
		if err != nil {
			// log.Error(err)
		}

		user1 := domain.User{
			Name:      "Bondhan Novandy",
			Mobile:    "1234567890",
			Email:     "bondhan@github.com",
			Address:   "RT 5 RW 3 No 10 Jalan Cinta",
			ZipCodeID: zc1.ID,
		}

		err = u.InsertAUser(user1)
		if err != nil {
			// log.Error(err)
		}

		bondhan, err := u.GetAUserFromMobile("1234567890")
		if err != nil {
			// log.Error(err)
		}

		buyer := domain.Buyer{
			UserID: bondhan.ID,
		}

		err = b.InsertABuyer(buyer)
		if err != nil {
			// log.Error(err)
		}
	})
}

// insertSeller ...
func insertSeller() {
	manager.GetContainer().Invoke(func(s repository.SellerRepository, u repository.UserRepository, z repository.ZipCodeRepository) {

		zc2, err := z.GetZipCode("12346")
		if err != nil {
			// log.Error(err)
		}

		user2 := domain.User{
			Name:      "Muhammad Hisyam",
			Mobile:    "32947384732",
			Email:     "hisyam@github.com",
			Address:   "RT 11 RW 13 No 110 Jalan Keadilan",
			ZipCodeID: zc2.ID,
		}

		err = u.InsertAUser(user2)
		if err != nil {
			// log.Error(err)
		}

		hisyam, err := u.GetAUserFromMobile("32947384732")
		if err != nil {
			// log.Error(err)
		}

		seller := domain.Seller{
			UserID: hisyam.ID,
		}

		err = s.InsertASeller(seller)
		if err != nil {
			// log.Error(err)
		}
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
			// log.Error(err)
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
			// log.Error(err)
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
