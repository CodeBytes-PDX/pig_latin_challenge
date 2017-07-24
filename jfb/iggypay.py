#!/usr/bin/env python

import re
import sys

sep = " "
string_in = sep.join(sys.argv[1:])

words_in = string_in.split()

string_out = []
for word in words_in:
    if re.search('[^a-z\']$', word):
        end_punct_start = re.search('[^a-z\']$', word.lower()).start()
    else:
        end_punct_start = -1
    n_consonants = re.search('[aeiou]', word.lower()).start()
    if n_consonants == 0:
        if end_punct_start >= 0:
            word_out = word[:end_punct_start] + "way" + word[end_punct_start:]
        else:
            word_out = word + "way"
    else:
        if word[0] >= 'A' and word[0] <= 'Z':
            cap = 1
        else:
            cap = 0
        if end_punct_start >= 0:
            word_out = word[n_consonants:end_punct_start]
        else:
            word_out = word[n_consonants:]
        word_out += word[:n_consonants] + "ay"
        if end_punct_start >= 0:
            word_out += word[end_punct_start:]

        if cap == 1:
            word_out = word_out.lower().capitalize()

    string_out.append(word_out)

print sep.join(string_out)
