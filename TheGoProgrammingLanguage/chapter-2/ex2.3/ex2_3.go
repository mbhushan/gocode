package popcount

// pc[i] is the population count of i.
var pcount [256]byte

func init() {
	for i := range pcount {
		pcount[i] = pcount[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pcount[byte(x>>(0*8))] +
		pcount[byte(x>>(1*8))] +
		pcount[byte(x>>(2*8))] +
		pcount[byte(x>>(3*8))] +
		pcount[byte(x>>(4*8))] +
		pcount[byte(x>>(5*8))] +
		pcount[byte(x>>(6*8))] +
		pcount[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var sum byte
	for i := uint(0); i < 8; i++ {
		sum += pcount[byte(x>>(i*8))]
	}
	return int(sum)
}
