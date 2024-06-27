func findCenter(edges [][]int) int {
    
    a, b := edges[0][0], edges[0][1]
	if a == edges[1][0] || a == edges[1][1] {
		return a
	}
	return b
}       
