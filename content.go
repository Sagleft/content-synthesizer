package contentsynthesizer

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func CreateTextContent(client AI, prompt string) (Content, error) {
	data, err := client.GetPlainAnswer(prompt)
	if err != nil {
		return nil, fmt.Errorf("get answer: %w", err)
	}

	return NewTextContent(client, data), nil
}

func (c *textContent) String() string {
	return c.data
}

func (c *textContent) Join(ct Content, joint Prompt) (Content, error) {
	newPrompt := fmt.Sprintf(
		"%s: \n\ntext1:\n%s\n\ntext2:\n%s",
		joint, c.String(), ct.String(),
	)

	answer, err := c.client.GetPlainAnswer(newPrompt)
	if err != nil {
		return nil, fmt.Errorf("get answer: %w", err)
	}

	return NewTextContent(c.client, answer), nil
}

func (c *textContent) Modify(p Prompt) (Content, error) {
	newPrompt := fmt.Sprintf(
		`%s: \n\n"%s"`,
		p, c.String(),
	)

	answer, err := c.client.GetPlainAnswer(newPrompt)
	if err != nil {
		return nil, fmt.Errorf("get answer: %w", err)
	}

	return NewTextContent(c.client, answer), nil
}

func (c *textContent) CreateStructure(dataExample any, destPointer any) error {
	exampleBytes, err := json.MarshalIndent(
		dataExample,
		jsonMarshalPrefix,
		jsonMarshalIndent,
	)
	if err != nil {
		return fmt.Errorf("encode example: %w", err)
	}

	requestPrompt := fmt.Sprintf(
		promptGetJSON,
		c.String(), string(exampleBytes),
	)

	plainAnswer, err := c.client.GetPlainAnswer(requestPrompt)
	if err != nil {
		return fmt.Errorf("get answer: %w", err)
	}

	plainAnswer = fixJSON(plainAnswer)

	if err := json.Unmarshal([]byte(plainAnswer), destPointer); err != nil {
		return fmt.Errorf("decode answer from JSON: %w", err)
	}
	return nil
}

func (c *textContent) YesOrNo(q Question) (bool, error) {
	reqPrompt := fmt.Sprintf(
		promptGetBinary,
		c.String(), q,
	)

	response, err := c.client.GetPlainAnswer(reqPrompt)
	if err != nil {
		return false, fmt.Errorf("get answer: %w", err)
	}

	response = strings.Trim(response, ".")
	if response != responseYes && response != responseNo {
		return false, fmt.Errorf("invalid binary response: %q", response)
	}

	return response == responseYes, nil
}
