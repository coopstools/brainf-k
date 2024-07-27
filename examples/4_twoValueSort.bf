this code block needs 6 spaces in the stack with the first space being zeroed
>,>,<                    takes input into s1 and s2
+>+<                     does an increment necessary for if either value is 0
[-   >>>+<<<   >-[>]<<]  standard magic loop but makes copy of smallest number into s4
>>>[<]                   hones in on s3 which is a 0 before the smaller value
<<[>+<-]                 moves the diff to s2 if not already there
>>>[<<+<+>>>-]           copies s4 ie smallest value onto s1 and s2 making s2 the larger initial number
<<-<-.>.                 remove 0 protection increments from line 2 and outputs values