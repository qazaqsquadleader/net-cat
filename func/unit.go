package cat

import (
	"bufio"
	"cat/model"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	MessageChannel = make(chan model.Message)
	Clients        = make(map[net.Conn]string)
	mu             sync.Mutex
)

func Run(l net.Listener) error {
	history, err := os.Create("../files/history.txt")
	if err != nil {
		return err
	}
	defer func() error {
		if err = os.Remove("../files/history.txt"); err != nil {
			return err
		}
		return nil
	}()
	go SendMessageAll()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error in connect: %v", err)
			continue
		}
		go Handle(conn, history)
	}
}

func Handle(conn net.Conn, history *os.File) {
	defer conn.Close()
	fmt.Fprintln(conn, "Welcome to TCP-chat!")
	logo, err := Logotype()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(conn, logo)

	scan := bufio.NewScanner(conn)
	var name string

	for {
		fmt.Fprintf(conn, "[ENTER YOUR NAME]:")
		if scan.Scan() {
			name = scan.Text()
		}
		if err = CheckConnection(conn, name, &mu, Clients); err != nil {
			fmt.Fprintln(conn, fmt.Sprint(err)+"Please thry again.")
			continue
		}
		mu.Lock()
		Clients[conn] = name
		mu.Unlock()
		break
	}

	fileData, err := os.ReadFile(history.Name())
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(conn, string(fileData))
	fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-1-2 15:4:5"), name)
	MessageChannel <- model.NewMessage(name, fmt.Sprintf("%s has joined our chat...", name))

	for scan.Scan() {
		fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-1-2 15:4:5"), name)
		if len(strings.TrimSpace(scan.Text())) == 0 {
			continue
		}
		text := strings.TrimSpace(scan.Text())
		MessageChannel <- model.NewMessage(name, fmt.Sprintf("[%s][%s]:%s", time.Now().Format("2006-1-2 15:4:5"), name, text))
		history.WriteString(fmt.Sprintf("[%s][%s]:%s\n", time.Now().Format("2006-1-2 15:4:5"), name, text))
	}
	MessageChannel <- model.NewMessage(name, fmt.Sprintf("%s has left our chat...", name))
	mu.Lock()
	delete(Clients, conn)
	mu.Unlock()
}

func SendMessageAll() {
	for {
		select {
		case msg := <-MessageChannel:
			mu.Lock()
			for conn, name := range Clients {
				if name == msg.Name {
					continue
				}
				timeN := time.Now().Format("2006-1-2 15:4:5")
				fmt.Fprintf(conn, "\n%s\n", msg.Text)
				fmt.Fprintf(conn, "[%s][%s]:", timeN, name)
			}
			mu.Unlock()
		}
	}
}
