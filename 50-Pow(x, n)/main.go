package main

func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	half := myPow(x, n/2)
	if n%2 == 0 {
		return half * half
	}

	if n%2 == 1 {
		return half * half * x
	}

	return myPow(x, n)
}
