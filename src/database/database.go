package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"sync"
	"time"
)

//var conf = config.LoadConfig()

//var redisEable = conf.RedisConfig.Enable

var mutex sync.Mutex

func initDatabase() *gorm.DB {
	conf := config.LoadConfig()
	db, err := gorm.Open("sqlite3", conf.Database)
	if conf.Debug {
		db.LogMode(true)
	}
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}

/*
func initRedis() *redis.Client {
	conf := config.LoadConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.RedisConfig.Addr,
		Password: conf.RedisConfig.Password, // no password set
		DB:       conf.RedisConfig.DB,       // use default DB
	})
	return rdb
}

func redisGet(key string) string {
	rdb := initRedis()
	ctx := context.Background()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "key does not exists"
		}
		log.Print(err)
		return "err"
	}
	return val
}

func redisSet(key string, value string, exp time.Duration) string {
	rdb := initRedis()
	ctx := context.Background()
	_, err := rdb.Set(ctx, key, value, exp).Result()
	if err != nil {
		if redisGet(key) != value {
			return "data set fail"
		}
		return "err"
	}
	return "ok"
}*/

//Server
func AddServer(server model.Server) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return false
	}

	tx.AutoMigrate(&model.Server{})
	if err := tx.Create(&server).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func UpdateServer(where model.Server, server model.Server) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Server{})
	if err := tx.Model(&server).Where(where).Update(server).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetServer(server model.Server) []model.Server {
	/*data := "key does not exists"
	if redisEable {
		serverjson, _ := json.Marshal(server)
		data = redisGet(string(serverjson)) //check cache
	}
	if data == "key does not exists" {*/
	db := initDatabase()
	defer db.Close()
	defer db.Close()
	db.AutoMigrate(&model.Server{})
	servers := []model.Server{}
	db.Where(server).Find(&servers)
	return servers
	/*} else {
		servers := []model.Server{}
		json.Unmarshal([]byte(data), &servers)
		return servers
	}*/
}

func DelServer(id int64, uid int64) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Server{})
	server := model.Server{ID: id, UID: uid}
	if err := tx.Where(server).Delete(model.Server{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//User
func AddUser(user model.User) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.User{})
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func UpdateUser(where model.User, user model.User) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.User{})
	if err := tx.Model(&user).Where(where).Update(user).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetUser(user model.User) []model.User {
	db := initDatabase()
	defer db.Close()
	defer db.Close()
	db.AutoMigrate(&model.User{})
	users := []model.User{}
	db.Where(user).Find(&users)
	return users
}

func DelUser(user model.User) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.User{})
	if err := tx.Where(user).Delete(model.User{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//Domain
/*func addDomain(domain model.Domain) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Domain{})
	if err := tx.Create(&domain).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}*/

func GetDomain(domain model.Domain) []model.Domain {
	db := initDatabase()
	defer db.Close()
	defer db.Close()
	db.AutoMigrate(&model.Domain{})
	domains := []model.Domain{}
	db.Where(domain).Find(&domains)
	return domains
}

func UpdateDomain(where model.Domain, domain model.Domain) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Domain{})
	if err := tx.Model(&domain).Where(where).Update(domain).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//Site
func AddSite(site model.Site) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Site{})
	if err := tx.Create(&site).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetSite(site model.Site) []model.Site {
	db := initDatabase()
	defer db.Close()
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
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Site{})
	if err := tx.Where(site).Delete(model.Site{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//Cer
func AddCer(certificate model.Certificate) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Certificate{})
	if err := tx.Create(&certificate).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DelCer(certificate model.Certificate) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Certificate{})
	if err := tx.Where(certificate).Delete(model.Certificate{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func UpdateCer(where model.Certificate, certificate model.Certificate) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Certificate{})
	if err := tx.Model(&certificate).Where(where).Update(certificate).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetCer(certificate model.Certificate) []model.Certificate {
	db := initDatabase()
	defer db.Close()
	defer db.Close()
	db.AutoMigrate(&model.Certificate{})
	certificates := []model.Certificate{}
	db.Where(certificate).Find(&certificates)
	return certificates
}

// LinkCer link
func LinkServer(serverLink model.ServerLink) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.ServerLink{})
	if err := tx.Create(&serverLink).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetLinkCer(serverLink model.ServerLink) []model.ServerLink {
	db := initDatabase()
	defer db.Close()
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
	defer db.Close()
	db.AutoMigrate(&model.ServerLink{})
	ServerLinks := []model.ServerLink{}
	serverLink.CertificateID = 0
	db.Where(serverLink).Find(&ServerLinks)
	return ServerLinks
}

func UnLinkServer(serverLink model.ServerLink) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.ServerLink{})
	if err := tx.Where(serverLink).Delete(model.ServerLink{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

// TODO: sql事务改造
//System Config
func SetConfig(config model.SysConfig) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if len(GetConfig(model.SysConfig{Name: config.Name})) == 0 {
		return addConfig(config)
	}
	db := initDatabase()
	defer db.Close()
	defer db.Close()
	db.AutoMigrate(&model.SysConfig{})
	db.Model(&config).Where(model.SysConfig{Name: config.Name}).Update(config)
	if len(GetConfig(config)) == 0 {
		return false
	}
	return true
}

func addConfig(config model.SysConfig) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
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
	defer db.Close()
	db.AutoMigrate(&model.Docker{})
	dockers := []model.Docker{}
	db.Where(docker).Find(&dockers)
	return dockers
}

// 要传入Userid
func EditDocker(docker model.Docker) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if len(GetDocker(model.Docker{ID: docker.ID})) != 0 {
		db := initDatabase()
		defer db.Close()
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
/*func delDocker(docker model.Docker) bool {
	if len(GetDocker(model.Docker{ID: docker.ID})) != 0 {
		db := initDatabase()
	defer db.Close()
		defer db.Close()
		db.AutoMigrate(&model.Docker{})
		db.Model(&docker).Where(model.Docker{ID: docker.ID}).Update(docker)
		if len(GetDocker(docker)) == 0 {
			return true
		}
		return false
	}
	return true
}*/

// 要传入Useriddocker
func AddDocker(docker model.Docker) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if len(GetDocker(model.Docker{Name: docker.Name, UID: docker.UID})) == 0 {
		db := initDatabase()
		defer db.Close()
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

func AddLog(service string, event string, level int64) bool {
	mutex.Lock()
	defer mutex.Unlock()
	newLog := model.LogInfo{Service: service, Info: event, Level: level, CreatedAt: time.Now()}
	db := initDatabase()
	defer db.Close()
	defer db.Close()
	db.AutoMigrate(&model.LogInfo{})
	db.Create(&newLog)
	if !(db.NewRecord(newLog)) {
		return true
	}
	return false
}

func AddEvent(eventType int64, target int64, code int64, info string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Event{})
	if len(GetEvent(eventType, target, code, info, 0)) != 0 {
		if err := tx.Model(&model.Event{}).Where(model.Event{Type: eventType, TargetID: target, Code: code, Info: info}).Update(model.Event{Active: 1}).Error; err != nil {
			tx.Rollback()
			return false
		}
	} else {
		if err := tx.Create(&model.Event{Type: eventType, TargetID: target, Code: code, Info: info, Active: 1}).Error; err != nil {
			tx.Rollback()
			panic(err)
			return false
		}
	}
	tx.Commit()
	return true
}

func GetEvent(eventType int64, target int64, code int64, info string, active int64) []model.Event {
	db := initDatabase()
	defer db.Close()
	defer db.Close()
	db.AutoMigrate(&model.Event{})
	events := []model.Event{}
	db.Where(model.Event{Type: eventType, TargetID: target, Code: code, Info: info, Active: active}).Find(&events)
	return events
}

func FinishEvent(id int64) bool {
	mutex.Lock()
	defer mutex.Unlock()
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	tx.AutoMigrate(&model.Event{})
	event := model.Event{Active: 2}
	if err := tx.Model(&event).Where(model.Event{ID: id, Active: 1}).Update(event).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
