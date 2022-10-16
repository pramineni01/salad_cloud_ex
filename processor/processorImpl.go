package processor

import (
	"context"
	"fmt"
	"log"
)

func (p *processor) Process(ctx context.Context) {
	// while infinite
	for true {

			c, err := p.connect(ctx)
			if err != nil {
				log.Fatal("TCP connection error: ", err)
			}

			defer c.Close()
			msg, err := p.listen(ctx, c)
			if err != nil {
			processMessage(msg)
		}
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
	// process message
}