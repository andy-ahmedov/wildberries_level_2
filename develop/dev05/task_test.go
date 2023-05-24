package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type grepTest struct {
	A, B, C        int
	c, i, v, F, n  bool
	data, pattern  string
	expectedString string
	expectedError  error
}

var grepTests = []grepTest{
	{0, 0, 0, false, false, false, false, false, "hello\nworld\nworst enemy\nWor$$ friend\nWorst enemy\nJohn Smith\nSteve Jobs\nLizzie Allen\nlizzy allen", "Worst", "Worst enemy\n", nil},
	{0, 0, 0, false, true, false, false, false, "hello\nworld\nworst enemy\nWor$$ friend\nWorst enemy\nJohn Smith\nSteve Jobs\nLizzie Allen\nlizzy allen", "Worst", "worst enemy\nWorst enemy\n", nil},
	{2, 0, 0, true, true, false, false, true, "hello\nworld\nworst enemy\nWor$$ friend\nWorst enemy\nJohn Smith\nSteve Jobs\nLizzie Allen\nlizzy allen", "Worst", "found matches: 2\n2 worst enemy\n3 Wor$$ friend\n4 Worst enemy\n5 John Smith\n6 Steve Jobs\n", nil},
	{0, 0, 0, false, false, false, true, false, "hello\nworld\nworst enemy\nWor$$ friend\nWorst enemy\nJohn Smith\nWor.t enemy\nSteve Jobs\nLizzie Allen\nlizzy allen", "Wor.", "Wor.t enemy\n", nil},
	{0, 0, 0, true, true, true, false, true, "hello\nworld\nworst enemy\nWor$$ friend\nWorst enemy\nJohn Smith\nWor.t enemy\nSteve Jobs\nLizzie Allen\nlizzy allen", "Wor.", "found matches: 5\n0 hello\n5 John Smith\n7 Steve Jobs\n8 Lizzie Allen\n9 lizzy allen\n", nil},
}

func TestGrep(t *testing.T) {
	for _, test := range grepTests {
		output, err := grep(test.data, test.pattern, test.A, test.B, test.C, test.c, test.i, test.v, test.F, test.n)
		assert.Equal(t, test.expectedError, err)
		assert.Equal(t, test.expectedString, output)
	}
}
