"""AoC 2024 - Day 1 tests"""

import unittest
from unittest.mock import mock_open, patch

from day01 import calcdistance, partone, parttwo


class Day1Tests(unittest.TestCase):
    """Day 1 tests"""

    def test_calcdistance(self) -> None:
        """Verify calculate distance function"""
        self.assertEqual(calcdistance(1, 3), 2)
        self.assertEqual(calcdistance(2, 3), 1)
        self.assertEqual(calcdistance(3, 3), 0)
        self.assertEqual(calcdistance(3, 4), 1)
        self.assertEqual(calcdistance(3, 5), 2)
        self.assertEqual(calcdistance(4, 9), 5)

    def test_partone(self) -> None:
        """Verify part one solution"""
        data: str = "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
        with patch("builtins.open", mock_open(read_data=data)):
            res: int = partone("foo")
            self.assertEqual(res, 11)

    def test_parttwo(self) -> None:
        """Verify part two solution"""
        data: str = "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
        with patch("builtins.open", mock_open(read_data=data)):
            res: int = parttwo("bar")
            self.assertEqual(res, 31)


if __name__ == "__main__":
    unittest.main()
