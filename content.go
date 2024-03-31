package contentsynthesizer

type textContent struct {
	client AI
	data   string
}

func NewTextContent(client AI, data string) Content {
	return &textContent{
		client: client,
		data:   data,
	}
}

func (c *textContent) String() string {
	return c.data
}

func (c *textContent) Join(Content, Prompt) (Content, error) {
	// TODO
	return nil, nil
}

func (c *textContent) Modify(Prompt) (Content, error) {
	// TODO
	return nil, nil
}

func (c *textContent) CreateStructure(dataExample any, destPointer any) error {
	// TODO
	return nil
}
