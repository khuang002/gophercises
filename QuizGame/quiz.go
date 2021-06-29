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
	duration := flag.Int64("t", 30, "Duration (in second) of the quiz. Default is 30 sec.")
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
	problems := parseLines(fileDate)

	timer := time.NewTimer(time.Second * time.Duration(*duration))

problemLoop:
	for _, problem := range problems {
		fmt.Println(problem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("Time exceeded.")
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.a {
				points++
				fmt.Println("you got it correct!")
			} else {
				fmt.Println("you got it wrong!")
			}
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
