package main

// standarf 'pdqsort' partition, adopted for data slice of int 
func qsPartitionInt(data []int, a, b, pivot int) (newpivot int, alreadyPartitioned bool) {
	data[a], data[pivot] = data[pivot], data[a]
	i, j := a + 1, b - 1
	for i <= j && data[i] < data[a] {
		i++
	}
	for i <= j && !(data[j] < data[a]) {
		j--
	}

	if i > j {
		data[j], data[a] = data[a], data[j]
		return j, true
	}

	data[i], data[j] = data[j], data[i]
	i++
	j--

	for {
		for i <= j && data[i] < data[a] {
			i++
		}
		for i <= j && !(data[j] <data[a]) {
			j--
		}
		if i > j {
			break
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	data[a], data[j] = data[j], data[a]
	return j, false

}