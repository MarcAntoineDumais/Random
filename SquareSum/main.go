package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    for i := 2; i < 100; i++ {
        g := createGraph(i)
        t := time.Now()
        solution := solve(g)
        timeTrack(t, fmt.Sprintf("Solving problem %d", i))
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

// Solution #3 from https://www.hackerearth.com/practice/algorithms/graphs/hamiltonian-path/tutorial/
func solve(g graph) string {
    visited := make([]bool, g.n)
    stack := make([]int, 0, g.n)
    for i := 0; i < g.n; i++ {
        visited[i] = true
        stack = append(stack, i)
        solution, found := dfs(g, i, stack, visited)
        if found {
            for i, _ := range solution {
                solution[i]++
            }
            return fmt.Sprintf("%v", solution)
        }
        visited[i] = false
        stack = stack[:0]
    }
    return "";
}

func dfs(g graph, v int, stack []int, visited []bool) ([]int, bool) {
    if len(stack) == g.n {
        return stack, true
    }
    for i := 0; i < g.n; i++ {
        if g.e[v][i] && !visited[i] {
            visited[i] = true
            stack = append(stack, i)
            solution, found := dfs(g, i, stack, visited)
            if found {
                return solution, true
            }
            visited[i] = false
            stack = stack[:len(stack)-1]
        }
    }
    return nil, false
}