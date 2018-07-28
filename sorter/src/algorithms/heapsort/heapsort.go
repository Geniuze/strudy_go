package heapsort

//import "fmt"


func adjustHeap(a []int, i int, length int) {
    tmp := a[i]

    for k:=i*2+1; k<length; k=k*2+1 {
        if k+1 < length && a[k] < a[k+1] {
            k++
        }
        if tmp < a[k] {
            a[i] = a[k]
            i = k
        } else {
            break
        }
    }
    a[i] = tmp
}

func HeapSort(a []int) {
    for i:=len(a)/2-1; i>=0; i-- {
        // 构建初始大顶椎 (a[k]>a[k*2+1] && a[k]>a[k*2+2])
        adjustHeap(a, i, len(a))
    }

    for i:=len(a)-1; i>=0; i-- {
        a[0], a[i] = a[i], a[0]

        adjustHeap(a, 0, i)
    }
}
