#[derive(Debug)]
struct Card {
    id: i16,
    winning_nums: Vec<i32>,
    have_nums: Vec<i32>,
}

fn get_input() -> &'static str {
    "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36   
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
}

fn main() {
    let _ = get_input()
        .lines()
        .map(|s| s.trim())
        .map(|l| {
            let parts = l.split_once(":").map(|(card_name, rest)| {
              let (_, id) = card_name.split_once(" ").unwrap();
              return (id, rest.trim());
            }).unwrap();

            let (id, raw_nums) = parts;


            let nums = raw_nums.split_once("|").map(|(winning, have)| {
                let winning_nums = winning.split_whitespace().filter_map(|num| num.parse().ok()).collect();
                let have_nums = have.split_whitespace().filter_map(|num| num.parse().ok()).collect();
                return (winning_nums, have_nums);
            });


            return nums.map(|(winning_nums, have_nums)| {
                let c = Card {
                    id: id.parse().expect("Couldn't parse id for card"),
                    winning_nums,
                    have_nums,
                };
                return c;
            });
        }).for_each(|x| println!("{:?}", x));

}
