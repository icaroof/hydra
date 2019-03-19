package hydrachat

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func chatServerFunc(t *testing.T) func() {
	return func() {
		t.Log("Starting Hydra Chat Server...")

		if err := Run(":2300"); err != nil {
			t.Error("Could not start chat server", err)
			return
		}

		t.Log("Started Hydra Chat Server")
	}
}

func connectToServer(t *testing.T) (conn net.Conn) {
	conn, err := net.Dial("tcp", "127.0.0.1:2300")

	if err != nil {
		t.Fatal("Could not connect to Hydra chat system", err)
	}

	return
}

func TestRun(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	t.Log("Testing Hydra Chat sending and receiving")

	go once.Do(chatServerFunc(t))

	time.Sleep(1 * time.Second)

	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))

	t.Logf("Hello %s, connecting to the Hydra Chat System...\n", name)

	conn := connectToServer(t)

	t.Log("Connected")
	defer conn.Close()

	name += ": "

	msgCh := make(chan string)

	go func() {
		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			recvMsg := scanner.Text()
			sentMsg := <-msgCh

			if strings.Compare(recvMsg, sentMsg) != 0 {
				t.Errorf("Chat message %s does not match with %s", recvMsg, sentMsg)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		msgBody := fmt.Sprintf("Random Message %d", rand.Intn(400))
		msg := name + msgBody

		_, err := fmt.Fprintf(conn, msg+"\n")

		if err != nil {
			t.Error(err)
			return
		}

		msgCh <- msg
	}
}

func TestServerConnection(t *testing.T) {
	t.Log("Testing Hydra Chat Server connection")

	f := chatServerFunc(t)
	go once.Do(f)

	time.Sleep(1 * time.Second)

	conn := connectToServer(t)

	conn.Close()
}
