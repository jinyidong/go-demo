package channel

import "testing"

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
