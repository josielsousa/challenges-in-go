// This templates is based on  https://www.codementor.io/@tucnak/using-golang-for-competitive-programming-h8lhvxzt3
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{})                   { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) (n int, err error) { return fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	for {
		var a, b uint32
		_, err := scanf("%d %d\n", &a, &b)
		if err != nil && err == io.EOF {
			break
		}

		printf("%d\n", a^b)
	}
}
