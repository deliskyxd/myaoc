package main

//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
//Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
//Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
//Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
//Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
    red int = 12
    green int = 13
    blue int = 14
)

func calculate(filename string) int {
    // Zmienne
    var suma int
    suma = 0
    //var nr_gry int
    gra := map[string]int{
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    var power int
    power = 1

    plik, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    defer plik.Close()

    skaner := bufio.NewScanner(plik)

    // działanie w pliku
    for skaner.Scan(){
        tekst := strings.SplitAfter(skaner.Text(),":")

        // Zapisz numer gry
        //nr := strings.Split(tekst[0]," ") 
        //nr[1] = strings.Trim(nr[1],":")
        //nr_gry,_ = strconv.Atoi(nr[1])

        tekst = strings.Split(tekst[1],";")
        
        // Sprawdzanie pojedynczej gry
        for i := 0 ;i < len(tekst); i++ { // przechodzneie przez każde wyjęcie piłki
            single := strings.Split(tekst[i],",") // wynik pojedynczego wyjęcia

            // wyczyśc mapę przed przypisaniem 
            //delete(gra,"red")
            //delete(gra,"green")
            //delete(gra,"blue")

            for j := 0 ; j < len(single) ; j++ {
                temp := strings.Split(single[j]," ")

                //fmt.Println("temp")
                //for k,v := range temp {
                //    fmt.Println(k,"=",v)
                //}

                // SPRAWDZANIE POJEDYŃCZYCH WARTOŚCI I PRZYPISYWANIE DO MAPY
                // jeśli większa niż maksymalna
                // temp ma 2 miejsca w tabeli [ 1 = wynik, 2 = kolor]
                num,_ := strconv.Atoi(temp[1])
                if ( gra [temp[2]] < num){
                    gra[temp[2]] = num
                }
            }

            //fmt.Println("Wartość mapy")
            //for key,value := range gra {
            //    fmt.Println(key," = ",value)
            //}
            
        }

        // jeśli wszystkie wartości nie przekraczają granicznych, dodaj gre do sumy
        //if( gra["red"] <= red && gra["green"] <= green && gra["blue"] <= blue ){ 
        //    suma += nr_gry
        //}

        // power = wymnożone wszystkie kostki
        power = 1
        for _,v := range gra {
            power *= v
        }

        suma += power

        // wyczyść tabelę przed grą
        for k := range gra {
            delete(gra,k)
        }
    }
    // 

    return suma
}


func main() {
    // Zmienna ilość gier 
    filename := "input.txt"
    wynik := calculate(filename)
    fmt.Println("wynik = ",wynik)
}

