package main

import "golesson/interfaces"

func main() {
	//variables.Demo1()
	//fmt.Print()

	//conditionals.Demo1()
	//conditionals.Demo2()
	//conditionals.Demo3()

	//loops.Demo1()

	//arrays.Demo4()

	//slices.Demo2()

	//functions.SayHello("Cevdet")
	//var result = functions.Add(11, 19)
	//fmt.Println(result * 11)
	//var result2 = functions.Division(93, 31)
	//fmt.Println("Division result is :", result2)

	// result1, result2, result3, resul4 := functions.FourTrans(15, 4) // If we want dont take the any result we can replace the appointment _
	// fmt.Println("Addition is:", result1)
	// fmt.Println("Substraction is: ", result2)
	// fmt.Println("Multiplication is: ", result3)
	// fmt.Println("Division is: ", resul4)

	// var result = functions.AddVariadic(1, 4, 6, 8, 3)
	// fmt.Println("Addition is: ", result)

	// var result2 = functions.AddVariadic(5, 3, 9, 7)
	// fmt.Println("Addition is:", result2)

	// exArray := []int{2, 8, 7, 3, 9, 10}
	// var result3 = functions.AddVariadic(exArray...)
	// fmt.Println("Addition is: ", result3)

	//maps.Demo1()

	//Range.Demo1()
	//Range.Demo2()
	//Range.Demo3()

	//num := 18
	//pointers.Demo1(&num)
	//fmt.Println("Result in main: ", num)

	// numbers := []int{0, 1, 2}
	// pointers.Demo2(numbers)
	// fmt.Println("Result in main :", numbers[0])

	//structs.Demo1()

	//structs.Demo2()

	/*	 Routines
	//go goroutines.EvenNumbers()
	//go goroutines.OddNumbers()

	//time.Sleep(7 * time.Second)
	//fmt.Println("Main over")
	*/

	/*	Channels
		evenNumberCn := make(chan int)
		oddNumberCn := make(chan int)

		go channels.EvenNumbers(evenNumberCn)
		go channels.OddNumbers(oddNumberCn)

		evenNumberSum, oddNumberSum := <-evenNumberCn, <-oddNumberCn

		multi := evenNumberSum * oddNumberSum
		fmt.Println("Multiplication: ", multi)
	*/

	interfaces.Demo1()

	interfaces.Demo2()

	interfaces.Demo3()
}
