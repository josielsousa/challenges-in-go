package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

/**

1-- 2 -- 3 -- 4 -- 5 -- 6 -- 7 -- 8

farms := 8
sheepsInFarm := []int64{1, 3, 5, 7, 11, 13, 17, 19}
sheepsInFarm := []int64{1, 3, 5, 7, 11, 13, 16, 19}

8
1 3 5 7 11 13 16 19

-- Case shees on Star i == impar odd, then Star i + 1
-- Case shees on Star i == par even, then Star i - 1
**/

func main() {
	defer writer.Flush()

	var farms int64
	scanf("%d\n", &farms)

	totalNonStolenSheep := int64(0)
	sheepsInFarm := make([]int64, farms)

	i := int64(0)
	for i < farms {
		scanf("%d", &sheepsInFarm[i])
		totalNonStolenSheep += sheepsInFarm[i]
		i++
	}

	farmActual := int64(0)
	nonStolenSheep := int64(0)
	totalFarmsvisited := int64(0)

	totalStolen := int64(0)

	run := true
	for run {
		totalFarmsvisited++

		//When passes by the farm Star i, he steals only one sheep (if there is any)
		sheepsNow := sheepsInFarm[farmActual]
		if sheepsNow > 0 {
			sheepsInFarm[farmActual]-- // rouba
			totalStolen++
		}

		nonStolenSheep += sheepsInFarm[farmActual]
		if sheepsNow%2 == 0 {
			farmActual -= 1

			if farmActual >= 0 {
				totalFarmsvisited--
			}

		} else {
			farmActual += 1
		}

		//If there is not the Star to which he wants to go, he halts his trek.
		run = farmActual >= 0 && farmActual < farms
	}

	totalNonStolenSheep -= totalStolen
	printf("%d %d\n", totalFarmsvisited, totalNonStolenSheep)
}
