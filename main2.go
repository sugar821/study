package main
import(
	"fmt"
	"math/rand"
)
func main(){


fmt.Println("Spaceline         Days  TripType  Price     ")
fmt.Println("-------------------------------------------------")


for count := 0; count < 10; count++{
// spaceline
var check_spaceline = rand.Intn(3)
var spaceline = ""

// 飛行速度１６km〜３０kmのランダム
var speed = rand.Intn(14)+16
var days = 62100000/(speed*60*60*24)

// 0なら片道、1なら往復
var trip_type = rand.Intn(2)

// 料金は速度比例で３６〜５０m$
var price = (speed + 20)

switch{
	case check_spaceline == 0 :
			spaceline = "Virgin Galactic"
	case 	check_spaceline == 1 :
			spaceline = "SpaceX"
	case 	check_spaceline == 2 :
			spaceline = "spaceX adventures"
}
fmt.Printf("%-20v/",spaceline)

fmt.Printf("%-4vdays/",days)

if trip_type == 0{
	fmt.Printf("%-15v/","one-way")
}else{
	fmt.Printf("%-15v/","round-trip")
}
fmt.Printf("$%4v",price)
fmt.Println("")

}
}
