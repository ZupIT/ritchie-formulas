use colored::*;

pub fn run(sample_text: String, sample_list: String, sample_bool: String) {
    println!("Hello World!");
    println!("{}", format!("You receive {} in text.", sample_text).blue());
    println!("{}", format!("You receive {} in list.", sample_list).green());
    println!("{}", format!("You receive {} in boolean.", sample_bool).red());
}
