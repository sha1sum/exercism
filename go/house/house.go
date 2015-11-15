package house

type Phrase struct {
	subject string
	action  string
}

// Define all lines' subjects and how they relate to their objects
var phrases = []Phrase{
	Phrase{subject: "horse and the hound and the horn", action: "belonged to"},
	Phrase{subject: "farmer sowing his corn", action: "kept"},
	Phrase{subject: "rooster that crowed in the morn", action: "woke"},
	Phrase{subject: "priest all shaven and shorn", action: "married"},
	Phrase{subject: "man all tattered and torn", action: "kissed"},
	Phrase{subject: "maiden all forlorn", action: "milked"},
	Phrase{subject: "cow with the crumpled horn", action: "tossed"},
	Phrase{subject: "dog", action: "worried"},
	Phrase{subject: "cat", action: "killed"},
	Phrase{subject: "rat", action: "ate"},
	Phrase{subject: "malt", action: "lay in"},
}

// Return a a phrase with the noun on the end.
func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

// Given a subject, a list of relational phrases, and a noun, make a string for a verse.
func Verse(subject string, relPhrases []string, nounPhrase string) string {
	return subject + " " + recursiveEmbed(relPhrases, nounPhrase)
}

// Recurse through the relational phrases, using the first as the subject to embed and the recursion as the rest of
// the string
func recursiveEmbed(relPhrases []string, nounPhrase string) string {
	if len(relPhrases) > 0 {
		return Embed(relPhrases[0], recursiveEmbed(relPhrases[1:], nounPhrase))
	}
	return nounPhrase
}

// Return the full song
func Song() string {
	// First line is simple enough
	song := "This is the house that Jack built."
	// An array to hold all of our relational phrases
	relPhrases := make([]string, 0)
	// Loop through the subjects and actions and make relational phrases out of them
	for i := 0; i < len(phrases); i++ {
		relPhrases = append(relPhrases, "the "+phrases[i].subject+"\nthat "+phrases[i].action)
	}
	// Move backwards through relPhrases and make a verse from the last n phrases
	for i := len(relPhrases) - 1; i >= 0; i-- {
		// Split between the verses
		song += "\n\n"
		song += Verse("This is", relPhrases[i:], "the house that Jack built.")
	}
	return song
}
