package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
)

func mapReducer(mapper func(int) int, reducer func(int, int) int, initial int) func([]int) int {
    return func(args []int) int {
        for _, v := range args {
            initial = reducer(initial, mapper(v))
        }

        return initial
    }
}

func fuel(mass int) int {
    return mass / 3 - 2
}

func moduleTotal(mass int) int {
    total := 0
    mt := fuel(mass)

    for mt >= 0 {
        total += mt

        mt = fuel(mt)
    }

    return total
}

func readModules(fileName string) []int {
    var modules []int
    file, err := os.Open(fileName)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        mass, err := strconv.Atoi(scanner.Text())

        if err != nil {
            log.Fatal(err)
        }

        modules = append(modules, mass)
    }

    return modules
}

func main() {
    fuelCalc := mapReducer(
        func (mass int) int { return moduleTotal(mass) },
        func (a, v int) int { return a + v },
        0,
    )

    modules := readModules("input.txt")

    fmt.Println("Total fuel needed:", fuelCalc(modules))
}

