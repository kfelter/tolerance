# tolerance
fault tolerant math operations

Inspired by [this](https://www.youtube.com/watch?v=N5faA2MZ6jY) video about how spaces handles bit flips caused by radiation while doing calculations in space.

They handle the faults by having 3x dual core cpus and redundant RAM. Each cpu is run in single core mode, so if a math calculation needs to be done, it runs on all 6 cpus at the same time then the output is compared. This tolerance package simulates that by running concurrent go routines and storing the results in a voting map. Then we check the values in the voting map, if a value has more than 50% of the votes it is returned, retry otherwise.

# cli
Summing
`go run cmd/math/main.go add 1 2 3 4`
Result
```
2020/06/12 01:17:11 Using 8 redundant processes
2020/06/12 01:17:11 System is 50 % faulty
2020/06/12 01:17:11 worker [3] votes 56.607648896720676
2020/06/12 01:17:11 worker [4] votes 97.55460600371605
2020/06/12 01:17:11 worker [0] votes 133.07510238842693
2020/06/12 01:17:11 worker [2] votes 10
2020/06/12 01:17:11 worker [1] votes 76.37160421475915
2020/06/12 01:17:11 worker [7] votes 47.18367995791324
2020/06/12 01:17:11 worker [5] votes 0.9303107067502085
2020/06/12 01:17:11 worker [6] votes 98.87671900825038
2020/06/12 01:17:11 fualt on more than 3 / 8 redundant operations - retry
...
2020/06/12 01:19:14 result 10 has 5 votes
result [10] attempts [3] faulty% [50] time [467.632µs]
```

Subtract b from a
`go run cmd/math/main.go sub 6 4`
Result
```
2020/06/12 01:19:14 args [/var/folders/xq/y3xb0ppd0998dtk6wb0b6gtc0000gp/T/go-build634907364/b001/exe/main sub 6 4]
2020/06/12 01:19:14 worker [0] votes -54.10765798457902
2020/06/12 01:19:14 worker [4] votes 10.129002518490317
2020/06/12 01:19:14 worker [2] votes 13.007430254329215
2020/06/12 01:19:14 worker [6] votes 82.39600257429571
2020/06/12 01:19:14 worker [1] votes -1.391810642436198
2020/06/12 01:19:14 worker [3] votes 13.419719834413598
2020/06/12 01:19:14 worker [7] votes 2
2020/06/12 01:19:14 worker [5] votes -20.821590204446842
2020/06/12 01:19:14 fualt on more than 3 / 8 redundant operations - retry
...
2020/06/12 01:19:14 result 2 has 5 votes
result [2] attempts [3] faulty% [50] time [686.352µs]
```

Factorial N
`go run cmd/math/main.go fact 5`
Result
```
2020/06/12 01:20:42 Using 8 redundant processes
2020/06/12 01:20:42 System is 50 % faulty
2020/06/12 01:20:42 args [/var/folders/xq/y3xb0ppd0998dtk6wb0b6gtc0000gp/T/go-build344055400/b001/exe/main fact 5]
2020/06/12 01:20:42 worker [3] votes 8.017659978030324e+139
2020/06/12 01:20:42 worker [4] votes 1.6126493842642177e+22
2020/06/12 01:20:42 worker [6] votes 120
2020/06/12 01:20:42 worker [7] votes 120
2020/06/12 01:20:42 worker [1] votes 4.5714423403637087e+151
2020/06/12 01:20:42 worker [0] votes 194652.0018425049
2020/06/12 01:20:42 worker [5] votes 5.4216144359793714e+156
2020/06/12 01:20:42 worker [2] votes 120
2020/06/12 01:20:42 fualt on more than 3 / 8 redundant operations - retry
2020/06/12 01:20:42 worker [3] votes 4.687578040005233e+53
2020/06/12 01:20:42 worker [5] votes 3.6004846883267537e+110
2020/06/12 01:20:42 worker [4] votes 120
2020/06/12 01:20:42 worker [2] votes 120
2020/06/12 01:20:42 worker [1] votes 1.4470745703570386e+18
2020/06/12 01:20:42 worker [6] votes 6.040334037793743e+23
2020/06/12 01:20:42 worker [0] votes 1.4805423328394172e+12
2020/06/12 01:20:42 worker [7] votes 7.385452901542006e+134
2020/06/12 01:20:42 fualt on more than 3 / 8 redundant operations - retry
2020/06/12 01:20:42 worker [1] votes 2.382263804067538e+94
2020/06/12 01:20:42 worker [4] votes 9.006615120197429e+147
2020/06/12 01:20:42 worker [0] votes 0.8133489942197941
2020/06/12 01:20:42 worker [2] votes 120
2020/06/12 01:20:42 worker [6] votes 120
2020/06/12 01:20:42 worker [7] votes 1.0333633489152708e+89
2020/06/12 01:20:42 worker [5] votes 2.0041476747139956e+83
2020/06/12 01:20:42 worker [3] votes 3.0089077389882377e+120
2020/06/12 01:20:42 fualt on more than 3 / 8 redundant operations - retry
2020/06/12 01:20:42 worker [6] votes 120
2020/06/12 01:20:42 worker [1] votes 4.2331898878731865e+118
2020/06/12 01:20:42 worker [3] votes 120
2020/06/12 01:20:42 worker [2] votes 9.6383316855897e+20
2020/06/12 01:20:42 worker [0] votes 120
2020/06/12 01:20:42 worker [5] votes 120
2020/06/12 01:20:42 worker [7] votes 2.3205191424360877e+120
2020/06/12 01:20:42 worker [4] votes 4.267598099484513e+144
2020/06/12 01:20:42 fualt on more than 3 / 8 redundant operations - retry
2020/06/12 01:20:42 worker [0] votes 1.3452075222709003e+65
2020/06/12 01:20:42 worker [3] votes 2.391294850546806e+96
2020/06/12 01:20:42 worker [2] votes 8.451726935279835e+64
2020/06/12 01:20:42 worker [1] votes 120
2020/06/12 01:20:42 worker [5] votes 120
2020/06/12 01:20:42 worker [4] votes 1.099478317188936e+105
2020/06/12 01:20:42 worker [6] votes 5631.050649841869
2020/06/12 01:20:42 worker [7] votes 1.4042554811835394e+19
2020/06/12 01:20:42 fualt on more than 3 / 8 redundant operations - retry
2020/06/12 01:20:42 worker [0] votes 7.376332685677801e+87
2020/06/12 01:20:42 worker [6] votes 1.6754177221353673e+124
2020/06/12 01:20:42 worker [7] votes 120
2020/06/12 01:20:42 worker [5] votes 120
2020/06/12 01:20:42 worker [3] votes 9.666950876005684e+138
2020/06/12 01:20:42 worker [2] votes 120
2020/06/12 01:20:42 worker [4] votes 120
2020/06/12 01:20:42 worker [1] votes 120
2020/06/12 01:20:42 result 120 has 5 votes
result [120] attempts [6] faulty% [50] time [1.084204ms]
```

