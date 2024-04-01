package contentsynthesizer

// special characters
const (
	specialBracketSymbol = "`"
	threeBrackets        = specialBracketSymbol + specialBracketSymbol + specialBracketSymbol
	responseYes          = "Yes"
	responseNo           = "No"
)

// prompts
const (
	promptGetArray = "The following data is available:\n\n%s\n\nPlease provide the data in CSV, each heading on a new line, without numbering, without commas at the end of each line"

	promptGetJSON = "The following data is available:\n\n%s\n\nPlease provide data in JSON, example:\n" + threeBrackets + "%s" + threeBrackets

	promptGetBinary = `The following data is available:\n\n%s\n\nMy question is:\n%s\n\nReturn Yes or No in your answer. Say "Yes" if the answer is yes, say "No" if the answer is no.`
)

// params
const (
	jsonMarshalPrefix = ""
	jsonMarshalIndent = "	"
)
