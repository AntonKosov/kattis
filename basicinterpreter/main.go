package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var writer = bufio.NewWriterSize(os.Stdout, 1_000_000)

func main() {
	lines := read()
	vars := NewVariables()
	mapLines := mapLines(lines)
	commands := parseCommands(lines, mapLines, &vars)
	for cl := 0; cl < len(commands); cl = commands[cl].Exec(cl) {
	}
	writer.Flush()
}

func mapLines(codeLines []CodeLine) map[int]int {
	ml := make(map[int]int, len(codeLines))
	for _, line := range codeLines {
		ml[line.lineNumber] = 0
	}
	lineNumbers := make([]int, 0, len(ml))
	for lineNumber := range ml {
		lineNumbers = append(lineNumbers, lineNumber)
	}
	sort.Slice(lineNumbers, func(i, j int) bool { return lineNumbers[i] < lineNumbers[j] })
	for i, ln := range lineNumbers {
		ml[ln] = i
	}

	return ml
}

func parseCommands(lines []CodeLine, mapLines map[int]int, vars *Variables) map[int]Command {
	commands := make(map[int]Command, len(lines))
	for _, line := range lines {
		commands[mapLines[line.lineNumber]] = parseCommand(line.command, mapLines, vars)
	}

	return commands
}

func parseCommand(st string, mapLines map[int]int, vars *Variables) Command {
	parts := strings.SplitN(st, " ", 2)
	assertEq(2, len(parts))
	st = parts[1]
	switch parts[0] {
	case "LET":
		return NewLetCommand(st, vars)
	case "IF":
		return NewIfCommand(st, mapLines, vars)
	case "PRINT":
		return NewPrintCommand(st, vars)
	case "PRINTLN":
		return NewPrintlnCommand(st, vars)
	}

	panic("incorrect input")
}

type Command interface {
	Exec(currentLine int) int
}

type LetCommand struct {
	vars      *Variables
	target    rune
	statement ArithmeticStatement
}

func (c LetCommand) Exec(cl int) int {
	c.vars.Set(c.target, c.statement.Calc())
	return cl + 1
}

func NewLetCommand(st string, vars *Variables) LetCommand {
	parts := strings.SplitN(st, " ", 3)
	assertEq(3, len(parts))

	return LetCommand{
		vars:      vars,
		target:    rune(parts[0][0]),
		statement: NewAStatement(parts[2], vars),
	}
}

type IfCommand struct {
	condition  Condition
	lineNumber int
}

func (c IfCommand) Exec(cl int) int {
	if c.condition.Calc() {
		return c.lineNumber
	}

	return cl + 1
}

func NewIfCommand(st string, mapLines map[int]int, vars *Variables) IfCommand {
	parts := strings.Split(st, " ")
	assertEq(6, len(parts))

	return IfCommand{
		condition:  NewCondition(parts[0], parts[1], parts[2], vars),
		lineNumber: mapLines[strToInt(parts[5])],
	}
}

type PrintCommand struct {
	statement PrintStatement
}

func (c PrintCommand) Exec(cl int) int {
	writer.WriteString(c.statement.Calc())
	return cl + 1
}

func NewPrintCommand(st string, vars *Variables) PrintCommand {
	return PrintCommand{statement: NewPrintStatement(st, vars)}
}

type PrintlnCommand struct {
	statement PrintStatement
}

func (c PrintlnCommand) Exec(cl int) int {
	writer.WriteString(c.statement.Calc() + "\n")
	return cl + 1
}

func NewPrintlnCommand(st string, vars *Variables) PrintlnCommand {
	return PrintlnCommand{statement: NewPrintStatement(st, vars)}
}

type ArithmeticStatement interface {
	Calc() int32
}

type OneValueArithmeticStatement struct {
	value ValueStatement
}

func (s OneValueArithmeticStatement) Calc() int32 {
	return s.value.Calc()
}

type OperationArithmeticStatement struct {
	op func() int32
}

func (s OperationArithmeticStatement) Calc() int32 {
	return s.op()
}

func NewAStatement(st string, vars *Variables) ArithmeticStatement {
	parts := strings.Split(st, " ")
	if len(parts) == 1 {
		return OneValueArithmeticStatement{value: NewValueStatement(parts[0], vars)}
	}
	assertEq(3, len(parts))
	a, b := NewValueStatement(parts[0], vars), NewValueStatement(parts[2], vars)
	var op func() int32
	switch oper := parts[1]; oper {
	case "+":
		op = func() int32 { return a.Calc() + b.Calc() }
	case "-":
		op = func() int32 { return a.Calc() - b.Calc() }
	case "*":
		op = func() int32 { return a.Calc() * b.Calc() }
	case "/":
		op = func() int32 { return a.Calc() / b.Calc() }
	default:
		panic("unknown operation: " + oper)
	}

	return OperationArithmeticStatement{op: op}
}

type Condition struct {
	cond func() bool
}

func (c Condition) Calc() bool {
	return c.cond()
}

func NewCondition(first, op, second string, vars *Variables) Condition {
	a, b := NewValueStatement(first, vars), NewValueStatement(second, vars)
	var cond func() bool
	switch op {
	case "=":
		cond = func() bool { return a.Calc() == b.Calc() }
	case ">":
		cond = func() bool { return a.Calc() > b.Calc() }
	case "<":
		cond = func() bool { return a.Calc() < b.Calc() }
	case "<>":
		cond = func() bool { return a.Calc() != b.Calc() }
	case "<=":
		cond = func() bool { return a.Calc() <= b.Calc() }
	case ">=":
		cond = func() bool { return a.Calc() >= b.Calc() }
	default:
		panic("invalid condition")
	}

	return Condition{cond: cond}
}

type PrintStatement interface {
	Calc() string
}

type PrintLiteralStatement struct {
	literal string
}

func (s PrintLiteralStatement) Calc() string {
	return s.literal
}

type PrintVariableStatement struct {
	value ValueStatement
}

func (s PrintVariableStatement) Calc() string {
	return strconv.Itoa(int(s.value.Calc()))
}

func NewPrintStatement(st string, vars *Variables) PrintStatement {
	if rune(st[0]) == '"' {
		return PrintLiteralStatement{
			literal: st[1 : len(st)-1],
		}
	}

	return PrintVariableStatement{value: NewValueStatement(st, vars)}
}

type ValueStatement interface {
	Calc() int32
}

type ConstantValueStatement struct {
	value int32
}

func (s ConstantValueStatement) Calc() int32 {
	return s.value
}

type VariableValueStatement struct {
	value func() int32
}

func (s VariableValueStatement) Calc() int32 {
	return s.value()
}

func NewValueStatement(st string, vars *Variables) ValueStatement {
	v, err := strconv.Atoi(st)
	if err != nil {
		return VariableValueStatement{value: func() int32 { return vars.Get(rune(st[0])) }}
	}

	return ConstantValueStatement{value: int32(v)}
}

func strToInt(v string) int {
	vi, err := strconv.Atoi(v)
	must(err)
	return vi
}

func read() []CodeLine {
	reader := bufio.NewReaderSize(os.Stdin, 1_000_000)
	var lines []CodeLine
	for {
		line, err := reader.ReadBytes('\n')
		if len(line) > 0 {
			line = line[:len(line)-1]
		}
		if len(line) > 0 {
			parts := strings.SplitN(string(line), " ", 2)
			assertEq(2, len(parts))
			lines = append(lines, CodeLine{
				lineNumber: strToInt(parts[0]),
				command:    parts[1],
			})
		}
		if err != nil {
			if err == io.EOF {
				return lines
			}
			panic(err)
		}
		if len(line) == 0 {
			return lines
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func assertEq[T comparable](expected, actual T) {
	if expected != actual {
		panic("incorrect input")
	}
}

type Variables struct {
	values []int32
}

func NewVariables() Variables {
	return Variables{
		values: make([]int32, 'Z'-'A'+1),
	}
}

func (v *Variables) Get(variable rune) int32 {
	return v.values[variable-'A']
}

func (v *Variables) Set(variable rune, value int32) {
	v.values[variable-'A'] = value
}

type CodeLine struct {
	lineNumber int
	command    string
}
