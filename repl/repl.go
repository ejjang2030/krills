package repl

import (
	"bufio"
	"fmt"
	"io"
	"krills/evaluator"
	"krills/lexer"
	"krills/object"
	"krills/parser"
)

const PROMPT = ">> "

const KRILLS_SKIN = `

___  __    ________  ___  ___       ___       ________      
|\  \|\  \ |\   __  \|\  \|\  \     |\  \     |\   ____\     
\ \  \/  /|\ \  \|\  \ \  \ \  \    \ \  \    \ \  \___|_    
 \ \   ___  \ \   _  _\ \  \ \  \    \ \  \    \ \_____  \   
  \ \  \\ \  \ \  \\  \\ \  \ \  \____\ \  \____\|____|\  \  
   \ \__\\ \__\ \__\\ _\\ \__\ \_______\ \_______\____\_\  \ 
    \|__| \|__|\|__|\|__|\|__|\|_______|\|_______|\_________\
                                                 \|_________|
                                                             
                                                             

`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	isContinue := false
	line := ""
	for {
		if !isContinue {
			fmt.Fprintf(out, PROMPT)
		}
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		if isContinue {
			line += scanner.Text()
		} else {
			line = scanner.Text()
		}
		if line[len(line) - 1] == '\\' {
			isContinue = true
			continue
		} else {
			isContinue = false
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, KRILLS_SKIN)
	io.WriteString(out, "Woops! We ran into some Krills business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
