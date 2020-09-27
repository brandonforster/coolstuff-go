package quizzer

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"log"
	"os"
)

// Main entry point
func Main() {
	file := pullFile()

	answers := []bool{}

	for _, line := range file {
		answers = append(answers, parseLine(line))
	}

	for i := range answers {
		fmt.Println(answers[i])
	}
}

func pullFile() []string{
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	return lines
}

func parseLine(line string) bool {
	tokens := strings.Split(line, ",")

	question := tokens[0]
	answer, err := strconv.Atoi(tokens[1])
	if err != nil {
		log.Fatal(err)
	}

	operators := []string{"+", "-", "*", "/"}
	for _, op := range operators {
		isOp, members := parseOp(question, op)
		if isOp {
			return doOperation(members, op, answer)
		}
	}

	return false
}

func parseOp(question string, operator string) (bool, []int) {
	tokens := strings.Split(question, operator)

	if len(tokens) == 1 {
		return false, nil
	}

	members := []int{}
	for _, token := range tokens {
		t, err := strconv.Atoi(token)
		if err != nil {
			log.Fatal(err)
		}

		members = append(members, t)
	}

	return true, members
}

func doOperation(members []int, operation string, answer int) bool {
	if operation == "+" {
		solution :=  members[0] + members[1]
		return solution == answer
	}

	if operation == "-" {
		solution :=  members[0] - members[1]
		return solution == answer
	}

	if operation == "*" {
		solution :=  members[0] * members[1]
		return solution == answer
	}

	if operation == "/" {
		solution :=  members[0] / members[1]
		return solution == answer
	}

	return false
}
