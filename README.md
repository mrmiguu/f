# f
f is a functional slice library

```
increaseInt := func(i int) int {
    return i + 1
}

increaseIntSlice := func(i []int) []int {
    return f.MapInt(i, increaseInt)
}

print := func(i int) {
    fmt.Print(i, " ")
}

printIntSlice := func(i []int) {
    f.Each(i, print)
}

intSliceSlice := [][]int{
    {0, 1, 2},
    {3, 4, 5},
    {6, 7, 8},
}

jntSliceSlice := f.Map(intSliceSlice, increaseIntSlice)
f.Each(jntSliceSlice, printIntSlice)
```