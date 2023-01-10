package shop

import (
	"fmt"

	guuid "github.com/google/uuid"
)

// This function helps in creating new lines
func Newline(num int) {
	for num > 0 {
		fmt.Println()
		num--
	}
}

// This function is used to generate a short ID
func GenUUID() string {
	id := guuid.New()
	Id := fmt.Sprintf("%v", id)
	return Id[:6]
}
