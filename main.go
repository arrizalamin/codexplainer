package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	c "github.com/arrizalamin/codexplainer/context"
	"github.com/arrizalamin/codexplainer/explainer/java"
)

var (
	wg sync.WaitGroup
)

func main() {
	var languageFlag string
	flag.StringVar(&languageFlag, "lang", "java", "-lang=[java|php]")
	flag.Parse()

	filenames := flag.Args()

	for _, filename := range filenames {
		wg.Add(1)
		go performExplainSource(filename, languageFlag)
	}
	fmt.Println("Generating...")
	wg.Wait()
}

func performExplainSource(filename string, lang string) {
	defer wg.Done()
	source := scanFile(filename)
	explainer := getExplainer(lang)
	ctx := c.NewContext(filename, source, explainer)
	ctx.Execute()
	ctx.PrintOutput()
}

func scanFile(filename string) string {
	sourceByte, err := ioutil.ReadFile(filename)
	fmt.Println("scanning " + filename)
	if err != nil {
		log.Fatalln("File " + filename + " not found")
	}
	return string(sourceByte)
}

func getExplainer(lang string) c.ExplainerFunc {
	switch lang {
	case "java":
		return java.JavaExplainer
	default:
		log.Fatalf("Explainer for %s not found", lang)
	}
	return nil
}
