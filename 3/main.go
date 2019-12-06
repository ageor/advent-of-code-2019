package main

import (
    "strings"
    "strconv"
    "fmt"
    "os"
    "bufio"
    "log"
)

type Line struct {
    x0, y0, x1, y1, l int
}

type Point struct {
    x, y, l int
}

func atoi(str string) int {
    x, _ := strconv.Atoi(str)

    return x
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func readWires(fileName string) []string {
    var wires []string
    file, err := os.Open(fileName)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        wires = append(wires, scanner.Text())
    }

    return wires
}

func intersect(h, v Line) bool {
    return h.y0 > min(v.y0, v.y1) && h.y0 < max(v.y0, v.y1) && v.x0 > min(h.x0, h.x1) && v.x0 < max(h.x0, h.x1)
}

func cross(h0, v0, h1, v1 []Line) (crossings []Point) {
    for _, h := range h0 {
        for _, v := range v1 {
            if intersect(h, v) {
                p := Point{v.x0, h.y0, h.l + v.l}
                p.l += abs(p.x - h.x0)
                p.l += abs(p.y - v.y0)
                crossings = append(crossings, p)
            }
        }
    }

    for _, h := range h1 {
        for _, v := range v0 {
            if intersect(h, v) {
                p := Point{v.x0, h.y0, h.l + v.l}
                p.l += abs(p.x - h.x0)
                p.l += abs(p.y - v.y0)
                crossings = append(crossings, p)
            }
        }
    }

    return
}

func runWire(wire string) (horizontal []Line, vertical []Line) {
    x, y := 0, 0
    total := 0

    moves := strings.Split(wire, ",")

    for _, move := range moves {
        dir := move[0]
        rng := atoi(move[1:])

        switch dir {
        case 'R':
            line := Line{x, y, x + rng, y, total}
            x += rng
            horizontal = append(horizontal, line)
        case 'L':
            line := Line{x, y, x - rng, y, total}
            x -= rng
            horizontal = append(horizontal, line)
        case 'U':
            line := Line{x, y, x, y - rng, total}
            y -= rng
            vertical = append(vertical, line)
        case 'D':
            line := Line{x, y, x, y + rng, total}
            y += rng
            vertical = append(vertical, line)
        }

        total += rng
    }

    return
}

func main() {
    wires := readWires("input.txt")

    h0, v0 := runWire(wires[0])
    h1, v1 := runWire(wires[1])

    crossings := cross(h0, v0, h1, v1)

    minDist := 100000
    for _, c := range crossings {
        if c.x == 0 && c.y == 0 { continue }
        dist := abs(c.x) + abs(c.y)

        if dist < minDist {
            minDist = dist
        }
    }

    minPath := 1000000
    for _, c := range crossings {
        if c.x == 0 && c.y == 0 { continue }

        if c.l < minPath {
            minPath = c.l
        }
    }

    fmt.Println(minDist)
    fmt.Println(minPath)
}
