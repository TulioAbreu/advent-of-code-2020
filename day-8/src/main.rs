use std::fs;
#[derive(Clone, Debug)]
struct Opcode {
    code: String,
    value: i32,
}

fn read_file(filepath: String) -> Vec<Opcode> {
    let file_content = fs::read_to_string(filepath).expect("Unable to read file");
    let file_lines = file_content.lines();
    let mut opcodes = Vec::new();
    for line in file_lines {
        let line_chunks: Vec<&str> = line.split(" ").collect();
        let opcode = Opcode {
            code: line_chunks[0].into(),
            value: line_chunks[1].parse().unwrap(),
        };
        opcodes.push(opcode);
    }
    opcodes
}

fn execute_opcode(opcode: &Opcode, acc: &mut i32)->i32{
    if opcode.code == "acc" {
        *acc += opcode.value;
        return 1;
    }
    else if opcode.code == "jmp" {
        return opcode.value;
    }
    else {
        return 1;
    }
}

fn run_program(opcodes: Vec<Opcode>) -> (i32, bool) {
    let mut acc: i32 = 0;
    let mut current_line: usize = 0;
    let mut count = vec![0; opcodes.len()];
    let mut full_execution = false;
    loop {
        if current_line >= opcodes.len() {
            full_execution = true;
            break;
        }

        if count[current_line] >= 1 {
            full_execution = false;
            break;
        }

        count[current_line] += 1;
        let line_increase: i32 = execute_opcode(&opcodes[current_line], &mut acc);
        current_line = (current_line as i32 + line_increase) as usize;
    }
    (acc, full_execution)
}

fn problem_1() {
    let opcodes: Vec<Opcode> = read_file(String::from("./input.txt"));
    let (acc, _) = run_program(opcodes);
    println!("Problem 1 ] acc = {}", acc);
}

fn problem_2() {
    let opcodes: Vec<Opcode> = read_file(String::from("input.txt"));
    for (index, opcode) in opcodes.iter().enumerate() {
        if opcode.code == "acc" {
            continue;
        }

        let mut programClone: Vec<Opcode> = opcodes.clone();
        if opcode.code == "jmp" {
            programClone[index] = Opcode {
                code: String::from("nop"),
                value: opcode.value,
            };
        } else {
            programClone[index] = Opcode {
                code: String::from("jmp"),
                value: opcode.value,
            };
        }

        let (acc, complete_run) = run_program(programClone);
        if complete_run {
            println!("Problem 2 ] acc = {}", acc);
        }
    }
}

fn main() {
    problem_1();
    problem_2();
}

