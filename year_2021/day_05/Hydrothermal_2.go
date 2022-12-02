package main

import (
  "os"
  "log"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  Vent_Locations, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }
  defer Vent_Locations.Close()

  scanner := bufio.NewScanner(Vent_Locations)

  Vent_Map := [][]int{}
  for i := 0; i != 998; i++ {

    Vent_Map_Row := []int{}
    for i_2 := 0; i_2 != 998; i_2++ {

      Vent_Map_Row = append(Vent_Map_Row, 0)
    }

    Vent_Map = append(Vent_Map, Vent_Map_Row)
  }

//  log.Printf("Vent_Map: %v", Vent_Map)

  for scanner.Scan() {
    Vent_Location_Line := scanner.Text()
    x1y1_x2y2 := strings.Split(Vent_Location_Line, " -> ")
    x1y1 := strings.Split(x1y1_x2y2[0], ",")
    x2y2 := strings.Split(x1y1_x2y2[1], ",")
    x1, _ := strconv.Atoi(x1y1[0])
    y1, _ := strconv.Atoi(x1y1[1])
    x2, _ := strconv.Atoi(x2y2[0])
    y2, _ := strconv.Atoi(x2y2[1])
//    if x1 == x2 || y1 == y2 {
//      log.Printf("%v,%v -> %v,%v", x1, y1, x2, y2)
    Mark_Vent_Map(x1, y1, x2, y2, &Vent_Map)
//    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

//  log.Printf("Vent_Map: %v", Vent_Map)

  Overlaps := Find_Overlaps(&Vent_Map)
  log.Printf("Overlaps: %v", Overlaps)
}

func Mark_Vent_Map(x1, y1, x2, y2 int, Vent_Map *[][]int) {
  (*Vent_Map)[y1][x1]++
  if x1 == x2 {
    for y1 != y2 {
      if y1 > y2 {
        y1--
        (*Vent_Map)[y1][x1]++
      } else {
        y1++
        (*Vent_Map)[y1][x1]++
      }
    }
  } else if y1 == y2 {
    for x1 != x2 {
      if x1 > x2 {
        x1--
        (*Vent_Map)[y1][x1]++
      } else {
        x1++
        (*Vent_Map)[y1][x1]++
      }
    }
  } else {
    for x1 != x2 && y1 != y2 {
      if x1 > x2 { x1--
      } else { x1++ }

      if y1 > y2 { y1--
      } else { y1++ }

      (*Vent_Map)[y1][x1]++
    }
  }
}

func Find_Overlaps(Vent_Map *[][]int) int {
  var Overlaps int
  for _, Row := range *Vent_Map {
    for _, Value := range Row {
      if Value > 1 { Overlaps++ }
    }
  }

  return Overlaps
}

