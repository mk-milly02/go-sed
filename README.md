# go-sed

`go-sed` is a simple and efficient command-line tool written in Go for performing basic text transformations, similar to the Unix `sed` command.

## Features

- Search and replace text in files
- In-place file editing
- Addition of double spacing
- Removal of trailing blank lines
- Suppress automatic printing of pattern space

## Usage

Basic usage of `go-sed`:

```sh
$ go run main.go -h
Usage of: go run main.go [options] [script] [filename]
Options:
    -i string
        edit in-place (default "s/The/Code/g")
    -n string
        only output a range of lines from the file (default "1p")
Script:
    - s/this/that/g
    - G
    - /^$/d
Filename:
    quotes.txt
```

```sh
go run main.go 's/old/new/g' input.txt
```

- `s/old/new/g`: Substitute `old` with `new` globally in the file

```sh
go run main.go -n '2,4p' input.txt
```

- `-n '2,4p'`: Only output a range of lines from the file i.e line 2 to line 4

```sh
go run main.go -n /code/p input.txt
```

- `-n /code/p`: Output only lines containing a specific pattern `code`

```sh
go run main.go G input.txt
```

- `G`: Support double spacing a file

```sh
go run main.go /^$/d input.txt
```

- `/^$/d`: Strip trailing blank lines from a file

```sh
go run main.go -i 's/Life/Code/g' intput.txt
```

- `-i 's/Life/Code/g'`: Edit in place

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
