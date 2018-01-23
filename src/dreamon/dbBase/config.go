package dbBase

import (
	"dreamon/config"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type configuration struct {
	Mysql   config.MySqlConfig
	MongoDB config.MongoDBConfig
}

var C configuration

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	//viper.AddConfigPath("./src/dreamon/dbBase")
	viper.AddConfigPath("./dbBase")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("Config file changed:", e.Op)
		fmt.Println("Config file changed:", e.String())
	})

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	fmt.Print(viper.Get("mongo.addr"))

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
func watchConfig(f func()) {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("Config file changed:", e.Op)
		fmt.Println("Config file changed:", e.String())
		f()
	})
}
