package functions

import "fmt"

func HandlerError(err error, msg string) {

	if err != nil {

		panic(fmt.Sprintf("Error: %s [%s]", msg, err))
	}
}
