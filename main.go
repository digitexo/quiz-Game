package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "csv file in the format of question ,,,answer")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
	flag.Parse()

	_ = csvFilename

	file, err := os.Open(*csvFilename)

	if err != nil {

		exit(fmt.Sprintf("Failed to open csv file : %s \n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("failed to parse csv file")
	}

	problems := parseLines(lines)
	//fmt.Println(problems[1])
	//fmt.Println(problem)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//<-timer.C
	//fmt.Println(timer)

	correct := 0
	wrong := 0
	for k, v := range problems {
		select {
		case <-timer.C:
			fmt.Println("you failed with time")
			return
		default:
			fmt.Printf("Probleem #%d :  %s = \n", k+1, v.question)

			tulem := ""
			fmt.Scanf("%s\n", &tulem)
			if tulem == v.answer {
				fmt.Println("true")
				correct++
			} else {
				fmt.Println("false")
				wrong++
			}
		}

	}

	fmt.Printf("Oigeid vastuseid on %d ja valesid %d\n", correct, wrong)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}

	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
