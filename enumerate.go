package main

import (
	"math"
	"math/rand"
	"time"
)

// Enumerate returns random permutations of the given alphabet for words of the
// given length. The permutations are enumerated in a string channel.
func Enumerate(length uint, alphabet string) <-chan string {
	var out = make(chan string)

	// Here is an optimized loop that generate random strings for the given
	// length and alphabet.
	go func() {
		// The random source will be used to get random bits.
		var src = rand.NewSource(time.Now().UnixNano())

		// Caching the size of the alphabet for performance reasons.
		var alphabetLength = len(alphabet)

		// maskLength is the number of bits necessary to represent the number
		// of letters in the alphabet, used to compute the size of the mask to
		// apply to the random numbers generated.
		var maskLength = uint(math.Log2(float64(alphabetLength)))

		// The mask is used to take only the necessary part of the randomly
		// generated number in a way more efficient and statistically valid
		// than doing a modulo.
		var mask = int64(1<<maskLength - 1)

		// poolSize is the number of mask that can be contained in a 63bits
		// random integer, that is the number of times the same integer can be
		// used to generate an random index.
		var poolSize = 63 / maskLength

		// Generate strings indefinitely.
		var b = make([]byte, length)
		for {
			var pool = int64(0)
			var remainder = uint(0)
			for i := uint(0); i < length; {
				// If there is no remainder, regenerate the pool.
				if remainder == 0 {
					pool, remainder = rand.Int63(), poolSize
				}

				// Generate a random int and use the lower bits of the pool to
				// get a random rune to add to the string.
				if idx := int(src.Int63() & mask); idx < alphabetLength {
					b[i] = alphabet[idx]
					i++
				}

				// Remove the used part of the pool and decrement the remainder.
				pool >>= maskLength
				remainder--
			}

			out <- string(b)
		}
	}()

	return out
}
