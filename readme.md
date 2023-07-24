# Word and Line Counter

This Go program is a simple word and line counter that reads input from the standard input and calculates either the number of words or lines, based on the user's choice.

## Usage

To use the program, run it from the command line and provide the `-w` flag if you want to count words. By default, it will count lines.

```bash
go run main.go -w
```

The program will then wait for input from the user. Enter your text and end the input with the word "exit" to terminate the input and see the count result.

## How it works

The program uses the `flag` package to parse command-line arguments and determine whether to count words or lines. It then launches a goroutine to display a message about the type of count (words or lines) while the user provides input. The `count` function reads the input, calculates the number of words or lines, and returns the count.

## Example

```bash
$ go run main.go -w
The words is true
reading data
This is a simple example for word counting.
Please type "exit" to end the input.
exit
The words count is: 10
```

```bash
$ go run main.go
The words is false
**********Enter Data**********
This is a simple example for line counting.
Please type "exit" to end the input.
exit
The lines count is: 2
```

## Dependencies

The program uses the following standard libraries from Go:

- "bufio" for buffered input/output operations.
- "flag" for parsing command-line arguments.
- "fmt" for formatted output.
- "io" for I/O operations.
- "os" for interacting with the operating system.

## Author

This program was created by [krishnaps7] ([krishnaps909@example.com](mailto:krishnaps909@example.com)).

Feel free to modify and use this program according to your needs. Happy coding!