package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var TEST_CASE_DIR = "input_data/in"

	cmdsBytes, err := os.ReadFile(fmt.Sprintf("%s/commands", TEST_CASE_DIR))
	if err != nil {
		panic(fmt.Errorf("reading in/commands file: %v", err))
	}

	dataBytes, err := os.ReadFile(fmt.Sprintf("%s/data", TEST_CASE_DIR))
	if err != nil {
		panic(fmt.Errorf("reading in/data file: %v", err))
	}

	cmdStrings := strings.Split(string(cmdsBytes), "\n")
	dataStrings := strings.Split(string(dataBytes), "\n")

	cmds := make(chan string, 25000)
	data := make(chan int, 25000)

	go ProduceCmds(cmds, cmdStrings)
	go ProduceData(data, dataStrings)

	var obj MKAverage
	var res int

	for cmd := range cmds {
		switch cmd {
		case "MKAverage":
			obj = Constructor(<-data, <-data)
		case "addElement":
			obj.AddElement(<-data)
		case "calculateMKAverage":
			res = obj.CalculateMKAverage()
			fmt.Println(res)
		}
	}

}

func ProduceCmds(cmds chan<- string, cmdsSlice []string) {
	for _, cmd := range cmdsSlice {
		cmds <- cmd
	}
	close(cmds)
}

func ProduceData(data chan<- int, dataSlice []string) {
	for _, item := range dataSlice {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		data <- num
	}
	close(data)
}

// ------------------------------------ Solution ------------------------------------

const CAPACITY = 100

type MKAverage struct {
	m, k              int
	stream, container []int
	tree              BinTree
}

func Constructor(m int, k int) MKAverage {
	var obj = MKAverage{m: m, k: k}
	obj.stream = make([]int, 0, m+100)
	obj.container = make([]int, 0, m)
	return obj
}

func (avgObject *MKAverage) AddElement(num int) {
	avgObject.stream = append(avgObject.stream, num)
	if len(avgObject.stream) > avgObject.m {
		avgObject.stream = avgObject.stream[1:]
	}
}

func (avgObject *MKAverage) CalculateMKAverage() int {
	if len(avgObject.stream) < avgObject.m {
		return -1
	}
	avgObject.tree.Reset()
	items := avgObject.stream[len(avgObject.stream)-avgObject.m:]
	for _, item := range items {
		avgObject.tree.Insert(item)
	}
	avgObject.container = avgObject.tree.Root.GetSortedArray()
	avgObject.container = avgObject.container[avgObject.k : len(avgObject.container)-avgObject.k]
	return int(float64(Sum(avgObject.container)) / float64(len(avgObject.container)))
}

func Sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
