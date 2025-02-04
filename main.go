package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var script string
	var filename string
	var input []byte
	var output []byte
	n := flag.String("n", "1p", "only output a range of lines from the file")
	i := flag.String("i", "s/The/Code/g", "edit in-place")
	flag.Usage = func() {
		println("Usage of: go run main.go [options] [script] [filename]")
		println("Options:")
		flag.PrintDefaults()
	}
	flag.Parse()

	//when no flag is specified
	if flag.NFlag() == 0 {
		script = flag.Arg(0)
		filename = flag.Arg(1)
	} else {
		filename = flag.Arg(0)
	}

	//when no filename is specified
	if filename == "" {
		input = readFromStdin()
	} else {
		input = readFromFile(filename)
	}

	//when no script is specified
	if script == "" {
		switch {
		case strings.HasPrefix(*n, "/"):
			output = filterByPartern(input, *n)
		case strings.HasPrefix(*i, "s/"):
			substituteInPlace(input, *i, filename)
		default:
			output = filter(input, *n)
		}
	} else {
		switch {
		case strings.HasPrefix(script, "s/"):
			output = substitute(input, script)
		case script == "G":
			output = doubleSpace(input)
		case strings.HasSuffix(script, "/d"):
			output = removeBlankLines(input)
		default:
			panic("invalid script")
		}
	}
	println(string(output))
}

func readFromFile(name string) []byte {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return content
}

func readFromStdin() []byte {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return content
}

func substitute(content []byte, script string) (result []byte) {
	script = strings.TrimPrefix(script, "s/")
	script = strings.TrimSuffix(script, "/g")
	parts := strings.Split(script, "/")
	if len(parts) != 2 {
		panic("invalid substitution")
	}
	result = []byte(strings.ReplaceAll(string(content), parts[0], parts[1]))
	return result
}

func filter(content []byte, n string) (result []byte) {
	n = strings.TrimSuffix(n, "p")
	parts := strings.Split(n, ",")
	r := bytes.NewBuffer(content)
	if len(parts) == 1 {
		line, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		for i := 0; i < line; i++ {
			l, err := r.ReadString('\n')
			if err != nil {
				panic(err)
			}
			result = append(result, []byte(l)...)
		}
	} else if len(parts) == 2 {
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for i := start; i <= end; i++ {
			l, err := r.ReadString('\n')
			if err != nil {
				panic(err)
			}
			result = append(result, []byte(l)...)
		}
	}
	return result
}

func filterByPartern(content []byte, pattern string) (result []byte) {
	pattern = strings.TrimPrefix(pattern, "/")
	pattern = strings.TrimSuffix(pattern, "/p")
	r := bytes.NewBuffer(content)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(l, pattern) {
			result = append(result, []byte(l)...)
		}
	}
	return result
}

func doubleSpace(content []byte) (result []byte) {
	r := bytes.NewBuffer(content)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			break
		}
		result = append(result, []byte(l)...)
		result = append(result, []byte("\r\n")...)
	}
	return result
}

func removeBlankLines(content []byte) (result []byte) {
	r := bytes.NewBuffer(content)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if l != "\n" && l != "\r\n" {
			result = append(result, []byte(l)...)
		}
	}
	return result
}

func substituteInPlace(content []byte, script, filename string) {
	var result []byte
	script = strings.TrimPrefix(script, "s/")
	script = strings.TrimSuffix(script, "/g")
	parts := strings.Split(script, "/")
	if len(parts) != 2 {
		panic("invalid substitution")
	}
	result = []byte(strings.ReplaceAll(string(content), parts[0], parts[1]))
	file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(result)
	if err != nil {
		panic(err)
	}
}
