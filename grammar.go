package main

const Masculine = 0
const Neuter = 1
const Feminine = 2
const Plural = 3

const Nominative = "nominative"
const Accusative = "accusative"
const Dative = "dative"

type Noun struct {
	word   string
	gender int8
}

type Cases struct {
	nominative []string
	accusative []string
	dative     []string
	genitive   []string
}

type Case struct {
	masculine string
	neutral   string
	feminine  string
	plural    string
}

var DefiniteArticlesCases = Cases{
	[]string{
		"der", "das", "die", "die",
	},
	[]string{
		"den", "das", "die", "die",
	},
	[]string{
		"dem", "dem", "der", "den",
	},
	[]string{
		"des", "des", "der", "der",
	},
}

var IndefiniteArticlesCases = Cases{
	[]string{
		"ein", "ein", "eine",
	},
	[]string{
		"einen", "ein", "eine",
	},
	[]string{
		"einem", "einem", "einer",
	},
	[]string{
		"eines", "eines", "einer",
	},
}

var articles = []Cases{
	DefiniteArticlesCases,
	IndefiniteArticlesCases,
}
