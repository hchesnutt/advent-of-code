package utils

import "strconv"

func ParseInt(s string) int {
	var integer, err = strconv.Atoi(s)
	if err != nil {
		panic(s)
	}

	return integer
}
