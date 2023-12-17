
fn get_input() -> String {
    let file = std::fs::read_to_string("input.txt").unwrap();
    return file;
}

#[derive(Debug)]
struct FoundSymbol {
    ln: usize,
    rn: usize,
}

#[derive(Debug)]
struct Adjacent {
    tl: Option<((usize, usize), char)>,
    tm: Option<((usize, usize), char)>,
    tr: Option<((usize, usize), char)>,
    
    ml: Option<((usize, usize), char)>,
    mm: Option<((usize, usize), char)>, 
    mr: Option<((usize, usize), char)>,

    bl: Option<((usize, usize), char)>,
    bm: Option<((usize, usize), char)>,
    br: Option<((usize, usize), char)>,
}

impl Adjacent {
    fn get_numbers(&self) -> Vec<((usize, usize), char)> {
        let mut adj_numbers: Vec<((usize, usize), char)> = Vec::new();
        
        match self.tl {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        match self.tm {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        match self.tr {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        match self.ml {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        match self.mm {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        match self.ml {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        match self.br {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        match self.bm {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        match self.bl {
            Some((p, c)) => { 
                if c.is_numeric() {
                    adj_numbers.push((p, c))
                }
            },
            None => ()
        }
        return adj_numbers
    }
}


impl FoundSymbol {
    fn get_around_target(&self, context: &Vec<Vec<char>>) -> Adjacent {
        let max_ln = context.len();
        let max_rn = context[0].len();
        let mut adj = Adjacent {
            tl: None,
            tm: None,
            tr: None,
            ml: None,
            mm: None,
            mr: None,
            bl: None,
            bm: None,
            br: None,
        };

        // Top row
        if self.ln > 0 {
          if self.rn > 0 {
            adj.tl = Some(((self.ln-1, self.rn-1), context[self.ln-1][self.rn-1]));
          }
         adj.tm = Some(((self.ln-1,self.rn), context[self.ln-1][self.rn]));
          if self.rn < max_rn - 1 {
            adj.tr = Some(((self.ln-1, self.rn+1), context[self.ln-1][self.rn+1]));
          }
        }

        // Middle row
        if self.rn > 0 {
          adj.ml = Some(((self.ln, self.rn-1), context[self.ln][self.rn-1]));
        }
        adj.mm = Some(((self.ln, self.rn), context[self.ln][self.rn]));
        if self.rn < max_rn - 1 {
          adj.mr = Some(((self.ln, self.rn+1), context[self.ln][self.rn+1]));
        }

        // Bottom row
        if self.ln < max_ln - 1 {
          if self.rn > 0 {
            adj.bl = Some(((self.ln+1,self.rn-1), context[self.ln+1][self.rn-1]));
          }
          adj.bm = Some(((self.ln+1, self.rn), context[self.ln+1][self.rn]));
          if self.rn < max_rn - 1 {
            adj.br = Some(((self.ln+1, self.rn+1), context[self.ln+1][self.rn+1]));
          }
        }
        return adj;
    }
}

fn is_within_one(num1: usize, num2: usize) -> bool {
    if num1 > num2 {
        num1 - num2 <= 1
    } else {
        num2 - num1 <= 1
    }
}


fn main() {
    let mut outer: Vec<Vec<char>> = Vec::new();
    for (_, content) in get_input().lines().enumerate() {
        let mut inner: Vec<char> = Vec::new();
        for (_, c) in content.char_indices() {
            inner.push(c);
        }
        outer.push(inner);
    }
    
    let mut adjs: Vec<Adjacent> = Vec::new();
    for (ln, line) in outer.iter().enumerate() {
        for (rn, c) in line.iter().enumerate() {
          if !c.is_alphanumeric() && *c != '.' {
            let t = FoundSymbol { ln, rn };
            adjs.push(t.get_around_target(&outer));
          }
        }
    }

    let sum: i32 = adjs.iter()
        .map(|a| a.get_numbers())
        .filter(|a| !a.is_empty())
        .flatten()
        .map(|(pos, _)| get_num(pos, &outer))
        .sum();

    print!("Sum is {}\n", sum);

}

fn get_num((ln,rn): (usize, usize), f: &Vec<Vec<char>>) -> i32 {

    let row = &f[ln];
    let mut head: bool = false;
    let mut tail: bool = false;
    
    let mut tidx = rn+1;
    let mut hidx = rn;
    let mut num: String = String::from("");
    
    while !head {
        if hidx >= 0 {
            let v = row[hidx]; 
            if v.is_numeric() {
                num.push(v);
            } else {
                break;
            }
            
        } else {
            head = true;
            break;
        }
        if hidx > 0 {
            hidx-=1;
        } else {
            head = true;
        }
    }

    num = num.chars().rev().collect();

    while !tail {
        if tidx < row.len() {
            let v = row[tidx]; 
            if v.is_numeric() {
                num.push(v);
            } else {
                break; 
            }
        } else {
            tail = true;
            break;
        }
        tidx+=1;
    }

    return num.parse::<i32>().unwrap_or(0);
}




