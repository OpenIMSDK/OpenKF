// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

var Config *config

func ConfigInit(configPath string) {
	// init viper
	initViper(configPath)

	// init Configuration
	Config = &config{
		App: App{
			Version: GetString("app.version"),
			Debug:   GetBool("app.debug"),
			LogFile: GetString("app.log_file"),
		},
		JWT: JWT{
			Secret:     GetString("jwt.secret"),
			Issuer:     GetString("jwt.issuer"),
			ExpireDays: GetInt("jwt.expire_days"),
		},
		Server: Server{
			Ip:   GetString("server.ip"),
			Port: GetInt("server.port"),
		},
		Mysql: Mysql{
			Ip:       GetString("mysql.ip"),
			Port:     GetInt("mysql.port"),
			Username: GetString("mysql.username"),
			Password: GetString("mysql.password"),
			Database: GetString("mysql.database"),
		},
		Redis: Redis{
			Ip:       GetString("redis.ip"),
			Port:     GetInt("redis.port"),
			Password: GetString("redis.password"),
			Database: GetInt("redis.database"),
		},
		Minio: Minio{
			Ip:              GetString("minio.ip"),
			Port:            GetInt("minio.port"),
			AccessKeyId:     GetString("minio.access_key_id"),
			SecretAccessKey: GetString("minio.secret_access_key"),
			Bucket:          GetString("minio.bucket"),
			AppBucket:       GetString("minio.app_bucket"),
			Location:        GetString("minio.location"),
		},
		Email: Email{
			Host:     GetString("email.host"),
			Port:     GetInt("email.port"),
			From:     GetString("email.from"),
			Nickname: GetString("email.nickname"),
			Password: GetString("email.password"),
		},
	}
}

type config struct {
	App    App
	JWT    JWT
	Server Server
	Mysql  Mysql
	Redis  Redis
	Minio  Minio
	Email  Email
}

type App struct {
	Version string `mapstructure:"version"`
	Debug   bool   `mapstructure:"debug"`
	LogFile string `mapstructure:"log_file"`
}

type JWT struct {
	Secret     string `mapstructure:"secret"`
	Issuer     string `mapstructure:"issuer"`
	ExpireDays int    `mapstructure:"expire_days"`
}

type Server struct {
	Ip   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`
}

type Mysql struct {
	Ip       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Redis struct {
	Ip       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type Minio struct {
	Ip              string `mapstructure:"ip"`
	Port            int    `mapstructure:"port"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Bucket          string `mapstructure:"bucket"`
	AppBucket       string `mapstructure:"app_bucket"`
	Location        string `mapstructure:"location"`
}

type Email struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	From     string `mapstructure:"from"`
	Nickname string `mapstructure:"nickname"`
	Password string `mapstructure:"password"`
}
