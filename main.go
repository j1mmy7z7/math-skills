package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"

	skills "skills/functions"
)

func main() {
	digits := new(big.Float)
	temp := new(big.Float)
	count := new(big.Float)
	Average := new(big.Float)
	var Numbers []*big.Float
	Args := os.Args[1:]
	if len(Args) != 1 {
		fmt.Println("Error with the number of arguments")
		return
	}
	check , err := skills.IsEmpty(Args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	if check{
		fmt.Println("the file is empty")
		return
	}
	file, err := os.Open(Args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	// we range through the lines in the files, appending the converting values while tracking the lines
	for filescanner.Scan() {
		temp, ok := temp.SetString(filescanner.Text())
		if !ok {
			fmt.Println("Error converting the data : setstring error")
			return
		}
		digits.Add(digits, temp)
		Numbers = append(Numbers, new(big.Float).Set(temp))
		count.Add(count, big.NewFloat(1))
	}

	file.Close()
	//we calculate the average here
	Average.Quo(digits, count)
	//we sort a pointer of the array 
	skills.QuickSort(Numbers, 0, len(Numbers)-1)
	median, err:= skills.Median(Numbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	Variance := skills.Variance(Numbers, count, Average)
	//we round of the average to the next integer
	Average.Add(Average, big.NewFloat(0.5))
	finalAverage := new(big.Int)
	//convert the float average to Int after rounding off
	Average.Int(finalAverage)
	finalVariance := new(big.Int)
	//round of the variance
	Variance.Add(Variance, big.NewFloat(0.5))
	//convert the rounded off variace to an integer
	Variance.Int(finalVariance) 
	StandardDeviaton := skills.StandardDeviation(Variance)
	fmt.Println("Average :",finalAverage)
	fmt.Println("Median :" ,median)
	fmt.Println("Variance :",finalVariance)
	fmt.Println("Standard Deviation :",StandardDeviaton)
}
