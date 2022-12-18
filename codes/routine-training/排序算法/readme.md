## benchmark result for bubbleSort & bubbleSortPro


### 结论
- bubbleSortPro 比 bubbleSort 快 2 倍左右
- bubbleSort和bubbleSortPro，多CPU内核对其执行时间无影响


```bash

(base)  ~/algorithm/Golang-algorithm/codes/routine-training/排序算法/ [main*] go test  -bench=. -benchtime=2s -count=2 -cpu=1,2,4 
goos: darwin
goarch: arm64
BenchmarkBubbleSort                     56781673                42.11 ns/op
BenchmarkBubbleSort                     55646796                41.90 ns/op
BenchmarkBubbleSort-2                   55514433                42.14 ns/op
BenchmarkBubbleSort-2                   55747617                41.91 ns/op
BenchmarkBubbleSort-4                   55622454                42.03 ns/op
BenchmarkBubbleSort-4                   55734456                41.92 ns/op
BenchmarkBubbleSortSorted               66508324                32.98 ns/op
BenchmarkBubbleSortSorted               70149030                32.94 ns/op
BenchmarkBubbleSortSorted-2             70846702                32.93 ns/op
BenchmarkBubbleSortSorted-2             69791025                32.92 ns/op
BenchmarkBubbleSortSorted-4             70404200                33.17 ns/op
BenchmarkBubbleSortSorted-4             70758888                32.98 ns/op
BenchmarkBubbleSortPro                  1000000000               0.0000001 ns/op
BenchmarkBubbleSortPro                  1000000000
BenchmarkBubbleSortPro-2                1000000000               0.0000001 ns/op
BenchmarkBubbleSortPro-2                1000000000               0.0000001 ns/op
BenchmarkBubbleSortPro-4                1000000000               0.0000001 ns/op
BenchmarkBubbleSortPro-4                1000000000               0.0000000 ns/op
BenchmarkBubbleSortProSorted            1000000000               0.0000035 ns/op
BenchmarkBubbleSortProSorted            1000000000               0.0000002 ns/op
BenchmarkBubbleSortProSorted-2          1000000000               0.0000001 ns/op
BenchmarkBubbleSortProSorted-2          1000000000               0.0000001 ns/op
BenchmarkBubbleSortProSorted-4          1000000000               0.0000000 ns/op
BenchmarkBubbleSortProSorted-4          1000000000               0.0000000 ns/op
PASS

```