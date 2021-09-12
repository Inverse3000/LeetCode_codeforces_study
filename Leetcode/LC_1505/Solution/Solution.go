package main

type node struct {
	l, r, val int
	lo, ro    *node
}

func (o *node) insert(index, a int) {
	o.val += a
	if o.l == o.r {
		return
	}
	m := (o.l + o.r) >> 1
	if index <= m {
		if o.lo == nil {
			o.lo = &node{l: o.l, r: m, val: 0}
		}
		o.lo.insert(index, a)
	} else {
		if o.ro == nil {
			o.ro = &node{l: m + 1, r: o.r, val: 0}
		}
		o.ro.insert(index, a)
	}
}
func (o *node) query(l, r int) int {
	if o.l >= l && o.r <= r {
		return o.val
	}
	if o == nil || r < o.l || l > o.r {
		return 0
	}
	return o.lo.query(l, r) + o.ro.query(l, r)
}
func minInteger(num string, k int) string {
	o := &node{l: 0, r: 300005}
	s := []byte(num)
	f := make([][]int, 10)
	for i, v := range s {
		f[int(v)] = append(f[int(v)], i)
	}
	res := []byte{}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= 9; j++ {
			if len(f[j]) != 0 {
				d := f[j][0] - i + o.query(f[j][0], len(s)-1)
				if d <= k {
					res = append(res, s[f[j][0]])
					o.insert(f[j][0], 1)
					f[j] = f[j][1:]
					k -= d
					break
				}
			}
		}
	}
	return string(res)
}
func main() {

}
