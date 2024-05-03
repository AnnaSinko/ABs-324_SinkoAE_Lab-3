package main

import (
	"fmt"
	"math"
	"os"
)

func firstFragm(x float64) float64 {
	res := -0.3333333333333333 * x - 0.6666666666666667
	if math.Abs(res) < 0.000001 {
		res = 0.0
	}
	return res
}

func secondFragm(x float64) float64 {
	res := math.Log(math.Abs(1 / math.Tan(x/2)))
	return res
}

func thirdFragm(x float64) float64 {
	res := math.Log(math.Abs(1 / math.Tan(x/2)))
	return res
}

func fourthFragm(x float64) float64 {
	res := 0.3333333333333333 * x - 0.6666666666666667
	if math.Abs(res) < 0.000001 {
		res = 0.0
	}
	return res
}

func main() {
	xStart := -5.0
	xEnd := 5.0
	dx := 0.5

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("При открытии файла произошла ошибка")
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "%10s %10s\n", "x", "y")
	for x := xStart; x <= xEnd; x += dx {
		if math.Abs(x) < 0.000001 {
			x = 0.0
		}

		if x >= -5.0 && x <= -2.0 {
			fmt.Fprintf(file, "%10.1f %10.1f\n", x, firstFragm(x))
		}
		if x > -2.0 && x < 0.0 {
			fmt.Fprintf(file, "%10.1f %10.1f\n", x, secondFragm(x))
		}
		if x > 0.0 && x < 2.0 {
			fmt.Fprintf(file, "%10.1f %10.1f\n", x, thirdFragm(x))
		}
		if x >= 2.0 && x <= 5.0 {
			fmt.Fprintf(file, "%10.1f %10.1f\n", x, fourthFragm(x))
		}
	}
}