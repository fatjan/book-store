package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //import driver
	"github.com/spf13/viper"
)

var DB *gorm.DB

// MaxAgeSessionTTL this is for store maxage session
type MaxAgeSessionTTL struct {
	Days   int
	Months int
	Years  int
}

// Database this is for store DB
type Database struct {
	Driver            string
	Host              string
	User              string
	Password          string
	DBName            string
	DBNumber          int
	Port              int
	APIUrl            string
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
	MaxAge            MaxAgeSessionTTL
}

// MysqlConnect connect to mysql using config name. return *gorm.DB incstance
func MysqlConnect(configName string) *gorm.DB {
	mysql := LoadDBConfig(configName)
	connectionString := mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + strconv.Itoa(mysql.Port) + ")/" + mysql.DBName + "?charset=utf8&parseTime=True&loc=Local"
	connection, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Established")

	if mysql.DebugMode {
		return connection.Debug()
	}

	return connection
}

// OpenDbPool for open DB Pool
func OpenDbPool() {
	DB = MysqlConnect("mysql")
	pool := viper.Sub("database.mysql.pool")
	DB.DB().SetMaxOpenConns(pool.GetInt("maxOpenConns"))
	DB.DB().SetMaxIdleConns(pool.GetInt("maxIdleConns"))
	DB.DB().SetConnMaxLifetime(pool.GetDuration("maxLifetime") * time.Second)
	DB.SingularTable(true)
	DB.Callback().Create().Remove("gorm:update_time_stamp")
}

func getLocalIPAddress() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}
