package shop

import (
	"fmt"

	guuid "github.com/google/uuid"
)

func Newline(num int) {
	for num > 0 {
		fmt.Println()
		num--
	}
}

func GenUUID() string {
	id := guuid.New()
	Id := fmt.Sprintf("%v", id)
	return Id[:10]
}

func GenLongUUID() string {
	id := guuid.New()
	Id := fmt.Sprintf("%v", id)
	return Id[:16]
}
