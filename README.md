# sat-solver
SAT (satisfiability) solver based on https://sahandsaba.com/understanding-sat-by-implementing-a-simple-sat-solver-in-python.html

```
$ go run *.go <(echo -e 'a b c\na ~b\n~b c')
2020/08/13 20:12:24 processing line = "a b c"
2020/08/13 20:12:24 processing line = "a ~b"
2020/08/13 20:12:24 processing line = "~b c"
2020/08/13 20:12:24 problem contains 3 variables (search space: 8) and 3 clauses
2020/08/13 20:12:24 solution = ~a ~b c
2020/08/13 20:12:24 solution = a ~b ~c
2020/08/13 20:12:24 solution = c a ~b
2020/08/13 20:12:24 solution = a b c
```
