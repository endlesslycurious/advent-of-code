"""Advent of Code 2024 - Day 2 solution"""

from typing import Iterator


def analyse_safety_infractions(report: list[int]) -> tuple[bool, int]:
    """Analyse the safety of a report, true for safe with -1 or false with index of failing reading"""
    increasing: str = " "

    for index, reading in enumerate(report):
        if index == 0:
            continue

        previous: int = report[index - 1]

        # Constant readings are unsafe
        if reading == previous:
            return (False, index)

        # Large swings in readings are unsafe
        if abs(reading - previous) > 3:
            return (False, index)

        if reading > previous:
            increasing = "+" if increasing != "-" else "bad"
        else:
            increasing = "-" if increasing != "+" else "bad"

        if increasing == "bad":
            return (False, index)

    return (True, -1)


def analyse_safety(report: list[int]) -> bool:
    """Verify an report is safe, returns true if safe"""
    safe, _ = analyse_safety_infractions(report)
    return safe is True


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

    # second pass remove first bad reading from unsafe report then retest
    for report in filter(lambda r: not analyse_safety(r), reports):
        _, index = analyse_safety_infractions(report)

        del report[index]

        if analyse_safety(report):
            safe += 1

    return safe


def main() -> None:
    """Main function"""
    filename: str = "input_day02.txt"
    safe: int = part_one(filename)
    print(f"Part One - {safe} safe repors!")

    safe = part_two(filename)
    print(f"Part Two - {safe} safe repors!")


if __name__ == "__main__":
    main()
