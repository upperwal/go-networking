package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("udp")

func reader(c *net.UDPConn) {
	buffer := make([]byte, 1000)
	for {
		i, err := c.Read(buffer)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("Recent read: ", string(buffer[:i]))
	}
}

func main() {

	logging.SetLogLevel("udp", "DEBUG")

	p := flag.String("p", "0", "port to listen on")
	flag.Parse()

	laddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:"+*p)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		panic(err)
	}

	go reader(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter UDP address: ")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")

		raddr, err := net.ResolveUDPAddr("udp4", text)
		if err != nil {
			panic(err)
		}

		for i := 0; i < 5; i++ {
			_, err = conn.WriteToUDP([]byte("hello"), raddr)
			if err != nil {
				log.Error(err)
			}
			time.Sleep(1 * time.Second)
		}
	}

}
