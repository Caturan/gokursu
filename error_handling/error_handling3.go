package errorhandling

import "fmt"

type BorderException struct {
	parameter int
	message   string
}

func (b *BorderException) Error() string { // If any struct we implement it can do error task
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func Quess1(quess int) (string, error) {
	if quess < 1 || quess > 100 {
		return "", &BorderException{quess, "Except the border"}
	}
	return "Correct", nil
}
