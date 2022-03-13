fn main() {
    println!("{}", grow(vec![1, 2, 3]));
    println!("{}", grow(vec![4, 1, 1, 1, 4]));
    println!("{}", grow(vec![2, 2, 2, 2, 2, 2]));
}

fn grow(nums: Vec<i32>) -> i32 {
    let mut total = 1;
    for num in nums {
        total = total * num;
    }
    return total;
}

// a better solution
// fn grow(array: Vec<i32>) -> i32 {
//     array.iter().product()
// }