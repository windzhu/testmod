package testmod

import (
	"fmt"
)

func SayHello(name, str string) string {
	return fmt.Sprintf("您好, %s, %s", name, str)
}
