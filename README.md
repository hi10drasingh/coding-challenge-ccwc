# ccwc

A custom implementation of the Unix `wc` (word count) tool in Go.

## Features

- Counts lines, words, characters, and bytes in a file or from standard input.
- Mimics the behavior of the classic `wc` utility.
- Simple and efficient, suitable for learning and practical use.

## Usage

Build the project:

```sh
go build -o ccwc ./cmd/ccwc
```

Run the tool:

```sh
./ccwc [flags] [file]
```

### Flags

- `-l` : Print the newline count
- `-w` : Print the word count
- `-c` : Print the byte count
- `-m` : Print the character count

If no flags are provided, all counts are displayed.

### Examples

Count lines, words, and bytes in a file:

```sh
./ccwc test.txt
```

Count only lines:

```sh
./ccwc -l test.txt
```

Read from standard input:

```sh
cat test.txt | ./ccwc
```

## Project Structure

```
.
├── cmd/
│   └── ccwc/
│       └── main.go
├── internal/
│   └── service/
│       ├── ccwc.go
│       └── ccwc_test.go
├── test.txt
├── go.mod
├── go.sum
└── README.md
```

- `cmd/ccwc/main.go`: Command-line interface entry point.
- `internal/service/ccwc.go`: Core logic for counting.
- `internal/service/ccwc_test.go`: Unit tests.
- `test.txt`: Sample input file.

## License

This project is licensed under the MIT License. See
