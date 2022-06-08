package connect

import (
	"time"
)

type Mesage struct {
	Id       uint      `json:"id" gorm:"primary_key"`
	Title    string    `json:"title"`
	Message  string    `json:"message"`
	Postdate time.Time `json:"postdate"`
}
