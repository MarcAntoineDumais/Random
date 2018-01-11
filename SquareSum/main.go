package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    for i := 1; i < 10; i++ {
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
    return fmt.Sprintf("%v", g)
}