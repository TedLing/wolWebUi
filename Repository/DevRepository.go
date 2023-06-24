package Repository

import (
	"encoding/json"
	"fmt"
	"wolWebUi/Tools"
)
import "wolWebUi/Model"

// GetDevList 获取所有设备
func GetDevList() *[]Model.Dev {
	//获取DB
	db := Tools.GetDB()
	//执行数据库查询操作
	var devlist []Model.Dev
	db.Find(&devlist)
	return &devlist
}

func AddDev(data []byte) error {

	var dev Model.Dev

	fmt.Println(string(data))
	err := json.Unmarshal(data, &dev)
	if err != nil {

		fmt.Println(err.Error())
		return err
	}

	fmt.Println(dev)

	//获取DB
	db := Tools.GetDB()
	result := db.Create(&dev)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
