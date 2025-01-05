"""AoC 2024 - Day 2 tests"""

import unittest

from day02 import analyse_safety


class DayTwoTests(unittest.TestCase):
    """Day two tests"""

    def test_analyse_safety(self) -> None:
        """Verify the analyse_safety method"""
        inputs: list[tuple[list[int], bool]] = [
            ([7, 6, 4, 2, 1], True),
            ([1, 2, 7, 8, 9], False),
            ([9, 7, 6, 2, 1], False),
            ([1, 3, 2, 4, 5], False),
            ([8, 6, 4, 4, 1], False),
            ([1, 3, 6, 7, 9], True),
        ]
        for report, expected in inputs:
            res: bool = analyse_safety(report)
            self.assertEqual(res, expected, msg=f"{report} -> {res}, expected {expected}!")


if __name__ == "__main__":
    unittest.main()
