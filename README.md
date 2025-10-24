Tiny Go tool that shows how any character is represented in Unicode, UTF-8, UTF-16, and UTF-32 encodings.

### 🧩 Install

```bash
go install github.com/fmo/unicode-inspector@latest

## Usage

`go run main.go 👍🏾`

Output:

```
Character: 👍🏾

Unicode code point(s): [U+1F44D U+1F3FE] <br />
UTF8 bytes: 0XF0 0X9F 0X91 0X8D 0XF0 0X9F 0X8F 0XBE <br />
UTF16 units: 0x[d83d dc4d d83c dffe] <br />
UTF32 code points: [U+1F44D U+1F3FE]
```
