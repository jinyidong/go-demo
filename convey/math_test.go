package convey

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)


func TestFib0(t *testing.T) {
	convey.Convey("FibTest 0=0",t, func() {
		convey.So(Fib(0),convey.ShouldEqual,0)
	})
}

func TestFib1(t *testing.T) {
	convey.Convey("FibTest 1=1",t, func() {
		convey.So(Fib(1),convey.ShouldEqual,1)
	})
}

func TestFib2(t *testing.T) {
	convey.Convey("FibTest 0=0",t, func() {
		convey.So(Fib(2),convey.ShouldEqual,1)
	})
}

func TestFib3(t *testing.T) {
	convey.Convey("FibTest 0=0",t, func() {
		convey.So(Fib(3),convey.ShouldEqual,2)
	})
}


func TestFib8(t *testing.T) {
	convey.Convey("FibTest 0=0",t, func() {
		convey.So(Fib(8),convey.ShouldEqual,21)
	})
}

func TestFib80(t *testing.T) {
	convey.Convey("FibTest 0=0",t, func() {
		convey.So(Fib(80),convey.ShouldEqual,23416728348467685)
	})
}

