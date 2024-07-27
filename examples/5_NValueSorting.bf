>>>>>>>,+[>,+]<[<]      	init input needs at least one value and a trailing 255

>[							start full loop
[<<<+>>>-]             		set n0
>[							start single pass sorting
[<<<+>>>-]<<<<           	mv n1 for sorting
[- >>>+<<< >-[>]<<]>>>[<] 	magic sort ln1
<<[>+<-]>>>[<<+<+>>>-] <<< 	magic sort ln2
[<+>-]            		    shifted smaller value
>>>>>]						go to next element in array
<<<<[<<[<]<+>>[>]>-]        shift high value to beginning
<<[[>>>>>+<<<<<-]<]>>>>>>]	shift the array over into the starting position
<<<<<<<[-.<]				remove 0 safe check on values and print output