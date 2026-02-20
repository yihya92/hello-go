package greet

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello %s:\n", name)
}
