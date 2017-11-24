package main

import "time"

func main() {
	t := time.NewTimer(1 * time.Second)
	//for {
	time.Sleep(time.Second)
	select {
	case <-t.C:
		print("cc")
	default:
		t.Stop()
	}
	//}
	print("end")
}
