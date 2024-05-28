# Coding Challenges wc tool

I'm in the process of learning Go and I came across the [Coding Challenges](https://codingchallenges.fyi/challenges/intro/) website.
This challenge is to build a version of the Unix wc command line tool. [Wikipedia](https://en.wikipedia.org/wiki/Wc_(Unix))

# Usage
```
// build binary
go build .
```

```
// from the project root

./ccwc path-to-file
```

```
./ccwc ./test.txt

// expected output
// 7145 58164 342190 ./test.txt
```

Supports flags for byte count `-c`, line count `-l`, character count `-m`, and word count `-w`. If no flag is passed the tool returns `-c`, `-l`, and `-w` options.

Also supports `-help` flag.

```
Usage of ./ccwc:
  -c    Returns byte count for file
  -l    Returns line count for file
  -m    Returns character count for file
  -w    Returns word count for file

```
