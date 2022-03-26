package go_channel

import (
	"fmt"
	"testing"
	"time"
)

func TestCloseChannel(t *testing.T) {
	channel := make(chan time.Time, 100)
	closed := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-closed:
				t.Log("sender quit")
				close(channel)
				return
			default:
				channel <- time.Now()
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			now := <-channel
			fmt.Println(i, now.Format("2006-01-02 15:04:05"))
			if i == 50 {
				closed <- true

				return
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
