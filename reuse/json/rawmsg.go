package json

import (
	"encoding/json"
)

const (
	Mounts    TerrainType = "mounts"
	Highlands TerrainType = "highlands"
	Lowlands  TerrainType = "lowlands"
	Mixed     TerrainType = "mixed"
	Other     TerrainType = "other"
)

type TerrainType string

type LandPlot struct {
	Location string      `json:"location"`
	Area     float32     `json:"area"`
	Terrain  TerrainType `json:"terrain"`
}

type RawData struct {
	LP json.RawMessage `json:"landplot"`
}
