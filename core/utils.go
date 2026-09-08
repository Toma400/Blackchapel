package core

import (
	// "gopkg.in/yaml.v3"
	// "https://github.com/goccy/go-yaml"
)

const TILE_SIZE = 32

// type Coord struct { // operates on px
//     X, Y int
// }

// type TileCoord struct { // operates on ~Tile
// 	X, Y int
// }

// func tileToPx(tc TileCoord) Coord { // converts TC to C, _will round the value_
// 	return Coord{tc.X * TILE_SIZE,
// 				 tc.Y * TILE_SIZE}
// }

// func SetOptions(scale int, movq Coord) *ebiten.DrawImageOptions {
//     opt := &ebiten.DrawImageOptions{}
//     if scale != 0            { opt.GeoM.Scale(float64(scale), float64(scale)) }
//     if !(movq == Coord{0,0}) { opt.GeoM.Translate(float64(movq.X),
//                                                   float64(movq.Y)) }
//     return opt
// }

// YAML structure for unmarshalling
// type MapLayerData struct {
// 	translation map[string]string
// 	coords      string
// }
// type InteriorMapData struct {
// 	BG MapLayerData
// 	// more layers to be implemented
// }

// func readInteriorMap(map_path string) {

// }
