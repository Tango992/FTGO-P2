package handler

import (
	"ugc-9/dto"
	"ugc-9/entity"
	"ugc-9/utils"

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

func (db DbHandler) AddStoreToDb(u entity.Store) *utils.ErrResponse {
	result := db.Create(&u)

	if result.Error != nil {
		err := utils.ErrInternalServer
		err.Description = result.Error.Error()
		return &err
	}
	
	return nil
}

func (db DbHandler) FindStoreInDb(credential *dto.Credential) (entity.Store, *utils.ErrResponse) {
	var store entity.Store

	res := db.Where("email = ?", credential.Email).First(&store)

	if res.RowsAffected == 0 {
		resErr := utils.ErrUnauthorized
		resErr.Description = "Invalid credentials"
		return entity.Store{}, &resErr
	}

	if res.Error != nil {
		resErr := utils.ErrInternalServer
		resErr.Description = res.Error.Error()
		return entity.Store{}, &resErr
	}

	return store, nil
}


func (db DbHandler) AddProductToDb(product *entity.Product) *utils.ErrResponse {
	result := db.Create(product)

	if result.Error != nil {
		resErr := utils.ErrInternalServer
		resErr.Description = result.Error.Error()
		return &resErr
	}

	return nil
}

func (db DbHandler) FindAllProductFromDb() ([]entity.Product, *utils.ErrResponse) {
	var products []entity.Product

	result := db.Find(&products)
	if result.Error != nil {
		resErr := utils.ErrInternalServer
		resErr.Description = result.Error.Error()
		return products, &resErr
	}
	return products, nil
}

func (db DbHandler) FindProductInDb(product *entity.Product) (*utils.ErrResponse) {
	result := db.First(product)
	if result.Error != nil {
		resErr := utils.ErrDataNotFound
		resErr.Description = result.Error.Error()
		return &resErr
	}
	return nil
}

func (db DbHandler) DeleteProductInDb(product *entity.Product) (*utils.ErrResponse) {

	result := db.Where("store_id = ?", product.StoreID).Delete(product)
	if result.Error != nil {
		resErr := utils.ErrInternalServer
		resErr.Description = result.Error.Error()
		return &resErr
	}

	if result.RowsAffected == 0 {
		resErr := utils.ErrDataNotFound
		resErr.Description = result.Error.Error()
		return &resErr
	}
	return nil
}

func (db DbHandler) UpdateProductInDb(product *entity.Product) (*utils.ErrResponse) {
	result := db.Model(product).Where("store_id = ?", product.StoreID).Updates(product)
	if result.Error != nil {
		resErr := utils.ErrInternalServer
		resErr.Description = result.Error.Error()
		return &resErr
	}

	if result.RowsAffected == 0 {
		resErr := utils.ErrDataNotFound
		resErr.Description = result.Error.Error()
		return &resErr
	}
	return nil
}