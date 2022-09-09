# Golang-Auto-Tune-Concurrency

## How it works?
Concurrency pool size will be configured every time. It will remember the number of running tasks and if it is too much to handle with current concurrency pool, then increase the number of pool and vice versa.

## Result

```
Running tasks:  0
Running tasks:  1
Running tasks:  2
Running tasks:  3
Running tasks:  4
Running tasks:  5
Running tasks:  6
Running tasks:  7
Running tasks:  8
Running tasks:  9
Pool cap configured to  15
Pool cap configured to  10
Pool cap configured to  15
Pool cap configured to  10
Pool cap configured to  15
Pool cap configured to  10
Pool cap configured to  15
Pool cap configured to  10
Pool cap configured to  15
Running tasks:  9
Running tasks:  6
Running tasks:  7
Running tasks:  3
Running tasks:  4
Running tasks:  5
Running tasks:  6
Running tasks:  7
Running tasks:  8
Running tasks:  9
Running tasks:  10
Running tasks:  11
Running tasks:  12
Running tasks:  13
Running tasks:  14
Pool cap configured to  20
Pool cap configured to  15
Pool cap configured to  20
Pool cap configured to  15
Pool cap configured to  20
Pool cap configured to  15
Pool cap configured to  20
Pool cap configured to  15
Pool cap configured to  20
Running tasks:  14
Running tasks:  1
Running tasks:  2
Running tasks:  3
Running tasks:  4
Running tasks:  5
Running tasks:  6
Running tasks:  7
Running tasks:  8
Running tasks:  9
Running tasks:  10
Running tasks:  11
Running tasks:  12
Running tasks:  13
Running tasks:  14
Running tasks:  15
Running tasks:  16
Running tasks:  17
Running tasks:  18
Running tasks:  19
Pool cap configured to  25
Pool cap configured to  20
Pool cap configured to  25
Pool cap configured to  20
Pool cap configured to  25
Pool cap configured to  20
Pool cap configured to  25
Pool cap configured to  20
Pool cap configured to  25
Running tasks:  19
Running tasks:  1
Running tasks:  2
Running tasks:  3
Running tasks:  4
Running tasks:  5
Running tasks:  6
Running tasks:  7
Running tasks:  8
Running tasks:  9
Running tasks:  10
Running tasks:  11
Running tasks:  12
Running tasks:  13
Running tasks:  14
Running tasks:  15
Running tasks:  16
Running tasks:  17
Running tasks:  18
Running tasks:  19
Running tasks:  20
Running tasks:  21
Running tasks:  22
Running tasks:  23
Running tasks:  24
Pool cap configured to  30
Pool cap configured to  25
Pool cap configured to  30
Pool cap configured to  25
Pool cap configured to  30
Pool cap configured to  25
Pool cap configured to  30
Pool cap configured to  25
Pool cap configured to  30
Pool cap configured to  25
Running tasks:  24
Running tasks:  1
Running tasks:  2
Running tasks:  3
Running tasks:  4
Running tasks:  5
Running tasks:  6
Running tasks:  7
Running tasks:  8
Running tasks:  9
Running tasks:  10
Running tasks:  11
Running tasks:  12
Running tasks:  13
Running tasks:  14
Running tasks:  15
Running tasks:  16
Running tasks:  17
Running tasks:  18
Running tasks:  19
Running tasks:  20
Running tasks:  21
Running tasks:  22
Running tasks:  23
Running tasks:  24
Pool cap configured to  30
Pool cap configured to  25
Pool cap configured to  30
Pool cap configured to  25
Pool cap configured to  30
Pool cap configured to  25
Pool cap configured to  30
Pool cap configured to  25
Pool cap configured to  30
Running tasks:  24
Running tasks:  1
Running tasks:  2
Running tasks:  3
Running tasks:  4
Pool cap configured to  25
Pool cap configured to  20
Pool cap configured to  15
Pool cap configured to  10
```