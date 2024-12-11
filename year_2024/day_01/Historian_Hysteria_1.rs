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

  left_numbers.sort();
  right_numbers.sort();

  let mut total = 0;
  for i in 0..left_numbers.len() {
    let temp = left_numbers[i] - right_numbers[i];
    total = total + temp.abs();
  }

  println!("Total: {}", total)
}
