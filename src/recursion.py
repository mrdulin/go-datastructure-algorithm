import unittest


def sum(arr):
    """递归求和，但是这种实现方式是mutable的，即会改变传入的数组

    Arguments:
        arr {[list]} -- [待求和的list]

    Returns:
        [int] -- [和]
    """
    item = arr.pop(0)
    if len(arr):
        item += sum(arr)
    return item


class TestSum(unittest.TestCase):
    def test_sum(self):
        actual = sum([1, 2, 3, 4, 5])
        self.assertEqual(actual, 15)

    def test_sum_mutable(self):
        arr = [1, 2, 3]
        actual = sum(arr)
        self.assertEqual(actual, 6)
        self.assertListEqual(arr, [])


def elements_counter(arr, count=0):
    """递归计算列表包含的元素数

    Arguments:
        arr {[list]} -- [列表]

    Keyword Arguments:
        count {int} -- [列表包含的元素数] (default: {0})

    Returns:
        [int] -- [列表包含的元素数]
    """
    if len(arr):
        arr.pop(0)
        count += 1
        return elements_counter(arr, count)
    return count


class TestElementsCounter(unittest.TestCase):
    def test_list_has_elements(self):
        actual = elements_counter([1, 2, 3, 4, 5, 6])
        self.assertEqual(actual, 6)

    def test_empty_list(self):
        actual = elements_counter([])
        self.assertEqual(actual, 0)


def find_max_num(arr):
    if len(arr) == 1:
        return arr[0]
    else:
        m = find_max_num(arr[1:])
        return m if m > arr[0] else arr[0]


class TestFindMaxNum(unittest.TestCase):
    def test_list_has_elements(self):
        actual = find_max_num([5, 2, 7, 4, 10])
        self.assertEqual(actual, 10)


if __name__ == '__main__':
    unittest.main()
