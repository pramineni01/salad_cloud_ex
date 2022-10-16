package processor

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net"
	"time"

	log "github.com/golang/glog"
)

const (
	SizeofUint32 uint32 = 4
	SizeofFloat64 = 8
	
	min_msg_size_const = 31
	header_size_const = 3
)

var header_content  = []byte {0x41, 0x49 , 0x52}	// 01000001, 01001001, 01010010	

func (p *processor) Process(ctx context.Context) {
	// while infinite
	for true {
		c, err := p.connect(ctx)
		if err != nil {
			log.Fatal("TCP connection error: ", err)
		}

		defer c.Close()
		p.listenAndProcess(ctx, c)
	}
}

func (p *processor) connect(ctx context.Context) (*net.TCPConn, error) {
	// use source and port from p and connect tcp
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
		log.Warningf("Unable to set keepalive - %s", err)
		tcpConn.Close()
		return nil, err
	}
	
	// return the connection
	return tcpConn, nil
}

func (p *processor) listenAndProcess(ctx context.Context, c *net.TCPConn) {
	// listen on the connection indefinitely and return on message
	for true {
		select{
		case <-ctx.Done():
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
			} else {
				processMessage(ctx, buf)
			}
		}
	}
}

func processMessage(ctx context.Context, msg []byte) {
	if bytes.Compare(msg[:3], header_content) != 0 {
		log.Fatal("Message header mismatch")
	}
	fp := co2footprint{}

	// header
	copyBytes(fp.Header, msg[:3], header_size_const)
	
	curr_idx := uint32(header_size_const)

	// fetch TailNumber info
	fp.TailNumberSize, fp.TailNumberValue = extractTailNumberInfo(msg[curr_idx:])
	curr_idx = uint32(curr_idx) + SizeofUint32 + fp.TailNumberSize

	// fetch EngineCount, EngineNameSize, EngineNameValue
	fp.EngineCount, fp.EngineNameSize, fp.EngineNameValue = extractEngineInfo(msg[curr_idx:])
	curr_idx = curr_idx + SizeofUint32 + SizeofUint32 + fp.EngineNameSize

	// fetch Latitude
	fp.Latitude =  extractLatitude(msg[curr_idx:])
	curr_idx = curr_idx + SizeofFloat64

	// fetch Longitude
	fp.Longitude = extractLongitude(msg[curr_idx:])
	curr_idx = curr_idx + SizeofFloat64

	// fetch Altitude
	fp.Altitude = extractAltitude(msg[curr_idx:])
	curr_idx = curr_idx + SizeofFloat64

	// fetch Temperature
	fp.Temperature = extractTemperature(msg[curr_idx:])
	curr_idx = curr_idx + SizeofFloat64

	fmt.Printf("Received message:\n%v", fp)
}

func copyBytes(dst [3]byte, src []byte, size int) {
	if len(dst) != len(src) {
		log.Exit("Header size does not match")
	}
	for i:= 0; i < size; i++ {
		dst[i] = src[i]
	}
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
func extractEngineInfo(msg []byte) (uint32, uint32, string) {
	engCount := binary.BigEndian.Uint32(extractBytes(msg, 4))
	engNum := binary.BigEndian.Uint32(extractBytes(msg, 4))
	engNumVal := extractStringOfSize(msg[4:], engNum)
	return engCount, engNum, engNumVal
}

func extractLatitude(msg []byte) float64 {
	bits := binary.BigEndian.Uint64(extractBytes(msg, 8))
	return math.Float64frombits(bits)
}

func extractLongitude(msg []byte) float64 {
	bits := binary.BigEndian.Uint64(extractBytes(msg, 8))
	return math.Float64frombits(bits)
}

func extractAltitude(msg []byte) float64 {
	bits := binary.BigEndian.Uint64(extractBytes(msg, 8))
	return math.Float64frombits(bits)
}

func extractTemperature(msg []byte) float64 {
	bits := binary.BigEndian.Uint64(extractBytes(msg, 8))
	return math.Float64frombits(bits)
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
	if uint32(len(msg)) < size {
		log.Fatal("Message size small to fetch tail number value")
	}

	return string(msg[4:size])
}