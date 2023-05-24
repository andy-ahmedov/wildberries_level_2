package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type sortTest struct {
	r, u, n       bool
	k             int
	data          []byte
	expected      string
	expectedError bool
}

var sortTests = []sortTest{
	{false, false, false, 0, []byte("Lizzie Blose lblose0@apache.org Female 29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59\nHarriet Brewitt hbrewitt2@go.com Female 31\nYurik Stockley ystockley3@themeforest.net Male 35\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nAlva McGaughey amcgaughey5@walmart.com Male 38\nRuggiero Newens rnewens6@histats.com Male 70\nShelden Rabbitt srabbitt7@census.gov Male 80\nLily Beatens lbeatens8@privacy.gov.au Female 95\nKonstance Gristock kgristock9@last.fm Female 81"), "Alva McGaughey amcgaughey5@walmart.com Male 38\nHarriet Brewitt hbrewitt2@go.com Female 31\nJerome Severs jsevers1@utexas.edu Male 59\nKonstance Gristock kgristock9@last.fm Female 81\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nLily Beatens lbeatens8@privacy.gov.au Female 95\nLizzie Blose ablose0@apache.org Female 29\nLizzie Blose lblose0@apache.org Female 29\nRuggiero Newens rnewens6@histats.com Male 70\nShelden Rabbitt srabbitt7@census.gov Male 80\nYurik Stockley ystockley3@themeforest.net Male 35\n", false},
	{true, false, false, 0, []byte("Lizzie Blose lblose0@apache.org Female 29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59\nHarriet Brewitt hbrewitt2@go.com Female 31\nYurik Stockley ystockley3@themeforest.net Male 35\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nAlva McGaughey amcgaughey5@walmart.com Male 38\nRuggiero Newens rnewens6@histats.com Male 70\nShelden Rabbitt srabbitt7@census.gov Male 80\nLily Beatens lbeatens8@privacy.gov.au Female 95\nKonstance Gristock kgristock9@last.fm Female 81"), "Yurik Stockley ystockley3@themeforest.net Male 35\nShelden Rabbitt srabbitt7@census.gov Male 80\nRuggiero Newens rnewens6@histats.com Male 70\nLizzie Blose ablose0@apache.org Female 29\nLizzie Blose lblose0@apache.org Female 29\nLily Beatens lbeatens8@privacy.gov.au Female 95\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nKonstance Gristock kgristock9@last.fm Female 81\nJerome Severs jsevers1@utexas.edu Male 59\nHarriet Brewitt hbrewitt2@go.com Female 31\nAlva McGaughey amcgaughey5@walmart.com Male 38\n", false},
	{false, true, false, 0, []byte("Lizzie Blose lblose0@apache.org Female 29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59\nHarriet Brewitt hbrewitt2@go.com Female 31\nYurik Stockley ystockley3@themeforest.net Male 35\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nAlva McGaughey amcgaughey5@walmart.com Male 38\nRuggiero Newens rnewens6@histats.com Male 70\nShelden Rabbitt srabbitt7@census.gov Male 80\nLily Beatens lbeatens8@privacy.gov.au Female 95\nKonstance Gristock kgristock9@last.fm Female 81"), "Alva McGaughey amcgaughey5@walmart.com Male 38\nHarriet Brewitt hbrewitt2@go.com Female 31\nJerome Severs jsevers1@utexas.edu Male 59\nKonstance Gristock kgristock9@last.fm Female 81\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nLily Beatens lbeatens8@privacy.gov.au Female 95\nLizzie Blose ablose0@apache.org Female 29\nLizzie Blose lblose0@apache.org Female 29\nRuggiero Newens rnewens6@histats.com Male 70\nShelden Rabbitt srabbitt7@census.gov Male 80\nYurik Stockley ystockley3@themeforest.net Male 35\n", false},
	{false, true, false, 4, []byte("Lizzie Blose lblose0@apache.org Female 29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59\nHarriet Brewitt hbrewitt2@go.com Female 31\nYurik Stockley ystockley3@themeforest.net Male 35\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nAlva McGaughey amcgaughey5@walmart.com Male 38\nRuggiero Newens rnewens6@histats.com Male 70\nShelden Rabbitt srabbitt7@census.gov Male 80\nLily Beatens lbeatens8@privacy.gov.au Female 95\nKonstance Gristock kgristock9@last.fm Female 81"), "Lizzie Blose ablose0@apache.org Female 29\nLizzie Blose lblose0@apache.org Female 29\nHarriet Brewitt hbrewitt2@go.com Female 31\nYurik Stockley ystockley3@themeforest.net Male 35\nAlva McGaughey amcgaughey5@walmart.com Male 38\nJerome Severs jsevers1@utexas.edu Male 59\nRuggiero Newens rnewens6@histats.com Male 70\nShelden Rabbitt srabbitt7@census.gov Male 80\nKonstance Gristock kgristock9@last.fm Female 81\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nLily Beatens lbeatens8@privacy.gov.au Female 95\n", false},
	{false, true, false, 5, []byte("Lizzie Blose lblose0@apache.org Female 29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59\nHarriet Brewitt hbrewitt2@go.com Female 31\nYurik Stockley ystockley3@themeforest.net Male 35\nLibbi Menlove lmenlove4@deliciousdays.com Polygender 86\nAlva McGaughey amcgaughey5@walmart.com Male 38\nRuggiero Newens rnewens6@histats.com Male 70\nShelden Rabbitt srabbitt7@census.gov Male 80\nLily Beatens lbeatens8@privacy.gov.au Female 95\nKonstance Gristock kgristock9@last.fm Female 81"), "", true},
}

func TestMySort(t *testing.T) {
	for _, test := range sortTests {
		output, err := MySort(test.data, test.r, test.n, test.u, test.k)
		if test.expectedError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.expected, output)
	}
}