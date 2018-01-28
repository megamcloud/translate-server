package main

import (
	"encoding/gob"
	"io"
)

var phraseMap map[string]map[string]string

func init() {
	phraseMap = make(map[string]map[string]string)
}

// AddTranslation stores a translation in the memory cache
func AddTranslation(targetLang, sourcePhrase, targetPhrase string) {

	if phraseMap[sourcePhrase] == nil {
		phraseMap[sourcePhrase] = make(map[string]string)
	}

	phraseMap[sourcePhrase][targetLang] = targetPhrase
}

// GetTranslation returns a translation from cache, or nil, if it's not present
func GetTranslation(targetLang, sourcePhrase string) (targetPhrase *string) {
	if _, ok := phraseMap[sourcePhrase]; !ok {
		return nil
	}

	phrase, ok := phraseMap[sourcePhrase][targetLang]
	if !ok {
		return nil
	}

	return &phrase
}

// StoreCache writes the current phrase cache to the writer.
func StoreCache(writer io.Writer) {
	enc := gob.NewEncoder(writer)
	enc.Encode(phraseMap)
}

// LoadCache reads the phrase cache from the reader.
func LoadCache(reader io.Reader) {
	dec := gob.NewDecoder(reader)
	dec.Decode(phraseMap)
}