# advent-of-code

My [advent of code](https://adventofcode.com) solutions in [Go](https://go.dev). 

## Structure
Code is organised by year then day:
```
.\[year]
  .\day[dd]
    .\day[dd]_input.txt
    .\day[dd]_test.go
    .\day[dd].go
    .\day[dd].txt
```

### Creating days
Use the provided `create-day.sh` shell script to create a new day consistently, e.g. for day 1 of 2021:
```
./create-day.sh 2021 01
```

## Testing
It is intended you run the tests from the repository root:
```
go test ./2021/01/.
```

## Benchmarking
You also run the benchmarks from the repository root:
```
go test ./2021/01/. --bench .
```

## Progress
Hoping to get at least 25 of 50 possible stars!

| Day | Part 1 | Part 2 | Notes |
| :---: | :---: | :---: | :---- |
| 01 | ⭐️ | ⭐️ | Signal monitoring|
| 02 | ⭐️ | ⭐️ | Instruction running |
| 03 | ⭐️ | ⭐️ | Bit twiddling |
| 04 | ⭐️ | ⭐️ | Bingo simulation|
| 05 | ⭐️ | ⭐️ | Line intersections on grid|
| 06 | ⭐️ | ⭐️ | Population growth modelling |
| 07 | ⭐️ | ⭐️ | Calculate optimal costs|
| 08 | ⭐️ | ⭐️ | Decoding output from input |
| 09 | ⭐️ | ⭐️ | 2D neighbour comparison |
| 10 | ⭐️ | ⭐️ | Braces checking |
| 11 | ⭐️ | | Grid analysis? |