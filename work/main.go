package main

import "fmt"
import "math"

func main() {
	buff := []float64{}
	flag := true
	for flag {
		var input float64
		fmt.Println("数字を入力してください（終了の場合は-1を入力）：")
		fmt.Scan(&input)
		if input < 0 {
			fmt.Println("終了")
			break
		} else {
			buff = append(buff, input)
		}
	}

	fmt.Printf("平均値は%gです．\n", CalcMean(buff))
	fmt.Printf("標準偏差は%gです．\n", CalcSd(buff))
}

func CalcMean(buff []float64) float64 {
	var sum float64
	for _, x := range buff {
		sum += x
	}
	n := float64(len(buff))
	mean := sum / n
	return mean
}

func CalcSd(buff []float64) float64 {
	n := float64(len(buff))
	mean := CalcMean(buff)
	s := 0.0
	for _, x := range buff{
		s += (x - mean) * (x - mean)
	}
	var sd float64
	sd = math.Sqrt(s / (n-1))
	return sd
}
