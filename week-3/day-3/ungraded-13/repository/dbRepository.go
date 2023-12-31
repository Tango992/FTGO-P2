package repository

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"ungraded-13/dto"
	"ungraded-13/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DbHandler struct {
	*gorm.DB
}

func NewDbHandler(db *gorm.DB) DbHandler {
	return DbHandler{
		DB: db,
	}
}

func (db DbHandler) AddUserIntoDb(data dto.RegisterUser) (entity.User, error) {
	user := entity.User{
		Username: data.Username,
		Password: data.Password,
		DepositAmount: data.DepositAmount,
	}

	res := db.Create(&user)
	if res.Error != nil {
		return entity.User{}, echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	user.Password = ""
	return user, nil
}

func (db DbHandler) FindUserInDb(username string) (entity.User, error){
	var user entity.User

	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}
		return entity.User{}, echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	return user, nil
}

func (db DbHandler) GetAllProducts() ([]entity.Product, error) {
	products := []entity.Product{}

	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	stats := sqlDB.Stats()
	fmt.Println("Postgres maximum open connection:", stats.MaxOpenConnections)
	
	
	res := db.Find(&products)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []entity.Product{}, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}
		return []entity.Product{}, echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}
	return products, nil
}

func (db DbHandler) EstablishTransactions(requestData *entity.Transaction) error {
	var userData entity.User
	if err := db.Where("id = ?", requestData.UserID).First(&userData).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Cannot find requested user")
	} 

	var productData entity.Product
	if err := db.Where("id = ?", requestData.ProductID).First(&productData).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Cannot find requested product")
	} 

	subTotal := productData.Price * float32(requestData.Quantity)
	requestData.TotalAmount = subTotal
	
	transactionErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&productData).Update("stock", (productData.Stock - requestData.Quantity)).Error; err != nil {
			return err
		}

		if err := tx.Model(&userData).Update("deposit_amount", (userData.DepositAmount - subTotal)).Error; err != nil {
			return err
		}

		if err := tx.Create(&requestData).Error; err != nil {
			return err
		}
		
		return nil
	})
	if transactionErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, transactionErr.Error())
	}
	
	return nil
}

func (db DbHandler) FindAllStoresInDb() ([]entity.Store, error) {
	stores := []entity.Store{}

	if err := db.Select("id", "name", "address").Find(&stores).Error; err != nil {
		return []entity.Store{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return stores, nil
}

func (db DbHandler) FindStoreInDb(storeId int) (dto.StoreWithSales, error) {
	store := dto.StoreWithSales{}

	res :=  db.Table("transactions t").Select("COALESCE(SUM(t.total_amount), 0) AS total_sales, s.name, s.rating, s.address, s.lat, s.long, s.id").Joins("RIGHT JOIN stores s ON s.id = t.store_id").Where("s.id = ?", storeId).Group("t.store_id").Group("s.id").Scan(&store)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dto.StoreWithSales{}, echo.NewHTTPError(http.StatusNotFound, "Store Id not found")
		}
		return dto.StoreWithSales{}, echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return dto.StoreWithSales{}, echo.NewHTTPError(http.StatusNotFound, "Store Id not found")
	}
	return store, nil
}