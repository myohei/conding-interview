package main
import (
    "log"
)

func main() {
    targetStr := "abcdefghijklmnopqrstuxayz"
    if check(targetStr) {
        log.Print("ok")
    } else {
        log.Print("ng")
    }
}

func check(target string) bool {
    result := make([]bool, 256)
    for i := 0; i<len(target); i++ {
        point := target[i]
        if result[point] {
            return false
        }
        result[point] = true
    }
    return true
}
