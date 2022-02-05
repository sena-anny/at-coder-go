package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
		path, err := filepath.Abs("./input.txt")
		in, err = os.Open(path)
		if err != nil {
			log.Printf("ファイル読み込み失敗:%v", err)
		}
	}

	customIo = NewCustomIo(in, out)
	defer customIo.Writer.Flush()
	solve()
}

func getCutPoint(cutSize int, startPoint int) int {
	if 360-startPoint-cutSize > 0 {
		return startPoint + cutSize
	} else {
		return startPoint + cutSize - 360
	}
}

func solve() {
	N := customIo.GetNextInt()
	Point := make([]int, N+2)
	Point[0] = 0
	Point[N+1] = 360
	for i := 0; i < N; i++ {
		Point[i+1] = getCutPoint(customIo.GetNextInt(), Point[i])
	}
	sort.Ints(Point)

	maxSize := 0
	tmpSize := 0
	for _, p := range Point {
		if p-tmpSize > maxSize {
			maxSize = p - tmpSize
		}
		tmpSize = p
	}
	customIo.Println(maxSize)
}

func cntDigits(n int) int {
	digits := 0
	for n > 0 {
		n = n / 10
		digits++
	}
	return digits
}

func maxInt(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
