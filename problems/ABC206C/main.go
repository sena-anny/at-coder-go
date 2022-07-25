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
	defer func() {
		customIo.Writer.Flush()
	}()
	solve()
}

func solve() {
	N := customIo.GetNextInt()
	X := customIo.GetNextInt()
	Y := customIo.GetNextInt()

	j := &jewel{red: make([]int, N+1), blue: make([]int, N+1), X: X, Y: Y}
	j.changeRed(N)
	customIo.Println(j.blue[1])
}

type jewel struct {
	red  []int
	blue []int
	X    int
	Y    int
}

func (j *jewel) changeRed(level int) {
	if level-1 == 0 {
		return
	}
	j.red[level-1] += 1
	j.blue[level] += j.X
	j.changeRed(level - 1)
	for i := 0; i < j.X; i++ {
		j.changeBlue(level)
	}
}

func (j *jewel) changeBlue(level int) {
	if level-1 == 0 {
		return
	}
	j.red[level-1] += 1
	j.blue[level-1] += j.Y
	j.changeRed(level - 1)
	for i := 0; i < j.Y; i++ {
		j.changeBlue(level - 1)
	}
}
