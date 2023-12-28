#[derive(Debug)]
struct Card {
    id: i16,
    winning_nums: Vec<i32>,
    have_nums: Vec<i32>,
}


impl Card {
    fn get_points(&self) -> i32 {
        let mut points: i32 = 0;
        let mut multi: i32 = 0;
        self.winning_nums.iter().for_each(|wn| {
            if self.have_nums.contains(wn) {
                if points > 0 {
                    multi += 1;
                }
                points = 1;
            }
        });

        for _i in 0..multi {
            points = points * 2
        }

        return points
    }
}

fn get_input() -> String {
    
    let file = std::fs::read_to_string("src/input.txt").unwrap();
    return file;
    /*
    "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36   
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"*/
}




fn main() {
    let s: i32 = get_input()
        .lines()
        .map(|s| s.trim())
        .map(|l| {
            let parts = l.split_once(":").map(|(card_name, rest)| {
              let (_, id) = card_name.split_once(" ").unwrap();
              return (id.trim(), rest.trim());
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
        }).filter_map(|x| {
            let v = x?;
            return Some(v.get_points());
        }).sum();
    println!("The sum should be {}", s);
//    .for_each(|p| println!("{}",p));

}
