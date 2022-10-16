package processor


type processor struct {
	source string
	port uint
}

type co2footprint struct {
	Header [3]byte
	TailNumberSize uint32
	TailNumberValue string
	EngineCount uint32
	EngineNameSize uint32
	EngineNameValue string
	Latitude float64
	Longitude float64
	Altitude float64
	Temperature float64
}
