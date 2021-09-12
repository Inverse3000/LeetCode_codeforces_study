package main

func myPow(x float64, n int) float64 {
	var r float64
	r = 1.0
	y := x
	for n > 0 {
		if n&1 == 1 {
			r = r * y
		}
		y *= y
		n = n >> 1
	}
	return r
}
func main() {

}
