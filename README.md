# f
f is a functional slice library

```go
import "github.com/mrmiguu/f"

increaseIntSlice := func(i []int) []int {
    // int slice map    
    return f.MapInt(i, increaseInt)
}

printIntSlice := func(i []int) {
    // int slice each
    f.EachInt(i, print)
}

increaseInt := func(i int) int {
    return i + 1
}

print := func(i int) {
    fmt.Print(i, " ")
}

intSliceSlice := [][]int{
    {0, 1, 2},
    {3, 4, 5},
    {6, 7, 8},
}

// generic map
jntSliceSlice := f.Map(intSliceSlice, increaseIntSlice)

// generic each
f.Each(jntSliceSlice, printIntSlice)
```