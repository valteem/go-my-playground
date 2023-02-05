package reuse

func SliceOfSquares(inp []int32) []int32 {
	outp := []int32{}
	for _, v := range inp {
		outp = append(outp, v*v)
	}
	return outp
}