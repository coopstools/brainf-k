# brainf-k
A library containing an interpreter and examples

To run the BrainF-k REPL, just use the command
```bash
make
```

To run the code with one of the provided examples, make sure you are using th correct version of the go compiler and run
```bash
make run FILE=examples/3_fibSeq.bf ARGS="5"
```

To compile `.bf` files:
```bash
make compile FILE=examples/3_fibSeq.bf
or
make build
./bf -c examples/3_fibSeq.bf
```
and then the code can be run as
```bash
./3_fibSeq 12
```

## Variations

### Output Formatting

*WARNING* The starting index when running the code will be 4, not 0. The first 4 bytes in the stack are used for system settings.
```
byte 0 - determines whether prints, `.`, are sent to stdout (0) or stderr (not 0)
byte 1 - determines whether prints, `.`, are printed as ascii characters (1), hexadecimal values (2) or decimal values (0 or >2).
byte 2 - determines whether dumps, `#`, are sent to stdout (0) or stderr (not 0)
byte 3 - determines whether dumps, `#`, are printed as ascii characters (1), hexadecimal values (2) or decimal values (0 or >2). The preceding index in a dump will always be a decimal value.
```

Examples:
```bash
$ make compile FILE=examples/3_fibSeq.bf && ./3_fibSeq 12 2> err.txt > out.txt
Cleaning...
go clean
Building...
go build -o bf ./main
chmod +x bf
Compiling...
./bf -c examples/3_fibSeq.bf
gcc 3_fibSeq.c -o 3_fibSeq

$ cat err.txt
12 [233  144  89  55  34  21  13  8  5  3  2  1  1]

$ cat out.txt
1 1 2 3 5 8 13 21 34 55 89 144 233
```

Side Note: because of how stdout buffers work, dumps `#` will always be flushed/printed at the moment they are called, while prints `.` may not be flushed/printed until the end of execution of the program.

### Stack Dump
In the BrainF**k code, the `#` can used to print a portion of the stack, centered around the current index, to stdout. Stringing multiple `#` will increase the width of the stack output. The total number of elements printed will be 2*<number of #> + 1.

### Immediate Addressing
When using the `,` operator in the BrainF**k code, it will default to drawing values from the argument list supplied at run time. The values can be comma or space delimited. The `,` also has an immediate mode where a value can be supplied in the code, such as in `<<,15>+`. Here, rather than taking from argument list, it will pull in the value `15`.


## Examples
In the example directory are 5 examples that show the progress of how I learned BF.
1. HI
    This code was a simplified version of Hello,World. Basically, if you can type Hi, you know enough to expand to hello world.
2. Single Fib
    This code was the first steps in actual arithmatic, and involved core points of knowledge, such as moving numbers around on the stack.
3. Fib Sequence
    This code was the first foray into dealing with arrays, including finding first and last element.
4. Two Value Sort
    This was the first step into creating decision trees and branching operations. It makes use of the magic loop, `[->-[>]<<]`, but modified for this specific application.
5. N Value Sorting
    This code took the concepts from 3 and 4 and combined them into a full on, multi branching sequence. This code implements the bubble sort algorithm. Note that the array needs to terminate with a 255.
    ```shell
    $ go run main/main.go examples/5_NValueSorting.bf 25 32 41 35 255
    [25 32 35 41]
   ```
6. Optimized Sorting This works like the above program but runs in half as many operations