package handler

import (
	"fmt"
	"net/http"
	"preview-week3/dto"
	"preview-week3/entity"
	"preview-week3/helpers"
	"preview-week3/utils"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	DbHandler
}

func NewPhotoController(dbHandler DbHandler) PhotoController {
	return PhotoController{
		DbHandler: dbHandler,
	}
}

func (pc PhotoController) PostPhoto(c *gin.Context) {
	var photoTemp dto.PhotoData
	
	if err := c.ShouldBindJSON(&photoTemp); err != nil {
		helpers.ErrJsonWriter(c, utils.ErrBadRequest, err.Error())
		return
	}
	
	claimsTemp, exists := c.Get("user")
	if !exists {
		helpers.ErrJsonWriter(c, utils.ErrUnauthorized, "Token not found")
		return
	}
	claims := helpers.AssertClaims(claimsTemp)

	photoData := entity.Photo{
		Title: photoTemp.Title,
		Caption: photoTemp.Caption,
		PhotoUrl: photoTemp.PhotoUrl,
		UserID: claims.Id,
	}

	if dbErr := pc.DbHandler.AddPhotoIntoDb(&photoData); dbErr != nil {
		helpers.ErrJsonWriter(c, *dbErr, nil)
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Message: "Posted",
		Data: photoData,
	})
}

func (pc PhotoController) GetPhotos(c *gin.Context) {
	claimsTemp, exists := c.Get("user")
	if !exists {
		helpers.ErrJsonWriter(c, utils.ErrUnauthorized, "Token not found")
		return
	}
	claims := helpers.AssertClaims(claimsTemp)

	userPhotos, dbErr := pc.DbHandler.GetPhotosInDb(claims.Id)
	if dbErr != nil {
		helpers.ErrJsonWriter(c, *dbErr, nil)
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Message: fmt.Sprintf("Get all photos posted by %v", claims.Username),
		Data: userPhotos,
	})
}