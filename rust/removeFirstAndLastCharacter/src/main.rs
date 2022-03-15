fn main() {
    println!("{}", remove_char("eloquent")); //loquen
    println!("{}", remove_char("country")); //ountr
    println!("{}", remove_char("person")); //erso
    println!("{}", remove_char("place")); //lac
    println!("{}", remove_char("ok")); //""
    println!("{}", remove_char("ooopsss")); //oopss
}


pub fn remove_char(s: &str) -> String {
    let first_last_off: &str = &s[1..s.len() - 1];
    return first_last_off.to_string();
}
