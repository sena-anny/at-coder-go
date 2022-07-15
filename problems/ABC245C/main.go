package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
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
	defer func() {
		customIo.Writer.Flush()
	}()
	solve()
}

func solve() {
	N := customIo.GetNextInt()
	K := customIo.GetNextInt()

	A := make([]int, N+1)
	B := make([]int, N+1)
	DP := make([]bool, N+1)
	EP := make([]bool, N+1)
	for i := 1; i < N+1; i++ {
		A[i] = customIo.GetNextInt()
	}
	for i := 1; i < N+1; i++ {
		B[i] = customIo.GetNextInt()
	}

	// 動的計画法
	DP[1] = true
	EP[1] = true

	for i := 2; i < N+1; i++ {
		if DP[i-1] {
			if math.Abs(float64(A[i]-A[i-1])) <= float64(K) {
				DP[i] = true
			}
			if math.Abs(float64(B[i]-A[i-1])) <= float64(K) {
				EP[i] = true
			}
		}
		if EP[i-1] {
			if math.Abs(float64(A[i]-B[i-1])) <= float64(K) {
				DP[i] = true
			}
			if math.Abs(float64(B[i]-B[i-1])) <= float64(K) {
				EP[i] = true
			}
		}
	}
	if DP[N] == true || EP[N] == true {
		customIo.Println("Yes")
	} else {
		customIo.Println("No")
	}
}
