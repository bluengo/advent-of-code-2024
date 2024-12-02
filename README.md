# Advent of Code 2024

Solutions writen in Go for [adventofcode2024](https://adventofcode.com/2024)

## Layout
This repository has the below layout:
```bash
.
├── cmd
│   └── main.go           # Entrypoint
├── go.mod
├── go.sum
├── inputs                # Your personal inputs here
│   ├── dayXX.txt
│   └── test_day02.txt
└── solutions             # Solutions per each day
    ├── day01
    │   └── day01.go
    └── day02
    │   └── day01.go
    ├── ...
    └── dayXX
        └── dayXX.go
```

## How to use it
You can run the entrypoint, providing the `day`, `part` and `inputs` filepath as flags:
```bash
# i.e. Running the second part of the first day challenge
go run cmd/main.go --day 1 --part 2 --inputs inputs/day01.txt
```

## Usage:
```
Flags:
  -h, --help             Show context-sensitive help.
      --day=INT          The day of the challenge to execute
      --part=INT         The part of the challenge to execute
      --inputs=STRING    Path to the inputs file for the challenge
```
