package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{})                   { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) (n int, err error) { return fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var qtd int
	scanf("%d\n", &qtd)

	//scanner := bufio.NewScanner(os.Stdin)

	count := 0
	for count < qtd {
		count++

		// scanner.Scan()
		val, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error on get text: ", err)
		}

		text := string(val)
		// scanner.Scan()
		// word := scanner.Text()

		val, _, err = reader.ReadLine()
		if err != nil {
			fmt.Println("Error on get word: ", err)
		}

		word := string(val)

		//printf("%s\n", fmt.Sprintf("\\b%s\\b", word))
		exp := regexp.MustCompile(fmt.Sprintf("\\b%s\\b", word))
		matchs := exp.FindAllStringIndex(text, -1)

		if len(matchs) <= 0 {
			printf("%d\n", -1)
			continue
		}

		exit := ""
		for i, match := range matchs {
			if i == 0 {
				exit += fmt.Sprintf("%d", match[0])
			} else {
				exit += fmt.Sprintf(" %d", match[0])
			}
		}

		printf("%s\n", exit)
	}
}

/**
6
43 53
63
*/
