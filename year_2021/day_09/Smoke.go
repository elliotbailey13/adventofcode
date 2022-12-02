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

  var Converted_Int_Array [][]int
  scanner := bufio.NewScanner(Raw_Input)
  for scanner.Scan() {
    Raw_String_Row := scanner.Text()
    Raw_String_Split := strings.Split(Raw_String_Row, "")
    var Converted_Int_Row_Array []int
    for _, Raw_String_Value := range Raw_String_Split {
      Raw_Int_Value, _ := strconv.Atoi(Raw_String_Value)
      Converted_Int_Row_Array = append(Converted_Int_Row_Array, Raw_Int_Value)
    }

    Converted_Int_Array = append(Converted_Int_Array, Converted_Int_Row_Array)
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }

  Raw_Input.Close()

  var Total_Lowest_Numbers int
  for Line, Row := range Converted_Int_Array {
    for Column, Value := range Row {
      Lowest := true
      if Line != 0 {
        if Value > Converted_Int_Array[ Line - 1 ][ Column ] { Lowest = false }
        if Value == Converted_Int_Array[ Line - 1 ][ Column ] { Lowest = false }
      }
      if Column != 0 {
        if Value > Converted_Int_Array[ Line ][ Column - 1 ] { Lowest = false }
        if Value == Converted_Int_Array[ Line ][ Column - 1 ] { Lowest = false }
      }
      if Column != 99 {
        if Value > Converted_Int_Array[ Line ][ Column + 1 ] { Lowest = false }
        if Value == Converted_Int_Array[ Line ][ Column + 1 ] { Lowest = false }
      }
      if Line != 99 {
        if Value > Converted_Int_Array[ Line + 1 ][ Column ] { Lowest = false }
        if Value == Converted_Int_Array[ Line + 1 ][ Column ] { Lowest = false }
      }

      if Lowest == true {
        if Line != 0 && Column != 0 && Column != 99 && Line != 99 {
          log.Printf(" %v ", Converted_Int_Array[ Line - 1 ][ Column ])
          log.Printf("%v%v%v", Converted_Int_Array[ Line ][ Column - 1 ], Value,  Converted_Int_Array[ Line ][ Column + 1 ])
          log.Printf(" %v ", Converted_Int_Array[ Line + 1 ][ Column ])
        }
        Total_Lowest_Numbers += Value + 1
      }
    }
//    log.Printf("Line: %v, %+v", Line, Value)
  }

  log.Printf("Total of all low points: %v", Total_Lowest_Numbers)
}
