import unittest
from .main import *

class TestHello(unittest.TestCase):
    def test_hello(self):
        self.assertEqual(hello("world"), "hello world!")
