package ordering

func BubbleSort(arr []int){
    if (len(arr) == 1){
        return
    }
    var order bool = false
    for !order {
        for i:=1; i<len(arr);i++ {
            order=true
            if arr[i]<arr[i-1]{
                order = false
                var aux = arr[i]
                arr[i]=arr[i-1]
                arr[i-1]=aux
            }              
        }  
    }
        
}
