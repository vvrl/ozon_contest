package firsttaskwhomeowed

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strings"
)

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
	return slices.Contains(words, str)
}

func findMax(score map[string]int) []string {

	max := math.MinInt
	var result []string

	for name, value := range score {
		if value > max {
			max = value
			result = []string{name}
		} else if value == max {
			result = append(result, name)
		}
	}

	sort.Strings(result)

	return result

}

func Task() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numOfSets, numOfStates, counter int

	fmt.Fscan(in, &numOfSets)

	for counter < numOfSets {

		fmt.Fscan(in, &numOfStates)
		var action string
		_, _ = in.ReadString('\n') // пропуск \n

		score := make(map[string]int, numOfStates)

		for i := 0; i < numOfStates; i++ {

			statesment, err := in.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			words := clearingWords(statesment)

			if len(words) == 0 {
				continue
			}

			flagNot := containsString(words, "not")
			var ind int

			if flagNot {
				ind = 4
			} else {
				ind = 3
			}

			action = strings.Join(words[ind:], " ")

			switch words[2] {
			case "am":

				if flagNot {
					score[words[0]] -= 1
				} else {
					score[words[0]] += 2
				}

			case "is":
				if flagNot {
					score[words[1]] -= 1
				} else {
					score[words[1]] += 1
				}

				_, ok := score[words[0]]
				if !ok {
					score[words[0]] = 0
				}

			}

		}

		result := findMax(score)

		for _, name := range result {
			fmt.Fprintf(out, "%s is %s.\n", name, action)
		}

		counter++

	}
}
