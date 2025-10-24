package main

import (
    "fmt"
    "log"
    "os"
    "unicode/utf16"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("no args")        
    }

    c := os.Args[1]

    r := []rune(c)
    
    encoded16 := utf16.Encode(r)

    fmt.Printf("Character: %s\n\n", c)
    fmt.Printf("Unicode code point(s): %U\n", r)
    fmt.Printf("UTF8 bytes: % #02X\n", []byte(c))
    fmt.Printf("UTF16 units: 0x%x\n", encoded16)
    fmt.Printf("UTF32 code points: %U\n", r)
}
