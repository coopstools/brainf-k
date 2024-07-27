This code finds the Nth value in the fibonocci sequence for values of N less than 13
,               set index of fibonocci desired
-               first value is already found; so start with decrement of index
>+              set first value
>+              set second value  #end initialization
<<              rit0 or return index to 0

[               start main loop
>[>>+<<-]<      add first value to result; rit0
>>[<+>>+<-]<<   add second value to first and result; rit0
>>>[<+>-]<<<    move result to second position; rit0
-]              decrement the index
>>.             output result (note this is byte value and not printable ascii)