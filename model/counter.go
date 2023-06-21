package model

import (
	"gorm.io/gorm"
	"time"
)

// FetchCounterInfoById ...
func FetchCounterInfoById(db *gorm.DB, id int64) (*TCounters, error) {
	result := &TCounters{}
	if err := db.Find(&result, "id = ? ", id).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// InsertCounter ...
func InsertCounter(db *gorm.DB, msg *TCounters, dup string) error {
	if len(msg.Createdat) <= 0 {
		msg.Createdat = time.Now().Format("2006-01-02 15:04:05")
	}
	if len(msg.Updatedat) <= 0 {
		msg.Updatedat = time.Now().Format("2006-01-02 15:04:05")
	}
	if len(dup) <= 0 {
		return db.Create(msg).Error
	}
	return db.Set("gorm:insert_option", dup).Create(msg).Error
}

// UpdateOrderCreateCounter ...
func UpdateOrderCreateCounter(db *gorm.DB, msg *TCounters) error {
	dup := "ON DUPLICATE KEY UPDATE count = VALUES(count), updatedAt = VALUES(updatedAt)"
	return InsertCounter(db, msg, dup)
}
