"""Advent of Code 2024 - Day 2 solution"""

from typing import Iterator


def analyse_safety(report: list[int]) -> bool:
    """Verify an report is safe, returns true if safe"""
    increasing: bool = False

    for index, reading in enumerate(report):
        if index == 0:
            continue

        previous: int = report[index - 1]

        # direction is set by first delta
        if index == 1:
            increasing = reading > previous

        # Constant readings are unsafe
        if reading == previous:
            return False

        # Large swings in readings are unsafe
        if abs(reading - previous) > 3:
            return False

        if reading < previous and increasing:
            return False

        if reading > previous and not increasing:
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


def part_two(filename: str) -> int:
    """Reads reports from files, attempts to correct first bad reading and returns number of safe reports"""
    reports: list[list[int]] = read_inputs(filename)
    safe: int = 0

    # first pass get the safe reports requiring no modification
    for _ in filter(analyse_safety, reports):
        safe += 1

    # second pass: brute force, in sequence remove single element from an unsafe report then retest
    for report in filter(lambda r: not analyse_safety(r), reports):

        for index in range(len(report)):
            copy: list[int] = report.copy()

            del copy[index]

            if analyse_safety(copy):
                safe += 1
                break

    return safe


def main() -> None:
    """Main function"""
    filename: str = "input_day02.txt"
    safe: int = part_one(filename)
    print(f"Part One - {safe} safe reports!")

    safe = part_two(filename)
    print(f"Part Two - {safe} safe reports!")


if __name__ == "__main__":
    main()
