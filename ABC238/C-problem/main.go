package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	defer customIo.Writer.Flush()
	solve()
}

func solve() {
	N := customIo.GetNextInt()
	const M = 998244353

	ans := 0
	for i := 1; i <= N; i *= 10 {
		x := minInt(N, i*10-1)
		y := x - i + 1
		// mod Mにおける等差数列の和
		ans += (y + 1) % M * (y % M) % M * invMod(2, M) % M
	}
	ans = (ans%M + M) % M

	customIo.Println(ans)

}

func cntDigits(n int) int {
	digits := 0
	for n > 0 {
		n = n / 10
		digits++
	}
	return digits
}

func maxInt(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// mod Mにおけるaの逆元を求める（非再帰拡張 Euclid の互除法）
func invMod(a, M int) int {
	p, x, u := M, 1, 0
	for p != 0 {
		t := a / p
		a, p = p, a-t*p
		x, u = u, x-t*u
	}
	if x < 0 {
		x += M
	}
	return x
}
