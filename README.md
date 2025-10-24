Tiny Go tool that shows how any character is represented in Unicode, UTF-8, UTF-16, and UTF-32 encodings.

### 🧩 Install

```bash
go install github.com/fmo/unicode@latest
```

## Usage

`go run main.go 🚀`

or from the binary

`./unicode 🚀`

Output:

```
Character: 🚀

Unicode code point(s): [U+1F680]
UTF8 bytes: 0XF0 0X9F 0X9A 0X80
UTF16 units: 0x[d83d de80]
UTF32 code points: [U+1F680]
```
