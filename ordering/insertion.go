package ordering


func InsertionSort(arr []int){
    for i:=1;i<len(arr); i++{
        var j,key int=i-1,arr[i]

        for (j>=0 && key<arr[j]){
            arr[j+1]=arr[j]    
            j--
        }
        arr[j+1]=key

    }
}


