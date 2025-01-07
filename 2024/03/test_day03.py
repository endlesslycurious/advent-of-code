"""AoC 2024 - Day 3 tests"""

import unittest
from unittest.mock import mock_open, patch

from day03 import part_one, process_mul_instructions, read_inputs


class Day03Tests(unittest.TestCase):
    """Day 3 tests"""

    def test_read_inputs(self) -> None:
        """Verify read_inputs works"""
        data: str = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

        with patch("builtins.open", mock_open(read_data=data)):
            res: list[str] = read_inputs("test")
            self.assertEqual(res, [data])

    def test_process_mul_instructions(self) -> None:
        """Verify process_mul_instructions"""
        data: str = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
        expected: int = 161
        res: int = process_mul_instructions(data)
        self.assertEqual(res, expected)

    def test_part_one(self) -> None:
        """Verify part one solution"""
        data: str = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
        expected: int = 161

        with patch("builtins.open", mock_open(read_data=data)):
            res: int = part_one("foo")
            self.assertEqual(res, expected)


if __name__ == "__main__":
    unittest.main()
