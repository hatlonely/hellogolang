package buildin

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect1(t *testing.T) {
	t1 := time.NewTicker(400 * time.Millisecond)
	t2 := time.NewTicker(300 * time.Millisecond)
	defer t1.Stop()

	end := time.After(3 * time.Second)
out:
	for {
		select {
		case <-end:
			fmt.Println("break")
			break out
		case <-t1.C:
			fmt.Println("hello world 1")
		case <-t2.C:
			fmt.Println("hello world 2")
		}
	}
}

func TestSelect2(t *testing.T) {
	fmt.Println(time.Now())
	// delay 3 second
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("after 3 second")
	}
	fmt.Println(time.Now())
}
