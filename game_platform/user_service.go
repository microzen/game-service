package game_platform

import (
	"hash/fnv"
	"net"
)

type User struct {
	Name string
	conn net.Conn
}

func NewUser(conn net.Conn, userName string) *User {
	return &User{
		conn: conn,
		Name: userName,
	}
}

func (u *User) HashCode() uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(u.Name))
	hash.Write([]byte(u.conn.RemoteAddr().String()))
	return hash.Sum32()
}
