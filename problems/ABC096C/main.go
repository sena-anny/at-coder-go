package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	H := customIo.GetNextInt()
	W := customIo.GetNextInt()

	S := make([][]string, H+2)
	S[0] = make([]string, W+2)
	S[H+1] = make([]string, W+2)
	for i := 1; i <= H; i++ {
		S[i] = make([]string, W+2)
		input := customIo.Text()
		splitInput := strings.Split(input, "")
		for j := 1; j <= W; j++ {
			S[i][j] = splitInput[j-1]
		}
	}

	cnt := 0
	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			if S[i][j] == "#" && S[i-1][j] != "#" && S[i][j-1] != "#" && S[i][j+1] != "#" && S[i+1][j] != "#" {
				cnt++
			}
		}
	}

	if cnt > 0 {
		customIo.Println("No")
	} else {
		customIo.Println("Yes")
	}

}
