# sat-solver
SAT (satisfiability) solver based on https://sahandsaba.com/understanding-sat-by-implementing-a-simple-sat-solver-in-python.html

```
$ go run ./cmd/sat-solver <(echo -e 'a b c\na ~b\n~b c')
2020/08/13 20:12:24 processing line = "a b c"
2020/08/13 20:12:24 processing line = "a ~b"
2020/08/13 20:12:24 processing line = "~b c"
2020/08/13 20:12:24 problem contains 3 variables (search space: 8) and 3 clauses
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
2020/08/14 10:28:04 problem contains 8 variables (search space: 256) and 18 clauses
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
