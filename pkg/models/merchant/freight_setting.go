package merchant

type FreightSetting struct {
	ShopId         uint64 `gorm:"column:shopId;primary_key;auto_increment:false"`         // 店铺ID
	IsFreePostage  bool   `gorm:"column:isFreePostage;default:'0'"`                       // 是否包邮
	ConditionPrice uint64 `gorm:"column:conditionPrice;type:bigint unsigned;default:'0'"` // 包邮金额条件（单位：分）
	Freight        uint64 `gorm:"column:freight;type:bigint unsigned;default:'0'"`        // 邮费金额（单位：分）
}

func (FreightSetting) TableName() string {
	return "merchant_freight_setting"
}
