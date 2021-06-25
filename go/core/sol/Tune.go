package sol

type CoverageTune = struct {
	Markets  []uint32
	VolLV    string
	CapLV    string
	VolValue float64
	CapValue float64
}
type ImpulseTune = struct {
	LV     string
	K      int
	Period int
	Rate   float64
}
type CorridorTune = struct {
	LV     string
	K      int
	Period int
	Up     float64
	Down   float64
}

type UserTune = struct {
	Author      int
	ID          string
	Time        int
	Name        string
	Active      bool
	Description string
	Coverage    CoverageTune
	Impulse     ImpulseTune
	Corridor    CorridorTune
}
