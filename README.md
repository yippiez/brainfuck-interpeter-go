# Brainfuck interpeter in go

Brainfuck interpeter written in go. 

## Usage

There are examples in the [examples](./.examples) category use them to run interpeter with the following structure.

```bash
$ go run execute.go .examples\helloworld.brainfuck
> Hello World!
```

## Buffer
Buffer is fixed to 64KB (64K x 1B) to change bufferSize in `execute.go` and if your buffer is bigger than 2^32 i.e 4GB change the `bufferPointer` in `brainfuck.go`.

> This is also true for the program size if it's bigger than 4GB change `programPointer` and `lenProgram`.

