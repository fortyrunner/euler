package main

import (
        "euler"
        "fmt"
        "time"
)

func main() {

        fmt.Println("Problem 8. What is the largest product of 13 adjacent digits?")
        start := time.Now()

        const max = 1000
        const adjacent = 13
        const number = "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"

        if len(number) != max {
                println("Missing digit")
        }

        answer := 0

        for i := 0; i < max; i++ {

                // respect the end boundary of the array (last digit is zero - don't need it)
                // Note that checking for 0 speeds up from 19 microseconds to 16
                // Checking for 1 makes it slower... if conditions are pricey

                if (i + adjacent) < max {

                        inner := 1
                        for offset := 0; offset < adjacent; offset++ {

                                // Use Ascii magic to convert a character into a digit
                                // Tricks first learned in 1984 still work...

                                digit := euler.Ctoi(number[i+offset])

                                if digit == 0 {
                                        inner = 0
                                        break
                                }

                                inner = inner * digit
                                //fmt.Printf("%d\n", ch)

                        }

                        if inner > answer {
                                answer = inner
                        }

                        //fmt.Printf("%d\n", inner)

                }

        }

        fmt.Printf("-->Answer: The largest product is %d\n", answer)
        fmt.Printf("took=%s\n\n", time.Since(start))

}
