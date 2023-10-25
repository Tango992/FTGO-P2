package handler

import (
	"net/http"
	"ugc-8/dto"
	"ugc-8/entity"

	"gorm.io/gorm"
)

type DbHandler struct {
	*gorm.DB 
}

func NewDbHandler(db *gorm.DB ) *DbHandler {
	return &DbHandler{
		DB: db,
	}
}

func (db DbHandler) AddStoreToDb(u entity.Store) *dto.Response {
	result := db.Create(&u)
	
	if result.Error != nil {
		return &dto.Response{
			Code: http.StatusInternalServerError,
			Message: result.Error.Error(),
			Data: nil,
		}
	}
	return nil
}

func (db DbHandler) FindStoreInDb(credential *dto.Credential) (entity.Store, *dto.Response) {
	var store entity.Store

	res := db.Where("email = ?", credential.Email).First(&store)

	if res.RowsAffected == 0 {
		return entity.Store{}, &dto.Response{
			Code: http.StatusUnauthorized,
			Message: "Invalid credentials",
			Data: nil,
		}
	}

	if res.Error != nil {
		return entity.Store{}, &dto.Response{
			Code: http.StatusInternalServerError,
			Message: res.Error.Error(),
			Data: nil,
		}
	}

	return store, nil
}


func (db DbHandler) AddProductToDb(product *entity.Product) *dto.Response {
	result := db.Create(product)

	if result.Error != nil {
		return &dto.Response{
			Code: http.StatusInternalServerError,
			Message: result.Error.Error(),
			Data: nil,
		}
	}

	return nil
}

func (db DbHandler) FindAllProductFromDb() ([]entity.Product, *dto.Response) {
	var products []entity.Product

	result := db.Find(&products)
	if result.Error != nil {
		return products, &dto.Response{
			Code: http.StatusInternalServerError,
			Message: result.Error.Error(),
			Data: nil,
		}
	}
	return products, nil
}

func (db DbHandler) FindProductInDb(product *entity.Product) (*dto.Response) {
	result := db.First(product)
	if result.Error != nil {
		return &dto.Response{
			Code: http.StatusNotFound,
			Message: result.Error.Error(),
			Data: nil,
		}
	}
	return nil
}

func (db DbHandler) DeleteProductInDb(product *entity.Product) (*dto.Response) {

	result := db.Where("store_id = ?", product.StoreID).Delete(product)
	if result.Error != nil {
		return &dto.Response{
			Code: http.StatusInternalServerError,
			Message: result.Error.Error(),
			Data: nil,
		}
	}

	if result.RowsAffected == 0 {
		return &dto.Response{
			Code: http.StatusNotFound,
			Message: "Product not found",
			Data: nil,
		}
	}
	return nil
}

func (db DbHandler) UpdateProductInDb(product *entity.Product) (*dto.Response) {
	result := db.Model(product).Where("store_id = ?", product.StoreID).Updates(product)
	if result.Error != nil {
		return &dto.Response{
			Code: http.StatusInternalServerError,
			Message: result.Error.Error(),
			Data: nil,
		}
	}

	if result.RowsAffected == 0 {
		return &dto.Response{
			Code: http.StatusNotFound,
			Message: "Product not found",
			Data: nil,
		}
	}
	return nil
}