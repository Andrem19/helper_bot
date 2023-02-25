package convert

import (
	"fmt"
	"strconv"
)

func Conv(message []string) (string, error) {
	num, err := strconv.ParseFloat(message[1], 32)
		if err != nil {
			return "", err
		}

	switch message[2] {
	case "ftcm":
		res := num * 30.48
		return fmt.Sprintf("%f", res), nil
	case "cmft":
		res := num / 30.48
		return fmt.Sprintf("%f", res), nil
	case "cmin":
		res := num / 2.54
		return fmt.Sprintf("%f", res), nil
	case "incm":
		res := num * 2.54
		return fmt.Sprintf("%f", res), nil
	default:
		return "unknown request", nil
	}
}
