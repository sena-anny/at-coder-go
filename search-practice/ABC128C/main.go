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
	M := customIo.GetNextInt()

	var ktmp int
	// 各電球とそれにつながるスイッチ群の関係
	SwitchList := make([][]int, M)
	// 電球の点灯条件
	P := make([]int, M)
	for i := 0; i < M; i++ {
		// スイッチの数
		ktmp = customIo.GetNextInt()
		// 電球iにつながるスイッチ群
		SwitchList[i] = make([]int, ktmp)
		for j := 0; j < ktmp; j++ {
			// 電球iにつながるスイッチjの番号
			SwitchList[i][j] = customIo.GetNextInt()
		}
	}
	for i := 0; i < M; i++ {
		P[i] = customIo.GetNextInt()
	}

	count := 0
	// スイッチが全てオフ（00000）-全てオン（11111）の状態
	for bits := 0; bits < 1<<N; bits++ {
		on := true
		for light := 0; light < M; light++ {
			sum := 0
			for _, switchNumber := range SwitchList[light] {
				sum += bits >> (switchNumber - 1) & 1
			}
			if (sum % 2) != P[light] {
				on = false
				break
			}
		}
		if on {
			count++
		}
	}
	customIo.Println(count)

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
