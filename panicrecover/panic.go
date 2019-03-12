package panicrecover

import "fmt"

func Panic()  {
	defer func() {
		if error:=recover();error!=nil {
			fmt.Print(error)
		}
	}()

	p()
}

func p()  {
	panic("hanic lll")
}
