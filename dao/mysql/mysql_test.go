package mysql

import "testing"

func TestMysql(t *testing.T) {
	db, err := Connect("root:123456@tcp(192.168.119.128:3306)/poker?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}

	// CfgCurrency mapped from table <cfg_currency>
	type CfgCurrency struct {
		ID        int64  `gorm:"column:id;primaryKey" json:"id"`  // 币种类型
		Currency  string `gorm:"column:currency" json:"currency"` // 币别
		Icon      string `gorm:"column:icon" json:"icon"`
		Precision int8   `gorm:"column:precision;not null" json:"precision"` // 保留小数位数
	}

	var result []CfgCurrency
	err = db.ExecRawSql(`
	select * from cfg_currency
`, nil).Find(&result).Error
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(result))
}
