package handler

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("bio")
	if result != "Hello bam" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorld Done")
}
