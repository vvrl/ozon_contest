package thirdtask

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convArrayToInt(arr []string) []int {
	intArray := make([]int, len(arr))

	for i, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		intArray[i] = num
	}

	return intArray
}

func Task() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	directionsEarthquake := [][]int{{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	directionsHouse := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	var numOfSets int
	fmt.Fscan(in, &numOfSets)
	//_, _ = in.ReadString('\n')

	for i := 0; i < numOfSets; i++ {

		var n, m int
		var x, y, p int

		fmt.Fscan(in, &n, &m)
		_, _ = in.ReadString('\n')

		housesGrid := make([][]int, n)

		// матрица с домами
		for i := 0; i < n; i++ {
			row, _ := in.ReadString('\n')
			row = strings.TrimSpace(row)
			rowElems := strings.Split(row, "")

			housesGrid[i] = convArrayToInt(rowElems)
		}

		fmt.Fscan(in, &x, &y, &p)
		_, _ = in.ReadString('\n')

		x, y = x-1, y-1

		earthquakeGrid := make([][]int, n)

		for i := range earthquakeGrid {
			earthquakeGrid[i] = make([]int, m)
		}

		earthquakeGrid[x][y] = p
		queue := [][]int{{x, y}}

		// сетка распространения землетрясения
		for len(queue) > 0 {
			currentPos := queue[0]
			queue = queue[1:]
			currentPower := earthquakeGrid[currentPos[0]][currentPos[1]]

			if currentPower <= 1 {
				continue
			}

			for _, direction := range directionsEarthquake {
				nx, ny := currentPos[0]+direction[0], currentPos[1]+direction[1]
				if nx >= 0 && ny >= 0 && nx < n && ny < m {
					if earthquakeGrid[nx][ny] < currentPower-1 {
						earthquakeGrid[nx][ny] = currentPower - 1
						queue = append(queue, []int{nx, ny})
					}
				}
			}
		}

		isVisited := make([][]bool, n)
		for i := range isVisited {
			isVisited[i] = make([]bool, m)
		}

		var destroyed int

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if housesGrid[i][j] != 0 && !isVisited[i][j] {
					house := make([][]int, 0)
					queue := [][]int{{i, j}}
					isVisited[i][j] = true
					//isDestroyed := false

					for len(queue) > 0 {
						currentPos := queue[0]
						queue = queue[1:]
						house = append(house, currentPos)

						for _, direction := range directionsHouse {
							nx, ny := currentPos[0]+direction[0], currentPos[1]+direction[1]
							if nx >= 0 && ny >= 0 && nx < n && ny < m {
								if housesGrid[nx][ny] != 0 && !isVisited[nx][ny] {
									isVisited[nx][ny] = true
									queue = append(queue, []int{nx, ny})
								}
							}
						}
					}

					for _, housePart := range house {
						if earthquakeGrid[housePart[0]][housePart[1]] > housesGrid[housePart[0]][housePart[1]] {
							destroyed++
							break

						}
					}
				}
			}
		}

		fmt.Println(destroyed)

	}

}
