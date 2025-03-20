<<+>>           set dumps to be sent to stderr
,               set fib index to user input; N less than 13
-               first value is already found; so start with decrement of index
>>+             set s2 value
>+              set s3 value
<<<             rit0 or return index to 0; end initialization

[>>             go to s2
[>]<            go to last element in array
[[>+<-]<]       shift element right
>>[<+<+>>-]     mv s3 into s1 and s2
<<[>>+<<-]      mv s1 into s3
>>>[<<+<+>>>-]  mv s4 into s1 and s2
<<<[>>>+<<<-]   mv s1 into s4
<-]             dec index

>>[>]<          goto end of array
[.<]            print each element in array
>>>>>>>######
