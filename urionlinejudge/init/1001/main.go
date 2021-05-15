// This templates is based on  https://www.codementor.io/@tucnak/using-golang-for-competitive-programming-h8lhvxzt3
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

func main() {
	defer writer.Flush()

	var a, b int
	scanf("%d\n%d", &a, &b)

	printf("X = %d\n", a+b)
}
