package model

import (
	"demo/conf"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

// UpdateWorksTypes  更新作品集状态信息
// @param works   string   作品集
func UpdateWorksTypes(works, types, form string, price float64, belong int) error {
	logs.Info("更新作品状态信息")
	time := time.Now()
	// 更新作品状态信息
	if err := conf.DB1.Model(WorkTypes{}).
		Where("works = ?", works).
		Where("belong = ?", belong).
		Updates(map[string]interface{}{"types": types, "price": price, "time": time, "form": form}).
		Error; err != nil {
		logs.Error("更新作品状态信息失败，err:%s", err)
		return err
	}
	return nil
}

// CreateMoldRecord 创建铸造记录
func CreateMoldRecord(works, hash string, number int) error {
	var work Work
	if err := conf.DB1.Model(Work{}).Where("works = ?", works).Group("works").Find(&work).Error; err != nil {
		logs.Error("创建铸造记录错误，err:%s", err)
		return err
	}
	moldRecord := MoldRecord{
		Name:   work.Name,
		Works:  works,
		Number: number,
		Hash:   hash,
	}
	restlu := conf.DB1.Create(&moldRecord)
	if restlu.Error != nil {
		logs.Error("创建铸造记录失败，err:%v", restlu.Error)
		return restlu.Error
	}
	return nil
}
