package secondtask

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Statement struct {
	firstName, secondName string
	delta                 int
	isEqual               bool
}

func clearingWords(str string) []string {

	statesment := strings.TrimSpace(str)
	words := strings.Fields(statesment)
	clearArray := make([]string, len(words))

	for ind, word := range words {
		clearArray[ind] = strings.Trim(word, ":!.,?")
	}
	return clearArray
}

func containsString(words []string, str string) bool {
	for _, v := range words {
		if v == str {
			return true
		}
	}
	return false
}

func solveTask(statesments []Statement, targetName string) int {

	fmt.Println(statesments)

	age := make(map[string]int)
	known := make(map[string]bool)

	working := true

	for working {
		working = false
		for _, state := range statesments {
			if state.secondName == "" && !state.isEqual {
				age[state.firstName] = state.delta
				known[state.firstName] = true

			}
			if state.isEqual {
				if known[state.firstName] && !known[state.secondName] {
					age[state.secondName] = age[state.firstName]
					known[state.secondName] = true
					working = true
				} else if !known[state.firstName] && known[state.secondName] {
					age[state.firstName] = age[state.secondName]
					known[state.firstName] = true
					working = true
				}
			} else {
				if known[state.firstName] && !known[state.secondName] {
					age[state.secondName] = age[state.firstName] - state.delta
					known[state.secondName] = true
					working = true
				} else if !known[state.firstName] && known[state.secondName] {
					age[state.firstName] = age[state.secondName] + state.delta
					known[state.firstName] = true
					working = true

				}
			}
		}
	}

	fmt.Println(age)

	return age[targetName]
}

func Task() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numOfSets int
	fmt.Fscan(in, &numOfSets)

	_, _ = in.ReadString('\n')

	for i := 0; i < numOfSets; i++ {
		askString, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		targerName := clearingWords(askString)[3]

		statesments := make([]Statement, 3)

		for j := 0; j < 3; j++ {
			state, err := in.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			words := clearingWords(state)

			if containsString(words, "than") {
				delta, _ := strconv.Atoi(words[2])
				statesments[j] = Statement{firstName: words[0], secondName: words[6], delta: delta, isEqual: false}
				if containsString(words, "younger") {
					statesments[j].delta *= -1
				}
			} else if containsString(words, "same") {
				statesments[j] = Statement{firstName: words[0], secondName: words[6], delta: 0, isEqual: true}
			} else {
				delta, _ := strconv.Atoi(words[2])
				statesments[j] = Statement{firstName: words[0], secondName: "", delta: delta, isEqual: false}
			}

		}

		fmt.Println(statesments)

		fmt.Fprintln(out, solveTask(statesments, targerName))

	}

}
