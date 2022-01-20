package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

var customIo *CustomIo

type CustomIo struct {
	Scanner *bufio.Scanner
	Writer  *bufio.Writer
}

func NewCustomIo(in io.Reader, out io.Writer) *CustomIo {
	const BufSize = 2000005
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, BufSize), BufSize)
	return &CustomIo{Scanner: scanner, Writer: bufio.NewWriter(out)}
}

func (i *CustomIo) Text() string {
	if !i.Scanner.Scan() {
		panic("scan failed")
	}
	return i.Scanner.Text()
}
func (i *CustomIo) Atoi(s string) int                 { x, _ := strconv.Atoi(s); return x }
func (i *CustomIo) GetNextInt() int                   { return i.Atoi(i.Text()) }
func (i *CustomIo) Atoi64(s string) int64             { x, _ := strconv.ParseInt(s, 10, 64); return x }
func (i *CustomIo) GetNextInt64() int64               { return i.Atoi64(i.Text()) }
func (i *CustomIo) Atof64(s string) float64           { x, _ := strconv.ParseFloat(s, 64); return x }
func (i *CustomIo) GetNextFloat64() float64           { return i.Atof64(i.Text()) }
func (i *CustomIo) Print(x ...interface{})            { fmt.Fprint(i.Writer, x...) }
func (i *CustomIo) Printf(s string, x ...interface{}) { fmt.Fprintf(i.Writer, s, x...) }
func (i *CustomIo) Println(x ...interface{})          { fmt.Fprintln(i.Writer, x...) }
func isLocal() bool                                   { return os.Getenv("AT_CODER") == "KJ" }

func main() {
	in := os.Stdin
	out := os.Stdout

	if isLocal() {
		path, _ := filepath.Abs("./input.txt")
		in, _ = os.Open(path)
	}

	customIo = NewCustomIo(in, out)
	defer customIo.Writer.Flush()
	solve()
}

func solve() {
	A := customIo.GetNextInt()
	B := customIo.GetNextInt()
	K := customIo.GetNextInt()

	divA := make(map[int]int)
	divB := make(map[int]int)

	for i := 1; i <= A; i++ {
		if A%i == 0 {
			divA[i] = i
		}
	}
	for i := 1; i <= B; i++ {
		if B%i == 0 {
			divB[i] = i
		}
	}

	var divL []int

	for _, a := range divA {
		if val, ok := divB[a]; ok {
			divL = append(divL, val)
		}
	}
	sort.Ints(divL)

	customIo.Println(divL[len(divL)-K])
}
