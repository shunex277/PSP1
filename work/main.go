// Program Assignment: program3
// Name: Inoue Shunei
// Date: 2020/06/9
// Description: this is a program that calculate the linear regression parameters,
// 				and correlation coefficients for a set of n pairs of data
// func declarations:
// 		main  
// 		CalcMean  
// 		CalcVarience
// 		CalcCov
// 		CorrCoef

package main

import "fmt"
import "math"

func main() {
	buff := [][]float64{}
	flag := true
	for flag {
		var input float64
		fmt.Println("1つ目のデータを1つずつ入力してください（終了の場合は-1を入力）：")
		fmt.Scan(&input)
		if input < 0 {
			fmt.Println("1つ目のデータ終了")
			break
		} else {
			add_data := [][]float64{{input,0}}
			buff = append(buff, add_data...)
		}
	}
	for i, _ := range buff {
		var input float64
		fmt.Println("2つ目のデータを1つずつ入力してください")
		fmt.Scan(&input)
		buff[i][1] = input
	}

	x := []float64{}
	y := []float64{}
	for _, v := range buff {
		x = append(x, v[0])
		y = append(y, v[1])
	}

// 回帰パラメータBeta_1の計算
	beta_1 := CalcCov(x, y) / CalcVarience(x)
// 回帰パラメータBeta_0の計算
	beta_0 := CalcMean(y) - beta_1 * CalcMean(x)

	fmt.Printf("回帰パラメータBeta_0は%.3f\n", beta_0)
	fmt.Printf("回帰パラメータBeta_1は%.3f\n", beta_1)
	fmt.Printf("相関係数r_xyは%.3f\n", CorrCoef(x, y))
	fmt.Printf("r^2は%.3f\n", CorrCoef(x,y) * CorrCoef(x,y))

// 見積もり値を受け取って見込み値を計算し表示する
	var input float64
	fmt.Println("x_kの値を入力してください：")
	fmt.Scan(&input)
	fmt.Printf("y_kの値は%.3f\n", beta_1 * input + beta_0)
}

// Reuse instructins
// 		func CalcMean(buff []float64) float64
// 		purpose: to calculate mean value
// 		return: mean value of list
func CalcMean(buff []float64) float64 {
	var sum float64
	for _, x := range buff {
		sum += x
	}
	n := float64(len(buff))
	mean := sum / n
	return mean
}

// Reuse instructins
// 		func CalcVarience(buff []float64) float64
// 		purpose: to calculate varience
// 		return: varience value of list
func CalcVarience(buff []float64) float64 {
	n := float64(len(buff))
	mean := CalcMean(buff)
	s := 0.0
	for _, x := range buff{
		s += (x - mean) * (x - mean)
	}
	return s / n
}

// Reuse instructins
// 		func CalcVarience(x,y []float64) float64
// 		purpose: to calculate covarience
// 		return: covarience value of list
func CalcCov(x, y []float64) float64 {
	n := float64(len(x))
	mean_x := CalcMean(x)
	mean_y := CalcMean(y)
	multi_xy := 0.0
	for i, _ := range x {
		multi_xy += x[i] * y[i]
	}
	return multi_xy/n - mean_x * mean_y
}

// Reuse instructins
// 		func CorrCoef(x,y []float64) float64
// 		purpose: to calculate correlation coefficient
// 		return: correlation coefficient value of list
func CorrCoef(x, y []float64) float64 {
	covarience_xy := CalcCov(x, y)
	sd_x := math.Sqrt(CalcVarience(x))
	sd_y := math.Sqrt(CalcVarience(y))
	return covarience_xy/(sd_x * sd_y)
}
