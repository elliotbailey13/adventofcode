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

  Gamma := ""
  Epsilon := ""
  var Readings_Array [][]string
  scanner := bufio.NewScanner(Readings)

  for scanner.Scan() {
    Reading_Array := strings.Split(scanner.Text(), "")
    Readings_Array = append(Readings_Array, Reading_Array)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  Gamma = Calculate_Code(Readings_Array, 0)
  Epsilon = Calculate_Code(Readings_Array, 1)

  log.Printf("Gamma Reading: %v", Gamma)
  log.Printf("Epsilon Reading: %v", Epsilon)

  Gamma_Int, _ := strconv.ParseInt(Gamma, 2, 64)
  Epsilon_Int, _ := strconv.ParseInt(Epsilon, 2, 64)

  log.Printf("Gamma Int Reading: %v", Gamma_Int)
  log.Printf("Epsilon Int Reading: %v", Epsilon_Int)

  Power_Consumption := Gamma_Int * Epsilon_Int
  log.Printf("Submarine power consumption: %v", Power_Consumption)
}

func Calculate_Code(Readings [][]string, Mode int) string {
  Final_Code := ""

  Bit_Location := 0
  for Bit_Location != 12 {
    Ones := 0
    Zeros := 0
    for _, Value := range Readings {
      if Value[Bit_Location] == "1" { Ones++
      } else { Zeros++ }
    }
    if Mode == 0 {
      if Ones > Zeros { Final_Code += "1"
      } else { Final_Code += "0" }
    } else {
      if Ones > Zeros { Final_Code += "0"
      } else { Final_Code += "1" }
    }

    Bit_Location++
  }

  return Final_Code
}
