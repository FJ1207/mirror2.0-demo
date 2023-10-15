package model

import (
	"demo/conf"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

// 查找用户信息
func FindUser(name string) (*UserDemo, error) {
	logs.Info("查找用户信息")
	var user UserDemo
	if err := conf.DB1.Model(UserDemo{}).Where("name = ?", name).Find(&user).Error; err != nil {
		logs.Error("保存铸造信息错误，err:%s", err)
		return nil, err
	}
	return &user, nil
}

// 更新用户金额
func UpdatePurUser(name string, price int64) error {
	user, err := FindUser(name)
	if err != nil {
		return err
	}
	amountleft := user.Amount - price
	if err := conf.DB1.Model(UserDemo{}).
		Where("name = ?", name).
		Updates(map[string]interface{}{"amount": amountleft}).
		Error; err != nil {
		logs.Error("更新购买用户金额信息错误，err:%s", err)
		return err
	}
	return nil
}

// 更新用户金额
func UpdateSellUser(name string, price int64) error {
	user, err := FindUser(name)
	if err != nil {
		return err
	}
	amountleft := user.Amount + price
	if err := conf.DB1.Model(UserDemo{}).
		Where("name = ?", name).
		Updates(map[string]interface{}{"amount": amountleft}).
		Error; err != nil {
		logs.Error("更新售卖用户金额信息错误，err:%s", err)
		return err
	}
	return nil
}

// CreateMInt 创建商品信息
func CreateWork(work *WorkDemo) error {
	logs.Info("创建商品信息")
	if err := conf.DB1.Model(WorkDemo{}).Create(&work).Error; err != nil {
		logs.Error("保存铸造信息错误，err:%s", err)
		return err
	}
	return nil
}

func UpdateWork(token_id, owner string) error {
	logs.Info("更新作品信息")
	if err := conf.DB1.Model(WorkDemo{}).
		Where("token_id = ?", token_id).
		Updates(map[string]interface{}{"owner": owner}).
		Error; err != nil {
		logs.Error("更新作品信息失败，err:%s", err)
		return err
	}
	return nil
}

// UpdateWorksTypes  更新作品集状态信息
// @param works   string   作品集
func UpdateWorksTypes(works, types, form string, price float64, belong int) error {
	logs.Info("更新作品状态信息")
	time := time.Now()
	// 更新作品状态信息
	if err := conf.DB1.Model(WorkTypesDemo{}).
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
