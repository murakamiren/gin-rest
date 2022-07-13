package util

import "strconv"

func BoolCheck(isPublished string) bool {

	toInt, _ := strconv.Atoi(isPublished)

	if toInt == 0 {
		return false
	} else {
		return true
	}
}