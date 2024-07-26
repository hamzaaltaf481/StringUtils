package main

import (
    "fmt"
    "io/ioutil"
	"time"
)
func Word_Spaces_Counter(s string){

	wordCounts := make(map[string]int)
	word := ""
	total_words := 0
	spaces := 0

	for _,char := range s {

		if char == ' '{
			spaces++
			if word != ""{
				wordCounts[word]++
				total_words ++
				word = ""
			} else {
				word+= string(char)
			}
		}
		
	}
	if word!= "" {
		wordCounts[word]++
		total_words++
	}

	fmt.Printf("Total Words: %v\nTotal Spaces: %v\n",total_words,spaces)

}
func Cpsmlt_Counter(s string){
	Cp := 0
	sm := 0
	for _,char := range s{

		if char >= 'A' && char <= 'Z'{
			Cp++
		}
		if char >= 'a' && char <= 'z'{
			sm++
		}
	}
	fmt.Printf("Capital Letter : %v\nSmall Letter: %v\n",Cp,sm)

}
func vowels_Counter(s string){

	vowel := 0
	for _,char := range s{

		if char == 'A' || char == 'a'|| char == 'E' || char == 'e' || char == 'I' || char == 'i' ||
		char == 'O' || char == 'o' || char == 'U' || char == 'u' {
			vowel++
		}

	}
	fmt.Printf("Total Vowels: %v\n", vowel)

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

    Word_Spaces_Counter(text)
	Cpsmlt_Counter(text)
	vowels_Counter(text)

	end := time.Now()
	Total_time := end.Sub(start)
	fmt.Printf("Execution time: %v\n", Total_time)
}