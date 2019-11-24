package mq

import (
	"fmt"
	"testing"
)

func TestStartConsume(t *testing.T) {
	defer StopConsume()
	Init()
	StartConsume(TransOSSQueueName, "transfer_oss", func(msg []byte) bool {
		fmt.Println(string(msg))
		return true
	})
}
