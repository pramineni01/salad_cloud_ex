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

// type co2footprint struct {
// 	Header [3]byte
// 	TailNumberSize [4]byte // unsigned integer
// 	TailNumberValue string
// 	EngineCount [4]byte // unsigned integer
// 	EngineNameSize [4]byte // unsigned integer
// 	EngineNameValue string
// 	Latitude [4]byte	// IEEE-754 64-bit floating-point number
// 	Longitude [4]byte	// IEEE-754 64-bit floating-point number
// 	Altitude [4]byte	// IEEE-754 64-bit floating-point number
// 	Temperature [4]byte	// IEEE-754 64-bit floating-point number
// }