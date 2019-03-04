package injector

import (
	"fmt"
	"github.com/codegangsta/inject"
	"reflect"
	"testing"
)

type Stu struct {
	Id int
}

func TestChannelSync(t *testing.T) {
	inj:=inject.New()
	stu:=Stu{
		Id:1,
	}
	inj.Map(stu)

	fmt.Println(inj.Get(reflect.TypeOf(Stu{})))
}