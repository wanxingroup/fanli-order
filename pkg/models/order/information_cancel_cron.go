package order

import (
	"time"
)

type InformationCancelCron struct {
	CronId    string    `gorm:"column:cronId;type:varchar(64);primary_key;comment:'cronId'"`
	CreatedAt time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;comment:'createdAt'"`
}

func (InformationCancelCron) TableName() string {
	return "order_cancel_cron"
}
