package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  Drawings_Input, err := os.Open("./Drawings.txt")
  if err != nil { log.Fatal(err) }

  scanner := bufio.NewScanner(Drawings_Input)

  var Drawings []string
  for scanner.Scan() {
    Drawings = strings.Split(scanner.Text(), ",")
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }
  Drawings_Input.Close()

  Cards_Input, err := os.Open("./Cards.txt")
  if err != nil { log.Fatal(err) }
  defer Cards_Input.Close()

  scanner = bufio.NewScanner(Cards_Input)

  var Card [][]string
  var Cards [][][]string
  var Match_Card [][]int
  var Match_Cards [][][]int
  for scanner.Scan() {
    Line := scanner.Text()
    if Line == "" {
      Cards = append(Cards, Card)
      Match_Cards = append(Match_Cards, Match_Card)
      Card = [][]string{}
      Match_Card = [][]int{}
      continue
    }

    Card = append(Card, strings.Fields(Line))
    Match_Card = append(Match_Card, []int{0, 0, 0, 0, 0})
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }
  Cards_Input.Close()

  var Bingo bool
  var Last_Bingo int
  var Card_Number int
  var Bingo_Number int
  var Last_Bingo_Card [][]string
  var Last_Bingo_Match_Card [][]int
  for _, Drawing := range Drawings {
    Call_Bingo_Number(&Cards, &Match_Cards, Drawing)
    log.Printf("Number: %v", Drawing)


    Bingo, Card_Number = Check_Cards(&Match_Cards)

    for Bingo == true {
      Last_Bingo = Card_Number
      Last_Bingo_Card = Cards[Card_Number]
      Last_Bingo_Match_Card = Match_Cards[Card_Number]
      Bingo_Number, _ = strconv.Atoi(Drawing)

//      Cards = append(Cards[:Card_Number], Cards[Card_Number + 1:]...)
//      Match_Cards = append(Match_Cards[:Card_Number], Match_Cards[Card_Number + 1:]...)

      copy(Cards[Card_Number:], Cards[Card_Number+1:])
      Cards[len(Cards) - 1] = [][]string{}
      Cards = Cards[:len(Cards)-1]
      copy(Match_Cards[Card_Number:], Match_Cards[Card_Number+1:])
      Match_Cards[len(Match_Cards) - 1] = [][]int{}
      Match_Cards = Match_Cards[:len(Match_Cards)-1]

      log.Print("")
      log.Printf("Bingo Number: %v", Bingo_Number)
      log.Printf("Card Values: %v", Last_Bingo_Card)
      log.Printf("Match Card Values: %v", Last_Bingo_Match_Card)
      log.Printf("Bingo on card: %v", Card_Number)
      log.Printf("Bingo on cards left: %v", len(Cards))

      Bingo, Card_Number = Check_Cards(&Match_Cards)
    }
  }

  log.Printf("Bingo!")
  log.Printf("Bingo Number: %v", Bingo_Number)
  log.Printf("Card: %v", Last_Bingo)
  log.Printf("Card Values: %v", Last_Bingo_Card)
  log.Printf("Match Card Values: %v", Last_Bingo_Match_Card)
  Final_Score := Final_Score_Logic(Last_Bingo_Card, Last_Bingo_Match_Card, Bingo_Number)
  log.Printf("Final Score: %v", Final_Score)
}

func Call_Bingo_Number(Cards *[][][]string, Match_Cards *[][][]int, Number string) {

  for Card_Number, Card := range *Cards {

    for Row_Number, Row := range Card {

      for Column_Number, Column := range Row {

	if Column == Number {

          (*Match_Cards)[Card_Number][Row_Number][Column_Number] = 1

        }
      }
    }
  }
}

func Check_Cards(Match_Cards *[][][]int) (bool, int) {

  for Card_Number, Card := range *Match_Cards {

    for Row_Number, Row := range Card {

      Row_Total := 0
      Column_Total := 0
      for Column_Number, Column := range Row {
        Row_Total += Column
        Column_Total += (*Match_Cards)[Card_Number][Column_Number][Row_Number]
      }
      if Row_Total == 5 || Column_Total == 5 { return true, Card_Number }
    }
  }

  return false, 0
}

func Final_Score_Logic(Card [][]string, Match_Card [][]int, Bingo_Number int) int {
  Column_Total := 0
  for Row_Number, Row := range Match_Card {

    for Column_Number, Column := range Row {
      if Column == 0 {
        Int_Value, _ := strconv.Atoi(Card[Row_Number][Column_Number])
        Column_Total += Int_Value
      }
    }
  }

  return Column_Total * Bingo_Number
}
