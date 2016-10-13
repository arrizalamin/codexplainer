package context

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type ExplainerFunc func(*Context)

type Context struct {
	filename  string
	Input     []string
	Explainer ExplainerFunc
	BlockHeap BlockHeap

	Mu     sync.Mutex
	Output []string
}

func NewContext(filename string, source string, explainer ExplainerFunc) Context {
	input := TrimInput(strings.Split(source, "\n"))

	return Context{
		filename:  filename,
		Input:     input,
		Output:    make([]string, len(input)),
		Explainer: explainer,
		BlockHeap: BlockHeap{},
	}
}

func TrimInput(input []string) []string {
	for i, line := range input {
		input[i] = strings.TrimSpace(line)
	}
	return input
}

func (c *Context) Execute() {
	c.Explainer(c)
}

func (c *Context) PrintOutput() {
	f, err := os.Create(c.filename + ".txt")
	if err != nil {
		panic(err)
	}
	for i, line := range c.Output {
		f.WriteString(fmt.Sprintf("%d. %s\n", (i + 1), line))
	}
	f.Sync()
}
