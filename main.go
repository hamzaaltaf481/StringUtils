package main

import (
    "fmt"
    "io/ioutil"
    "sync"
    "time"
    "strings"
)

type Results struct {
    TotalWords  int
    Spaces      int
    Capitals    int
    Small       int
    Vowels      int
}

func Word_Spaces_Counter(s string, wg *sync.WaitGroup, res *Results, ch chan *Results) {
    defer wg.Done()

    total_words := 0
    spaces := 0
    Cp := 0
    sm := 0
    vowel := 0

    a:= strings.Split(s," ")
    total_words = len(a)

    for _, char := range s {
        if char == ' ' {
            spaces++
        } 
        if char >= 'A' && char <= 'Z' {
            Cp++
        }
        if char >= 'a' && char <= 'z' {
            sm++
        }
        if char == 'A' || char == 'a' || char == 'E' || char == 'e' || char == 'I' || char == 'i' ||
            char == 'O' || char == 'o' || char == 'U' || char == 'u' {
            vowel++
        }
    }

    res.TotalWords = total_words
    res.Spaces = spaces
    res.Capitals = Cp
    res.Small = sm
    res.Vowels = vowel

    ch <- res
}

func main() {
    filePath := "./Test.txt"
    start := time.Now()

    fileContent, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    text := string(fileContent)
    var wg sync.WaitGroup
    wg.Add(1)

    ch := make(chan *Results, 3)
    res := &Results{}

    go Word_Spaces_Counter(text[:len(text)], &wg, res, ch)

        <-ch


    wg.Wait()

    fmt.Printf("Total Words: %v\nTotal Spaces: %v\nCapital Letters: %v\nSmall Letters: %v\nTotal Vowels: %v\n", res.TotalWords, res.Spaces, res.Capitals, res.Small, res.Vowels)

    end := time.Now()
    total_time := end.Sub(start)
    fmt.Printf("Execution time: %v\n", total_time)
}
