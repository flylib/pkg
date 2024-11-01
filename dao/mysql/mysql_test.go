package mysql

import "testing"

func TestMysql(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatal(err)
	}

	// CfgCurrency mapped from table <cfg_currency>
	type CfgCurrency struct {
		ID        int64  `gorm:"column:id;primaryKey" json:"id"`  // 币种类型
		Currency  string `gorm:"column:currency" json:"currency"` // 币别
		Icon      string `gorm:"column:icon" json:"icon"`
		Precision int8   `gorm:"column:decimals;not null" json:"precision"` // 保留小数位数
	}

	var result []CfgCurrency
	err = db.Raw(`
	select * from cfg_currency
`).Find(&result).Error
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(result))
}
