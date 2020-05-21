package problems 

func replaceElements(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }
    
    var currMax = -1
    for i := len(arr)-1; i >= 0; i-- {
        if arr[i] > currMax {
            arr[i], currMax = currMax, arr[i]
        }else {
            arr[i] = currMax
        }
    }
    return arr
}