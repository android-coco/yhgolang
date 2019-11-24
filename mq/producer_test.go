package mq

import (
	"fmt"
	"testing"
)

func TestPublish(t *testing.T) {
	var msg = "hello world"
	publish := Publish(TransExchangeName, TransOSSRoutingKey, []byte(msg))
	fmt.Println(publish)
}
