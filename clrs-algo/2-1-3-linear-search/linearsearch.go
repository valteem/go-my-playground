// 2.1-3 Linear Search

package linearsearch

func LinearSearch(slc []int, lookupvalue int) (int, bool) {
	for i, v := range slc {
		if v == lookupvalue {
			return i, true
		}
	}
	return -1, false
}