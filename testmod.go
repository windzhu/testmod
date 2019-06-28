package testmod

import (
	"fmt"
)

func SayHello(name string) string {
	return fmt.Sprintf("您好, %s", name)
}
