package main
import "fmt"
import "time"
import "sync"

var wg sync.WaitGroup
var mu sync.Mutex
var balance float64 =1000

func deposit(amount float64){
mu.Lock()
fmt.Println("deposit started ")
amt1 := balance
time.Sleep(1 * time.Second)
amt1 = amt1 + amount 
balance = amt1
mu.Unlock()
fmt.Println("total deposited amount ", amt1)
fmt.Println("deposit ended")
wg.Done()
}
func withdraw(amount float64){
mu.Lock()
fmt.Println("********************")
fmt.Println("withdraw started")
amt1 := balance
time.Sleep(1 * time.Second)
amt1 = amt1 - amount 
balance = amt1
mu.Unlock()
fmt.Println("total withdraw amount ", amt1)
fmt.Println("withdraw ended")
wg.Done()
}

func main(){
fmt.Println("initial balance :", balance)
fmt.Println("********************")
wg.Add(1)
go deposit(100)
wg.Add(1)
go withdraw(100)
fmt.Println("********************")

fmt.Println("final balance :", balance)

}