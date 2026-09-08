package core

import (
	"gopkg.in/yaml.v3"
	"fmt"
	"os"
)

type ObjectData struct {
	Tex string `yaml:"tex"`
	Col bool   `yaml:"col"`
}
type ObjectsTable struct {
	// try to not call explicitly, as funcs below allow key protection
	Definitions map[string]ObjectData
}

func RegisterObjects() ObjectsTable {
	fread, err := os.ReadFile("world/objects.yaml")
    if err != nil {
        panic("Couldn't read -objects.yaml- file! Check if it can be found in `world` directory.")
    }
    ret := ObjectsTable{}
    err = yaml.Unmarshal(fread, &ret.Definitions)
    if err != nil {
    	fmt.Println("Error unmarshalling the -objects.yaml- file!") // todo: put better error here one day
     	// default values below should handle broken values so that we can spare `panic()` here
    }
    return ret
}

func GetTexture(ot ObjectsTable, id string) string {
	val, ok := ot.Definitions[id]
	if ok {
		return val.Tex
	}
	return "" // default (no draw)
}

func GetCollision(ot ObjectsTable, id string) bool {
	if val, ok := ot.Definitions[id]; ok {
		return val.Col
	}
	return false // default (no collision)
}
