package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano",
		"Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin(users []string, coins int) int {
	for _,user := range users {
		length:=len(user)
		var count int = 0
		for i:=0;i<length;i++ {
			switch i {
			case 'e', 'E':
				count++
			case 'i', 'I':
				count = count + 2
			case  'o', 'O':
				count = count + 3
			case 'u', 'U':
				count = count + 4
			default:
				count = count
			}
		}
		distribution[user] = count
		coins = coins - count
	}
	return coins
}


func main() {
	var c rune
	c = '爱'
	fmt.Println(c)
	fmt.Printf("%T ,%c ,%v ,%d",c,c,c,c)
	var d byte
	d=24
	fmt.Printf("%T",d)
}

