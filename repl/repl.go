package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/d2fong/GIGO/lexer"
	"github.com/d2fong/GIGO/token"
)

// PROMPT is the prefix for each line of the print loop
const PROMPT = ">> "

// Start the read eval print loop
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
