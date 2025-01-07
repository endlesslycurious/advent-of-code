"""AoC 2024 - Day 3 tests"""

import unittest
from unittest.mock import mock_open, patch

from day03 import part_one


class Day03Tests(unittest.TestCase):
    """Day 3 tests"""

    def test_part_one(self) -> None:
        """Verify part one solution"""
        data: str = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
        expected: int = 161

        with patch("builtins.open", mock_open(read_data=data)):
            res: int = part_one("foo")
            self.assertEqual(res, expected)


if __name__ == "__main__":
    unittest.main()
