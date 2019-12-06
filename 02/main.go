package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
    "io/ioutil"
)

func parseData(data string) []int {
    s := strings.Split(data, ",")
    var res []int

    for _, x := range s {
        i, err := strconv.Atoi(x)

        if err != nil {
            log.Fatal(err)
        }

        res = append(res, i)
    }

    return res
}

func execute(memory []int) {
    address := 0

    next := func() int {
        defer func() { address += 1 }()
        return address
    }

    for address < len(memory) {
        instruction := memory[next()]

        if instruction == 99 {
            break
        }

        a := memory[next()]
        b := memory[next()]
        res := memory[next()]

        switch instruction {
        case 1:
            memory[res] = memory[a] + memory[b]
        case 2:
            memory[res] = memory[a] * memory[b]
        }
    }
}

func bootProgram(fileName string, first, second int) []int {
    content, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }

    instructions := parseData(strings.TrimSpace(string(content)))
    instructions[1] = first
    instructions[2] = second

    return instructions
}

func main() {
    instructions := bootProgram("input.txt", 12, 2)
    execute(instructions)
    fmt.Println(instructions[0])

    for x := 0; x < 100; x++ {
        for y := 0; y < 100; y++ {
            instructions = bootProgram("input.txt", x, y)
            execute(instructions)

            if instructions[0] == 19690720 {
                fmt.Println(100 * x + y)
                return
            }
        }
    }
}

