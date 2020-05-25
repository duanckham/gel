package main

import (
	"fmt"
	"os"
	"time"

	"github.com/duanckham/gel/agent"
	"github.com/duanckham/gel/gel"
	"github.com/duanckham/gel/server"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "agent" {
		fmt.Println("* running agent...")
		go runAgent()
	} else {
		fmt.Println("* running server...")
		go runServer()
	}

	select {}
}

func runServer() {
	server.New(5024)
}

func runAgent() {
	g := agent.New("127.0.0.1", 5024, time.Duration(5)*time.Second)

	testing(g)

	select {}
}

func testing(g gel.Gel) {
	go func() {
		fmt.Println("p1...")

		strA := "this is a log containing variable ?? and variable ?? and something"
		strB := "this is a log containing variable ?? and variable ?? and ??"
		strC := "this is a log containing variable ?? and variable ?? and ?? and ??"

		for {
			g.Log(strA, "test_1", 1)
			g.Log(strB, "test_2", 1, 2)
			g.Log(strC, "test_3", 1, 2, 3)
			g.Log(strC, "test_4", 1, 2, 3)

			time.Sleep(time.Duration(5) * time.Millisecond)
		}
	}()

	go func() {
		fmt.Println("p2...")

		for {
			g.Increment("dc", 1)

			time.Sleep(time.Duration(5) * time.Millisecond)
		}
	}()

	go func() {
		fmt.Println("p3...")

		for {
			g.Gauge("dc", 5.0)

			time.Sleep(time.Duration(5) * time.Millisecond)
		}
	}()

}
