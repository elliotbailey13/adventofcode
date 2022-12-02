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

  _ = Drawings
//  log.Printf("Drawings: %v", Drawings)


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
  var Card_Number int
  var Bingo_Number int
  for _, Drawing := range Drawings {
    Bingo, Card_Number = Call_Bingo_Number(&Cards, &Match_Cards, Drawing)
    if Bingo == true {
      Bingo_Number, _ = strconv.Atoi(Drawing)
      break
    }
  }

  log.Printf("Bingo!")
  log.Printf("Bingo Number: %v", Bingo_Number)
  log.Printf("Card: %v", Card_Number)
  log.Printf("Card Values: %v", Cards[Card_Number])
  log.Printf("Match Card Values: %v", Match_Cards[Card_Number])
  Final_Score := Final_Score_Logic(Cards[Card_Number], Match_Cards[Card_Number], Bingo_Number)
  log.Printf("Final Score: %v", Final_Score)
}

func Call_Bingo_Number(Cards *[][][]string, Match_Cards *[][][]int, Number string) (bool, int) {
  log.Printf("Number: %v", Number)

  for Card_Number, Card := range *Cards {

    for Row_Number, Row := range Card {

      for Column_Number, Column := range Row {

	if Column == Number {

          (*Match_Cards)[Card_Number][Row_Number][Column_Number] = 1

	  Bingo := Check_Card(Card_Number, Match_Cards)
	  if Bingo == true { return true, Card_Number }
        }
      }
    }
  }

  return false, 0
}

func Check_Card(Card_Number int, Match_Cards *[][][]int) bool {

  for Card_Number, Card := range *Match_Cards {

    for Row_Number, Row := range Card {

      Row_Total := 0
      Column_Total := 0
      for Column_Number, Column := range Row {
        Row_Total += Column
        Column_Total += (*Match_Cards)[Card_Number][Column_Number][Row_Number]
      }
      if Row_Total == 5 || Column_Total == 5 { return true }
    }
  }

  return false
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
