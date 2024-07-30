package main

import (
    "fmt"
    "io/ioutil"
    "sync"
    "time"
)

type Results struct {
    TotalWords  int
    Spaces      int
    Capitals    int
    Small       int
    Vowels      int
}

func Word_Spaces_Counter(s string, wg *sync.WaitGroup, ch chan *Results) {
    defer wg.Done()
    res := &Results{}

    for _, char := range s {
        if char == ' ' {
            res.Spaces++
        } 
        if char >= 'A' && char <= 'Z' {
            res.Capitals++
        }
        if char >= 'a' && char <= 'z' {
            res.Small++
        }
        if char == 'A' || char == 'a' || char == 'E' || char == 'e' || char == 'I' || char == 'i' ||
            char == 'O' || char == 'o' || char == 'U' || char == 'u' {
            res.Vowels++
        }
    }
    res.TotalWords = res.Spaces + 1   

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

    ch := make(chan *Results)
    numChunk := 4
    chunkSize := len(text)/numChunk    

    for i := 0; i < numChunk ; i++ {
        start := i * chunkSize
        end := start + chunkSize
        if end == len(text){
            end = len(text) +1
        }
        chunk := text[start:end]

        wg.Add(1)
        go Word_Spaces_Counter(chunk, &wg, ch)
        
    }
    go func() {
        wg.Wait()
        close(ch)
    }()

    finalRes := &Results{}
        for res := range ch{
        finalRes.TotalWords += res.TotalWords
        finalRes.Spaces += res.Spaces
        finalRes.Capitals += res.Capitals
        finalRes.Small += res.Small
        finalRes.Vowels += res.Vowels
    }

    


    fmt.Printf("Total Words: %v\nTotal Spaces: %v\nCapital Letters: %v\nSmall Letters: %v\nTotal Vowels: %v\n", finalRes.TotalWords, finalRes.Spaces, finalRes.Capitals, finalRes.Small, finalRes.Vowels)

    end := time.Now()
    total_time := end.Sub(start)
    fmt.Printf("Execution time: %v\n", total_time)
}
