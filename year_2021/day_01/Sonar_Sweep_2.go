package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
)

func main() {
  Sonar_Readings, err := os.Open("../input.txt")
  if err != nil { log.Fatal(err) }
  defer Sonar_Readings.Close()

  Runs := 0
  Sum_A := []int{}
  Sum_B := []int{}
  Sum_C := []int{}
  Sum_D := []int{}
  Compare_Mode := 0
  Increased_Counter := 0
  scanner := bufio.NewScanner(Sonar_Readings)

  for scanner.Scan() {
    Sonar_Reading, _ := strconv.Atoi(scanner.Text())
    Runs++

    if len(Sum_A) < 3 { Sum_A = append(Sum_A, Sonar_Reading) }

    if Runs > 1 {
      if len(Sum_B) < 3 { Sum_B = append(Sum_B, Sonar_Reading) }
    }

    if Runs > 2 {
      if len(Sum_C) < 3 { Sum_C = append(Sum_C, Sonar_Reading) }
    }

    if Runs > 3 {
      if len(Sum_D) < 3 { Sum_D = append(Sum_D, Sonar_Reading) }
    }

    if (len(Sum_A) + len(Sum_B)) == 6 && Compare_Mode == 0 {
      if Compare(Sum_A, Sum_B) { Increased_Counter++ }
      Compare_Mode = 1
      Sum_A = []int{}
    }

    if (len(Sum_B) + len(Sum_C)) == 6 && Compare_Mode == 1 {
      if Compare(Sum_B, Sum_C) { Increased_Counter++ }
      Compare_Mode = 2
      Sum_B = []int{}
    }

    if (len(Sum_C) + len(Sum_D)) == 6 && Compare_Mode == 2 {
      if Compare(Sum_C, Sum_D) { Increased_Counter++ }
      Compare_Mode = 3
      Sum_C = []int{}
    }

    if (len(Sum_D) + len(Sum_A)) == 6 && Compare_Mode == 3 {
      if Compare(Sum_D, Sum_A) { Increased_Counter++ }
      Compare_Mode = 0
      Sum_D = []int{}
    }
}

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  log.Printf("Total Number of Increased readings: %v", Increased_Counter)
}

func Compare(Array_1, Array_2 []int) bool {
//  log.Printf("Reading: %v : %v", Array_1, Array_2)

  if Sum_Array(Array_1) < Sum_Array(Array_2) {
    log.Printf("Reading: %v : %v (increased)", Sum_Array(Array_1), Sum_Array(Array_2))
    return true
  } else {
    log.Printf("Reading: %v : %v (decreased)", Sum_Array(Array_1), Sum_Array(Array_2))
    return false
  }
}

func Sum_Array(array []int) int {
 Result := 0
 for _, Value := range array { Result += Value }

 return Result
}
