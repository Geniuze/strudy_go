package qsort

func quicksort(values [] int) []int {
    if len(values) == 1 {
        return values
    }

    left := make([]int, 0)
    right := make([]int, 0)
    for _,value := range values {
        if value >= values[0] {
            right = append(right, value)
        } else {
            left = append(left, value)
        }
    }
    return append(quicksort(left), quicksort(right)...)
}

func quickSort(values []int, left, right int){
    tmp := values[left]
    i,j := left, right
    p := left

    for i<=j {
        if j>=p && values[j] >= tmp {
            j--
        }
        if j>=p {
            values[p] = values[j]
            p = j
        }

        if i<=p && values[i] <= tmp {
            i++
        }
        if i<=p {
            values[p] = values[i]
            p = i
        }
    }
    values[p] = tmp

    if p - left > 1 {
        quickSort(values, left, p - 1)
    }
    if right - p > 1 {
        quickSort(values, p + 1, right)
    }
}

func QuickSort(values [] int) {
    //return quicksort(values)
    quickSort(values, 0, len(values)-1)
}
