package exercise

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

type Adjective struct {
	word string
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

type Preposition struct {
	preposition string
	grammarCase string
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

var prepostitions = []Preposition{
	{"durch", Accusative},
	{"bis", Accusative},
	{"für", Accusative},
	{"ohne", Accusative},
	{"entlang", Accusative}, //along
	{"gegen", Accusative},
	{"um", Accusative}, //around
	{"mit", Dative},
	{"seit", Dative},
	{"zu", Dative},        //to
	{"aus", Dative},       //out
	{"außer", Dative},     //besides
	{"bei", Dative},       //next to
	{"nach", Dative},      //after
	{"seit", Dative},      //since
	{"von", Dative},       //from
	{"gegenüber", Dative}, //from
}

var DefiniteArticlesAdjectiveCaseEndings = Cases{
	[]string{
		"e", "e", "e", "en",
	},
	[]string{
		"en", "e", "e", "en",
	},
	[]string{
		"en", "en", "en", "en",
	},
	[]string{
		"en", "en", "en", "en",
	},
}

var IndefiniteArticlesAdjectiveCaseEndings = Cases{
	[]string{
		"er", "es", "e", "e",
	},
	[]string{
		"en", "es", "e", "e",
	},
	[]string{
		"en", "en", "en", "en",
	},
	[]string{
		"en", "en", "en", "er",
	},
}
