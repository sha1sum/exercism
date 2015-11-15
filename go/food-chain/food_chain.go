// +build !example

package foodchain

const TestVersion = 1

const firstLineOfEachStanza = "I know an old lady who swallowed a "
const lastLineOfEachStanza = "\nI don't know why she swallowed the fly. Perhaps she'll die."
const lastLines = "I know an old lady who swallowed a horse.\nShe's dead, of course!"

var animals = []string{
	"fly", "spider", "bird", "cat", "dog", "goat", "cow",
}

var secondLines = map[string]string{
	"cow":    "\nI don't know how she swallowed a cow!",
	"goat":   "\nJust opened her throat and swallowed a goat!",
	"dog":    "\nWhat a hog, to swallow a dog!",
	"cat":    "\nImagine that, to swallow a cat!",
	"bird":   "\nHow absurd to swallow a bird!",
	"spider": "\nIt wriggled and jiggled and tickled inside her.",
	"fly":    "",
}

func Song() string {
	return Verses(1, 2, 3, 4, 5, 6, 7, 8)
}

func Verses(verses ...int) string {
	output := ""
	for i := 1; i <= len(verses); i++ {
		output += Verse(i)
		if i != len(verses) {
			output += "\n\n"
		}
	}
	return output
}

func Verse(verse int) string {
	if verse == 8 {
		return lastLines
	}
	if verse < 1 {
		return ""
	}

	output := firstLineOfEachStanza + animals[verse-1] + "."
	output += secondLines[animals[verse-1]]
	for i := verse - 1; i > 0; i-- {
		if i > 0 {
			output += "\nShe swallowed the " + animals[i] + " to catch the " + animals[i-1]
		}
		if i == 2 {
			output += " that wriggled and jiggled and tickled inside her"
		}
		if i > 0 {
			output += "."
		}
	}
	return output + lastLineOfEachStanza
}
