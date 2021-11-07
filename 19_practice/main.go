package main

import "fmt"

func main() {

	// 1-) 1 ile 10 arasındaki sayıları if yapısıyla tek - çift olarak yazdırınız.

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "çittir")
		} else {
			fmt.Println(i, "tektir")
		}
	}

	// 2-) for yapısını kullanarak Go'da olmayan while döngüsüne örnek veriniz.

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// 3-) Switch fallthrough ifadesini açıklayınız.

	switch x := 75; {
	case x < 20:
		fmt.Printf("%d küçüktür 20\n", x)
		fallthrough
	case x < 50:
		fmt.Printf("%d küçüktür 50\n", x)
		fallthrough
	case x < 100:
		fmt.Printf("%d küçüktür 100\n", x)
		fallthrough
	case x < 200:
		fmt.Printf("%d küçüktür 200\n", x)
	}

	// 4-) Aşağıdaki if döngüsünü daha idiomatic hale getiriniz.

	/*
		if y := 20; y%2 == 0 {
			fmt.Println(y, "çifttir")
		} else {
			fmt.Println(y, "tektir")
		}
	*/

	y := 20
	if x%2 == 0 {
		fmt.Println(y, "çifttir")
		return
	}
	fmt.Println(y, "tektir")

	// 5-) 1 ile 50 arasındaki asal sayıları gösteren bir program yazınız.

	var i, j int

	for i = 2; i < 50; i++ {
		for j = 2; j < (i / j); j++ {
			if i%j == 0 {
				break
			}
		}

		if j > (i / j) {
			fmt.Printf("%d bir asal sayıdır\n", i)
		}
	}

}
