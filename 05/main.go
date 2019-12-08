package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
    "io/ioutil"
    "os"
)

func atoi(str string) int {
    x, err := strconv.Atoi(str)

    if err != nil {
        log.Fatal(err)
    }

    return x
}

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

func parseInstruction(instruction int) [4]int {
    return [4]int{0, (instruction / 1000) % 10, (instruction / 100) % 10, instruction % 100}
}

func execute(memory []int, args []string) {
    address := 0
    argi := 0

    next := func() int {
        defer func() { address += 1 }()
        return address
    }

    ab := func(instruction [4]int) (int, int) {
        a, b := memory[next()], memory[next()]

        if instruction[2] == 0 { a = memory[a] }
        if instruction[1] == 0 { b = memory[b] }

        return a, b
    }

    for address < len(memory) {
        instruction := parseInstruction(memory[next()])
        op := instruction[3]

        if (op == 99) {
            break;
        }

        switch op {
        case 1:
            a, b := ab(instruction)
            res := memory[next()]

            memory[res] = a + b
        case 2:
            a, b := ab(instruction)
            res := memory[next()]

            memory[res] = a * b
        case 3:
            if argi == len(args) {
                log.Fatal("Not enough arguments!")
            }

            input := atoi(args[argi])
            argi += 1

            res := memory[next()]
            memory[res] = input
        case 4:
            res := memory[next()]

            if instruction[2] == 0 {
                fmt.Println(memory[res])
            } else {
                fmt.Println(res)
            }

        case 5:
            a, b := ab(instruction)

            if a != 0 {
                address = b
            }
        case 6:
            a, b := ab(instruction)

            if a == 0 {
                address = b
            }
        case 7:
            a, b := ab(instruction)
            res := memory[next()]

            if a < b {
                memory[res] = 1
            } else {
                memory[res] = 0
            }
        case 8:
            a, b := ab(instruction)
            res := memory[next()]

            if a == b {
                memory[res] = 1
            } else {
                memory[res] = 0
            }
        }
    }
}

func bootProgram(fileName string) []int {
    content, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }

    instructions := parseData(strings.TrimSpace(string(content)))

    return instructions
}

func main() {
    instructions := bootProgram("input.txt")
    execute(instructions, os.Args[1:])
}

