package config

import (
	"core-server/util/json"
	"fmt"
	"os"
)
import "github.com/koding/multiconfig"

var (
	// C 全局配置
	C = new(Config)
)

func init() {
	loaders := []multiconfig.Loader{
		&multiconfig.TagLoader{},
		&multiconfig.EnvironmentLoader{},
	}

	loaders = append(loaders, &multiconfig.TOMLLoader{Path: "/config.toml"})

	m := multiconfig.DefaultLoader{
		Loader:    multiconfig.MultiLoader(loaders...),
		Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
	}

	m.MustLoad(C)

	PrintWithJSON()
}

// IsDebugMode 是否是debug模式
func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

// PrintWithJSON 基于JSON格式输出配置
func PrintWithJSON() {
	b, err := json.MarshalIndent(C, "", " ")
	if err != nil {
		os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
		return
	}
	os.Stdout.WriteString(string(b) + "\n")
}

// Config 配置参数
type Config struct {
	RunMode  string
	Root     Root
	GRPC     GRPC
	Gorm     Gorm
	MySQL    MySQL
	Postgres Postgres
}

// Root root用户
type Root struct {
	UserName string
	Password string
	RealName string
}

// GRPC 配置参数
type GRPC struct {
	Host string
	Port int
}

func (a GRPC) Addr() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

// Gorm gorm配置参数
type Gorm struct {
	Debug             bool
	DBType            string
	MaxLifetime       int
	MaxOpenConns      int
	MaxIdleConns      int
	TablePrefix       string
	EnableAutoMigrate bool
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

// Postgres postgres配置参数
type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// DSN 数据库连接串
func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s TimeZone=Asia/Shanghai",
		a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
}
