package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  Raw_Input, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }

  var Full_Input []string
  scanner := bufio.NewScanner(Raw_Input)
  for scanner.Scan() {
    Full_Input = append(Full_Input, scanner.Text())
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }

  Raw_Input.Close()

  var Octo_Values [][]int
  for _, Row := range Full_Input {
    var Line_Values []int
    String_Values := strings.Split(Row, "")
    for _, String_Value := range String_Values {
      Int_Value, _ := strconv.Atoi(String_Value)
      Line_Values = append(Line_Values, Int_Value)
    }

    Octo_Values = append(Octo_Values, Line_Values)
  }

  for Index, Row := range Octo_Values {
    log.Printf("Row #%v: %v", Index, Row)
  }

  var Steps int
  var Flashes int
  var All_Flash bool
  var Total_Flashes int
  for All_Flash == false {
    Flashes, All_Flash = Octo_Step(&Octo_Values)
    log.Printf("Flashes this step (%v): %v (All %v)", Steps + 1, Flashes, All_Flash)
    Total_Flashes += Flashes
    Steps++
  }

  log.Printf("Total flashes: %v", Total_Flashes)
}

func Octo_Step(Octo_Values *[][]int) (int, bool) {
  var Octo_Flashes [][]bool
  for len(Octo_Flashes) != 10 {
    Octo_Flashes = append(Octo_Flashes, []bool{false, false, false, false, false, false, false, false, false, false})
  }

  for Row_Number, Row := range *Octo_Values {
    for Column_Number, _ := range Row {
      (*Octo_Values)[Row_Number][Column_Number]++
    }
  }

  Flashing := true
  var Flash_Total int
  for Flashing == true {
    Flashing = false
    for Row_Number, Row := range *Octo_Values {
      for Column_Number, Value := range Row {
        if Value > 9 {
          if Octo_Flashes[Row_Number][Column_Number] == true { continue }

          Flash_Total++
          Flashing = true
          Octo_Flashes[Row_Number][Column_Number] = true
          if Row_Number - 1 != -1 && Column_Number - 1 != -1 {
            (*Octo_Values)[Row_Number - 1][Column_Number - 1]++
          }
          if Row_Number - 1 != -1 {
            (*Octo_Values)[Row_Number - 1][Column_Number]++
          }
          if Row_Number - 1 != -1 && Column_Number + 1 != 10 {
            (*Octo_Values)[Row_Number - 1][Column_Number + 1]++
          }
          if Column_Number - 1 != -1 {
            (*Octo_Values)[Row_Number][Column_Number - 1]++
          }
          if Column_Number + 1 != 10 {
            (*Octo_Values)[Row_Number][Column_Number + 1]++
          }
          if Row_Number + 1 != 10 && Column_Number - 1 != -1 {
            (*Octo_Values)[Row_Number + 1][Column_Number - 1]++
          }
          if Row_Number + 1 != 10 {
            (*Octo_Values)[Row_Number + 1][Column_Number]++
          }
          if Row_Number + 1 != 10 && Column_Number + 1 != 10 {
            (*Octo_Values)[Row_Number + 1][Column_Number + 1]++
          }
        }
      }
    }
  }

  for Row_Number, Row := range *Octo_Values {
    for Column_Number, Value := range Row {
      if Value > 9 {
        (*Octo_Values)[Row_Number][Column_Number] = 0
      }
    }
  }

  All_Flash := true
  for _, Octo_Flash_Row := range Octo_Flashes {
    for _, Octo_Flash_Column := range Octo_Flash_Row {
      if Octo_Flash_Column == false { All_Flash = false }
    }
  }

  return Flash_Total, All_Flash
}
