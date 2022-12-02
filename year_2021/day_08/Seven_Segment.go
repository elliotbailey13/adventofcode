package main

import (
  "os"
  "log"
  "bufio"
//  "strconv"
  "strings"
)

func main() {
  Raw_Input, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }

  var Raw_String_Array []string
  scanner := bufio.NewScanner(Raw_Input)
  for scanner.Scan() {
    Raw_String_Array = append(Raw_String_Array, scanner.Text())
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }

  Raw_Input.Close()

//  var First_Signal_Display [][]string
  var Output_Signal_Display [][]string
  for _, Raw_String_Row := range Raw_String_Array {
    Row_Split := strings.Split(Raw_String_Row, " | ")
    Output_Signal_Display = append(Output_Signal_Display, strings.Split(Row_Split[1], " "))
  }

  Digit_Count := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
  for _, Output_Signal_Display_Row := range Output_Signal_Display {
    for _, Output_Signal_Display := range Output_Signal_Display_Row {
      switch len(Output_Signal_Display) {
        case 2:
          Digit_Count[1]++

        case 4:
          Digit_Count[4]++

        case 3:
          Digit_Count[7]++

        case 7:
          Digit_Count[8]++
      }
    }
  }
  log.Printf("Digit_Count: %+v", Digit_Count)

  var Total_Count int
  for Number, Digit_Count_Value := range Digit_Count {
    log.Printf("Number of %v: %v", Number, Digit_Count_Value)
    Total_Count += Digit_Count_Value
  }

  log.Printf("Total Count: %v", Total_Count)
}
