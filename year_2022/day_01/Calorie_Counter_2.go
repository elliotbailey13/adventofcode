package main

import (
  "os"
  "log"
  "sort"
  "bufio"
  "strconv"
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

  var Elf_Totals []int
  var Calorie_Total int
  var Largest_Holder int
  for _, Full_Input_Line := range Full_Input {
    if Full_Input_Line == "" {
      if Calorie_Total > Largest_Holder {
        log.Printf("New elf with the most calories. %v", Calorie_Total)
	Largest_Holder = Calorie_Total
      }
      Elf_Totals = append(Elf_Totals, Calorie_Total)
      Calorie_Total = 0
    }

    Calorie_Value, _ := strconv.Atoi(Full_Input_Line)
    Calorie_Total += Calorie_Value
  }

  sort.Ints(Elf_Totals)
  Total_Elves := len(Elf_Totals)
  Total_Of_Top_Three := Elf_Totals[Total_Elves - 1] + Elf_Totals[Total_Elves - 2] + Elf_Totals[Total_Elves - 3]
  log.Printf("1st: %v | 2nd: %v | 3rd: %v | Total: %v", Elf_Totals[Total_Elves - 1], Elf_Totals[Total_Elves - 2], Elf_Totals[Total_Elves - 3], Total_Of_Top_Three)
}
