# Leetcode in Go - Cheet Sheet

Booklet with Go solutions for Yangshun Tay's [updated Blind 75 Leetcode](https://www.techinterviewhandbook.org/best-practice-questions/) list.

All problem descriptions are taken from [LeetCode](https://leetcode.com/).

Solutions were prepared in 2 steps:

1. LLM generated solution based on file and method names and description. The solution was verified on leetcode.com - all generated solutions passed the test.
2. Solution readability was manually improved (e.g. changes in variable names, use of Go idioms).

## Development

Use `make` (GNU or BSD):

- `make` - install dependencies
- `make preview` - preview booklet
- `make booklet` - creates booklet for distribution
- `make clean` - removes compilation artifacts
- `make test` - run source code tests
- `make check` - run unit tests of source code
- `make info` - print system info (useful for debugging).

## License

[0BSD](./LICENSE) (public domain)
