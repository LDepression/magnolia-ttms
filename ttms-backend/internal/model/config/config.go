/**
 * @Author: lenovo
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:36
 */

package config

import "time"

type AllConfig struct {
	Serve    Serve    `mapstructure:"Serve"`
	App      App      `mapstructure:"App"`
	Log      Log      `mapstructure:"Log"`
	Mysql    Mysql    `mapstructure:"Mysql"`
	Redis    Redis    `mapstructure:"Redis"`
	SMTPInfo SMTPInfo `mapstructure:"SMTPInfo"`
	Rule     Rule     `mapstructure:"Rule"`
	Work     Work     `mapstructure:"Work"`
}
type Serve struct {
	RunMode               string        `mapstructure:"RunMode"`
	Address               string        `mapstructure:"Address"`
	ReadTimeout           time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout          time.Duration `mapstructure:"WriteTimeout"`
	DefaultContextTimeout time.Duration `mapstructure:"DefaultContextTimeout"`
}

type App struct {
	Name    string `mapstructure:"Name"`
	Version string `mapstructure:"Version"`
}

type Log struct {
	Level         string `yaml:"Level"`
	LogSavePath   string `yaml:"LogSavePath"`
	LowLevelFile  string `yaml:"LowLevelFile"`
	LogFileExt    string `yaml:"LogFileExt"`
	HighLevelFile string `yaml:"HighLevelFile"`
	MaxSize       int    `yaml:"MaxSize"`
	MaxAge        int    `yaml:"MaxAge"`
	MaxBackups    int    `yaml:"MaxBackups"`
	Compress      bool   `yaml:"Compress"`
}

type Mysql struct {
	User     string `mapstructure:"user"`
	Password string ` mapstructure:"password"`
	Host     string ` mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string ` mapstructure:"dbName"`
}

type Redis struct {
	Addr     string ` mapstructure:"addr"`
	Password string ` mapstructure:"password"`
	PoolSize int    `mapstructure:"poolSize"`
}

type SMTPInfo struct {
	Host     string   `json:"host" mapstructure:"host"`
	Port     int      `json:"port" mapstructure:"port"`
	IsSSL    bool     `json:"isSSL" mapstructure:"isSSL"`
	UserName string   `json:"userName" mapstructure:"userName"`
	Password string   `json:"password" mapstructure:"password"`
	From     string   `json:"from" mapstructure:"from"`
	To       []string `json:"to" mapstructure:"to"`
}
type Rule struct {
	DelUserTime          time.Duration `json:"delUserTime" mapstructure:"delUserTime"`                   //延时删除用户的时间
	DelCodeTime          time.Duration `json:"delCodeTime" mapstructure:"delCodeTime"`                   //延时删除验证码的时间
	DefaultAccountAvatar string        `json:"DefaultAccountAvatar" mapstructure:"DefaultAccountAvatar"` //账户默认的头像
	DefaultClientTimeout time.Duration `json:"DefaultClientTimeout" mapstructure:"DefaultClientTimeout"` //客户端默认超时时间
	FileMaxSize          int64         `json:"FileMaxSize" mapstructure:"FileMaxSize"`
	DefaultInsertDataNum int           `json:"DefaultInsertDataNum" mapstructure:"DefaultInsertDataNum"`
}
type Work struct {
	TaskChanCapacity   int `json:"taskChanCapacity" mapstructure:"taskChanCapacity"`
	WorkerChanCapacity int `json:"workerChanCapacity" mapstructure:"workerChanCapacity"`
	WorkerNum          int `json:"workerNum" mapstructure:"workerNum"`
}
