package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type ServiceRecord struct {
	ID          uint      `gorm:"primaryKey;autoIncrement:true"`
	ServiceName string    `gorm:"type:string; index;" json:"service_name,omitempty"`
	LastUpdate  time.Time `gorm:"constraint:not null; default:current_timestamp;" json:"last_update,omitempty"`
	IpAddress   string    `gorm:"type:string; unique;" json:"ip_address,omitempty"`
}

func (service *ServiceRecord) BeforeUpdate(tx *gorm.DB) error {

	service.LastUpdate = time.Now().UTC()
	return nil
}

func (service *ServiceRecord) BeforeCreate(tx *gorm.DB) error {
	service.LastUpdate = time.Now().UTC()
	return nil
}

func MigrateDataBase() {

	dbcon := ReturnSession()
	fmt.Println("Connection Opened to Database")
	if err := dbcon.AutoMigrate(
		&ServiceRecord{},
	); err != nil {
		log.Fatalln(err)

	}

	fmt.Println("Database Migrated")
}
