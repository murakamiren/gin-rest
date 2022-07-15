package util

import "strconv"

func BoolCheck(intBool string) bool {

	toInt, _ := strconv.Atoi(intBool)

	if toInt == 0 {
		return false
	} else {
		return true
	}
}