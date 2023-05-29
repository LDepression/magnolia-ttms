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
	Serve Serve `mapstructure:"Serve"`
	App   App   `mapstructure:"App"`
	Log   Log   `mapstructure:"Log"`
	Mysql Mysql `mapstructure:"Mysql"`
	Redis Redis `mapstructure:"Redis"`
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
