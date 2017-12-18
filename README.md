# ll1-parser
An assignment in Theory of Computation subject

Do a parser for a languange generated by grammars

## Grammars
1. S -> ZS | ε
1. Z -> id = E;
1. E -> TE’
1. E’ -> +TE’ | -TE’ | ε
1. T -> FT’
1. T’ -> *FT’ | /FT’ | ε
1. F -> idA | int | real | (E)
1. A -> (E) | ε
