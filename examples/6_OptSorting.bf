Optimized to minimize number of operations; it uses 5N spaces in the stack
>,+[>,+]<                   init

[
  [>+<-]<                   shifted first
  [
    [>+<-]>>                shifted next in series
    [>>+<<-<-[<]>>]>[<]     find diff of top two
    >[>+<<<<+>>>-]<<[<+>-]  mv larger left and smaller right
    <<<                     reset pointer
  ]                         check if another element is in series
  >>-.                      print out largest value; largest value is forgotten
  >>>>[>]<                  reset pointer to end of next series
]