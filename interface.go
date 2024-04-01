package contentsynthesizer

type Prompt string

type Question string

type AI interface {
	GetPlainAnswer(prompt string) (string, error)
}

type Content interface {
	// String - represent this content as a string
	String() string

	// Join this content with another based on a joint prompt
	Join(Content, Prompt) (Content, error)

	// Modify this content by creating a new one based on it
	Modify(Prompt) (Content, error)

	/* CreateStructure - create a structure from this content
	based on the example structure model and return the result*/
	CreateStructure(dataExample any, destPointer any) error

	// YesOrNo - ask a question based on this content and get an answer - Yes or No.
	YesOrNo(Question) (bool, error)
}
