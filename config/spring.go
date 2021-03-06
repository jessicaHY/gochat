package config

import (
	"github.com/go-martini/martini"
	"log"
)

//type configType int
type configType int

const (
	Controller configType = iota
)

var configMap = make(map[configType]([]interface{}))

// Append a value to the key's slice
func AppendValue(key configType, value interface{}) {
	configMap[key] = append(configMap[key], value)
}

func GetSlice(key configType) []interface{} {
	return configMap[key]
}

func MappingController(m *martini.ClassicMartini) {
	log.Println("mapping controller")
	controllerMap := GetSlice(Controller)

	for _, ctn := range controllerMap {
		var a = ctn.(IController)
		a.SetRouter(m)
	}
}

type IController interface {
	SetRouter(m *martini.ClassicMartini)
}
