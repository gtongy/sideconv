package error

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
