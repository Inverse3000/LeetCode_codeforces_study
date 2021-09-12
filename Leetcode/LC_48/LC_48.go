package main

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < len(matrix)/2-1; i++ {
		for j := i; j < len(matrix)-i; j++ {
			x := i
			y := i
			lst := matrix[x][y]
			for k := 0; k <= 3; k++ {
				if k == 0 {
					x = y
					y = n - i
				}
				if k == 1 {
					y = n - x
					x = n - i
				}
				if k == 2 {
					x = y
					y = i
				}
				if k == 3 {
					y = n - x
					x = i
				}
				matrix[x][y] = lst
				lst = matrix[x][y]
			}
		}
	}
}
func main() {

}
