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
	g := NewGraph(n)

	for i := 0; i < n; i++ {
		_ = customIo.GetNextInt()
		k := customIo.GetNextInt()
		for j := 0; j < k; j++ {
			g.edges[i] = append(g.edges[i], Edge{
				from: i,
				to:   customIo.GetNextInt() - 1,
			})
		}
	}
	// 隣接リストを小さい順に並べる
	for i := 0; i < n; i++ {
		sort.Sort(g.edges[i])
	}
	for i := 0; i < n; i++ {
		if g.visited[i] {
			continue
		}
		// 0から開始
		g.dfs(i)
	}
	for i := 0; i < n; i++ {
		customIo.Printf("%d %d %d\n", i+1, g.foundTime[i], g.passTime[i])
	}

}

type Graph struct {
	edges     []Edges
	time      int
	visited   []bool
	foundTime []int
	passTime  []int
}

func NewGraph(n int) *Graph {
	return &Graph{
		edges:     make([]Edges, n),
		visited:   make([]bool, n),
		foundTime: make([]int, n),
		passTime:  make([]int, n),
	}
}

type Edges []Edge

func (e Edges) Len() int {
	return len(e)
}

func (e Edges) Less(i, j int) bool {
	return e[i].from < e[j].from
}

func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Edge struct {
	from int
	to   int
}

/*
	深さ優先探索(Depth First Search)
*/
func (g *Graph) dfs(i int) {
	g.visited[i] = true
	g.time++
	g.foundTime[i] = g.time
	for _, edge := range g.edges[i] {
		if g.visited[edge.to] {
			// 頂点iからアクセスできる頂点にアクセス済み
			continue
		}
		g.dfs(edge.to)
	}
	g.time++
	g.passTime[i] = g.time
}
