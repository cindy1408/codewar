fn main() {
    println!("{}", update_light("green")); //yellow
    println!("{}", update_light("yellow")); //red
    println!("{}", update_light("red")); //green
}

fn update_light(current: &str) -> String {
    match current {
        "red" => return "green".to_string(),
        "yellow" => return "red".to_string(),
        "green" => return "yellow".to_string(),
        x => panic!("Unexpected invalid current {:?}", x),
    }
}

// better 

// fn update_light(current: &str) -> String {
//     match current {
//         "green" => "yellow",
//         "yellow" => "red",
//         "red" => "green",
//         _ => panic!()
//     }.into()
// }