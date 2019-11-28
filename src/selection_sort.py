import unittest


def find_smallest(arr):
    smallest = arr[0]
    smallest_index = 0
    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i
    return smallest_index


def selection_sort(arr):
    """选择排序, 算法运行时间O(n * n)

    Arguments:
        arr {[list]} -- [待排序的数组]

    Returns:
        [list] -- [排序后的数组]
    """
    newArr = []
    for i in range(len(arr)):
        smallest_index = find_smallest(arr)
        newArr.append(arr.pop(smallest_index))
    return newArr


class TestSelectionSort(unittest.TestCase):
    def test_with_a_few_num_of_items(self):
        arr = [5, 3, 6, 2, 10]
        actual = selection_sort(arr)
        expect = [2, 3, 5, 6, 10]
        self.assertListEqual(actual, expect)
        self.assertEqual(actual, expect)


if __name__ == '__main__':
    unittest.main()
