package Repository

import (
	"encoding/json"
	"wolWebUi/Tools"
)
import "wolWebUi/Model"

// GetDevList 获取所有设备
func GetDevList() *[]Model.Dev {
	//获取DB
	db := Tools.GetDB()
	//执行数据库查询操作
	var devlist []Model.Dev
	db.Order("id").Find(&devlist)
	return &devlist
}

// AddDev 新增设备
func AddDev(data []byte) error {
	var dev Model.Dev
	err := json.Unmarshal(data, &dev)
	if err != nil {
		return err
	}
	//获取DB
	db := Tools.GetDB()
	result := db.Create(&dev)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// DelDev 删除设备
func DelDev(data []byte) error {

	var dev Model.Dev
	err := json.Unmarshal(data, &dev)
	if err != nil {
		return err
	}
	//获取DB
	db := Tools.GetDB()
	result := db.Delete(&dev)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateDev 更新设备
func UpdateDev(data []byte) error {

	var dev Model.Dev
	err := json.Unmarshal(data, &dev)
	if err != nil {
		return err
	}
	//获取DB
	db := Tools.GetDB()
	result := db.Save(&dev)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// 唤醒设备
func WeakUpDev(data []byte) {

}
