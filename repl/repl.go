package repl

import (
	"bufio"
	"fmt"
	"io"
	"krills/compiler"
	"krills/lexer"
	"krills/parser"
	"krills/vm"
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
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
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
