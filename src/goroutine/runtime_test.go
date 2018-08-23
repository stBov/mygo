package goroutine

import "testing"

func TestSay(t *testing.T){
	go say("world")
	say("hello")
}
