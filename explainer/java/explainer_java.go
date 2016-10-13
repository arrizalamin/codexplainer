package java

import (
	"fmt"
	"regexp"

	c "github.com/arrizalamin/codexplainer/context"
	e "github.com/arrizalamin/codexplainer/explainer"
)

func JavaExplainer(ctx *c.Context) {
	e.ExecuteExplainers(ctx)(
		e.Concurrent(
			ExplainPackageDeclaration,
			ExplainImportStatement,
			ExplainMethodDeclaration,
			ExplainBlankLine,
			ExplainConditional,
			ExplainVariableDeclaration,
			ExplainPrint,
			ExplainDoWhileLoop,
			ExplainForLoop,
			ExplainIncrementAndDecrement,
		),
		ExplainClosingBlock,
	)
}

func ExplainPackageDeclaration(ctx *c.Context) {
	reg := regexp.MustCompile("^package\\s+(.+);$")
	for i, line := range ctx.Input {
		if reg.MatchString(line) {
			ctx.Mu.Lock()
			ctx.Output[i] = "Deklarasi package " + reg.FindStringSubmatch(line)[1]
			ctx.Mu.Unlock()
		}
	}
}

func ExplainImportStatement(ctx *c.Context) {
	reg := regexp.MustCompile("^import\\s+(.+);$")
	for i, line := range ctx.Input {
		if reg.MatchString(line) {
			ctx.Mu.Lock()
			ctx.Output[i] = "Melakukan import package " + reg.FindStringSubmatch(line)[1]
			ctx.Mu.Unlock()
		}
	}
}

func ExplainMethodDeclaration(ctx *c.Context) {
	reg := regexp.MustCompile(
		"^(public|protected|private)\\s+(static?)\\s*(.+)\\s+(.+)\\((.+)\\)\\s\\{$",
	)
	for i, line := range ctx.Input {
		if reg.MatchString(line) {
			ctx.Mu.Lock()
			match := reg.FindStringSubmatch(line)
			ctx.BlockHeap.Push("method "+match[4], i)
			ctx.Output[i] = fmt.Sprintf(
				"Deklarasi method %s dengan kembalian %s dan parameter (%s)",
				match[4],
				match[3],
				match[5],
			)
			ctx.Mu.Unlock()
		}
	}
}

func ExplainBlankLine(ctx *c.Context) {
	for i, line := range ctx.Input {
		if line == "" {
			ctx.Mu.Lock()
			ctx.Output[i] = "Baris Kosong"
			ctx.Mu.Unlock()
		}
	}
}

func ExplainClosingBlock(ctx *c.Context) {
	for i, line := range ctx.Input {
		if line == "}" {
			ctx.Mu.Lock()
			ctx.Output[i] = "Penutup " + ctx.BlockHeap.Pop()
			ctx.Mu.Unlock()
		}
	}
}

func ExplainVariableDeclaration(ctx *c.Context) {
	reg := regexp.MustCompile("^(?:([A-Za-z0-9_\\[\\]]+)\\s+)?(\\w+)\\s*\\=\\s*(\\S.*);$")
	for i, line := range ctx.Input {
		if reg.MatchString(line) {
			ctx.Mu.Lock()
			matches := reg.FindStringSubmatch(line)
			ctx.Output[i] = fmt.Sprintf(
				"Set variabel %s dengan nilai %s",
				matches[2],
				matches[3],
			)
			ctx.Mu.Unlock()
		}
	}
}

func ExplainConditional(ctx *c.Context) {
	reg := regexp.MustCompile("^if\\s+\\((.+)\\)\\s+\\{?$")
	for i, line := range ctx.Input {
		if reg.MatchString(line) {
			ctx.Mu.Lock()
			match := reg.FindStringSubmatch(line)
			ctx.BlockHeap.Push("conditional", i)
			ctx.Output[i] = "Melakukan cek apakah " + match[1] + " bernilai true"
			ctx.Mu.Unlock()
		}
	}
}

func ExplainPrint(ctx *c.Context) {
	reg := regexp.MustCompile("^System.out.print(?:ln)?\\((.+)\\);$")
	for i, line := range ctx.Input {
		if reg.MatchString(line) {
			ctx.Mu.Lock()
			ctx.Output[i] = "Mencetak " + reg.FindStringSubmatch(line)[1]
			ctx.Mu.Unlock()
		}
	}
}

func ExplainDoWhileLoop(ctx *c.Context) {
	doRegex := regexp.MustCompile("^do\\s+\\{$")
	whileRegex := regexp.MustCompile("\\}\\s*while\\s*\\((.+)\\);$")
	for i, line := range ctx.Input {
		if doRegex.MatchString(line) {
			ctx.Mu.Lock()
			ctx.Output[i] = "Memulai looping do...while"
			ctx.Mu.Unlock()
		} else if whileRegex.MatchString(line) {
			ctx.Mu.Lock()
			matches := whileRegex.FindStringSubmatch(line)
			ctx.Output[i] = "Kembali ke statement do jika " + matches[1] + " = true"
			ctx.Mu.Unlock()
		}
	}
}

func ExplainForLoop(ctx *c.Context) {
	regularLoop := regexp.MustCompile("^for\\s*\\((.+;.+;.+)\\)(?:\\s*\\{)?$")
	forEachLoop := regexp.MustCompile("^for\\s*\\(\\w+\\s*\\w*\\s*\\:\\s*(\\w+)\\)(?:\\s*\\{)?$")
	for i, line := range ctx.Input {
		if regularLoop.MatchString(line) {
			ctx.Mu.Lock()
			ctx.BlockHeap.Push("looping", i)
			ctx.Output[i] = "Looping " + regularLoop.FindStringSubmatch(line)[1]
			ctx.Mu.Unlock()
		} else if forEachLoop.MatchString(line) {
			ctx.Mu.Lock()
			ctx.BlockHeap.Push("looping", i)
			ctx.Output[i] = "Looping sebanyak jumlah " + forEachLoop.FindStringSubmatch(line)[1]
			ctx.Mu.Unlock()
		}
	}
}

func ExplainIncrementAndDecrement(ctx *c.Context) {
	inc := regexp.MustCompile("^(\\w+)\\+\\+;$")
	dec := regexp.MustCompile("^(\\w+)--;$")
	for i, line := range ctx.Input {
		if inc.MatchString(line) {
			ctx.Mu.Lock()
			ctx.Output[i] = "Increment variable " + inc.FindStringSubmatch(line)[1]
			ctx.Mu.Unlock()
		} else if dec.MatchString(line) {
			ctx.Mu.Lock()
			ctx.Output[i] = "Decrement variable " + dec.FindStringSubmatch(line)[1]
			ctx.Mu.Unlock()
		}
	}
}
