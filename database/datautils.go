package database

import (
	"bluesdp/config"
	"fmt"
	"strconv"
	"time"
)

func CreateOrUpdate(service_name string, ip_address string) {
	dbcon := ReturnSession()
	var record ServiceRecord
	var create_record ServiceRecord
	var result_length int64
	if res := dbcon.Model(&ServiceRecord{}).Where("ip_address = ?", ip_address).Count(&result_length); res.Error != nil {
		fmt.Printf("Error when Checking : %v", res.Error.Error())

	}
	if result_length == 0 {
		create_record.IpAddress = ip_address
		create_record.ServiceName = service_name
		create_record.LastUpdate = time.Now().UTC()
		tx := dbcon.Begin()
		if err := tx.Create(&create_record).Error; err != nil {
			tx.Rollback()
			fmt.Printf("Error when Create : %v", err)
		}
		tx.Commit()

	} else {
		tx := dbcon.Begin()
		if err := dbcon.Model(&record).Where("ip_address = ?", ip_address).Update("last_update", time.Now().Unix()).Error; err != nil {
			tx.Rollback()
			fmt.Printf("Error when Update : %v", err)
		}
		tx.Commit()
	}

}

func DeleteOutdated() {
	var records []ServiceRecord
	dbcon := ReturnSession()
	clear_interval, _ := strconv.Atoi(config.Config("CLEAR_INTERVAL"))
	clear_interval = int(clear_interval)

	tx := dbcon.Begin()
	// setting expiration time criterion
	thres_hold := time.Now().UTC().Add((-time.Millisecond * time.Duration(clear_interval)))
	//  getting list of records with expired heart beat
	dbcon.Model(&ServiceRecord{}).Where("last_update <= ?", thres_hold).Find(&records)

	//  clearing the heart beat of the records if exists
	if len(records) > 0 {
		dbcon.Delete(&records)
	}
	tx.Commit()
}
