package passphrase

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// TODO: Better error handling


type WordList struct {
    digits int 
    table map[string]string
}

func NewWordList(path string, digits int) *WordList {
    table := make(map[string]string)
    readFile, err := os.Open(path)
    if err != nil {
        panic(err)
    }

    // TODO: Implement error handling. This should not be deferred without
    // error handling, must be checked explicitly.
    defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)

    for fileScanner.Scan() {
        line := fileScanner.Text()
        splits := strings.Fields(line)
        key, val := splits[0], splits[1]
        table[key] = val
    }

    wordList := WordList{
        digits: digits,
        table: table,
    }
    return &wordList
}

func generateWord(wl *WordList) string {
    wordList := wl.table
    digits := wl.digits

    var stringBuilder strings.Builder
    for i := 0; i < digits; i++ {
        // This assumes a 6-sided die.
        randint, err := rand.Int(rand.Reader, big.NewInt(6))
        if err != nil {
            panic(err)
        }

        // This is safe as the generated numbers are always in [1,6]
        randint = randint.Add(randint, big.NewInt(1))
        stringBuilder.WriteString(randint.String())
    }

    randomNumberAsStr := stringBuilder.String()
    // fmt.Println(randomNumberAsStr)
    // fmt.Println(wordList[randomNumberAsStr])
    word := wordList[randomNumberAsStr]
    return word
}

func GeneratePhraseWithCustomDelimiter(wl *WordList, length int, delimiter string) string {
    var sb strings.Builder
    for i := 0; i < length; i++ {
        word := generateWord(wl)
        sb.WriteString(word)
        if i < length - 1 {
            sb.WriteString(delimiter)
        }
    }
    return sb.String()
}

func GeneratePhrase(wl *WordList, length int) string {
    return GeneratePhraseWithCustomDelimiter(wl, length, "-")
}

func main() {
    // Ensure this is actually a random and cryptographically safe way to generate the number

    // wl := EffLarge()
    wl := EffSmallShortWords()
    phrase := GeneratePhrase(wl, 3)
    fmt.Println(phrase)
}
