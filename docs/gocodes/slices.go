package main

import "fmt"
import "reflect"

func main(){
	numsSlice:=make([]int,2,3);   /*[0,0]*/
	//var numsSlice []int
	fmt.Printf("%v %v %v\n",	numsSlice,len(numsSlice),cap(numsSlice));

	numsSlice=append(numsSlice,1);
	numsSlice=append(numsSlice,2);	
	fmt.Printf("%v %v %v\n",	numsSlice,len(numsSlice),cap(numsSlice));

	numsSlice=append(numsSlice,3);
	numsSlice=append(numsSlice,4);
	fmt.Printf("%v %v %v\n",	numsSlice,len(numsSlice),cap(numsSlice));

	reflect.ValueOf(&numsSlice).Elem().SetCap(6)	/*	SetCap(10) or SetCap(4) will give an error...as numsSlice contains 4 elements  */
	fmt.Printf("%v %v %v\n",	numsSlice,len(numsSlice),cap(numsSlice));
}