# Passphrase generator

This is a minimal passphrase generator based on the [diceware](https://theworld.com/~reinhold/diceware.html) method.

## Usage
### Using default wordlists
```go
import (
    "github.com/RasmusNygren/go-passphrase/passphrase"
)
// Generate a passphrase consisting of 4 words from the EffSmallShortWords wordlist
wordlist := passphrase.EffSmallShortWords()
phrase := passphrase.GeneratePhrase(wordlist, numWords)
```

### Using custom wordslists
Given a text file with whitespace-seperated columns of `id` to word. See example [Short wordlist](./passphrase/wordlists/eff_short_wordlist_1.txt)

```go
import (
    "github.com/RasmusNygren/go-passphrase/passphrase"
)
// digits is the length of the id-numbers
wordlist := passphrase.NewWordList(path, digits)
phrase := passphrase.GeneratePhrase(wordlist, numWords)
```
