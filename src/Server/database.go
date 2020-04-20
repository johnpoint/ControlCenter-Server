package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
func addServer(server Server) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Server{})
	db.Create(&server)
	if !(db.NewRecord(server)) {
		return true
	}
	return false
}

func updateServer(where Server, server Server) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Server{})
	db.Model(&server).Where(where).Update(server)
	if len(getServer(server)) != 0 {
		return true
	}
	return false
}

func getServer(server Server) []Server {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Server{})
	servers := []Server{}
	db.Where(server).Find(&servers)
	return servers
}

func delServer(id int64, uid int64) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Server{})
	server := Server{ID: id, UID: uid}
	db.Where(server).Delete(Server{})
	return true
}

//User
func addUser(user User) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&User{})
	db.Create(&user)
	if !(db.NewRecord(user)) {
		return true
	}
	return false
}

func updateUser(where User, user User) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&User{})
	db.Model(&user).Where(where).Update(user)
	if len(getUser(user)) != 0 {
		return true
	}
	return false
}
func getUser(user User) []User {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&User{})
	users := []User{}
	db.Where(user).Find(&users)
	return users
}

//Domain
func addDomain(domain Domain) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Domain{})
	db.Create(&domain)
	if !(db.NewRecord(domain)) {
		return true
	}
	return false
}

func getDomain(domain Domain) []Domain {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Domain{})
	domains := []Domain{}
	db.Where(domain).Find(&domains)
	return domains
}

func updateDomain(where Domain, domain Domain) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Domain{})
	db.Model(&domain).Where(where).Update(domain)
	if len(getDomain(domain)) != 0 {
		return true
	}
	return false
}

//Site
func addSite(site Site) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Site{})
	db.Create(&site)
	if !(db.NewRecord(site)) {
		return true
	}
	return false
}

func getSite(site Site) []Site {
	db := initDatabase()
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

func delSite(site Site) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Site{})
	db.Where(site).Delete(Site{})
	return true
}

//Cer
func addCer(certificate Certificate) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Certificate{})
	db.Create(&certificate)
	if !(db.NewRecord(certificate)) {
		return true
	}
	return false
}

// LinkCer link
func LinkServer(serverLink ServerLink) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&ServerLink{})
	db.Create(&serverLink)
	if !(db.NewRecord(serverLink)) {
		return true
	}
	return false
}

func getLinkCer(serverLink ServerLink) []ServerLink {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&ServerLink{})
	ServerLinks := []ServerLink{}
	serverLink.SiteID = 0
	db.Where(serverLink).Find(&ServerLinks)
	return ServerLinks
}

func getLinkSite(serverLink ServerLink) []ServerLink {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&ServerLink{})
	ServerLinks := []ServerLink{}
	serverLink.CertificateID = 0
	db.Where(serverLink).Find(&ServerLinks)
	return ServerLinks
}

func UnLinkServer(serverLink ServerLink) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&ServerLink{})
	db.Where(serverLink).Delete(ServerLink{})
	return true
}

func delCer(certificate Certificate) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Certificate{})
	db.Where(certificate).Delete(Certificate{})
	return true
}

func updateCer(where Certificate, certificate Certificate) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Certificate{})
	db.Model(&certificate).Where(where).Update(certificate)
	if len(getCer(certificate)) != 0 {
		return true
	}
	return false
}

func getCer(certificate Certificate) []Certificate {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Certificate{})
	certificates := []Certificate{}
	db.Where(certificate).Find(&certificates)
	return certificates
}

//System Config
func setConfig(config SysConfig) bool {
	if len(getConfig(SysConfig{Name: config.Name})) == 0 {
		return addConfig(config)
	}
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&SysConfig{})
	db.Model(&config).Where(SysConfig{Name: config.Name}).Update(config)
	if len(getConfig(config)) == 0 {
		return false
	}
	return true
}

func addConfig(config SysConfig) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&SysConfig{})
	db.Create(&config)
	if len(getConfig(config)) == 0 {
		return false
	}
	return true
}

func getConfig(config SysConfig) []SysConfig {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&SysConfig{})
	configs := []SysConfig{}
	db.Where(config).Find(&configs)
	return configs
}

// Docker
// 要传入Userid
func getDocker(docker Docker) []Docker {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&Docker{})
	dockers := []Docker{}
	db.Where(docker).Find(&dockers)
	return dockers
}

// 要传入Userid
func editDocker(docker Docker) bool {
	if len(getDocker(Docker{ID: docker.ID})) != 0 {
		db := initDatabase()
		defer db.Close()
		db.AutoMigrate(&Docker{})
		db.Model(&docker).Where(Docker{ID: docker.ID, UID: docker.UID}).Update(docker)
		if len(getDocker(docker)) == 0 {
			return false
		}
		return true
	}
	return false
}

// 要传入Userid
func delDocker(docker Docker) bool {
	if len(getDocker(Docker{ID: docker.ID})) != 0 {
		db := initDatabase()
		defer db.Close()
		db.AutoMigrate(&Docker{})
		db.Model(&docker).Where(Docker{ID: docker.ID}).Update(docker)
		if len(getDocker(docker)) == 0 {
			return true
		}
		return false
	}
	return true
}

// 要传入Useriddocker
func addDocker(docker Docker) bool {
	if len(getDocker(Docker{Name: docker.Name, UID: docker.UID})) == 0 {
		db := initDatabase()
		defer db.Close()
		db.AutoMigrate(&Docker{})
		db.Create(&docker)
		if !(db.NewRecord(docker)) {
			return true
		}
		return false
	}
	return false
}

func addLog(service string, even string, level int64) bool {
	log := LogInfo{Service: service, Info: even, Level: level, CreatedAt: time.Now()}
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&LogInfo{})
	db.Create(&log)
	if !(db.NewRecord(log)) {
		return true
	}
	return false
}
