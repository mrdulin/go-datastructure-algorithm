import unittest
import math


def binary_search(arr, item):
    """二分查找，一个包含n个元素的列表，最多需要以2为底n的对数次查找

    8个元素的列表，最多需要3步，即可找到元素在列表中的位置，即索引

    Arguments:
        arr {[list]} -- [有序列表]
        item {[type]} -- [要找的元素]

    Returns:
        [int] -- [要找的元素在列表中的索引]
    """
    minIndex = 0
    maxIndex = len(arr) - 1

    while minIndex <= maxIndex:
        midIndex = math.floor((minIndex + maxIndex) / 2)
        print('minIndex: {minIndex}, maxIndex: {maxIndex}, midIndex: {midIndex}'.format(
            minIndex=minIndex, maxIndex=maxIndex, midIndex=midIndex))
        guess = arr[midIndex]
        if guess == item:
            return midIndex
        if guess > item:
            maxIndex = midIndex - 1
        else:
            minIndex = midIndex + 1
    return None


class TestBinarySearch(unittest.TestCase):
    def test_list_with_a_few_num_of_items(self):
        print("===test_list_with_a_few_num_of_items===")
        my_list = [1, 2, 3, 4, 5, 6, 7, 8]
        actual = binary_search(my_list, 5)
        self.assertEqual(actual, 4)

    def test_list_with_a_large_num_of_items(self):
        print("===test_list_with_a_large_num_of_items===")
        my_list = []
        for i in range(1, 101):
            my_list.append(i)
        actual = binary_search(my_list, 99)
        self.assertEqual(actual, 98)


if __name__ == '__main__':
    unittest.main()
