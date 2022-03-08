package main

func main() {
	// assign1()
	// assign2()
	// assign3()
	// assign4()
	// assign5()
	// assign6()
	// assign7()
	// assign8()
	// assign9()
	// assign10()
	// assign11()
	// assign12()
	// assign13()
	// assign14()
	assign15()
}

func checkPrime(n int) int {
	var max int
	max = n
	for i := 2; i < max; i++ {
		if n%i == 0 {
			return 0
		}
		max = n / i
	}
	return 1
}
func splitBy(s string, n int) []string {
	var ss []string
	for i := 1; i < len(s); i++ {
		if i%n == 0 {
			ss = append(ss, s[:i])
			s = s[i:]
			i = 1
		}
	}
	ss = append(ss, s)
	return ss
}
