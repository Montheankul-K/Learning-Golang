package main

import "fmt"

func sumInt(nums []int) int {
	var sum int // Golang จะ initial variable ให้ = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumFloat64(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

type Number interface {
	int | float64
}

// Generics : function ที่รับ params ได้หลาย type
func sum[value Number](nums []value) value {
	// [value int | float64] = [value Number] ใช้งานร่วมกับ interface
	// รับ params เป็น int หรือ float64 ก็ได้
	var sum value
	for _, num := range nums {
		sum += num
	}
	return sum
}

type Game struct {
	Title    string
	Platform string
	Price    int
}

type Movie struct {
	Title string
	Price int
}

func (g Game) getPrice() int {
	return g.Price
}

func (m Movie) getPrice() int {
	return m.Price
}

type GameOrMovie interface {
	getPrice() int
}

func sumPrice[value GameOrMovie](objs []value) int {
	var result int
	for _, obj := range objs {
		result += obj.getPrice()
	}
	return result
}

func main() {
	// Not use generics
	numsInt := []int{1, 2, 3, 4, 5}
	numsFloat64 := []float64{1.1, 1.2, 3.3, 4.4, 5.5}
	fmt.Println(sumInt(numsInt))
	fmt.Println(sumFloat64(numsFloat64))

	// Use generics
	fmt.Println(sum(numsInt))
	fmt.Println(sum(numsFloat64))

	games := []Game{
		{
			Title:    "Minecraft",
			Platform: "PC",
			Price:    30},
		{
			Title:    "The Sims 4",
			Platform: "PC",
			Price:    40,
		},
	}
	movies := []Movie{
		{
			Title: "Archemy of Soul",
			Price: 50,
		},
		{
			Title: "The Jungle",
			Price: 50,
		},
	}

	fmt.Println(sumPrice(games))
	fmt.Println(sumPrice(movies))
}
