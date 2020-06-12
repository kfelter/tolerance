package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/felts94/tolerance"
)

var (
	subFunc = func(num ...float64) float64 {
		return num[0] - num[1]
	}

	sumFunc = func(num ...float64) float64 {
		sum := float64(0)
		for i := range num {
			sum += num[i]
		}
		return sum
	}

	factFunc = func(num ...float64) float64 {
		sum := float64(1)
		for cur := num[0]; cur > 0; cur-- {
			sum *= cur
		}
		return sum
	}
)

func usage() string {
	return "<add|sub|fact>"
}

func main() {
	tolerance.Fault = 50
	log.Printf("Using %d redundant processes\n", tolerance.Threads)
	log.Println("System is", tolerance.Fault, "% faulty")
	log.Printf("args %v", os.Args)
	rand.Seed(time.Now().Unix())

	if len(os.Args) < 2 {
		fmt.Println(usage())
		return
	}
	args := os.Args[1:]
	switch args[0] {
	case "add":
		args = args[1:]
		do(args, sumFunc)
	case "sub":
		args = args[1:]
		if len(args) != 2 {
			fmt.Println("Must have exactly 2 numbers: ", args)
			return
		}
		do(args, subFunc)
	case "fact":
		args = args[1:]
		if len(args) != 1 {
			fmt.Println("Must have exactly 1 number: ", args)
			return
		}
		do(args, factFunc)
	default:
		fmt.Println(usage())
	}
}

func do(args []string, mFunc func(...float64) float64) {
	nums := make([]float64, len(args))
	for i := range args {
		nums[i], _ = strconv.ParseFloat(args[i], 64)
	}
	result, attempts, dur := tolerance.Do(nums, mFunc)
	fmt.Printf("result [%v] attempts [%d] faulty%% [%d] time [%v]\n", result, attempts, tolerance.Fault, dur)
}
