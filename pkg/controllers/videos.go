package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"rtmp-recorder/pkg/models"
)

type VideoController struct{}

func (c *VideoController) GetMany(ctx *gin.Context) {
	s := models.GetVideos()
	ctx.JSON(http.StatusOK, gin.H{
		"total": len(s),
		"rows":  s,
	})
}

func (c *VideoController) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, models.GetVideo(id))
}

func (c *VideoController) Create(ctx *gin.Context) {
	var v models.Video
	if err := ctx.BindJSON(&v); err == nil {
		if e := v.Save(); e == nil {
			ctx.JSON(http.StatusCreated, v)
		} else {
			HandleError(ctx, e.Error())
		}
	} else {
		glog.Error(err.Error())
	}
}

func (c *VideoController) Update(ctx *gin.Context) {
	// id := ctx.Param("id")
	// var v models.Role
	// ctx.BindJSON(&v)
	// v.Id = id
	// if e := v.Update(); e == nil {
	// 	ctx.JSON(http.StatusCreated, v)
	// } else {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
	// }
}

func (c *VideoController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var v models.Video
	v.Id = id
	if e := v.Delete(); e == nil {
		ctx.JSON(http.StatusNoContent, nil)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
	}
}

func NewVideoController(r gin.IRouter) *VideoController {
	c := &VideoController{}
	r.Group("/videos").
		GET("", c.GetMany).
		GET("/:id", c.GetOne).
		POST("", c.Create).
		PUT("/:id", c.Update).
		DELETE("/:id", c.Delete).
		OPTIONS("/:id", nil)
	return c
}
