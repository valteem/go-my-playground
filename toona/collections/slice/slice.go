package slice

func Contains[V comparable](s []V, value V) bool {
	if len(s) > 0 {
		for _, v := range s {
			if v == value {
				return true
			}
		}
		return false
	} else {
		return false
	}
}