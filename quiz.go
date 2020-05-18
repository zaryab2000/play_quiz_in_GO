package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
	"time"
	"encoding/csv" 
)
func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}

func parseproblems(lines  [][]string) []problem{
	ret := make([]problem, len(lines))
	for index, line := range lines{
		ret[index] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct{
	question string
	answer string
}

func main() {
	file_name := flag.String("csv", "problems.csv","fetching problems")
	
	timeLimit := flag.Int("limit", 30,"time limit for the quiz game")

	flag.Parse()
	
	file, error := os.Open(*file_name)
	if error != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n",*file_name))
		
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to Open file")
	}
	probs := parseproblems(lines)

	timer := time.NewTimer(time.Duration(*timeLimit)*time.Second)
	correct := 0
	for index, problem := range probs{
		fmt.Printf("Problem #%d: %s= \n",index+1,problem.question)
		answerCh := make(chan string)
		go func(){
			var ans string
			fmt.Scanln(&ans)
			answerCh <- ans
		}()
		select {
		case <- timer.C:
			fmt.Printf("Hey! Your Score is %d out of %d",correct,len(probs))
			return
		case answer := <-answerCh:
			if answer == problem.answer{
				correct ++
			}
			
		}
	} 
}


