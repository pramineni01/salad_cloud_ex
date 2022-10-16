package processor


type processor struct {
	source string
	port uint
}

type co2footprint struct {
	Header [3]byte
	TailNumberSize [4]byte // unsigned integer
	TailNumberValue string
	EngineCount [4]byte // unsigned integer
	EngineNameSize [4]byte // unsigned integer
	EngineNameValue string
	Latitude [8]byte	// IEEE-754 64-bit floating-point number
	Longitude [8]byte	// IEEE-754 64-bit floating-point number
	Altitude [8]byte	// IEEE-754 64-bit floating-point number
	Temperature [8]byte	// IEEE-754 64-bit floating-point number
}