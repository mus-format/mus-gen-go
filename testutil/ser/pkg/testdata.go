package another

import (
	"fmt"

	com "github.com/mus-format/common-go"
)

const MyAwesomeIntDTM com.DTM = iota + 1

type MyInt int

func (i MyInt) Print() {
	fmt.Println("MyInt")
}
