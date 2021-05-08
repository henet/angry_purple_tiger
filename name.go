package angry_purple_tiger

import (
	"crypto/md5"
)

// Name returns the "angry purple tiger" Helium hotspot name for the given input.
func Name(input string) string {
	d := md5.Sum([]byte(input))
	return adjectives[d[0]^d[1]^d[2]^d[3]^d[4]] + "-" +
		colors[d[5]^d[6]^d[7]^d[8]^d[9]] + "-" +
		animals[d[10]^d[11]^d[12]^d[13]^d[14]^d[15]]
}
