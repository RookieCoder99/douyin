package common

import (
	"context"
	"douyin/config"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
)

type ErrNo int

const (
	OK             ErrNo = 0
	ParamInvalid   ErrNo = 1 // 参数不合法
	UserHasExisted ErrNo = 2 // 该 Username 已存在
	UserHasDeleted ErrNo = 3 // 用户已删除
	UserNotExisted ErrNo = 4 // 用户不存在
	WrongPassword  ErrNo = 5 // 密码错误
	LoginRequired  ErrNo = 6 // 用户未登录

	RepeatRequest ErrNo = 254 // 重复请求
	UnknownError  ErrNo = 255 // 未知错误

)

var RespMsg = map[ErrNo]string{
	OK:           "成功",
	ParamInvalid: "参数不合法",
}

func Init() {
	Db = initMysqlDB()
	Rdb = initRedisClient()
}

func initMysqlDB() *gorm.DB {
	dsn := Config.Mysql.Username + ":" + Config.Mysql.Password + "@tcp(" + Config.Mysql.Ip + ":" + Config.Mysql.Port + ")" +
		"/" + Config.Mysql.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("数据库连接失败 ", err)
		return nil
	}
	// 设置连接池，空闲连接
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	// 设置最大连接数
	sqlDb.SetMaxOpenConns(Config.Mysql.MaxOpenConns)
	// 设置空闲连接数
	sqlDb.SetMaxIdleConns(Config.Mysql.MaxIdleConns)
	db.AutoMigrate(&model.TUser{})
	db.AutoMigrate(&model.TVideo{})

	return db
}

func initRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Ip + ":" + Config.Redis.Port,
		Password: Config.Redis.Password, // no password set
		DB:       Config.Redis.Database, // use default DB
		PoolSize: Config.Redis.PoolSize, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatal("redis数据库连接失败 ", err)
		return nil
	}
	//time.Sleep(time.Second * 2)
	//rdb.Set(ctx, "test", "test", 10*time.Second)
	return rdb
}

func RespData(c *gin.Context, errNo ErrNo, v interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Code": errNo,
		"Data": v,
	})
}
