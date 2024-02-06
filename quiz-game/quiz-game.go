package quiz_game

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func RollQuizGame() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds.")

	flag.Parse()

	//Opening the file
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	//Reading the file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	//Parsing the slice into a slice of structs
	problems := parseLines(lines)

	//Setting up the timer in seconds. Int needs to be converted to Duration type
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//Count for correct answers.
	correct := 0
	
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		answerCh := make(chan string)
		//Loading answers into answer channel via go routine.
		// This is a closure using data defined outside of it.
		go func() {
			var answer string
			//Scanf gets rid of trailing spaces...
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		//Using a select statement to break out if time is up.
		//If the answer is correct, it will carry on.
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
			//Loading the value out of the channel
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("Correct")
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	//Breaking this down into a slice of structs.
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
