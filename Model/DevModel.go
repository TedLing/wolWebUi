package Model

type Tabler interface {
	TableName() string
}

func (Dev) TableName() string {
	return "devlist"
}

// Dev 设备列表
type Dev struct {
	Id      int `gorm:"primaryKey"`
	Devname string
	Mac     string
}
