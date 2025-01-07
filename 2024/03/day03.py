"""Advent of Code 2024 - Day 3 solution"""


def read_inputs(filename: str) -> list[str]:
    """Read inputs from file"""
    data: list[str] = []

    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            data.append(line)

    return data


def part_one(filename: str) -> int:
    """Solve part one"""
    data: list[str] = read_inputs(filename)

    return 0


def main() -> None:
    """Main function"""
    filename: str = "input_day03.txt"
    res: int = part_one(filename)
    print(f"Part One - Total {res}")


if __name__ == "__main__":
    main()
