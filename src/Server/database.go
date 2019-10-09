package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase(path string) *gorm.DB {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}

func addServer(server Server) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Server{})
	db.Create(&server)
	if !(db.NewRecord(server)) {
		return true
	}
	return false
}

func addUser(user User) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&User{})
	db.Create(&user)
	if !(db.NewRecord(user)) {
		return true
	}
	return false
}

func addDomain(domain Domain) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Domain{})
	db.Create(&domain)
	if !(db.NewRecord(domain)) {
		return true
	}
	return false
}

func addSite(site Site) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Site{})
	db.Create(&site)
	if !(db.NewRecord(site)) {
		return true
	}
	return false
}

func getSite(site Site) []Site {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Site{})
	sites := []Site{}
	if site.Name == "*" {
		db.Find(&sites)
	} else {
		db.Where(site).Find(&sites)
	}
	return sites
}

func addCer(certificate Certificate) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Certificate{})
	db.Create(&certificate)
	if !(db.NewRecord(certificate)) {
		return true
	}
	return false
}

func delCer(certificate Certificate) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Certificate{})
	db.Where("id LIKE ?", certificate.ID).Delete(Certificate{})
	return true
}

func updateCer(where Certificate, certificate Certificate) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Certificate{})
	db.Model(&certificate).Where(where).Update(certificate)
	if len(getCer(certificate)) != 0 {
		return true
	}
	return false
}

func getCer(certificate Certificate) []Certificate {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Certificate{})
	certificates := []Certificate{}
	if certificate.Name == "*" {
		db.Find(&certificates)
		return certificates
	}
	db.Where(certificate).Find(&certificates)
	return certificates
}

func getUser(user User) []User {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&User{})
	users := []User{}
	db.Where(user).Find(&users)
	return users
}

func getServer(server Server) []Server {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Server{})
	servers := []Server{}
	if server.Hostname == "*" {
		db.Find(&servers)
	} else {
		db.Where(server).Find(&servers)
	}
	return servers
}

func getDomain(domain Domain) []Domain {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Domain{})
	domains := []Domain{}
	db.Where(domain).Find(&domains)
	return domains
}

func updateDomain(where Domain, domain Domain) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Domain{})
	db.Model(&domain).Where(where).Update(domain)
	if len(getDomain(domain)) != 0 {
		return true
	}
	return false
}

func updateServer(where Server, server Server) bool {
	db := initDatabase("test.db")
	defer db.Close()
	db.AutoMigrate(&Server{})
	db.Model(&server).Where(where).Update(server)
	if len(getServer(server)) != 0 {
		return true
	}
	return false
}