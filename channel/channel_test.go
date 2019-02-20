package channel

import (
	"fmt"
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
