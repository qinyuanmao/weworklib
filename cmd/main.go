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
	chatDatas, err := client.GetChatData(0, 100, "", "", 30)
	if err != nil {
		log.Fatalf("get chat data failed: %v", err)
	}
	for _, chatData := range chatDatas {
		mp, e := client.DecryptData(chatData.PublickeyVer, chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if e != nil {
			log.Fatal(e.Error())
		}
		fmt.Println(mp["msgid"])
		for key, value := range mp {
			fmt.Printf("%s: %v\n", key, value)
		}
		fmt.Println()
	}
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
