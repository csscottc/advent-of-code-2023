use std::collections::HashMap;

fn main() {
    println!("Hello, world!");
    let h = Hand {
        cards: vec!['A', 'B', 'B', 'B', 'B'],
    };
    println!("{:?}", h);
    println!("{:?}", h.get_type())
}

#[derive(Debug)]
struct Hand {
    cards: Vec<char>,
}

impl Hand {
    fn get_type(&self) -> HandType {
        let mut match_map = HashMap::new();
        for (i, &c1) in self.cards.iter().enumerate() {
            match_map.insert(c1, 0);
            for (i, &c2) in self.cards.iter().enumerate() {
                println!("c1 {} c2 {}", c1, c2);
                if c1 == c2 {
                    if let Some(value) = match_map.get_mut(&c1) {
                        *value += 1
                    }
                }
            }
            println!("{} matches {:?}", c1, match_map);
        }

        let mut pairs_of: Vec<char> = vec![];
        let mut trips_of: Vec<char> = vec![];
        let mut singles_of: Vec<char> = vec![];
        for (key, value) in match_map {
            match value {
                4 => return HandType::FourOfAKind,
                5 => return HandType::FiveOfAKind,
                2 => {
                    if !pairs_of.contains(&key) {
                        pairs_of.push(key)
                    }
                }
                3 => {
                    if !trips_of.contains(&key) {
                        trips_of.push(key)
                    }
                }
                1 => {
                    if !singles_of.contains(&key) {
                        singles_of.push(key)
                    }
                }
                _ => {}
            }
        }
        
        if trips_of.len() == 1 && pairs_of.len() == 1 {
            return HandType::FullHouse;
        }

        if trips_of.len() == 1 && pairs_of.len() == 0 {
            return HandType::ThreeOfAKind;
        }

        if pairs_of.len() > 1 {
            return HandType::TwoPair; 
        }

        if pairs_of.len() == 1 {
            return HandType::OnePair;
        }

        return HandType::HighCard;
    }
}

#[derive(Debug)]
#[derive(PartialEq)]
enum HandType {
    FiveOfAKind,
    FourOfAKind,
    FullHouse,
    ThreeOfAKind,
    TwoPair,
    OnePair,
    HighCard,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_pair() {
        let x = Hand { cards: vec!['A','A','J','J','Q'], };
        assert_eq!(x.get_type(), HandType::TwoPair);
    }

    #[test]
    fn test_one_pair() {
        let x = Hand { cards: vec!['A','A','K','J','Q'], };
        assert_eq!(x.get_type(), HandType::OnePair);
    }

    #[test]
    fn test_five_of_a_kind() {
        let x = Hand { cards: vec!['K','K','K','K','K'], };
        assert_eq!(x.get_type(), HandType::FiveOfAKind);
    }

    #[test]
    fn test_four_of_a_kind() {
        let x = Hand { cards: vec!['K','K','Q','K','K'], };
        assert_eq!(x.get_type(), HandType::FourOfAKind);
    }

    #[test]
    fn test_three_of_a_kind() {
        let x = Hand { cards: vec!['A','K','Q','K','K'], };
        assert_eq!(x.get_type(), HandType::ThreeOfAKind);
    }

    #[test]
    fn test_high_card() {
        let x = Hand { cards: vec!['A','K','Q','1','9'], };
        assert_eq!(x.get_type(), HandType::HighCard);
    
    }
    #[test]
    fn test_full_house() {
        let x = Hand { cards: vec!['A','K','K','A','A'], };
        assert_eq!(x.get_type(), HandType::FullHouse);
    }
}
