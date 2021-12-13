package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
)

func main() {
  Sonar_Readings, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }
  defer Sonar_Readings.Close()

  Previous_Reading := 0
  Increased_Counter := 0
  scanner := bufio.NewScanner(Sonar_Readings)

  for scanner.Scan() {
    Sonar_Reading, _ := strconv.Atoi(scanner.Text())
    if Previous_Reading != 0 {
      if Previous_Reading < Sonar_Reading {
        Increased_Counter++
        log.Printf("Reading: %v (increased)", Sonar_Reading)
      } else {
        log.Printf("Reading: %v (decreased)", Sonar_Reading)
      }
    }

    Previous_Reading = Sonar_Reading
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  log.Printf("Total Number of Increased readings: %v", Increased_Counter)
}
