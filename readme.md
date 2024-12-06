# My `cat` Command for Windows
A simple implementation of the `cat` command (`my-cat`) for Windows, written in Go. It supports reading single or multiple files, with an optional line numbering feature.

## Features
- **Single File Reading**: Display the contents of a single file.
- **Multiple File Reading**: Concatenate and display contents of multiple files.
- **Line Numbering**: Add line numbers to the output using the `-n` flag.

## Usage
```bash
$ my-cat [-n] <file1> [file2 ...]
```
### Read
- To read a file execute `my-cat <file1>`. Example:
```bash
$ my-cat error.log

# Output:
HTTP ERR: Status: 500 at Wed, 04 Dec 2024 09:12:56
HTTP ERR: Status: 401 at Wed, 04 Dec 2024 11:32:31
HTTP ERR: Status: 418 at Wed, 04 Dec 2024 16:08:47
```

### Read Multiple Files
- To read multiple files list the file names after the `my-cat` command `my-cat <file1> [file2 file3 ...]`. Example:
```bash
$ my-cat error.log file.txt

# Output:
HTTP ERR: Status: 500 at Wed, 04 Dec 2024 09:12:56
HTTP ERR: Status: 401 at Wed, 04 Dec 2024 11:32:31
HTTP ERR: Status: 418 at Wed, 04 Dec 2024 16:08:47

Hello from txt file!
```

### Read with Line Numbers
- To read files with line numbers, simply add the `-n` flag.
```bash
$ my-cat -n error.log

# Output:
1 HTTP ERR: Status: 500 at Wed, 04 Dec 2024 09:12:56
2 HTTP ERR: Status: 401 at Wed, 04 Dec 2024 11:32:31
3 HTTP ERR: Status: 418 at Wed, 04 Dec 2024 16:08:47
```
