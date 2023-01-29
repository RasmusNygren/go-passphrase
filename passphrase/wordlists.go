package passphrase


func EffLarge() *WordList {
    wl := NewWordList("./wordlists/eff_large_wordlist.txt", 5)
    return wl
}

func EffSmallShortWords() *WordList {
    wl := NewWordList("./wordlists/eff_short_wordlist_1.txt", 4)
    return wl
}
