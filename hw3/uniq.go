package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func uniq(input io.Reader, output io.Writer, c, d, u, i bool, numFields, numChars int) error {
	in := bufio.NewScanner(input)
	var lines []string
	for in.Scan() {
		line := in.Text()
		lines = append(lines, line)
	}
	var prev string
	count := 0

	for j, l := range lines {
		txt := l
		if i {
			txt = strings.ToLower(txt)
		}
		if numFields > 0 && txt != "" {
			cutfield(&txt, numFields)
		}
		if numChars > 0 && txt != "" {
			if numChars < len(txt) {
				txt = txt[numChars:]
			}
		}
		if txt == prev {
			count++
			if !c && !u && !d || j != len(lines)-1 {
				continue
			}
		}
		if c && (j != 0) {
			fmt.Fprintln(output, count+1, " ", prev)
		}
		if d && (count > 0) {
			fmt.Fprintln(output, prev)
		}
		if u && (count == 0) && (j != 0) {
			fmt.Fprintln(output, prev)
		}

		if !c && !u && !d {
			fmt.Fprintln(output, l)
		}
		prev = txt
		count = 0

	}
	return nil

}

func cutfield(txt *string, nf int) {
	t := strings.Fields(*txt)
	if nf > len(t) {
		return
	}
	t = t[nf:]
	*txt = strings.Join(t, " ")
}

func main() {

	count := flag.Bool("c", false, "number of occurrences")
	dupl := flag.Bool("d", false, "print duplicate lines")
	unique := flag.Bool("u", false, "print unique lines")
	numFields := flag.Int("f", 0, "number of fields we should ignore")
	numChars := flag.Int("s", 0, "number of chars we should ignore")
	ignReg := flag.Bool("i", false, "ignore registr")

	flag.Parse()

	if (*count && *dupl) || (*count && *unique) || (*dupl && *unique) {
		fmt.Fprintln(os.Stderr, "uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		return
	}
	var input *os.File
	var output *os.File
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}
	if flag.NArg() > 1 {
		outfile, err := os.OpenFile(flag.Arg(1), os.O_WRONLY, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer outfile.Close()
		output = outfile
	} else {
		output = os.Stdout
	}

	err := uniq(input, output, *count, *dupl, *unique, *ignReg, *numFields, *numChars)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
