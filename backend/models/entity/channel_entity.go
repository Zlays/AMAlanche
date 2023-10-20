package entity

import "github.com/jinzhu/gorm"

type Channel struct {
	gorm.Model
	Name         string `json:"name"`
	OwnerId      uint   `json:"ownerId"`
	Participants []User `json:"participants" gorm:"many2many:channel_participants"`
}
