package game_platform

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

func Init(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()
	log.Printf("game service listening on port %d", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept: %v", err)
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("failed to read: %v", err)
		return
	}
	key := GetValues(message, "username")
	user := NewUser(conn, key["username"])
	defer func() {
		DeleteRoom(user.HashCode())
	}()
	sentTo(*user, string(user.HashCode()))

	for {
		message, err = reader.ReadString(0)
		if err != nil {
			log.Printf("failed to read: %v", err)
			continue
		}
		if len(message) > 4 && message[1:4] == "key" {
			key := GetValues(message, "key", "roomname", "size", "roomid")
			if key["key"] == "create" {
				name := key["roomname"]
				size, err := strconv.ParseInt(key["size"], 10, 32)
				if err != nil {
					log.Printf("failed to parse size: %v", err)
					continue
				}
				id := NewRoom(user, name, int(size))
				sentTo(*user, string(id))
			} else if key["key"] == "join" {
				roomID := key["roomid"]
				id, err := strconv.Atoi(string(roomID))
				if err != nil {
					log.Printf("failed to parse room id: %v", err)
					continue
				}
				room := GetRoom(uint32(id))
				room.AddUser(user)
				sentTo(*user, roomID)
			}
		}
	}
}
