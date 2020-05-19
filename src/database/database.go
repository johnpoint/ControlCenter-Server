package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"main/src/model"
	"time"
)

func initDatabase() *gorm.DB {
	conf := loadConfig()
	db, err := gorm.Open("sqlite3", conf.Database)
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}

//Server
func addServer(server model.Server) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	db.Create(&server)
	if !(db.NewRecord(server)) {
		return true
	}
	return false
}

func updateServer(where model.Server, server model.Server) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	db.Model(&server).Where(where).Update(server)
	if len(getServer(server)) != 0 {
		return true
	}
	return false
}

func getServer(server model.Server) []model.Server {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	servers := []model.Server{}
	db.Where(server).Find(&servers)
	return servers
}

func delServer(id int64, uid int64) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	server := model.Server{ID: id, UID: uid}
	db.Where(server).Delete(model.Server{})
	return true
}

//User
func addUser(user model.User) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.Create(&user)
	if !(db.NewRecord(user)) {
		return true
	}
	return false
}

func updateUser(where model.User, user model.User) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.Model(&user).Where(where).Update(user)
	if len(getUser(user)) != 0 {
		return true
	}
	return false
}
func getUser(user model.User) []model.User {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	users := []model.User{}
	db.Where(user).Find(&users)
	return users
}

//Domain
func addDomain(domain model.Domain) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Domain{})
	db.Create(&domain)
	if !(db.NewRecord(domain)) {
		return true
	}
	return false
}

func getDomain(domain model.Domain) []model.Domain {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Domain{})
	domains := []model.Domain{}
	db.Where(domain).Find(&domains)
	return domains
}

func updateDomain(where model.Domain, domain model.Domain) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Domain{})
	db.Model(&domain).Where(where).Update(domain)
	if len(getDomain(domain)) != 0 {
		return true
	}
	return false
}

//Site
func addSite(site model.Site) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Site{})
	db.Create(&site)
	if !(db.NewRecord(site)) {
		return true
	}
	return false
}

func getSite(site model.Site) []model.Site {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Site{})
	sites := []model.Site{}
	if site.Name == "*" {
		db.Find(&sites)
	} else {
		db.Where(site).Find(&sites)
	}
	return sites
}

func delSite(site model.Site) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Site{})
	db.Where(site).Delete(model.Site{})
	return true
}

//Cer
func addCer(certificate model.Certificate) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Certificate{})
	db.Create(&certificate)
	if !(db.NewRecord(certificate)) {
		return true
	}
	return false
}

// LinkCer link
func LinkServer(serverLink model.ServerLink) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.ServerLink{})
	db.Create(&serverLink)
	if !(db.NewRecord(serverLink)) {
		return true
	}
	return false
}

func getLinkCer(serverLink model.ServerLink) []model.ServerLink {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.ServerLink{})
	ServerLinks := []model.ServerLink{}
	serverLink.SiteID = 0
	db.Where(serverLink).Find(&ServerLinks)
	return ServerLinks
}

func getLinkSite(serverLink model.ServerLink) []model.ServerLink {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.ServerLink{})
	ServerLinks := []model.ServerLink{}
	serverLink.CertificateID = 0
	db.Where(serverLink).Find(&ServerLinks)
	return ServerLinks
}

func UnLinkServer(serverLink model.ServerLink) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.ServerLink{})
	db.Where(serverLink).Delete(model.ServerLink{})
	return true
}

func delCer(certificate model.Certificate) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Certificate{})
	db.Where(certificate).Delete(model.Certificate{})
	return true
}

func updateCer(where model.Certificate, certificate model.Certificate) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Certificate{})
	db.Model(&certificate).Where(where).Update(certificate)
	if len(getCer(certificate)) != 0 {
		return true
	}
	return false
}

func getCer(certificate model.Certificate) []model.Certificate {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Certificate{})
	certificates := []model.Certificate{}
	db.Where(certificate).Find(&certificates)
	return certificates
}

//System Config
func setConfig(config model.SysConfig) bool {
	if len(getConfig(model.SysConfig{Name: config.Name})) == 0 {
		return addConfig(config)
	}
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.SysConfig{})
	db.Model(&config).Where(model.SysConfig{Name: config.Name}).Update(config)
	if len(getConfig(config)) == 0 {
		return false
	}
	return true
}

func addConfig(config model.SysConfig) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.SysConfig{})
	db.Create(&config)
	if len(getConfig(config)) == 0 {
		return false
	}
	return true
}

func getConfig(config model.SysConfig) []model.SysConfig {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.SysConfig{})
	configs := []model.SysConfig{}
	db.Where(config).Find(&configs)
	return configs
}

// Docker
// 要传入Userid
func getDocker(docker model.Docker) []model.Docker {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Docker{})
	dockers := []model.Docker{}
	db.Where(docker).Find(&dockers)
	return dockers
}

// 要传入Userid
func editDocker(docker model.Docker) bool {
	if len(getDocker(model.Docker{ID: docker.ID})) != 0 {
		db := initDatabase()
		defer db.Close()
		db.AutoMigrate(&model.Docker{})
		db.Model(&docker).Where(model.Docker{ID: docker.ID, UID: docker.UID}).Update(docker)
		if len(getDocker(docker)) == 0 {
			return false
		}
		return true
	}
	return false
}

// 要传入Userid
func delDocker(docker model.Docker) bool {
	if len(getDocker(model.Docker{ID: docker.ID})) != 0 {
		db := initDatabase()
		defer db.Close()
		db.AutoMigrate(&model.Docker{})
		db.Model(&docker).Where(model.Docker{ID: docker.ID}).Update(docker)
		if len(getDocker(docker)) == 0 {
			return true
		}
		return false
	}
	return true
}

// 要传入Useriddocker
func addDocker(docker model.Docker) bool {
	if len(getDocker(model.Docker{Name: docker.Name, UID: docker.UID})) == 0 {
		db := initDatabase()
		defer db.Close()
		db.AutoMigrate(&model.Docker{})
		db.Create(&docker)
		if !(db.NewRecord(docker)) {
			return true
		}
		return false
	}
	return false
}

func addLog(service string, even string, level int64) bool {
	log := model.LogInfo{Service: service, Info: even, Level: level, CreatedAt: time.Now()}
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.LogInfo{})
	db.Create(&log)
	if !(db.NewRecord(log)) {
		return true
	}
	return false
}
