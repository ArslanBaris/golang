package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// 1 -) Iki rakam arasında toplama, çıkarma ve çarpma
	// işleminin yapıldığı bir fonkiyon yazınız.

	x, y := 10, 4
	sum1, dif1, prod1 := calculation(x, y)
	fmt.Println("Toplam: ", sum1)
	fmt.Println("Fark: ", dif1)
	fmt.Println("Çarpım: ", prod1)

	// 2 -) Kullanıcı tarafından girilen nota göre geçtiniz
	// veya kaldınız geri dönüşünü yazdırınız.

	fmt.Print("Lütfen aldığınız notu giriniz: ")
	grade, _ := getGrade()
	var result string
	if grade >= 50 {
		result = "Geçtin"
	} else {
		result = "Kaldın"
	}
	fmt.Println(result)

	// 3 -) 1 ile 100 arasındaki bir sayıyı tahmin etme uygulaması
	// yazınız. Toplam tahmin hakkınız 10 olsun.

	target := numRand(1, 100)

	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkınız kaldı")
		fmt.Print("Lütfen tahmininizi yapınız: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println("Tahmininiz daha büyük, daha küçük bir sayı giriniz.")
		} else if num < target {
			fmt.Println("Tahmininiz daha küçük, daha büyük bir sayı giriniz.")
		} else {
			fmt.Println("Doğru Tahmin, hedef sayı ", target, " idi ", attempts, " seferde bulundunuz")
			break
		}

	}

}

func calculation(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2
	return sum, dif, prod
}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input) // string to int

	if err != nil {
		fmt.Println(err)
	}
	return num, nil
}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
