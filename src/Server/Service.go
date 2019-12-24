package main

type Service struct {
	Id       int64 `gorm:"AUTO_INCREMENT"`
	Name     string
	Enable   string
	Disable  string
	Status   int64
	Serverid int64
}
