package util

// ExistsInStringArray .
func ExistsInStringArray(s string, array []string) (exist bool) {
	for _, item := range array {
		exist = s == item
		if exist {
			return
		}
	}
	return
}
