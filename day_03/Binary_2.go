package main

import (
  "os"
  "log"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  Readings, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }
  defer Readings.Close()

  CO2 := ""
  Oxygen := ""
  var Readings_Array [][]string
  scanner := bufio.NewScanner(Readings)

  for scanner.Scan() {
    Reading_Array := strings.Split(scanner.Text(), "")
    Readings_Array = append(Readings_Array, Reading_Array)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  CO2 = Calculate_Code(Readings_Array, 1)
  Oxygen = Calculate_Code(Readings_Array, 0)

  log.Printf("CO2 Reading: %v", CO2)
  log.Printf("Oxygen Reading: %v", Oxygen)

  CO2_Int, _ := strconv.ParseInt(CO2, 2, 64)
  Oxygen_Int, _ := strconv.ParseInt(Oxygen, 2, 64)

  log.Printf("CO2 Int Reading: %v", CO2_Int)
  log.Printf("Oxygen Int Reading: %v", Oxygen_Int)

  Life_Support := CO2_Int * Oxygen_Int
  log.Printf("Submarine Life Support: %v", Life_Support)
}

func Calculate_Code(Readings [][]string, Mode int) string {
  Bit_Location := 0
  var Working_Array [][]string

  for Bit_Location != 12 {
    Ones := 0
    Zeros := 0
    if len(Working_Array) > 0 {
      if len(Working_Array) == 1 { break }
      Readings = Working_Array
    }

    for _, Value := range Readings {
      if Value[Bit_Location] == "1" { Ones++
      } else { Zeros++ }
    }

    Keeper_Bit := ""
    if Mode == 0 {
      if Ones > Zeros { Keeper_Bit = "1"
      } else if Ones == Zeros { Keeper_Bit = "1"
      } else { Keeper_Bit = "0" }
    } else {
      if Ones > Zeros { Keeper_Bit = "0"
      } else if Ones == Zeros { Keeper_Bit = "0"
      } else { Keeper_Bit = "1" }
    }

    Working_Array = [][]string{}
    for _, Value := range Readings {
      if Value[Bit_Location] == Keeper_Bit {
        Working_Array = append(Working_Array, Value)
      }
    }

    Bit_Location++
  }

  Final_Code := ""
  for _, Value := range Working_Array[0] { Final_Code += Value }

  return Final_Code
}
