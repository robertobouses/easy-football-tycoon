package app

import "math/rand"

func ProbabilisticIncrement14() int {
	if rand.Intn(7) == 4 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement20() int {
	if rand.Intn(5) == 4 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement25() int {
	if rand.Intn(4) == 3 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement33() int {
	if rand.Intn(3) == 2 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement40() int {
	if rand.Intn(5) < 2 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement44() int {
	if rand.Intn(25) < 11 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement50() int {
	if rand.Intn(2) == 0 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement57() int {
	if rand.Intn(7) < 4 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement66() int {
	if rand.Intn(3) < 2 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement62() int {
	if rand.Intn(8) < 5 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement71() int {
	if rand.Intn(10) < 7 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement75() int {
	if rand.Intn(4) < 3 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement80() int {
	if rand.Intn(5) < 4 {
		return 1
	}
	return 0
}

func ProbabilisticIncrement90() int {
	if rand.Intn(10) < 9 {
		return 1
	}
	return 0
}
func ProbabilisticIncrement94() int {
	if rand.Intn(50) < 47 {
		return 1
	}
	return 0
}
