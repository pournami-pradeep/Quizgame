package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csv_file := flag.String("csv", "problems.csv", "csv file which contains questions and answers")
	shuffle := flag.Bool("s", false, "flag to decide whether to shuffle questions or not")
	timer := flag.Int("t", 30, "Timer for the quiz")
	flag.Parse()
	file, err := os.Open(*csv_file)
	if err != nil {
		fmt.Println("Error while opening the csv file.Please provide another file.")
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Invalid file!", err)
	}
	problems := parseLines(lines)
	fmt.Println(problems, "PRI")
	if *shuffle {
		shuffleList(problems)

	}
	t := time.NewTimer(time.Duration(*timer) * time.Second)
	var score int
	for _, p := range problems {
		select {
		case <-t.C:
			fmt.Println("Time out")
			return

		default:
			fmt.Println(p.q, " ?")
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				score += 1
			}

		}

	}

	fmt.Printf("You scored %d out of %d\n", score, len(problems))

}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		if len(line) < 2 {
			continue
		}

		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret

}
func shuffleList(problems []problem) {
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})

}
