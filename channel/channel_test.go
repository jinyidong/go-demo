package channel

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"testing"
)

func TestChannelSelect2(t *testing.T) {
	fmt.Println(1<<16)//将1的二进制位左移16位，右侧补0
}
func TestChannelSelect(t *testing.T) {
	ChannelSelect()
}
func TestDoubleCloseChan(t *testing.T) {
	DoubleCloseChan()
}

func TestWriteCloseChan(t *testing.T) {
	WriteCloseChan()
}

func TestReadCloseChanWithNoBuffer(t *testing.T) {
	ReadCloseChanWithNoBuffer()
}

func TestTwoConsumer(t *testing.T) {
	TwoConsumer()
}

func TestChannelSelect3(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		fmt.Println(i)
		v = append(v, i)
	}
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}

