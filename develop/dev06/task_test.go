package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCut struct {
	f, d           string
	s              bool
	data           string
	expectedString string
	expectedError  error
}

var testCuts = []testCut{
	{"0, 2,3", " ", false, "Lizzie_Blose_lblose0@apache.org_Female_29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59", "Lizzie_Blose_lblose0@apache.org_Female_29\nLizzie ablose0@apache.org Female\nJerome jsevers1@utexas.edu Male\n", nil},
	{"1,2", " ", false, "Lizzie_Blose_lblose0@apache.org_Female_29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59", "Lizzie_Blose_lblose0@apache.org_Female_29\nBlose ablose0@apache.org\nSevers jsevers1@utexas.edu\n", nil},
	{"1,2", "_", true, "Lizzie_Blose_lblose0@apache.org_Female_29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59", "Blose lblose0@apache.org\n", nil},
}

func TestCut(t *testing.T) {
	for _, test := range testCuts {
		output, err := Cut(test.data, test.f, test.d, test.s)
		assert.Equal(t, test.expectedError, err)
		assert.Equal(t, test.expectedString, output)
	}
}