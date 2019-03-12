package tdd

import (
	"github.com/prashantv/gostub"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	convey.Convey("1+1",t, func() {
		convey.So(Add(1,2),convey.ShouldEqual,3)
	})
}

func TestAdd2(t *testing.T) {
	convey.Convey("1+(-1)",t, func() {
		convey.So(Add(1,-1),convey.ShouldEqual,0)
	})
}

//stub使用时需要将方法改写为var funName=fun()
func TestAdd3(t *testing.T) {
	stubs:=gostub.StubFunc(&Add,1)//直接替代函数返回值
	defer stubs.Reset()

	convey.Convey("sub 2+1-1",t, func() {
		convey.So(Sub(2,1),convey.ShouldEqual,0)
	})
}