package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var answer int 

var numbers []string = []string{
    "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}


func znajdzDlugosc(nazwa string) {
    // Puszczam dwa skanery od początku i końca, które jak spotkają liczbę
    // to zapisują ją do zmiennych które potem łączę w jedną liczbę

    // Otwarcie pliku
    plik, err := os.Open(nazwa)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer plik.Close()

    // odczytanie linii z pliku, sprawdzenie rozmiaru
    skaner := bufio.NewScanner(plik)
    for skaner.Scan() {
        rozmiar := len(skaner.Text()) // rozmiar tekstu ze skanera
        odczytane := skaner.Text() // odczytany tekst -//-

        // Znalezione liczby
        var first string // pierwsza liczba int
        var second string // ostatnia liczba int

        // Znalezione stringi
        var min int // pozycja pierwszy string
        var num1 string // pierwszy 
        var max int // pozycja ostatni string
        var num2 string // ostatni

        min = rozmiar
        max = -1

        // Przeszukaj string w poszukiwaniu liczb , zapisz gdzie pierwsza i ostatnia
        for i := 0 ; i < 10 ; i++ { // pierwsza liczba (string)
            first := strings.Index(odczytane, numbers[i]) 
            last := strings.LastIndex(odczytane, numbers[i]) 
           // fmt.Println("first = ",first)
           // fmt.Println("last = ",last)
           // fmt.Println("number = ",numbers[i])
            if( first != -1 && first <= min  ) {
                min = first 
                num1 = strconv.Itoa(i)
                // Usuń tę liczbę z linijki żeby się nie zdublowała
            }

            if( last != -1 && last > max ) {
                max = last 
                num2 = strconv.Itoa(i)
            }
        }

       // fmt.Println("min = ",min)
       // fmt.Println("max = ",max)

        // Porównywanie pozycji pierwszego inta i stringa, który pierwszy
        //przekonwertowanie jeśli to string na liczbę

        for i := 0 ; i < rozmiar ; i++ {
            // jeśli to jest liczba -> dodaj ją
            // liczby w kodzie ascii mają numer dziesiętny od 49 (0) do 57 (9)
            if( int(odczytane[i]) >= 49 && int(odczytane[i]) <= 57) {
                //fmt.Println("First = ", string(odczytane[i]) )
                if ( min < i ) {
                    first = string(num1) 
                    break
                }
                first = string(odczytane[i]) 
                // Czy liczba jest pierwsza czy string?
                break
            }
        }

        for i := rozmiar-1 ; i >= 0 ; i-- {
            if( int(odczytane[i]) >= 49 && int(odczytane[i]) <= 57) {
                //fmt.Println("Second = ", string(odczytane[i]) )
                if ( max > i ) {
                    // jeśli min < i to pierwsza liczba to string
                     second = string(num2)
                     break
                }
                second = string(odczytane[i])
                // Czy liczba jest pierwsza czy string?
                break
            }
        }
        //fmt.Println("'",skaner.Text(),"'")

        // suma stringów
        sum := first + second
        //fmt.Println("sum = ",first,"+",second,"=",sum)
        // konwertowanie ASCII na liczbę
        temp,_ := strconv.Atoi(sum)
        //fmt.Println("wynik = ",temp)

        answer += temp
    }
}


func main() {
    filename := "input.txt"

    znajdzDlugosc(filename)
    fmt.Println(answer)
}
