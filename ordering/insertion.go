package ordering


func Insertion(param []int){
    
    var key,j int

    for i := 1; i < len(param); i++{
        key=param[i];
        j=i-1;

        for (j>=0 && param[j]>key){
            param[j+1]=param[j];
            j-=1
        }
        param[j+1]=key
    }
}
