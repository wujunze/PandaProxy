package password

import (
	"fmt"
	"testing"
)

func TestRandId(t *testing.T) {
	pwd :=RandPassword()
	fmt.Println(pwd)
	fmt.Println("Hello Panda")
}
