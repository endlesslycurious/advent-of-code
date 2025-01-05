"""AoC 2024 - Day 2 tests"""

import unittest
from unittest.mock import mock_open, patch

from day02 import analyse_safety_infractions, part_one, read_inputs


class DayTwoTests(unittest.TestCase):
    """Day two tests"""

    def test_analyse_safety_infractions(self) -> None:
        """Verify the analyse_safety method"""
        inputs: list[tuple[list[int], tuple[bool, int]]] = [
            ([7, 6, 4, 2, 1], (True, -1)),
            ([1, 2, 7, 8, 9], (False, 2)),
            ([9, 7, 6, 2, 1], (False, 3)),
            ([1, 3, 2, 4, 5], (False, 2)),
            ([8, 6, 4, 4, 1], (False, 3)),
            ([1, 3, 6, 7, 9], (True, -1)),
        ]
        for report, expected in inputs:
            res: tuple[bool, int] = analyse_safety_infractions(report)
            self.assertEqual(res, expected, msg=f"{report} -> {res}, expected {expected}!")

    def test_read_inputs(self) -> None:
        """Verify read input method"""
        data: str = "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"
        expected: list[list[int]] = [
            [7, 6, 4, 2, 1],
            [1, 2, 7, 8, 9],
            [9, 7, 6, 2, 1],
            [1, 3, 2, 4, 5],
            [8, 6, 4, 4, 1],
            [1, 3, 6, 7, 9],
        ]

        with patch("builtins.open", mock_open(read_data=data)):
            reports: list[list[int]] = read_inputs("foo")
            for report, answer in zip(reports, expected):
                self.assertEqual(report, answer)

    def test_part_one(self) -> None:
        """Verify part one soltuion"""
        data: str = "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"

        with patch("builtins.open", mock_open(read_data=data)):
            res: int = part_one("bar")
            self.assertEqual(res, 2)


if __name__ == "__main__":
    unittest.main()
