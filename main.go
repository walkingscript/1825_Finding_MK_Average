package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var TEST_CASE_DIR = "input_data/in2"

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
	m, k                int
	leftSide, rightSide []int
	mid                 BinTree
}

func Constructor(m int, k int) MKAverage {
	var obj = MKAverage{m: m, k: k}
	obj.leftSide = make([]int, 0, k)
	obj.rightSide = make([]int, 0, k)
	return obj
}

func (avgObject *MKAverage) AddElement(num int) {
	if avgObject.mid.ItemsCount == avgObject.m {
		midMin, midMax := avgObject.mid.Min(), avgObject.mid.Max()
		if num <= midMin {
			avgObject.leftSide = append(avgObject.leftSide, num)
		} else if num >= midMax {
			avgObject.rightSide = append(avgObject.rightSide, num)
		} else {
			if num < avgObject.mid.Root.Value {
				avgObject.leftSide = append(avgObject.leftSide, avgObject.mid.PopLeft())
			} else {
				avgObject.rightSide = append(avgObject.rightSide, avgObject.mid.PopRight())
			}
		}
	}
	avgObject.mid.Insert(num)
}

func (avgObject *MKAverage) CalculateMKAverage() int {
	if avgObject.mid.ItemsCount < avgObject.m {
		return -1
	}
	return avgObject.GetAverage()
}

func (avgObject *MKAverage) GetAverage() ItemType {
	return avgObject.mid.Sum() / avgObject.mid.ItemsCount
}
