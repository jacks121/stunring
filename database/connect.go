package database

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"swetelove/utils"

	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	RedisClient   *redis.Client
	MysqlDB       *gorm.DB
	ElasticClient *elastic.Client
	ctx           = context.Background()
)

func Init() {
	configPath := filepath.Join(".", "config.toml")
	config := utils.NewConfig(configPath)
	connectRedis(config)
	connectMysql(config)
	connectElasticsearch(config)
}

func connectRedis(config *utils.Config) {
	dbString := config.Get("Redis.DB")
	dbInt, err := strconv.Atoi(dbString)
	if err != nil {
		// 处理转换错误
		panic(err)
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Get("Redis.Addr"), // replace with your Redis address
		Password: "",                       // replace with your password if needed
		DB:       dbInt,                    // default DB
	})

	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func connectMysql(config *utils.Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.Get("Database.User"), config.Get("Database.Password"), config.Get("Database.Host"), config.Get("Database.Port"), config.Get("Database.Name"), config.Get("Database.Charset"))
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func connectElasticsearch(config *utils.Config) {
	var err error
	ElasticClient, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.Get("ES.Url"))) // replace with your Elasticsearch URL
	if err != nil {
		panic(err)
	}
}

// CloseConnections 关闭所有开启的数据库连接
func CloseConnections() {
	// Close Redis connection
	err := RedisClient.Close()
	if err != nil {
		utils.LogError.Fatalf("Failed to close Redis connection: %v", err)
	}

	// Close MySQL connection
	sqlDB, err := MysqlDB.DB()
	if err != nil {
		utils.LogError.Fatalf("Failed to retrieve sqlDB: %v", err)
	}
	err = sqlDB.Close()
	if err != nil {
		utils.LogError.Fatalf("Failed to close MySQL connection: %v", err)
	}
}
