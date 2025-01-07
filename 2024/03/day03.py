"""Advent of Code 2024 - Day 3 solution"""

from re import findall


def read_inputs(filename: str) -> list[str]:
    """Read inputs from file"""
    data: list[str] = []

    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            data.append(line)

    return data


def process_mul_instructions(line: str) -> int:
    """Parse line for instructions, returns total of multiplied values"""
    total: int = 0

    # https://www.regextester.com for the win!
    for a, b in findall(r"mul\(([0-9]+),([0-9]+)\)", line):
        total += int(a) * int(b)

    return total


def part_one(filename: str) -> int:
    """Solve part one"""
    data: list[str] = read_inputs(filename)
    total: int = 0

    for line in data:
        total += process_mul_instructions(line)

    return total


def main() -> None:
    """Main function"""
    filename: str = "input_day03.txt"
    res: int = part_one(filename)
    print(f"Part One - Total {res}")


if __name__ == "__main__":
    main()
