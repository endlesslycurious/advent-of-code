"""Advent of Code 2024 - Day 2 solution"""


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


def main() -> None:
    """Main function"""
    print("hello")


if __name__ == "__main__":
    main()
