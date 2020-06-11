package main

import "fmt"
import "math"

func main() {
	buff := [][]float64{}
	flag := true
	for flag {
		var input float64
		fmt.Println("数字を入力してください（終了の場合は-1を入力）：")
		fmt.Scan(&input)
		if input < 0 {
			fmt.Println("終了")
			break
		} else {
			s2 := [][]float64{{input,0}}
			buff = append(buff, s2...)
		}
	}

	for i, _ := range buff {
		var input float64
		fmt.Println("数字を入力してください（終了の場合は-1を入力）：")
		fmt.Scan(&input)
		buff[i][1] = input
		// fmt.Printf("%f\n", buff[i])
	}

	x := []float64{}
	y := []float64{}
	for _, v := range buff {
		x = append(x, v[0])
		y = append(y, v[1])
	}
	fmt.Printf("%f\n",x)
	fmt.Printf("%f\n",y)	
}