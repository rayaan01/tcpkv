package internal

import (
	"fmt"
	"os"
)

func HandleError(msg string, e error) {
	errMsg := fmt.Sprintf("ERROR: %s \n CAUSE: %s \n", msg, e.Error())
	fmt.Println(errMsg)
	os.Exit(1)
}
