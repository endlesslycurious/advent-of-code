"""Advent of Code 2024 - Day 2 solution"""

from typing import Iterator


def analyse_safety(report: list[int]) -> bool:
    """Analyse the safety of a report, true for safe"""
    increase: int = 0
    decrease: int = 0

    for index, reading in enumerate(report):
        if index == 0:
            continue

        previous: int = report[index - 1]

        # Constant readings are unsafe
        if reading == previous:
            return False

        # Large swings in readings are unsafe
        if abs(reading - previous) > 3:
            return False

        if reading > previous:
            increase += 1
        else:
            decrease += 1

    # if not all increasing or decreasing then unsafe!
    intervals: int = len(report) - 1
    if (increase != intervals) & (decrease != intervals):
        return False

    return True


def read_inputs(filename: str) -> list[list[int]]:
    """Read reports from file"""
    reports: list[list[int]] = []

    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            report: list[int] = [int(x) for x in line.split()]
            reports.append(report)

    return reports


def part_one(filename: str) -> int:
    """Reads reports from files and returns number of safe reports"""
    reports: list[list[int]] = read_inputs(filename)
    safe: int = 0

    for _ in filter(analyse_safety, reports):
        safe += 1

    return safe


def main() -> None:
    """Main function"""
    inputs: str = "input_day02.txt"
    safe: int = part_one(inputs)
    print(f"Part One - {safe} safe repors!")


if __name__ == "__main__":
    main()
