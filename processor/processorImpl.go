package processor

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	log "github.com/golang/glog"
)

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
	// process message - not implemented yet
}