package contentsynthesizer

type Prompt string

type AI interface {
	GetPlainAnswer(Prompt) (string, error)
}

type Content interface {
	String() string
	Join(Content, Prompt) (Content, error)
	Modify(Prompt) (Content, error)
	CreateStructure(dataExample any, destPointer any) error
}
