package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"time"
)

func initDatabase() *gorm.DB {
	conf := LoadConfig()
	db, err := gorm.Open("sqlite3", conf.Database)
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}

//Server
func AddServer(server model.Server) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	db.Create(&server)
	if !(db.NewRecord(server)) {
		return true
	}
	return false
}

func UpdateServer(where model.Server, server model.Server) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	db.Model(&server).Where(where).Update(server)
	if len(GetServer(server)) != 0 {
		return true
	}
	return false
}

func GetServer(server model.Server) []model.Server {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	servers := []model.Server{}
	db.Where(server).Find(&servers)
	return servers
}

func DelServer(id int64, uid int64) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	server := model.Server{ID: id, UID: uid}
	db.Where(server).Delete(model.Server{})
	return true
}

//User
func AddUser(user model.User) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.Create(&user)
	if !(db.NewRecord(user)) {
		return true
	}
	return false
}

func UpdateUser(where model.User, user model.User) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.Model(&user).Where(where).Update(user)
	if len(GetUser(user)) != 0 {
		return true
	}
	return false
}
func GetUser(user model.User) []model.User {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	users := []model.User{}
	db.Where(user).Find(&users)
	return users
}

func DelUser(user model.User) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.Where(user).Delete(model.User{})
	return true
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

func GetDomain(domain model.Domain) []model.Domain {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Domain{})
	domains := []model.Domain{}
	db.Where(domain).Find(&domains)
	return domains
}

func UpdateDomain(where model.Domain, domain model.Domain) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Domain{})
	db.Model(&domain).Where(where).Update(domain)
	if len(GetDomain(domain)) != 0 {
		return true
	}
	return false
}

//Site
func AddSite(site model.Site) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Site{})
	db.Create(&site)
	if !(db.NewRecord(site)) {
		return true
	}
	return false
}

func GetSite(site model.Site) []model.Site {
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

func DelSite(site model.Site) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Site{})
	db.Where(site).Delete(model.Site{})
	return true
}

//Cer
func AddCer(certificate model.Certificate) bool {
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

func GetLinkCer(serverLink model.ServerLink) []model.ServerLink {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.ServerLink{})
	ServerLinks := []model.ServerLink{}
	serverLink.SiteID = 0
	db.Where(serverLink).Find(&ServerLinks)
	return ServerLinks
}

func GetLinkSite(serverLink model.ServerLink) []model.ServerLink {
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

func DelCer(certificate model.Certificate) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Certificate{})
	db.Where(certificate).Delete(model.Certificate{})
	return true
}

func UpdateCer(where model.Certificate, certificate model.Certificate) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Certificate{})
	db.Model(&certificate).Where(where).Update(certificate)
	if len(GetCer(certificate)) != 0 {
		return true
	}
	return false
}

func GetCer(certificate model.Certificate) []model.Certificate {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Certificate{})
	certificates := []model.Certificate{}
	db.Where(certificate).Find(&certificates)
	return certificates
}

//System Config
func SetConfig(config model.SysConfig) bool {
	if len(GetConfig(model.SysConfig{Name: config.Name})) == 0 {
		return addConfig(config)
	}
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.SysConfig{})
	db.Model(&config).Where(model.SysConfig{Name: config.Name}).Update(config)
	if len(GetConfig(config)) == 0 {
		return false
	}
	return true
}

func addConfig(config model.SysConfig) bool {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.SysConfig{})
	db.Create(&config)
	if len(GetConfig(config)) == 0 {
		return false
	}
	return true
}

func GetConfig(config model.SysConfig) []model.SysConfig {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.SysConfig{})
	configs := []model.SysConfig{}
	db.Where(config).Find(&configs)
	return configs
}

// Docker
// 要传入Userid
func GetDocker(docker model.Docker) []model.Docker {
	db := initDatabase()
	defer db.Close()
	db.AutoMigrate(&model.Docker{})
	dockers := []model.Docker{}
	db.Where(docker).Find(&dockers)
	return dockers
}

// 要传入Userid
func EditDocker(docker model.Docker) bool {
	if len(GetDocker(model.Docker{ID: docker.ID})) != 0 {
		db := initDatabase()
		defer db.Close()
		db.AutoMigrate(&model.Docker{})
		db.Model(&docker).Where(model.Docker{ID: docker.ID, UID: docker.UID}).Update(docker)
		if len(GetDocker(docker)) == 0 {
			return false
		}
		return true
	}
	return false
}

// 要传入Userid
func delDocker(docker model.Docker) bool {
	if len(GetDocker(model.Docker{ID: docker.ID})) != 0 {
		db := initDatabase()
		defer db.Close()
		db.AutoMigrate(&model.Docker{})
		db.Model(&docker).Where(model.Docker{ID: docker.ID}).Update(docker)
		if len(GetDocker(docker)) == 0 {
			return true
		}
		return false
	}
	return true
}

// 要传入Useriddocker
func AddDocker(docker model.Docker) bool {
	if len(GetDocker(model.Docker{Name: docker.Name, UID: docker.UID})) == 0 {
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

func AddLog(service string, even string, level int64) bool {
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
