package main

import (
  "os"
  "log"
  "bufio"
  "slices"
  "strconv"
  "strings"
)

var (
  Hands []Hand
  Five_of_a_Kind []Hand
  Four_of_a_Kind []Hand
  Full_House []Hand
  Three_of_a_Kind []Hand
  Two_Pair []Hand
  One_Pair []Hand
  High_Card []Hand
  Sorted_Hands []Hand
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

  for _, Full_Input_Line := range Full_Input {
    err = Load_Data(Full_Input_Line)
    if err != nil {
      log.Printf("Load data error. %v - %v", Full_Input_Line, err)
      os.Exit(1)
    }
  }

  log.Printf("Length of hands: %v", len(Hands))
  err = Sort_Hand_Types()
  if err != nil {
    log.Printf("Sorting hand types error. %v", err)
    os.Exit(1)
  }

  log.Printf("Number of five of a kind: %v", len(Five_of_a_Kind))
  log.Printf("Number of four of a kind: %v", len(Four_of_a_Kind))
  log.Printf("Number of full house: %v", len(Full_House))
  log.Printf("Number of three of a kind: %v", len(Three_of_a_Kind))
  log.Printf("Number of two pair: %v", len(Two_Pair))
  log.Printf("Number of one pair: %v", len(One_Pair))
  log.Printf("Number of high card: %v", len(High_Card))

  Hands = []Hand{}

  Five_of_a_Kind, err = Order_Hands(Five_of_a_Kind)
  if err != nil {
    log.Printf("Ordering of five of a kind hands error. %v", err)
    os.Exit(1)
  }
  log.Print("Five of a Kind sorted.")

  Four_of_a_Kind, err = Order_Hands(Four_of_a_Kind)
  if err != nil {
    log.Printf("Ordering of four of a kind hands error. %v", err)
    os.Exit(1)
  }
  log.Print("Four of a Kind sorted.")

  Full_House, err = Order_Hands(Full_House)
  if err != nil {
    log.Printf("Ordering of full house hands error. %v", err)
    os.Exit(1)
  }
  log.Print("Full House sorted.")

  Three_of_a_Kind, err = Order_Hands(Three_of_a_Kind)
  if err != nil {
    log.Printf("Ordering of three of a kind hands error. %v", err)
    os.Exit(1)
  }
  log.Print("Three of a Kind sorted.")

  Two_Pair, err = Order_Hands(Two_Pair)
  if err != nil {
    log.Printf("Ordering of two pair hands error. %v", err)
    os.Exit(1)
  }
  log.Print("Two Pair sorted.")

  One_Pair, err = Order_Hands(One_Pair)
  if err != nil {
    log.Printf("Ordering of one pair hands error. %v", err)
    os.Exit(1)
  }
  log.Print("One Pair sorted.")

  High_Card, err = Order_Hands(High_Card)
  if err != nil {
    log.Printf("Ordering of high card hands error. %v", err)
    os.Exit(1)
  }
  log.Print("High Card sorted.")

  Hands = append(Hands, Five_of_a_Kind...)
  Hands = append(Hands, Four_of_a_Kind...)
  Hands = append(Hands, Full_House...)
  Hands = append(Hands, Three_of_a_Kind...)
  Hands = append(Hands, Two_Pair...)
  Hands = append(Hands, One_Pair...)
  Hands = append(Hands, High_Card...)

  log.Printf("Length of hands: %v", len(Hands))
  slices.Reverse(Hands)

  var Total_Winnings int
  for Index, Hand := range Hands {
    Winnings := ( Index + 1 ) * Hand.Bid
    Total_Winnings = Total_Winnings + Winnings
  }

  log.Printf("Total Winnings: %v", Total_Winnings)
}

func Load_Data(Line string) error {
  Line_Parts := strings.Split(Line, " ")
  Hand_Bid, err := strconv.Atoi(Line_Parts[1])
  if err != nil { return err }
  
  Wilds := strings.Count(Line_Parts[0], "J")

  Hands = append(Hands, Hand{ Raw_Hand: Line_Parts[0], Wilds: Wilds, Bid: Hand_Bid })

  return nil
}

func Sort_Hand_Types() error {
  var Cards []string
  var Raw_Cards []string
  var Compact_Cards []string
  for _, Current_Hand := range Hands {
    Cards = strings.Split(Current_Hand.Raw_Hand, "")
    Raw_Cards = strings.Split(Current_Hand.Raw_Hand, "")

    slices.Sort(Cards)
    Compact_Cards = slices.Compact(Cards)

    switch len(Compact_Cards) {
      case 5:
        if Current_Hand.Wilds == 1 {
          One_Pair = append(One_Pair, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
        } else {
          High_Card = append(High_Card, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
        }

      case 4:
        if Current_Hand.Wilds == 1 {
          Three_of_a_Kind = append(Three_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
	} else if Current_Hand.Wilds == 2 {
          Three_of_a_Kind = append(Three_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
        } else {
          One_Pair = append(One_Pair, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
        }

      case 3:
        var Matches int
        var Highest_Matches int

        for i := 0; i != len(Raw_Cards); i++ {
          Card_Compare := Raw_Cards[i]
          for _, Card := range Raw_Cards {
            if Card_Compare == Card { Matches++ }
          }

          if Matches > Highest_Matches { Highest_Matches = Matches }
          Matches = 0
        }

        if Highest_Matches == 3 {
          if Current_Hand.Wilds == 3 {
            Four_of_a_Kind = append(Four_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          } else if Current_Hand.Wilds == 1 {
            Four_of_a_Kind = append(Four_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          } else {
            Three_of_a_Kind = append(Three_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          }
        } else {
          if Current_Hand.Wilds == 2 {
            Four_of_a_Kind = append(Four_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          } else if Current_Hand.Wilds == 1 {
            Full_House = append(Full_House, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          } else {
            Two_Pair = append(Two_Pair, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          }
        }

      case 2:
        var Matches int
        var Highest_Matches int

	for i := 0; i < len(Raw_Cards); i++ {
          Card_Compare := Raw_Cards[i]
          for _, Card := range Raw_Cards {
            if Card_Compare == Card { Matches++ }
          }

          if Matches > Highest_Matches { Highest_Matches = Matches }
          Matches = 0
        }

        if Highest_Matches == 4 {
          if Current_Hand.Wilds == 4 {
            Five_of_a_Kind = append(Five_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          } else if Current_Hand.Wilds == 1 {
            Five_of_a_Kind = append(Five_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          } else {
            Four_of_a_Kind = append(Four_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          }
        } else {
          if Current_Hand.Wilds == 3 {
            Five_of_a_Kind = append(Five_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          } else if Current_Hand.Wilds == 2 {
            Five_of_a_Kind = append(Five_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          } else {
            Full_House = append(Full_House, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })
          }
        }

      case 1:
        Five_of_a_Kind = append(Five_of_a_Kind, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })

      default:
        Sorted_Hands = append(Sorted_Hands, Hand{ Raw_Hand: Current_Hand.Raw_Hand, Bid: Current_Hand.Bid })

    }
  }

  return nil
}

func Order_Hands(Current_Hands []Hand) ([]Hand, error) {
  if len(Current_Hands) == 1 { return Current_Hands, nil }

  Sorted := 1
  var Runs int
  var err error
  var Temp_Sorted_Hands []Hand
  for Sorted > 0 {
    Runs++

    var Hand_1 Hand
    Temp_Sorted_Hands = []Hand{}
    for Index, Current_Hand := range Current_Hands {
      if Index == 0 {
        Hand_1 = Current_Hand
        continue
      }

      if Sorted == 2 {
        Temp_Sorted_Hands = append(Temp_Sorted_Hands, Current_Hand)
        continue
      }

      var Hand_1_Cards_Converted []int
      var Hand_2_Cards_Converted []int
      Hand_1_Cards_Converted, err = Convert_Hand_String_Int(Hand_1.Raw_Hand)
      if err != nil { return []Hand{}, err }

      Hand_2_Cards_Converted, err = Convert_Hand_String_Int(Current_Hand.Raw_Hand)
      if err != nil { return []Hand{}, err }

      for i := 0; i < len(Hand_1_Cards_Converted); i++ {
        if Hand_1_Cards_Converted[i] > Hand_2_Cards_Converted[i] {
          Temp_Sorted_Hands = append(Temp_Sorted_Hands, Hand_1)
          Hand_1 = Current_Hand
          break
        }

        if Hand_1_Cards_Converted[i] < Hand_2_Cards_Converted[i] {
          Temp_Sorted_Hands = append(Temp_Sorted_Hands, Current_Hand)
          Temp_Sorted_Hands = append(Temp_Sorted_Hands, Hand_1)
          Sorted++
          break
        }
      }

      if Sorted == 1 && (Index + 1) == len(Current_Hands) {
        Temp_Sorted_Hands = append(Temp_Sorted_Hands, Current_Hand)
      }
    }

    Current_Hands = Temp_Sorted_Hands

    if Sorted > 1 { Sorted = 1
    } else { Sorted = 0 }

    if Runs > 1000000000 {
      log.Print("Ordering of hands took too long")
      os.Exit(1)
    }
  }

  return Temp_Sorted_Hands, nil
}

func Convert_Hand_String_Int(Hand_Cards string) ([]int, error) {
  var Temp_Hand_Cards_Converted []int

  Hand_Cards_Strings := strings.Split(Hand_Cards, "")

  for _, Hand_Card := range Hand_Cards_Strings {
    switch Hand_Card {
      case "A":
        Temp_Hand_Cards_Converted = append(Temp_Hand_Cards_Converted, 14)

      case "K":
        Temp_Hand_Cards_Converted = append(Temp_Hand_Cards_Converted, 13)

      case "Q":
        Temp_Hand_Cards_Converted = append(Temp_Hand_Cards_Converted, 12)

      case "J": 
        Temp_Hand_Cards_Converted = append(Temp_Hand_Cards_Converted, 1)

      case "T":
        Temp_Hand_Cards_Converted = append(Temp_Hand_Cards_Converted, 10)

      default:
        Temp_Int, err := strconv.Atoi(Hand_Card)
        if err != nil { return []int{}, err }

        Temp_Hand_Cards_Converted = append(Temp_Hand_Cards_Converted, Temp_Int)
    }
  }

  return Temp_Hand_Cards_Converted, nil
}

type Hand struct {
  Raw_Hand string
  Wilds int
  Bid int
}
