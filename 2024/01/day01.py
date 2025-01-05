""" Advent of code 2024 - Day 1 solution"""

from collections import defaultdict


def calcdistance(left: int, right: int) -> int:
    """Calculate distance"""
    return abs(left - right)


def partone(filename: str) -> int:
    """Reads file and calculates total distances"""
    left: list[int] = []
    right: list[int] = []

    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            leftnum, rightnum = map(int, line.split("   "))
            left.append(leftnum)
            right.append(rightnum)

    left.sort()
    right.sort()

    total: int = 0
    for numleft, numright in zip(left, right):
        total += calcdistance(numleft, numright)

    return total


def parttwo(filename: str) -> int:
    """Reads file and calcuates the similarirty score"""
    score: int = 0
    left: list[int] = []
    right: dict[int, int] = defaultdict(int)

    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            leftnum, rightnum = map(int, line.split("   "))
            left.append(leftnum)
            right[rightnum] += 1

    for num in left:
        if num in right:
            score += num * right[num]

    return score


def main() -> None:
    """Main function"""
    filename: str = "input_day01.txt"

    total: int = partone(filename)
    print(f"Part One - Total is {total}!")

    score: int = parttwo(filename)
    print(f"Part Two - Score is {score}!")


if __name__ == "__main__":
    main()
