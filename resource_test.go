package resource

import (
	"fmt"
	"testing"
)

func TestURLHash(t *testing.T) {
	file, err := Download("http://code.google.com/p/go/source/browse/src/pkg/net/http/request_test.go?spec=svn0dac18e695f3fdb448883cab04e962d876620fab&name=0dac18e695f3&r=0dac18e695f3fdb448883cab04e962d876620fab")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", file.Name())
}
