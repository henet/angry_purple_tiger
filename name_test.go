package angry_purple_tiger_test

import (
	"fmt"
	"testing"

	"github.com/henet/angry_purple_tiger"
)

func TestName(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{"112CuoXo7WCcp6GGwDNBo6H5nKXGH45UNJ39iEefdv2mwmnwdFt8", "feisty-glass-dalmatian"},
	}

	for i, test := range tests {
		output := angry_purple_tiger.Name(test.Input)
		if test.Expected != output {
			t.Errorf("%d) input %s, want %s, got %s", i, test.Input, test.Expected, output)
		}
	}
}

func ExampleName() {
	fmt.Println(angry_purple_tiger.Name("112CuoXo7WCcp6GGwDNBo6H5nKXGH45UNJ39iEefdv2mwmnwdFt8"))
	// Output: feisty-glass-dalmatian
}
