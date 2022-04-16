package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func main() {
	in := os.Stdin
	out := os.Stdout

	customIo = NewCustomIo(in, out)
	defer func() {
		customIo.Writer.Flush()
	}()
	for solve() {
	}
}

func solve() bool {
	w := customIo.GetNextInt()
	h := customIo.GetNextInt()
	if w == 0 && h == 0 {
		return false
	}
	g := NewGraph(w, h)

	for height := 0; height < h; height++ {
		g.land[height] = make([]int, w)
		for width := 0; width < w; width++ {
			g.land[height][width] = customIo.GetNextInt()
		}
	}

	ans := 0
	for height := 0; height < h; height++ {
		for width := 0; width < w; width++ {
			if g.land[height][width] == 0 {
				continue
			}
			ans++
			// 探索開始
			g.dfs(height, width)
		}
	}
	customIo.Println(ans)
	return true
}

type Graph struct {
	w    int
	h    int
	land [][]int
}

func NewGraph(w, h int) *Graph {
	return &Graph{
		w:    w,
		h:    h,
		land: make([][]int, h),
	}
}

/*
	深さ優先探索(Depth First Search)
*/
func (g *Graph) dfs(h, w int) {
	if h < 0 || h >= g.h || w < 0 || w >= g.w {
		// 外に出てしまった
		return
	}
	if g.land[h][w] == 0 {
		// 海 or 訪問済み
		return
	}
	g.land[h][w] = 0
	// 移動
	dh := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dw := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < 8; i++ {
		g.dfs(h+dh[i], w+dw[i])
	}
}
