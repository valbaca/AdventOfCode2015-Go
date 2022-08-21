# Advent of Code in Go

Solutions to Advent of Code problems in Go.

This currently contains year 2015

Each day contains a "TIL" (Today I Learned) message that explains something about Go that I learned.

Looking back, the most common themes are:

1. Conversion between types is cumbersome. A lot of Advent of Code is parsing various ad-hoc formats for numbers and text commands.
2. Go is *fast*. So fast it can make brute-force solutions viable. Similarly, if any code is ever taking longer than a few seconds, you know a better solution exists (as opposed to other languages where `O(n)` solutions can still execute slowly).
3. That said, bytes.Buffer appending is much faster than strings (as it is in most languages).
4. Go's built-in Regex is fast but limited; no back-references.
5. Finding the unit testing boundary is tricky. Nearly every question has some sort of "parse input into structs and execute logic." Sometimes the parse logic is a majority of the code, other times it's trivial, so unit tests may be the actual text input, or it could be using the intermediate structs.

For example:

```go
func PartX(s string) {
	/* ... */
	someStruct := parseInput(s)
	actual(someStruct)
}
```

## Tips

Use [watchexec](https://watchexec.github.io/) to automatically re-run the code on file save:
```shell
$ watchexec go run main.go [year] [day]
```