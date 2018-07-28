package heapsort

import "testing"

func TestHeapSort1(t *testing.T) {
    a := []int{5,6,3,4,2,8,9}
    HeapSort(a)

    if a[0] != 2 || a[1] != 3 || a[2] != 4 || a[3] != 5 || a[4] != 6 ||  a[5] != 8 || a[6] != 9 {
        t.Error("HeapSort failed. expect 2,3,4,5,6,8,9 Got", a)
    }
}

func TestHeapSort2(t *testing.T) {
    a := []int{4}
    HeapSort(a)

    if a[0] != 4 {
        t.Error("HeapSort failed. expect 4 Got", a)
    }
}
