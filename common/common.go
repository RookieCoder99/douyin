package common

import (
	"context"
	"douyin/config"
	"douyin/model"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
	"time"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
	Rmq    *utils.Session
)

// redis key前缀
const (
	UserLoginPrefix          = "USER_LOGIN:"
	UserFollowCountPrefix    = "USER_FOLLOW_COUNT:"
	UserFollowerCountPrefix  = "USER_FOLLOWER_COUNT:"
	VideoFavoriteCountPrefix = "VIDEO_FAVORITE_COUNT:"
	VideoCommentCountPrefix  = "VIDEO_COMMENT_COUNT:"

	UserHasVideoPrefix      = "USER_HAS_VIDEO_SET:"
	UserFavoriteVideoPrefix = "USER_FAVORITE_VIDEO_SET:"

	UserFollowPrefix   = "USER_FOLLOW_SET:"
	UserFollowerPrefix = "USER_FOLLOWER_SET:"

	UserTotalFavoritedPrefix = "USER_Total_Favorited:"
	UserFavoriteCountPrefix  = "USER_Favorite_Count:"
)

const (
	OK                  int32 = 0
	ParamInvalid        int32 = 1   // 参数不合法
	UserHasExisted      int32 = 2   // 该 Username 已存在
	UserHasDeleted      int32 = 3   // 用户已删除
	UserNotExisted      int32 = 4   // 用户不存在
	WrongPassword       int32 = 5   // 密码错误
	LoginRequired       int32 = 6   // 用户未登录
	UploadError         int32 = 7   //视频上传失败
	FavoriteActionError int32 = 8   //点赞失败
	FavoriteCancelError int32 = 9   //点赞取消失败
	CommentActionError  int32 = 10  //评论失败
	CommentDeleteError  int32 = 11  //评论取消失败
	FollowActionError   int32 = 12  //关注失败
	FollowCancelError   int32 = 13  //关注取消失败
	RepeatRequest       int32 = 254 // 重复请求
	UnknownError        int32 = 255 // 未知错误

)

var RespMsg = map[int32]string{
	OK:                  "成功",
	ParamInvalid:        "参数不合法",
	WrongPassword:       "密码错误",
	UploadError:         "视频上传失败",
	FavoriteActionError: "点赞失败",
	FavoriteCancelError: "点赞取消失败",
	CommentActionError:  "评论失败",
	CommentDeleteError:  "评论取消失败",
	FollowActionError:   "关注失败",
	FollowCancelError:   "关注取消失败",
}

func Init() {
	Db = initMysqlDB()
	Rdb = initRedisClient()
	//Rmq = initRabbitmq()
}

func initMysqlDB() *gorm.DB {
	dsn := Config.Mysql.Username + ":" + Config.Mysql.Password + "@tcp(" + Config.Mysql.Ip + ":" + Config.Mysql.Port + ")" +
		"/" + Config.Mysql.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})

	if err != nil {
		log.Fatal("数据库连接失败 ", err)
		return nil
	}
	// 设置连接池，空闲连接
	sqlDb, _ := db.DB()
	//defer sqlDb.Close()

	// 设置最大连接数
	sqlDb.SetMaxOpenConns(Config.Mysql.MaxOpenConns)
	// 设置空闲连接数
	sqlDb.SetMaxIdleConns(Config.Mysql.MaxIdleConns)
	db.AutoMigrate(&model.TUser{})
	db.AutoMigrate(&model.TVideo{})
	db.AutoMigrate(&model.TFavorite{})
	db.AutoMigrate(&model.TComment{})
	db.AutoMigrate(&model.TRelation{})
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

func initRabbitmq() *utils.Session {
	Rmq = utils.New(Config.RabbitMQ.Queue, Config.RabbitMQ.Url)
	//return Rmq

	//message := []byte("message")
	// Attempt to push a message every 2 seconds
	//for {
	//	time.Sleep(time.Second * 2)
	//	if err := queue.Push(message); err != nil {
	//		fmt.Printf("Push failed: %s\n", err)
	//	} else {
	//		fmt.Println("Push succeeded!")
	//	}
	//}
	return Rmq
}

func Resp(c *gin.Context, resMap map[string]interface{}) {
	c.JSON(http.StatusOK, resMap)
}

func Resp0(c *gin.Context, int32 int32, statusMsg string) {
	resMap := map[string]interface{}{
		"status_code": int32,
		"status_msg":  statusMsg,
	}
	c.JSON(http.StatusOK, resMap)
}

func Resp1(c *gin.Context, int32 int32, statusMsg string, key1 string, value1 interface{}) {
	resMap := map[string]interface{}{
		"status_code": int32,
		"status_msg":  statusMsg,
	}
	resMap[key1] = value1
	c.JSON(http.StatusOK, resMap)
}
func Resp2(c *gin.Context, int32 int32, statusMsg string, key1 string, value1 interface{}, key2 string, value2 interface{}) {
	resMap := map[string]interface{}{
		"status_code": int32,
		"status_msg":  statusMsg,
	}
	resMap[key1] = value1
	resMap[key2] = value2
	c.JSON(http.StatusOK, resMap)
}
