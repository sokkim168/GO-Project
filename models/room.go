package models

import (
	"fmt"
	"time"

	config "example.com/booking/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Room struct {
	ID         uint           `gorm:"unique;primaryKey;autoIncrement"`
	RoomNumber string         `gorm:"room_number;not null"`
	RoomSize   string         `gorm:"room_size;not null" `
	Floor      int            `gorm:"floor;not null" `
	Status     string         `gorm:"status"`
	CreatedAt  time.Time      `gorm:"autoCreateTime;type:TIMESTAMP(6)"`
	UpdatedAt  time.Time      `gorm:"autoCreateTime;type:TIMESTAMP(6)"`
	DeletedAt  gorm.DeletedAt `gorm:"type:TIMESTAMP(6)"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Room{})
}

func MigrateRoom() {
	fmt.Println("Room initilized")
}

func (r *Room) Create() *Room {
	db.Create(&r)
	return r
}

func (r *Room) All() []Room {
	var Rooms []Room
	db.Find(&Rooms)
	return Rooms
}

func Find(id int32) *Room {
	var Get Room
	db.Where("ID=?", id).Find(&Get)
	return &Get
}

func (r *Room) Update(id int32) *Room {
	var Get Room
	db.Where("ID=?", id).Updates(r).Find(&Get)
	return &Get
}

func Delete(id int32) *Room {
	var room Room
	db.Where("ID=?", id).Delete(&room)
	return &room
}
