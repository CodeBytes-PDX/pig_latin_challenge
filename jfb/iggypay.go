package main

import (
    "fmt"
    "os"
    "regexp"
    s "strings"
)

func main() {

    var end_punct_start int
    var cap_first bool
    var match_loc []int
    var word_out string

    /* combine all args into a single string and then split on spaces. this
    ** normalizes things in case a/the command-line paramater has a space
    ** in it.
    */
    string_in := s.Join(os.Args[1:], " ")
    words_in := s.Split(string_in, " ")

    words_out := make([]string, len(words_in))

    re_nonword_end := regexp.MustCompile(`[^a-z\']$`);
    re_vowel := regexp.MustCompile(`[aeiou]`);

    for i, word := range words_in {
	match_loc = re_nonword_end.FindIndex([]byte(s.ToLower(word)))
	if match_loc == nil {
	    end_punct_start = -1
	} else {
	    end_punct_start = match_loc[0]
	}
	n_consonants_head := re_vowel.FindIndex([]byte(s.ToLower(word)))[0]
	if n_consonants_head == 0 {
	    if end_punct_start >= 0 {
		word_out = word[:end_punct_start] + "way" + word[end_punct_start:]
	    } else {
		word_out = word + "way"
	    }
	} else {
	    if word[0] >= 'A' && word[0] <= 'Z' {
		cap_first = true
	    } else {
		cap_first = false
	    }
	    if end_punct_start >= 0 {
		word_out = word[n_consonants_head:end_punct_start]
	    } else {
		word_out = word[n_consonants_head:]
	    }
	    word_out += word[:n_consonants_head] + "ay"
	    if end_punct_start >= 0 {
		word_out += word[end_punct_start:]
	    }

	    if cap_first {
		word_out = s.ToUpper(string(word_out[0])) + s.ToLower(word_out[1:])
	    }
	}
	words_out[i] = word_out
    }

fmt.Println(s.Join(words_out, " "))

}
