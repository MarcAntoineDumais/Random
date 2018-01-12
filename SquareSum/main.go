package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    for i := 2; i < 50; i++ {
        t := time.Now()
        g := createGraph(i)
        solution := solve(g)
        timeTrack(t, fmt.Sprintf("Iteration %d", i))
        if solution == "" {
            fmt.Println("No solution")
        } else {
            fmt.Println(solution)
        }
    }
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}

type graph struct {
    n int
    e [][]bool
}

func createGraph(n int) graph {
    var g graph
    g.n = n
    g.e = make([][]bool, n)
    for i := 0; i < n; i++ {
        g.e[i] = make([]bool, n)
    }
    
    squares := make(map[int]bool)
    max := int(math.Ceil(math.Sqrt(float64(n*2))))
    for i := 1; i < max; i++ {
        squares[i*i] = true
    }
    
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if squares[i+j+2] {
                g.e[i][j] = true
                g.e[j][i] = true
            }
        }
    }
    
    return g
}

func solve(g graph) string {
    var recurse func(graph, []int) []int
    recurse = func(g graph, ans []int) []int {
        if len(ans) == g.n {
            return ans
        }
        g2 := make([][]bool, g.n)
        for i := 0; i < g.n; i++ {
            g2[i] = make([]bool, g.n)
            for j := 0; j < g.n; j++ {
                g2[i][j] = g.e[i][j]
            }
        }
        ans2 := make([]int, len(ans))
        copy(ans2, ans)
        //fmt.Println(g.e, ans2)
        
        last := ans[len(ans)-1]
        //fmt.Println("last ", last)
        edge, found := findEdge(g, last)
        for found {
            //fmt.Println("found edge ", edge)
            back := make([]bool, g.n * 2)
            for i := 0; i < g.n; i++ {
                back[i] = g.e[last][i]
                back[g.n + i] = g.e[i][last]
                g.e[last][i] = false
                g.e[i][last] = false
            }
            ans2 = append(ans2, edge)
            solution := recurse(g, ans2)
            if solution != nil {
                return solution
            }
            for i := 0; i < g.n; i++ {
                g.e[last][i] = back[i]
                g.e[i][last] = back[g.n + i]
            }
            g.e[last][edge] = false
            g.e[edge][last] = false
            edge, found = findEdge(g, last)
        }
        for i := 0; i < g.n; i++ {
            for j := 0; j < g.n; j++ {
                g.e[i][j] = g2[i][j]
            }
        }
        return nil
    }
    
    for i := 0; i < g.n; i++ {
        tmp := make([]int, 1)
        tmp[0] = i
        ans := recurse(g, tmp[:])
        if ans != nil {
            for j := 0; j < len(ans); j++ {
                ans[j]++
            }
            return fmt.Sprintf("%v", ans)
        }
    }
    
    return ""
}

func findEdge(g graph, v int) (int, bool) {
    if v >= g.n {
        return 0, false
    }
    for i := 0; i < g.n; i++ {
        if i != v {
            if g.e[i][v] {
                return i, true
            }
        }
    }
    return 0, false
}