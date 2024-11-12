#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'extraLongFactorials' function below.
#
# The function accepts INTEGER n as parameter.
#

def extraLongFactorials(n):
  if n < 0:
    return ValueError("n must be non-negative")
  elif n == 0:
    return 1
  else:
    return n * extraLongFactorials(n-1)


if __name__ == '__main__':
    n = int(input().strip())

    result = extraLongFactorials(n)
    print(result)
