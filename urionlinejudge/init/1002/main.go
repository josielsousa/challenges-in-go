// This templates is based on  https://www.codementor.io/@tucnak/using-golang-for-competitive-programming-h8lhvxzt3
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

const pi = 3.14159

func main() {
	defer writer.Flush()

	var r float64
	scanf("%f\n", &r)

	a := pi * (math.Pow(r, 2))
	printf("A=%.4f\n", a)
}
