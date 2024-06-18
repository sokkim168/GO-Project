package models

import (
	"time"

	config "example.com/booking/config"
	"gorm.io/gorm"
)

type RoomBooking struct {
	ID          uint           `gorm:"unique;primaryKey;autoIncrement"`
	RoomID      uint           `gorm:"not null"`
	BookingDate string         `gorm:"type:TIMESTAMP(6);not null"`
	StartDate   string         `gorm:"type:TIMESTAMP(6);not null"`
	EndDate     string         `gorm:"type:TIMESTAMP(6);not null"`
	Status      string         `gorm:"size:100;not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;type:TIMESTAMP(6)"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime;type:TIMESTAMP(6)"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Room        Room           `gorm:"foreignkey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&RoomBooking{})
}

func (r *RoomBooking) Create() *RoomBooking {
	db.Create(&r)
	return r
}

func (r *RoomBooking) All() []RoomBooking {
	var RB []RoomBooking
	db.Find(&RB)
	return RB
}

func (r *RoomBooking) Find(id int32) *RoomBooking {
	var Get RoomBooking
	db.Where("ID=?", id).Find(&Get)
	return &Get
}

func (r *RoomBooking) Update(id int32) *RoomBooking {
	var Get RoomBooking
	db.Where("ID=?", id).Updates(r).Find(&Get)
	return &Get
}

func (r *RoomBooking) Delete(id int32) *RoomBooking {
	var booking RoomBooking
	db.Where("ID=?", id).Delete(&booking)
	return &booking
}
