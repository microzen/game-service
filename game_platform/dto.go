package game_platform

import (
	"fmt"
	"strings"
)

type Response struct {
	message string
	status  int
}
type Request struct {
	message string
	status  int
}

func sendTo(user User, response Response) {
	res := fmt.Sprintf("-message=%v -status=%d", response.message, response.status)
	user.conn.Write([]byte(res))
}

func sentTo(user User, message string) {
	sendTo(user, Response{string(message), 200})
}

func GetValues(str string, names ...string) map[string]string {
	result := make(map[string]string)
	for _, name := range names {
		result[name] = ""
	}
	arguments := strings.Split(str, " ")
	for i := 0; i < len(arguments); i++ {
		if result[arguments[i]] == "" {
			result[arguments[i]] = arguments[i+1]
			i++
		}
	}
	return result
}
