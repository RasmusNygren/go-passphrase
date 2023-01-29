package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
	// "strconv"
)



func main() {

    wordList := make(map[string]string)
    readFile, err := os.Open("./wordlists/eff_large_wordlist.txt")
    if err != nil {
        panic(err)
    }
    defer readFile.Close()


    fileScanner := bufio.NewScanner(readFile)
    // fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        line := fileScanner.Text()
        splits := strings.Fields(line)
        key, val := splits[0], splits[1]
        wordList[key] = val
    }

    // Ensure this is actually a random and cryptographically safe way to generate the number
    for i := 0; i < 5; i++ {
        randint, err := rand.Int(rand.Reader, big.NewInt(6))
        if err != nil {
            panic(err)
        }
        // This is safe as the generated numbers are always in [1,6]
        fmt.Println(int(randint.Int64()) + 1)
    }

	// fmt.Println(wordList)
}
