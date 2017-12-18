package ll1

import "testing"

func TestParser(t *testing.T) {
	for _, test := range tests {
		if actual := Parse(test.input); actual != test.expected {
			t.Errorf("Parse(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}
