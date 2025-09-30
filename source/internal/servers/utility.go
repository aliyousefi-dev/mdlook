package servers

import (
	"strconv"
)

func IncrementPort(port string) string {
	portInt, _ := strconv.Atoi(port)

	portInt++

	// Convert the incremented port back to a string
	incPort := strconv.Itoa(portInt)
	return incPort
}
