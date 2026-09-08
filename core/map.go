package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gopkg.in/yaml.v3"
	// "fmt"
	"os"
)

type BGTile struct { // -bg- layer tiles (background)
	tex string // texture path
}
// type ObjTile struct { // -obj- layer tiles (static objects/activators)
// 	// col bool   // collision (tile inherits it upwards)
// }
// type ItTile struct { // -it- layer tiles (items)

// }
// type SpecTile struct { // -spec- layer tiles (specials)

// }

type Tile struct {
 	bgt BGTile
// 	obt ObjTile
// 	itt ItTile
// 	spt SpecTile
// 	// unique data
	col bool     // (summed up collisions)
}

type LayerData struct {
	Translation map[string]string `yaml:"translation"`
	Coordinates string            `yaml:"map"`
}

type InteriorMap struct {
	Name    string    `yaml:"name"`
	BG_Data LayerData `yaml:"bg"`
	// non-YAML data calculated later ('-' skips YAML parsing)
	Tiles   [][]Tile `yaml:"-"`
	// todo: check if sizes are equal (X/Y) between subtiles
}

func generateTiles(im *InteriorMap, ot ObjectsTable) {
	var line []Tile
	for _, char := range im.BG_Data.Coordinates {
		if char == '\n' {
			im.Tiles = append(im.Tiles, line)
			line     = nil // clears out the collection for going down
			continue
		}
		val, ok := im.BG_Data.Translation[string(char)]
		if !ok {
			panic("Tile " + string(char) + " not found in interior map" + im.Name + "definitions!")
		}
		t := Tile{
			bgt: BGTile{
				tex: GetTexture(ot, val),
			},
			col: GetCollision(ot, val), // OR other collisions
		}
		line = append(line, t)
	}
}

func ReadInteriorMap(name string, ot ObjectsTable) InteriorMap {
	fread, err := os.ReadFile("world/interiors/" + name + ".yaml")
	if err != nil {
		panic("Couldn't read -" + name + "- interior map! Check if the map exists.")
	}
	ret := InteriorMap{}
    err = yaml.Unmarshal(fread, &ret)
    if err != nil {
    	panic("Could't properly read -" + name + "- interior map! Check the file for syntax errors or missing fields.")
    }
    generateTiles(&ret, ot)
    return ret
}

func DrawInteriorMap(im InteriorMap, screen *ebiten.Image) {
	for _, line := range im.Tiles {
		for _, tile := range line {
			if tile.bgt.tex != "" { // skip if transparent/no texture
				screen.DrawImage()
				// https://github.com/Toma400/Ilmalaiva/blob/sky/core/table.go#L132
				// https://github.com/Toma400/Ilmalaiva/blob/sky/core/cae.go#L37
			}
		}
	}
}
