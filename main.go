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
		if *n != "" {
			output = filter(input, *n)
		}
	} else {
		switch {
		case strings.HasPrefix(script, "s/"):
			output = substitute(input, script)
		case strings.HasPrefix(script, "p"):
			break
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
