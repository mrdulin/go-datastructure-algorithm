import unittest


def quick_sort(arr):
    print('arr: {arr}'.format(arr=arr))
    if len(arr) < 2:
        return arr
    else:
        pivot = arr[0]
        less = [i for i in arr[1:] if i < pivot]
        greater = [i for i in arr[1:] if i > pivot]
        return quick_sort(less) + [pivot] + quick_sort(greater)


class TestQuickSort(unittest.TestCase):
    def test_list_has_elements(self):
        actual = quick_sort([6, 2, 0, 22, 1, 43])
        self.assertListEqual(actual, [0, 1, 2, 6, 22, 43])


if __name__ == '__main__':
    unittest.main()
