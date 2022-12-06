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

  Redundant_Pairs := 0
  var Elf_1_Sections, Elf_2_Sections []int
  for _, Full_Input_Line := range Full_Input {
    Sections := strings.Split(Full_Input_Line, ",")
    Elf_1_Sections_String := strings.Split(Sections[0], "-")
    Elf_2_Sections_String := strings.Split(Sections[1], "-")

    Elf_1_Sections_Temp_1, _ := strconv.Atoi(Elf_1_Sections_String[0])
    Elf_1_Sections_Temp_2, _ := strconv.Atoi(Elf_1_Sections_String[1])
    Elf_1_Sections = []int{ Elf_1_Sections_Temp_1, Elf_1_Sections_Temp_2 }

    Elf_2_Sections_Temp_1, _ := strconv.Atoi(Elf_2_Sections_String[0])
    Elf_2_Sections_Temp_2, _ := strconv.Atoi(Elf_2_Sections_String[1])
    Elf_2_Sections = []int{ Elf_2_Sections_Temp_1, Elf_2_Sections_Temp_2 }

    In_Range := false
    for _, Elf_1_Section := range Elf_1_Sections {
      if In_Range != true {
        if Elf_1_Section >= Elf_2_Sections[0] && Elf_1_Section <= Elf_2_Sections[1] {
          In_Range = true
          continue
        }
      }
    }

    if In_Range == true {
      Redundant_Pairs += 1
      continue
    }

    In_Range = false
    for _, Elf_2_Section := range Elf_2_Sections {
      if In_Range != true {
        if Elf_2_Section >= Elf_1_Sections[0] && Elf_2_Section <= Elf_1_Sections[1] {
          In_Range = true
          continue
        }
      }
    }

    if In_Range == true { Redundant_Pairs += 1 }
  }

  log.Printf("Redundant Pairs: %v", Redundant_Pairs)
}
