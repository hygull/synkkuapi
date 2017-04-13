package main

import "fmt"

func main() {
   var num,counter int
   totalPrimes := 0
   var primesSlice []int
   fmt.Print("Enter a number <...Upto which you wanna see all the primes...>  ");
   fmt.Scan(&num)
   if(num>=2){
        for i:=2;i<=num;i++{
            counter=0;
            for j:=2;j<=i;j++{   /*There's a different way to minimize time complexity of this program*/
                if i%j==0{       /*You have to implement that program...this is simple one*/
                    counter+=2;
                }
            }
            if (counter==2){
                fmt.Printf("%d%s",i," ");
                primesSlice=append(primesSlice,i);
                totalPrimes+=1
            }
        }
        fmt.Println("\n"); /* Newline */
        fmt.Printf("%v%s",primesSlice,"\n")
        fmt.Printf("\nTotal number of primes till %d : %d\n",num,totalPrimes);
   }else{
    fmt.Println("\nThis is not a proper number.Please specify a number n, where n>=2")
   }
}
