package list

func Contain(slices *[]string, element string) bool {
	if len(*slices) == 0 {
		return false
	}
	for _, v := range *slices {
		if v == element {
			return true
		}
	}
	return false
}
