use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn is_safe(line_parts: Vec<&str>) -> bool {
  let mut safe = true;
  let mut last_part = 0;
  let mut increasing = false;
  let mut decreasing = false;
  let mut current_part: i32;

  for part in line_parts {
    current_part = part.parse::<i32>().unwrap();

    if last_part == 0 {
      last_part = current_part;
      continue;
    }

    let difference = (last_part - current_part).abs();

    if increasing == false && decreasing == false {
      if last_part > current_part { decreasing = true; }
      if last_part < current_part { increasing = true; }
    } else {
      if last_part > current_part {
        if increasing == true { safe = false; }
      }
      if last_part < current_part {
        if decreasing == true { safe = false; }
      }
    }

    if difference > 3 { safe = false; }
    if difference == 0 { safe = false; }

    last_part = current_part;
  }

  return safe;
}

fn main() {
  let mut mem_lines = Vec::new();

  if let Ok(file_lines) = read_lines("./input.txt") {
    for line in file_lines.flatten() { mem_lines.push(line); }
  }

  let mut safe_lines = 0;

  for line in mem_lines {
    let line_parts: Vec<&str> = line.split(" ").collect();

    let mut safe = is_safe(line_parts.to_vec());

    if safe == false {
      let mut drop_index = 0;

      while drop_index != line_parts.len() {
        let mut new_parts = Vec::<&str>::new();

        for (pos, number) in line_parts.iter().enumerate() {
          if drop_index != pos { new_parts.push(number); }
        }

        safe = is_safe(new_parts);

        if safe == true { break; }

        drop_index += 1;
      }
    }

    if safe == true { safe_lines += 1; }
  }

  println!("Safe line: {}", safe_lines);
}
