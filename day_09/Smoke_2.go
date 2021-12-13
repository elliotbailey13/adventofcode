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

  var Height_Map [][]int
  scanner := bufio.NewScanner(Raw_Input)
  for scanner.Scan() {
    Raw_String_Row := scanner.Text()
    Raw_String_Split := strings.Split(Raw_String_Row, "")
    var Converted_Int_Row_Array []int
    for _, Raw_String_Value := range Raw_String_Split {
      Raw_Int_Value, _ := strconv.Atoi(Raw_String_Value)
      Converted_Int_Row_Array = append(Converted_Int_Row_Array, Raw_Int_Value)
    }

    Height_Map = append(Height_Map, Converted_Int_Row_Array)
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }

  Raw_Input.Close()

  var Basins [][][]int
  for Line, Row := range Height_Map {
    for Column, Value := range Row {
      Lowest := true
      if Line != 0 {
        if Value > Height_Map[ Line - 1 ][ Column ] { Lowest = false }
        if Value == Height_Map[ Line - 1 ][ Column ] { Lowest = false }
      }
      if Column != 0 {
        if Value > Height_Map[ Line ][ Column - 1 ] { Lowest = false }
        if Value == Height_Map[ Line ][ Column - 1 ] { Lowest = false }
      }
      if Column != 99 {
        if Value > Height_Map[ Line ][ Column + 1 ] { Lowest = false }
        if Value == Height_Map[ Line ][ Column + 1 ] { Lowest = false }
      }
      if Line != 99 {
        if Value > Height_Map[ Line + 1 ][ Column ] { Lowest = false }
        if Value == Height_Map[ Line + 1 ][ Column ] { Lowest = false }
      }

      if Lowest == true {
        log.Printf("Finding basin at %v:%v", Line, Column)
        Basin_Locations := Find_Basin(&Height_Map, []int{Line, Column})
        log.Printf("Locations of Basin: %v", Basin_Locations)
        log.Printf("Size of Basin: %v", len(Basin_Locations))
        Basins = append(Basins, Basin_Locations)
      }
    }
  }

  var Highest_Basins []int
  for len(Highest_Basins) != 3 {
    var Highest_Basin int
    var Highest_Basin_Id int
    for Basin_Id, Basin := range Basins {
      if len(Basin) > Highest_Basin {
        Highest_Basin_Id = Basin_Id
        Highest_Basin = len(Basin)
      }
    }
    Basins = append(Basins[:Highest_Basin_Id], Basins[Highest_Basin_Id + 1:]...)
    Highest_Basins = append(Highest_Basins, Highest_Basin)
  }

  Total_Size := Highest_Basins[0] * Highest_Basins[1] * Highest_Basins[2]

  log.Printf("Basins: %v", Highest_Basins)
  log.Printf("Total of Highest Basins: %v", Total_Size)
}

func Find_Basin(Height_Map *[][]int, Location []int) [][]int {
  Row_Line := Location[0]
  Column := Location[1]

  Basin_Locations := [][]int{}
  Basin_Locations = append(Basin_Locations, []int{Row_Line, Column})

  var Top_Temp_Locations [][]int
  var Bottom_Temp_Locations [][]int
  var Left_Temp_Locations [][]int
  var Right_Temp_Locations [][]int
  New_Location := true
  for New_Location == true {

    Top_Temp_Locations = [][]int{}
    Bottom_Temp_Locations = [][]int{}
    Left_Temp_Locations = [][]int{}
    Right_Temp_Locations = [][]int{}
    for _, Location := range Basin_Locations {
      Row_Line := Location[0]
      Column := Location[1]
//      log.Printf("Looking for more locations %v:%v", Row_Line, Column)
/*
      if Row_Line != 0 { Top_Temp_Locations = Run_Check(Height_Map, Location, 0) }
      if Row_Line != 99 { Bottom_Temp_Locations = Run_Check(Height_Map, Location, 1) }
      if Column != 0 { Left_Temp_Locations = Run_Check(Height_Map, Location, 2) }
      if Column != 99 { Right_Temp_Locations = Run_Check(Height_Map, Location, 3) }
*/
      if Row_Line != -1 { Top_Temp_Locations = Run_Check(Height_Map, Location, 0) }
      if Row_Line != 100 { Bottom_Temp_Locations = Run_Check(Height_Map, Location, 1) }
      if Column != -1 { Left_Temp_Locations = Run_Check(Height_Map, Location, 2) }
      if Column != 100 { Right_Temp_Locations = Run_Check(Height_Map, Location, 3) }
    }

    if len(Top_Temp_Locations) != 0 {
        log.Printf("Top %v", Top_Temp_Locations)
        Check_Dups(&Top_Temp_Locations, &Basin_Locations)
        if len(Top_Temp_Locations) != 0 {
          log.Printf("Adding top %v", Top_Temp_Locations)
          Basin_Locations = append(Basin_Locations, Top_Temp_Locations...)
        }
    }

    if len(Bottom_Temp_Locations) != 0 {
        log.Printf("Bottom %v", Bottom_Temp_Locations)
        Check_Dups(&Bottom_Temp_Locations, &Basin_Locations)
        if len(Bottom_Temp_Locations) != 0 {
          log.Printf("Adding bottom %v", Bottom_Temp_Locations)
          Basin_Locations = append(Basin_Locations, Bottom_Temp_Locations...)
        }
    }

    if len(Left_Temp_Locations) != 0 {
        log.Printf("Left %v", Left_Temp_Locations)
        Check_Dups(&Left_Temp_Locations, &Basin_Locations)
        if len(Left_Temp_Locations) != 0 {
          log.Printf("Adding left %v", Left_Temp_Locations)
          Basin_Locations = append(Basin_Locations, Left_Temp_Locations...)
        }
    }

    if len(Right_Temp_Locations) != 0 {
        log.Printf("Right %v", Right_Temp_Locations)
        Check_Dups(&Right_Temp_Locations, &Basin_Locations)
        if len(Right_Temp_Locations) != 0 {
          log.Printf("Adding right %v", Right_Temp_Locations)
          Basin_Locations = append(Basin_Locations, Right_Temp_Locations...)
        }
    }

    log.Printf("Final Locations: %v", Basin_Locations)

    if len(Top_Temp_Locations) == 0 &&
      len(Bottom_Temp_Locations) == 0 &&
      len(Left_Temp_Locations) == 0 &&
      len(Right_Temp_Locations) == 0 {
        New_Location = false
    }
  }

  return Basin_Locations
}

func Check_Dups(Temp_Locations *[][]int, Basin_Locations *[][]int) {
  var New_Locations_List [][]int
  for _, Locations := range *Temp_Locations {
    Row_Line := Locations[0]
    Column := Locations[1]
    New_Location := true

    for _, Basin_Location := range *Basin_Locations {
      Basin_Row_Line := Basin_Location[0]
      Basin_Column := Basin_Location[1]

      if Basin_Row_Line == Row_Line && Basin_Column == Column {
        New_Location = false
      }
    }

    if New_Location == true {
      New_Locations_List = append(New_Locations_List, []int{ Row_Line, Column })
    }
  }

  if len(New_Locations_List) != 0 { *Temp_Locations = New_Locations_List
  } else { *Temp_Locations = nil }
}


func Run_Check(Height_Map *[][]int, Location []int, Mode int) [][]int {
  Low_Point := (*Height_Map)[Location[0]][Location[1]]
  Row_Line := Location[0]
  Column := Location[1]

  Move := 1
  Value := Low_Point
  var Locations [][]int
  switch Mode {
    case 0:
      for Value != 9 {
        if Row_Line - Move != -1 {
          Value = (*Height_Map)[ Row_Line - Move ][ Column ]
          if Value != 9 {
            Locations = append(Locations, []int{ Row_Line - Move, Column })
          }
          Move++
        } else { break }
      }

    case 1:
      for Value != 9 {
        if Row_Line + Move != 100 {
          Value = (*Height_Map)[ Row_Line + Move ][ Column ]
          if Value != 9 {
            Locations = append(Locations, []int{ Row_Line + Move, Column })
          }
          Move++
        } else { break }
      }

    case 2:
      for Value != 9 {
        if Column - Move != -1 {
          Value = (*Height_Map)[ Row_Line ][ Column - Move ]
          if Value != 9 {
            Locations = append(Locations, []int{ Row_Line, Column - Move })
          }
          Move++
        } else { break }
      }

    case 3:
      for Value != 9 {
        if Column + Move != 100 {
          Value = (*Height_Map)[ Row_Line ][ Column + Move ]
          if Value != 9 {
            Locations = append(Locations, []int{ Row_Line, Column + Move })
          }
          Move++
        } else { break }
      }
  }

  return Locations
}
