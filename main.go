package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var TEST_CASE_DIR = "input_data/in3"

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

	cmds := make(chan string, 50000)
	data := make(chan int, 50000)

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
			fmt.Printf("%d ", res)
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
	m, k        int
	Left, Right []int
	Mid         BinTree
}

func Constructor(m int, k int) MKAverage {
	var obj = MKAverage{m: m, k: k}
	return obj
}

func (MKAvgObject *MKAverage) AddElement(num int) {
	MKAvgObject.Mid.Insert(num)
}

func (MKAvgObject *MKAverage) CalculateMKAverage() int {
	if MKAvgObject.Mid.Count < MKAvgObject.m {
		return -1
	}
	if MKAvgObject.Mid.Count >= MKAvgObject.m {
		for i := 0; i < MKAvgObject.k; i++ {
			MKAvgObject.Left = append(MKAvgObject.Left, MKAvgObject.Mid.PopLeft())
			MKAvgObject.Right = append(MKAvgObject.Right, MKAvgObject.Mid.PopRight())
		}
	}
	return MKAvgObject.GetAverage()
}

func (MKAvgObject *MKAverage) GetAverage() ItemType {
	return MKAvgObject.Mid.Sum() / MKAvgObject.Mid.Count
}

func Sum(items []ItemType) ItemType {
	var total ItemType
	for _, item := range items {
		total += item
	}
	return total
}
