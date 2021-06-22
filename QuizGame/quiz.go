package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := flag.String("f", "problems.csv", "Name of the problem file ")
	duration := flag.Int64("t", 100, "Duration (in second) of the quiz. Default is 30 sec.")
	flag.Parse()

	fmt.Println("Hit enter to begin.")
	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')
	for {
		if strings.HasSuffix(input, "\n") {
			break
		}
	}
	fmt.Println("Starting with file", *fileName, ", and duration", *duration, "second")

	file, err := os.Open(*fileName)
	check(err)
	r := csv.NewReader(file)
	fileDate, err := r.ReadAll()
	check(err)
	numLines, points := len(fileDate), 0

	time.AfterFunc(time.Second*time.Duration(*duration), func() {
		fmt.Println("Time exceeded.")
		fmt.Println(points, "/", numLines, "correct.")
		os.Exit(0)
	})
	problems := parseLines(fileDate)
	for _, problem := range problems {
		fmt.Println(problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.a {
			points++
			fmt.Println("you got it correct!")
		} else {
			fmt.Println("you got it wrong!")
		}
	}
	fmt.Println(points, "/", numLines, "correct.")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}
