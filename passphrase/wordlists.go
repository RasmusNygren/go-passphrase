package passphrase

import "embed"

//go:embed wordlists/*
var f embed.FS

func EffLarge() *WordList {
	wl := NewWordListCustomFS(f, "wordlists/eff_large_wordlist.txt", 5)
	return wl
}

func EffSmallShortWords() *WordList {
	wl := NewWordListCustomFS(f, "wordlists/eff_short_wordlist_1.txt", 4)
	return wl
}
