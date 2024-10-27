package game_platform

import (
	"sync"
)

type Room struct {
	size     int
	capacity int
	Name     string
	Host     *User
	Users    []*User
}

var rooms = make(map[uint32]*Room)
var clientMutex = sync.Mutex{}

func NewRoom(host *User, name string, size int) uint32 {
	room := Room{}
	room.size = 0
	room.capacity = size
	room.Name = name
	room.Host = host
	room.Users = make([]*User, size)
	clientMutex.Lock()
	rooms[host.HashCode()] = &room
	clientMutex.Unlock()
	room.AddUser(host)
	return host.HashCode()
}

func (r *Room) AddUser(user *User) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	if r.size < r.capacity {
		for i, u := range r.Users {
			if u == nil {
				r.Users[i] = user
				r.size++
				break
			}
		}
	}
}
func (r *Room) RemoveUser(user *User) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	for i, u := range r.Users {
		if u == user {
			r.Users[i] = nil
			r.size--
			break
		}
	}
}

func GetRoom(id uint32) *Room {
	if rooms[id] != nil {
		return rooms[id]
	} else {
		return nil
	}
}

func DeleteRoom(id uint32) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	if rooms[id] != nil {
		delete(rooms, id)
	}
}
