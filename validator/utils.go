package validator

import (
	"fmt"
)

func toInt(value string) int {
	i := 0
	fmt.Sscanf(value, "%d", &i)
	return i
}
