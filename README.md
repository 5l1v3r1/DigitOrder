# Problem

For any base N, figure out a number in base N such that the first x digits, written as a number, are divisible by x.

For example, in base 2 it's 10, in base 10 it's 3816547290, etc.

So far I've solved up to 14:

    base 2 list is [1, 0]
    base 4 list is [1, 2, 3, 0]
    base 4 list is [3, 2, 1, 0]
    base 6 list is [1, 4, 3, 2, 5, 0]
    base 6 list is [5, 4, 3, 2, 1, 0]
    base 8 list is [3, 2, 5, 4, 1, 6, 7, 0]
    base 8 list is [5, 2, 3, 4, 7, 6, 1, 0]
    base 8 list is [5, 6, 7, 4, 3, 2, 1, 0]
    base 10 list is [3, 8, 1, 6, 5, 4, 7, 2, 9, 0]
    base 14 list is [9, 12, 3, 10, 5, 4, 7, 6, 11, 8, 1, 2, 13, 0]
