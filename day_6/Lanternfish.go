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

  var Lanternfish []int
  for scanner.Scan() {
    Lanternfish_Line := scanner.Text()
    Lanternfish_Strings := strings.Split(Lanternfish_Line, ",")
    for _, Remaining_Days_String := range Lanternfish_Strings {
      Remaining_Days, _ := strconv.Atoi(Remaining_Days_String)
      Lanternfish = append(Lanternfish, Remaining_Days)
    }
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }
  Lanternfish_Input.Close()

  Days_Remaining := 80
  for Days_Remaining != 0 {
    Calculate_Daily_Fish_Change(&Lanternfish)
    Days_Remaining--
  }

  log.Printf("Total Number of Lanternfish: %v", len(Lanternfish))
}

func Calculate_Daily_Fish_Change(Lanternfish *[]int) {
  New_Fish := 0
  for Fish_Id, Fish_Day := range *Lanternfish {
    if Fish_Day == 0 {
      (*Lanternfish)[Fish_Id] = 6
      New_Fish++
    } else {
      (*Lanternfish)[Fish_Id]--
    }
//    log.Printf("Fish #%v - Days remaining %v", Fish_Id, Fish_Day)
  }

  for New_Fish != 0 {
    *Lanternfish = append(*Lanternfish, 8)
    New_Fish--
  }
}
