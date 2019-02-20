package interceptor

import (
	"fmt"
	"testing"
)

func init()  {
	New= func(id int,name string) IModel {
		p:=Proxy{
			&Account{id,name},
			"ddddd",
		}
		return &p
	}
}
func TestAccount_Query(t *testing.T) {
	a:=New(1,"金义冬")
	fmt.Println(a.Query(1))
}
