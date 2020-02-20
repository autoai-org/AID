import unittest

from cvpm.Server import get_available_port

class NetworkTester(unittest.TestCase):

    def test_available_port(self):
        print (get_available_port(8080))
        self.assertTrue(True)

if __name__ == "__main__":
    unittest.main()