# brainf-k
A library containing an interpreter and examples

To run the code, make sure you are using th correct version of the go compiler and run
```bash
go run main/main.go examples/1_hi.bf
```

## Options
### c - output as characters
The default output is an array of byte values. The `-c` option allows output as a string.
```bash
$ go run main/main.go examples/1_hi.bf
[72 73]

$ go run main/main.go examples/1_hi.bf -c
HI
```

### d - dump stack
If included, this option causes the stack to be dumped, inluding the final index, the first 17 values on the stack, and the number of operations performed during run time.
```shell
$ go run main/main.go examples/1_hi.bf -s
1 [0 73 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 126
[72 73]
```

### s - silent
If included, this silences the result output. Note that `-d` supersedes this option.

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