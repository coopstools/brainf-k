preprocessor
<<<< >++ >>++ >      Sets the output to use hexadecimal
code
,>,>>,>,<<<<  Takes in two two byte values; separated by a blank byte; MSB firsts
[->>>+<<<]    Add MSB together
>[
  ->>+>+>+<    dec addin; inc MSB; inc LSB; set flag; ret LSB
  [<->>-]      if LSB is nonzero: dec MSB; clr flag
  >[->]<<      clr flag; ret LSB
  <<<          ret addin
]
>>###
.>.          print MSB; print LSB