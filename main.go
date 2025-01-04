package main

import (
	"flag"
	"io"
	"os"
	"strings"
)

func main() {
    flag.Parse()
	substitution := flag.Args()[0]
    filename := flag.Args()[1]
    if filename == "" {
        flag.PrintDefaults()
        return
    }
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    content, err := io.ReadAll(file)
    if err != nil {
        panic(err)
    }
    result := substitute(content, substitution)
    println(string(result))
}

func substitute(content []byte, substitution string) (result []byte) {
    substitution = strings.TrimPrefix(substitution, "s/")
    substitution = strings.TrimSuffix(substitution, "/g")
    parts := strings.Split(substitution, "/")
    if len(parts) != 2 {
        panic("invalid substitution")
    }
    result = []byte(strings.ReplaceAll(string(content), parts[0], parts[1]))
    return result
}
