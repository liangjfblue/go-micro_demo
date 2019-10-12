package config

import (
	"errors"
	"log"

	"github.com/micro/go-micro/config"
)

func InitConfig() []string {
	var (
		nodesAddr = make([]string, 0)
	)

	if err := config.LoadFile("./config/config.json"); err != nil {
		log.Fatalln(err)
		return nil
	}
	conf := config.Map()
	for _, value := range conf["nodes"].(map[string]interface{}) {
		for _, value2 := range value.(map[string]interface{}) {
			nodesAddr = append(nodesAddr, value2.(string))
		}
	}
	if len(nodesAddr) <= 0 {
		panic(errors.New("nodesAddr empty"))
	}

	return nodesAddr
}
