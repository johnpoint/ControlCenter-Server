package main

// Service TODO
type Service struct {
	ID       int64 `gorm:"AUTO_INCREMENT"`
	Name     string
	Enable   string
	Disable  string
	Status   int64
	Serverid int64
}
