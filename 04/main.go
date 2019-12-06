package main

import (
    "fmt"
)

func norm(a [6]int) bool {
    for i := 0; i < len(a) - 1; i++ {
        if a[i] > a[i + 1] { return false }
    }

    return true
}

func doubles(a [6]int) bool {
    m := make(map[int]int)

    for i := 0; i < len(a); i++ {
        m[a[i]] += 1
    }

    for _, v := range m {
        if v == 2 { return true }
    }

    return false
}

func less(a, b [6]int) bool {
    for i := 0; i < len(a); i++ {
        if (a[i] < b[i]) { return true }
        if (a[i] > b[i]) { return false }
    }

    return false
}

func seq(a [6]int) bool {
    for i := 0; i < len(a) - 1; i++ {
        if a[i] == a[i + 1] { return true }
    }

    return false
}

func inc(a [6]int) [6]int {
    a[5] += 1

    for i := len(a) - 1; i >= 0; i-- {
        if a[i] > 9 && i > 0 {
            a[i - 1] += 1
            a[i] = 0
        }
    }

    return a
}

func main() {
    start, end := [6]int{1, 5, 2, 0, 8, 5}, [6]int{6, 7, 0, 2, 8, 3}

    counter := 0

    valid := func(pass [6]int) bool {
        return norm(pass) && seq(pass) && doubles(pass)
    }

    if valid(end) { counter += 1 }

    for less(start, end) {
        if valid(start) { counter += 1 }

        start = inc(start)
    }

    fmt.Println(counter)
}

