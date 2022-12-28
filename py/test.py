import unittest
from .main import *

class TestHello(unittest.TestCase):
    def test_hello(self):
        self.assertEqual(hello("dr666m1"), "hello dr666m1!")
