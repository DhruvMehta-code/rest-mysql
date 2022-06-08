package mesage

import (
	"net/http"
	"rest-go/rest-go-mysql/connect"

	"time"

	"github.com/gin-gonic/gin"
)

var db = connect.Connect()

func ReadMessage(ctx *gin.Context) {
	var msg connect.Mesage
	msg.Postdate = time.Now()
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = db.Create(&msg).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, msg)

}

func GetAll(ctx *gin.Context) {
	var msg []connect.Mesage
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = db.Find(&msg).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, msg)

}

func GetOne(ctx *gin.Context) {
	var msg connect.Mesage
	id, _ := ctx.Params.Get("id")
	err := ctx.ShouldBindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = db.Where("id = ?", id).First(&msg).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, msg)

}
