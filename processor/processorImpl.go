package processor

import (
	"context"
	"fmt"
	"log"
)

const (
	float64_len = 4
	uint32_len = 4

	min_msg_size_const = 31
	header_size_const = 3
	header_content = []byte{65, 73, 82}	// 01000001, 01001001, 01010010	
)

func (p *processor) Process(ctx context.Context) {
	// while infinite
	for true {
		c, err := p.connect(ctx)
		if err != nil {
			log.Fatal("TCP connection error: ", err)
		}

		msg, err := p.listen(ctx, c)
		c.Close()
	}
}

func (p *processor) connect(ctx contex.Context) (*TCPConn, error) {
	// use source and port from p and connect tcp
	if (ctx.Done()) {
		return errors.New("User interrupt")
	}
	
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", p.source, p.port))
	if err != nil {
		fmt.Printf("Unable to resolve IP")
	}

	// establish tcp connection, set keepAlive
	tcpConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal("TCP connection error. Error: ", err)
		return nil, err
	}

	err = tcpConn.SetKeepAlive(true)
	if err != nil {
		log.Warn("Unable to set keepalive - %s", err)	
		tcpConn.Close()
		return nil, err
	}
	
	// return the connection
	return tcpConn, nil
}

func (p *processor) listen(ctx contex.Context, c *TCPConn) ([]byte, err){
	// listen on the connection indefinitely and return on message
	status, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return nil, err
	}

	for true {
		select{
		case ctx.Done:
			break
		default:
			buf, err := io.ReadAll(c)
			if err != nil {
				log.Fatal("ReadAll returned error. Error: ", err)
			} 
			if len(buf) == 0 {
				log.Fatal("ReadAll returned empty buffer")
				time.Sleep(1000)
				continue
			}
			if len(buf) < min_msg_size_const {
				log.Fatal("Message length is smaller than expected")
			} else {
				processMessage(ctx, buf)
			}
		}
	}
}

func readToBuffer() ([]byte, err) {
	io.ReadAll()
	return nil, err
}

func processMessage(ctx context.Context, msg []byte) {
	if bytes.Compare(msg[:3], header_content) != 0 {
		log.Fatal("Message header mismatch")
	}
	fp := co2footprint{}

	// header
	n, err := io.CopyN(fp.Header, msg, header_size_const)
	if (err != nil) || (n != header_size_const) {
		log.Fatal("Error copying header size")
	}
	
	curr_idx := header_size_const

	// fetch TailNumber info
	fp.TailNumberSize, fp.TailNumberValue = extractTailNumberInfo(msg[curr_idx:])
	curr_idx = curr_idx + unsafe.SizeOf(uint32) + fp.TailNumberSize

	// fetch EngineCount, EngineNameSize, EngineNameValue
	fp.EngineCount, fp.EngineNameSize, fp.EngineNameValue = extractEngineInfo(msg[curr_idx:])
	curr_idx = curr_idx + unsafe.SizeOf(uint32) + unsafe.SizeOf(uint32) + fp.EngineNameSize

	// fetch Latitude
	fp.Latitude =  extractLatitude(msg[curr_idx:])
	curr_idx = curr_idx + unsafe.SizeOf(float64)

	// fetch Longitude
	fp.Longitude := extractLongitude(msg[curr_idx:])
	curr_idx = curr_idx + unsafe.SizeOf(float64)

	// fetch Altitude
	fp.Altitude := extractAltitude(msg[curr_idx:])
	curr_idx = curr_idx + unsafe.SizeOf(float64)

	// fetch Temperature
	fp.Temperature := extractTemperature(msg[curr_idx:])
	curr_idx = curr_idx + unsafe.SizeOf(float64)

	fmt.Printf("Received message:\n%v", fp)
}


// returns tail number and value
// panics if an issue with message
func extractTailNumberInfo(msg []byte) (uint32, string) {
	tailNum := binary.BigEndian.Uint32(extractBytes(msg, 4))
	tailNumVal := extractStringOfSize(msg[4:], tailNum)
	return tailNum, tailNumVal
}

// returns engine number and value
// panics if an issue with message
func extractEngineInfo(msg []byte) (uint32, string) {
	engNum := binary.BigEndian.Uint32(extractBytes(msg, 4))
	engNumVal := extractStringOfSize(msg[4:], engNum)
	return engNum, engNumVal
}

func extractLatitude(msg []byte) float64 {
	tmp := extractBytes(msg, 8)
	return math.Float64frombits(byte)
}

func extractLongitude(msg []byte) float64 {
	tmp := extractBytes(msg, 8)
	return math.Float64frombits(byte)
}

func extractAltitude(msg []byte) float64 {
	tmp := extractBytes(msg, 8)
	return math.Float64frombits(byte)
}

func extractTemperature(msg []byte) float64 {
	tmp := extractBytes(msg, 8)
	return math.Float64frombits(byte)
}

// extracts bytes, given slice and size required
func extractBytes(msg []byte, size int) []byte {
	if len(msg) < size {
		log.Fatal("Invalid message size")
	}
	return msg[:size]
}

// extracts bytes into string, given slice and size required
func extractStringOfSize(msg []byte, size uint32) string {
	if len(msg) < size {
		log.Fatal("Message size small to fetch tail number value")
	}

	return string(msg[4:size])
}