Go Runtime          | Goroutine 1           | Goroutine 2
                     mov   eax, 12345
                     add   eax, 1
interrupt:
save G1's state
restore G2's state
                                             mov   eax, 12345
                                             add   eax, 1
                                             mov   12345, eax
interrupt:
save G2's state
restore G1's state
                    
                     mov 12345, eax

