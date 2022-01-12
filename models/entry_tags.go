package models

type EntryTags struct {
	EntryId uint `gorm:"comment:'内容ID';primarykey;autoIncrement:false"`
	TagId   uint `gorm:"comment:'标签ID';primarykey;autoIncrement:false"`
}
