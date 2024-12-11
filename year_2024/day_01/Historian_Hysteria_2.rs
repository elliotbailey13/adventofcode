use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
  let mut left_numbers = Vec::new();
  let mut right_numbers = Vec::new();

  if let Ok(lines) = read_lines("./input.txt") {
    for line in lines.flatten() {
      let line_parts: Vec<&str> = line.split("   ").collect();

      left_numbers.push(line_parts[0].parse::<i32>().unwrap());
      right_numbers.push(line_parts[1].parse::<i32>().unwrap());
    }
  }

  let mut count: i32;
  let mut total = 0;
  for left_number in left_numbers {
    count = 0;

    for right_number in &right_numbers {
      if left_number == *right_number {
        count = count + 1;
      }
    }

    total = total + ( left_number * count )
  }

  println!("Total: {}", total)
}
