package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  Lanternfish_Input, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }

  scanner := bufio.NewScanner(Lanternfish_Input)

  Lanternfish := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}
  for scanner.Scan() {
    Lanternfish_Line := scanner.Text()
    Lanternfish_Strings := strings.Split(Lanternfish_Line, ",")
    for _, Remaining_Days_String := range Lanternfish_Strings {
      Remaining_Days, _ := strconv.Atoi(Remaining_Days_String)
      Lanternfish[Remaining_Days]++
    }
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }
  Lanternfish_Input.Close()

  log.Printf("Fish Count Array: %v:%v", Lanternfish, len(Lanternfish))

  var Total_Fish int64
  Days_Remaining := 256
  for Days_Remaining != 0 {
    Total_Fish = Calculate_Daily_Fish_Change(&Lanternfish)
    log.Printf("Day #%v - Total Fish: %v", Days_Remaining, Total_Fish)
    Days_Remaining--
  }

  log.Printf("Total Number of Lanternfish: %v", Total_Fish)
}

func Calculate_Daily_Fish_Change(Lanternfish *[]int64) int64 {
  New_Fish_Totals := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}

  var Reset_Fish int64
  for Fish_Day, Fish_Count := range *Lanternfish {
    if Fish_Day == 0 {
      New_Fish_Totals[8] = Fish_Count
      Reset_Fish = Fish_Count
    } else {
      if Fish_Day == 7 { Fish_Count += Reset_Fish }

      New_Fish_Totals[Fish_Day - 1] = Fish_Count
    }
  }

  var Total_Fish int64
  for _, Fish_Count := range New_Fish_Totals { Total_Fish += Fish_Count }

  *Lanternfish = New_Fish_Totals

  return Total_Fish
}
