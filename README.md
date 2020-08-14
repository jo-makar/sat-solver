# sat-solver
SAT (satisfiability) solver based on https://sahandsaba.com/understanding-sat-by-implementing-a-simple-sat-solver-in-python.html

```
$ go run ./cmd/sat-solver <(echo -e 'a b c\na ~b\n~b c')
2020/08/13 20:12:24 processing line = "a b c"
2020/08/13 20:12:24 processing line = "a ~b"
2020/08/13 20:12:24 processing line = "~b c"
2020/08/13 20:12:24 problem contains 3 variables and 3 clauses
2020/08/13 20:12:24 solution = ~a ~b c
2020/08/13 20:12:24 solution = a ~b ~c
2020/08/13 20:12:24 solution = c a ~b
2020/08/13 20:12:24 solution = a b c
```

## four-color

Convert the Four color theorem (https://en.wikipedia.org/wiki/Four_color_theorem) into a SAT problem.

Each region is numbered (1 to n) and neighbors are provided one per line.  The variables are ri, bi, gi, yi to indicate region i is colored red, blug, green or yellow.

```
$ go run ./cmd/sat-solver/ -s <(go run ./cmd/four-color/ <(echo -e '1 2'))
2020/08/14 10:28:04 processing line = "1 2"
2020/08/14 10:28:04 problem contains 2 regions and 1 neighbors
2020/08/14 10:28:04 processing line = "r1 b1 g1 y1"
2020/08/14 10:28:04 processing line = "~r1 ~b1"
2020/08/14 10:28:04 processing line = "~r1 ~g1"
2020/08/14 10:28:04 processing line = "~r1 ~y1"
2020/08/14 10:28:04 processing line = "~b1 ~g1"
2020/08/14 10:28:04 processing line = "~b1 ~y1"
2020/08/14 10:28:04 processing line = "~g1 ~y1"
2020/08/14 10:28:04 processing line = "r2 b2 g2 y2"
2020/08/14 10:28:04 processing line = "~r2 ~b2"
2020/08/14 10:28:04 processing line = "~r2 ~g2"
2020/08/14 10:28:04 processing line = "~r2 ~y2"
2020/08/14 10:28:04 processing line = "~b2 ~g2"
2020/08/14 10:28:04 processing line = "~b2 ~y2"
2020/08/14 10:28:04 processing line = "~g2 ~y2"
2020/08/14 10:28:04 processing line = "~r1 ~r2"
2020/08/14 10:28:04 processing line = "~b1 ~b2"
2020/08/14 10:28:04 processing line = "~g1 ~g2"
2020/08/14 10:28:04 processing line = "~y1 ~y2"
2020/08/14 10:28:04 problem contains 8 variables and 18 clauses
2020/08/14 10:28:04 solution = g2 y1
2020/08/14 10:28:04 solution = b2 y1
2020/08/14 10:28:04 solution = y1 r2
2020/08/14 10:28:04 solution = y2 g1
2020/08/14 10:28:04 solution = g1 b2
2020/08/14 10:28:04 solution = g1 r2
2020/08/14 10:28:04 solution = y2 b1
2020/08/14 10:28:04 solution = g2 b1
2020/08/14 10:28:04 solution = b1 r2
2020/08/14 10:28:04 solution = y2 r1
2020/08/14 10:28:04 solution = r1 g2
2020/08/14 10:28:04 solution = b2 r1
```

## sudoku

Convert a Sudoku puzzle into a SAT problem.

Let pijk represent row i, line j and value k (all of range [1,9]).  These also serve as inputs for the initial puzzle values.

```
$ go run ./cmd/sat-solver/ -s <(echo p318 p429 p621 p927 p732 p835 p934 p243 p846 p451 p552 p162 p564 p666 p765 p874 p975 p188 p381 p483 p393 p596 p798 | tr ' ' '\n'; go run ./cmd/sudoku/)
2020/08/14 14:27:14 processing line = "p318"
2020/08/14 14:27:14 processing line = "p429"
2020/08/14 14:27:14 processing line = "p621"
2020/08/14 14:27:14 processing line = "p927"
2020/08/14 14:27:14 processing line = "p732"
2020/08/14 14:27:14 processing line = "p835"
2020/08/14 14:27:14 processing line = "p934"
2020/08/14 14:27:14 processing line = "p243"
2020/08/14 14:27:14 processing line = "p846"
2020/08/14 14:27:14 processing line = "p451"
2020/08/14 14:27:14 processing line = "p552"
2020/08/14 14:27:14 processing line = "p162"
2020/08/14 14:27:14 processing line = "p564"
2020/08/14 14:27:14 processing line = "p666"
2020/08/14 14:27:14 processing line = "p765"
2020/08/14 14:27:14 processing line = "p874"
2020/08/14 14:27:14 processing line = "p975"
2020/08/14 14:27:14 processing line = "p188"
2020/08/14 14:27:14 processing line = "p381"
2020/08/14 14:27:14 processing line = "p483"
2020/08/14 14:27:14 processing line = "p393"
2020/08/14 14:27:14 processing line = "p596"
2020/08/14 14:27:14 processing line = "p798"
<go run ./cmd/sudoku lines elided>
2020/08/14 14:27:14 problem contains 729 variables and 11768 clauses
2020/08/14 14:27:14 solution = p337 p179 p344 p773 p216 p243 p596 p297 p913 p975 p612 p467 p828 p571 p942 p224 p666 p986 p719 p968 p451 p857 p564 p156 p478 p393 p261 p787 p483 p369 p492 p525 p272 p863 p436 p798 p741 p754 p533 p621 p874 p322 p147 p684 p548 p991 p882 p835 p445 p934 p732 p552 p959 p846 p726 p285 p649 p429 p589 p162 p638 p258 p376 p811 p115 p131 p194 p899 p188 p381 p695 p123 p355 p239 p318 p653 p765 p517 p927 p414 p677
```
