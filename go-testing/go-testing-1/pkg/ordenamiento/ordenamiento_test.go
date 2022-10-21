package ordering

import "testing"

func TestOrdenringGood(t *testing.T) {
	ordering(5, 4, 3, 2, 1)
}

func TestOrderingBad(t *testing.T) {
	ordering(0, 0, 0, -1)
}
