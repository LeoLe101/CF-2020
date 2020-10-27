package utils

import (
	"fmt"
	"math"
	"sort"
)

// Profile - hold of all the request time and information about the calculation
type Profile struct {
	numberRequest      int
	fastestTime        int
	slowestTime        int
	totalTime          int
	maxSize            int
	minSize            int
	failTime           int
	meanTime           int
	medianTime         int
	requestSuccessRate int
	requestTimeList    []int
	errorList          []error
}

// InitProfile - Profile for further calculation
func (p *Profile) InitProfile() {
	p.numberRequest = -1
	p.totalTime = -1
	p.failTime = -1
	p.fastestTime = math.MaxInt64
	p.slowestTime = math.MinInt64
	p.maxSize = math.MinInt64
	p.minSize = math.MaxInt64
	p.meanTime = -1
	p.medianTime = -1
	p.requestSuccessRate = 100
}

// HandleError - Keep track all the errors code
func (p *Profile) HandleError(err error) {
	p.errorList = append(p.errorList, err)
}

// CalculateRequest - Calculate all the information from each request for final Info Print out
func (p *Profile) CalculateRequest(requestTime int, dataSize int) {
	p.requestTimeList = append(p.requestTimeList, requestTime)
	p.totalTime += requestTime

	// Get Fastest and Slowest Time
	if p.fastestTime > requestTime {
		p.fastestTime = requestTime
	}

	if p.slowestTime < requestTime {
		p.slowestTime = requestTime
	}

	// Get Max and Min Size
	if p.maxSize < dataSize {
		p.maxSize = dataSize
	}

	if p.minSize > dataSize {
		p.minSize = dataSize
	}
}

// PrintInfo - Print out all information from calculation
func (p *Profile) PrintInfo() {

	// Calculate Mean & Median time
	p.meanTime = p.totalTime / p.numberRequest

	// Sort array of time requested to find median value
	sort.Ints(p.requestTimeList)
	arrLen := len(p.requestTimeList)
	var index int = 0

	// Check odd or even length in array time
	if (arrLen % 2) == 0 {
		index = (arrLen - 1) / 2
		if index < arrLen && index > -1 {
			p.medianTime = p.requestTimeList[index]
		}
	} else {
		index = arrLen / 2
		if index < arrLen && (index-1) > -1 {
			p.medianTime = (p.requestTimeList[index] + p.requestTimeList[index-1]) / 2
		}
	}

	// Calculate Success rate percentage
	p.requestSuccessRate = ((p.numberRequest - len(p.errorList)) / p.numberRequest) * 100

	// Print Result
	fmt.Println("Number of Request: ", p.numberRequest)
	fmt.Println("Fastest Time: ", p.fastestTime, "ms")
	fmt.Println("Slowest Time: ", p.slowestTime, "ms")
	fmt.Println("Mean of All Request Time: ", p.meanTime, "ms")
	fmt.Println("Median of All Request Time: ", p.medianTime, "ms")

	fmt.Print("Error Codes Captured: ")
	if len(p.errorList) < 1 {
		fmt.Println("No Error Captured! LUCKY!!!")
	} else {
		for i, currErr := range p.errorList {
			fmt.Println("- ", i, currErr)
		}
	}

	fmt.Println("Successful Rate: ", p.requestSuccessRate, "%")
	fmt.Println("Smallest Response Byte Size: ", p.minSize)
	fmt.Println("Largest Response Byte Size: ", p.maxSize)
}
