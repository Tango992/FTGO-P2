package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"ungraded-13/config"
	"ungraded-13/controller"
	"ungraded-13/repository"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	getProducts = `{"message":"Get all products","data":[{"product_id":1,"name":"Teh Pucuk","stock":1000,"price":100},{"product_id":2,"name":"Teh Botol","stock":500,"price":125},{"product_id":3,"name":"Es Tee","stock":740,"price":115}]}
`
	getStores = `{"message":"Get all stores","data":[{"store_id":1,"name":"Grand Indonesia","address":"Jl. Tlk. Betung I No.45A, Kb. Melati, Kecamatan Tanah Abang, Kota Jakarta Pusat,"},{"store_id":2,"name":"Plaza Senayan","address":"Jl. Asia Afrika No.8, RT.1/RW.3, Gelora, Kecamatan Tanah Abang, Kota Jakarta Pusat"},{"store_id":3,"name":"Summarecon Mall Bekasi","address":"Sentra Summarecon Bekasi, Jl. Boulevard Ahmad Yani, Marga Mulya, Kec. Bekasi Utara"},{"store_id":4,"name":"Plaza Ambarrukmo","address":"Jl. Laksda Adisucipto No.80, Ambarukmo, Caturtunggal, Daerah Istimewa Yogyakarta"},{"store_id":5,"name":"Pakuwon Mall Surabaya","address":"Jl. Mayjend. Jonosewojo No.2, Babatan, Kec. Wiyung, Surabaya"}]}
`
)

var (
	e *echo.Echo
	db *gorm.DB
	dbHandler repository.DbHandler
	productController controller.ProductController
	storeController controller.StoreController
)

func TestMain(m *testing.M) {
	e = echo.New()

	db = config.InitDB()
	dbHandler = repository.NewDbHandler(db)
	storeController = controller.NewStoreController(dbHandler)
	productController = controller.NewProductHandler(dbHandler)

	m.Run()
}

func TestGetStores(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, productController.GetProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, getProducts, rec.Body.String())
	}
}

func TestGetProducts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/stores", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, storeController.GetStores(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, getStores, rec.Body.String())
	}
}
