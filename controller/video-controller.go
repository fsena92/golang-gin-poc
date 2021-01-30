package controller

import (
	"net/http"
	"strconv"

	"github.com/fsena92/golang-gin-poc/entity"
	"github.com/fsena92/golang-gin-poc/service"
	"github.com/fsena92/golang-gin-poc/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) (entity.Video, error)
	Update(ctx *gin.Context) (entity.Video, error)
	Delete(ctx *gin.Context) (entity.Video, error)
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return video, err
	}
	err = validate.Struct(video)
	if err != nil {
		return video, err
	}
	c.service.Save(video)
	return video, nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *controller) Update(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return video, err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return video, err
	}
	video.ID = id

	err = validate.Struct(video)
	if err != nil {
		return video, err
	}
	c.service.Update(video)
	return video, nil
}

func (c *controller) Delete(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return video, err
	}
	video.ID = id
	c.service.Delete(video)
	return video, nil
}
