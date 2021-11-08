package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	result, err := evenNum(11)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Girdiğiniz sayı: ", result)
	}

	result1, err := sRoot(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1)
	}

	result3 := sRoot2(-4)
	{
		fmt.Println(result3)
	}

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız", file)
	}

}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("HATA : Cift sayi girilmedi !")
	}
	return num, nil // Baslangic degeri olmayan ifade
}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negatif sayıların karekökü alınamaz")
	}
	return math.Sqrt(num), nil
}

func sRoot2(num float64) float64 {
	return math.Sqrt(num)
}
