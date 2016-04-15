package main

// Enumerate returns all the possibles permutations of the given alphabet for
// words of the given length. The permutations are enumerated in a string
// channel to prevent storing a big slice of strings.
func Enumerate(length uint, alphabet []rune) <-chan string {
	// Here is the exit case, when the length is 1 we just have to send the
	// alphabet's chars.
	if length == 1 {
		var out = make(chan string)
		go func() {
			defer close(out)
			for _, r := range alphabet {
				out <- string(r)
			}
		}()
		return out
	}

	// Otherwise, enumerate the permutations for the length-1 words and for
	// each of them add each rune of the alphabet.
	var out = make(chan string)
	var in = Enumerate(length-1, alphabet)
	for p := range in {
		for _, r := range alphabet {
			out <- p + string(r)
		}
	}
	return out
}
