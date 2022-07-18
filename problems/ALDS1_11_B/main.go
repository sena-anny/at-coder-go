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
	defer func() {
		customIo.Writer.Flush()
	}()
	solve()
}

func solve() {
	n := customIo.GetNextInt()
	g := Graph{
		edge:         make([][]int, n+1),
		discoverTime: make([]int, n),
		finishTime:   make([]int, n),
		time:         0,
	}
	for i := 1; i < n+1; i++ {
		u := customIo.GetNextInt()
		k := customIo.GetNextInt()
		g.edge[u] = make([]int, k+1)
		for j := 1; j < k+1; j++ {
			g.edge[u] = append(g.edge[u], customIo.GetNextInt())
		}
	}

	// 頂点番号の小さい順に訪問する
	for i := range g.edge {
		sort.Ints(g.edge[i])
	}

	g.dfs(1)
	for i := 1; i < n+1; i++ {
		customIo.Printf("%d %d %d", i, g.discoverTime[i], g.finishTime[i])
	}

}

type Graph struct {
	edge [][]int
	// 頂点の発見時刻 初期値0
	discoverTime []int
	// 頂点の探索完了時刻 初期値0
	finishTime []int
	// 現在時刻
	time int
}

// DFS
func (g Graph) dfs(now int) {
	g.time++
	g.discoverTime[now] = g.time

	for i := range g.edge[now] {
		if g.discoverTime[i] == 0 {
			g.dfs(i)
		}
		g.finishTime[now] = g.time
	}
}
