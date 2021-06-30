import unittest
from unittest import TestCase

from mlpm_client import Client


class TestingPythonClient(TestCase):
    def test_malformedHost(self):
        with self.assertRaises(ValueError):
            c = Client("test_string")

        with self.assertRaises(ValueError):
            c = Client("autoai.org")

    def test_properHost(self):
        c = Client("http://aid.autoai.org")


if __name__ == "__main__":
    unittest.main()
