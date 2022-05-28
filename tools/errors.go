package tools

import "fmt"

//HandleError handles errors
func HandleError(err error) {
	errorMessage := "Error: "
	switch err.Error() {
	case "EOF":
		errorMessage += "Connection closed by client"
	default:
		errorMessage += err.Error()
	}
	fmt.Println(errorMessage)
}
