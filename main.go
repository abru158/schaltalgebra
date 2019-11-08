package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

func parse(str string) (expression, rune, rune, error) {
	s := &scanner{}
	s.init([]byte(str))
	p := &parser{}
	p.init(s)
	expr, err := p.parse()
	return expr, p.lowestVar, p.highestVar, err
}

func boolStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func main() {
	fmt.Print(`Logic Algebra Calc ~~

To express a negation, use '!' or '-' (minus/dash)
To express a logical AND, use '^' or '&'
To express a logical OR, use 'v' or '|'

Input needs to be lower case and between a and h inclusive

Type "exit" to end the program

`)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter formula: ")
		formula, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("Error: ", err)
		}
		formula = strings.TrimRight(formula, "\r\n")
		if formula == "exit" {
			return
		}

		expr, lvar, hvar, err := parse(formula)
		if err != nil {
			fmt.Println("Error while parsing:", err)
			continue
		}

		states := stateMachine{}
		states.init(lvar, hvar)

		w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', tabwriter.AlignRight)

		for i := lvar; i <= hvar; i++ {
			fmt.Fprintf(w, "%s\t", string(i))
		}
		fmt.Fprintf(w, "Q\t\n")

		for {
			for i := lvar; i <= hvar; i++ {
				fmt.Fprintf(w, "%s\t", boolStr(states.get(i)))
			}
			fmt.Fprintf(w, "%s\t\n", boolStr(expr.eval(&states)))

			w.Flush()

			if !states.increment() {
				break
			}
		}

		fmt.Print("\n\n")
	}
}
