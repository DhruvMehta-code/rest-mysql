package connect

import "time"

type Mesage struct {
	Id       uint      `json:"id" gorm:"primary_key"`
	Title    string    `json:"title" form:"title"`
	Message  string    `json:"message" form:"message"`
	Postdate time.Time `json:"postdate" form:"postdate"`
}
