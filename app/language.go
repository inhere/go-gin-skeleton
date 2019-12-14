package app

import (
	"fmt"
	"strings"

	"github.com/gookit/i18n"
)

func initLanguage() {
	// conf := map[string]string{
	// 	"langDir": "resource/lang",
	// 	"allowed": "en:English|zh-CN:简体中文",
	// 	"default": "en",
	// }
	conf := Config.StringMap("lang")
	fmt.Printf("language - %v\n", conf)

	// en:English|zh-CN:简体中文
	langList := strings.Split(conf["allowed"], "|")
	languages := make(map[string]string, len(langList))

	for _, str := range langList {
		item := strings.Split(str, ":")
		languages[item[0]] = item[1]
	}

	// init and load data
	i18n.Init(conf["langDir"], conf["default"], languages)

	I18n = i18n.Default()
}
