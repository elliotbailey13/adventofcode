use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
  let mut mem_lines = Vec::new();
  if let Ok(file_lines) = read_lines("./input.txt") {
    for line in file_lines.flatten() { mem_lines.push(line); }
  }

  for line in mem_lines {
    println!("{}", line);
  }
}
