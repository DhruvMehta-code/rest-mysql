package connect

import (
	"time"

	"gorm.io/gorm"
)

type Mesage struct {
	gorm.Model
		Id uint `json:"id" gorm:"primary_key"`
		Title string `json:"title"`
		Message string `json:"message"`
		Postdate time.Time `json:"postdate"`
}