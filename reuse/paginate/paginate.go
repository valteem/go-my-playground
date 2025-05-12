package paginate

func Paginate[T any](items []T, pageStart, pageSize int) []T {

	start := (pageStart - 1) * pageSize
	end := start + pageSize

	if start >= len(items) {
		return []T{}
	}

	if end > len(items) {
		end = len(items)
	}

	return items[start:end]

}
