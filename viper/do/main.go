package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type mySqlConfig struct {
	Source []string
	Replicas []string
	LogMode int
	Prefix string
	MaxIdle int
	MaxOpen int
}

func main() {
	config:=viper.New();
	config.AddConfigPath("./viper/config")
	config.SetConfigName("env")
	config.SetConfigType("yaml")
	if err:=config.ReadInConfig();err!=nil{
		panic(err)
	}
	mysql:=&mySqlConfig{}
	config.Sub("Mysql.Default").Unmarshal(mysql)
	fmt.Println(mysql)
	fmt.Println(config.Get("Server.http.Port"))

}
