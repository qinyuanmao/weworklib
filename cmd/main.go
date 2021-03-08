package main

import (
	"fmt"
	"log"

	"github.com/qinyuanmao/weworklib"
	"github.com/spf13/viper"
)

func main() {
	loadConfig("/workspaces/weworklib/config")
	client, err := weworklib.NewClient(
		viper.GetString("enterprise_wechat.crop_id"),
		viper.GetString("enterprise_wechat.crop_secret"),
		map[uint32]string{
			1: weworklib.ReadFile(viper.GetStringSlice("enterprise_wechat.private_pems")[0]),
			2: weworklib.ReadFile(viper.GetStringSlice("enterprise_wechat.private_pems")[1]),
		})
	if err != nil {
		log.Fatalf("init wework client failed: %v", err)
	}
	defer client.Free()
	messages, err := client.GetChatList(0, 100, "", "", 30)
	if err != nil {
		log.Fatalf("获取消息内容异常: %v", err)
	}
	fmt.Println(messages)
}

func loadConfig(path string) (err error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return
}
