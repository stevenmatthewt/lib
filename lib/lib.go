package lib_test

import (
	"fmt"

	"github.com/h2non/gock"
)

func Print() {
	gock.Off()
	fmt.Println("Hello")
}
