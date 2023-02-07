package cat

import (
	"errors"
	"net"
	"sync"
)

func CheckConnection(conn net.Conn, name string, mu *sync.Mutex, Client map[net.Conn]string) error {
	mu.Lock()
	defer mu.Unlock()
	if len(name) == 0 {
		return errors.New("The name is empty!")
	}
	for _, nameInCliets := range Clients {
		if nameInCliets == name {
			return errors.New("This name is already!")
		}
	}
	if len(Clients) > 9 {
		return errors.New("There are 10 clients in the chat room.")
	}
	for _, value := range name {
		if !(value >= 'a' && value <= 'z') && !(value >= 'A' && value <= 'Z') {
			return errors.New("The name haven`t a valid value! Type only letters.")
		}
	}
	return nil
}
