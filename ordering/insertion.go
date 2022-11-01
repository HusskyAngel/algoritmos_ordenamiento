package ordering


func Insertion(arr []int){
    
    for i:=1; i< len(arr);i++{
        var key,j=arr[i],i-1 

        for (j>= 0 && arr[j]>key){
            arr[j+1]=arr[j]
            j-=1
        }
        arr[j+1]=key
    }
    return
}
