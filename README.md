<p align="center">
<img src="https://files.ieeevit.org/Hacktoberfest-23/Pykitzoid.png">
</p>

# Hacktoberfest 2023 with IEEE-VIT :blue_heart:

This is a repository containing a Python package coded in Go, with the motive of providing ML support.

Support open source software by participating in [Hacktoberfest](https://hacktoberfest.digitalocean.com) and get goodies and digital badges! :blue_heart:

> Please check all issues labelled as `hacktoberfest` to start contributing!

Kindly consider leaving a :star: if you like the repository and our organisation.

## Getting Started

- Fork it.

- Clone your forked repo and move inside it:

`git clone https://github.com/IEEE-VIT/pykitzoid.git && cd pykitzoid`

- Checkout to a new branch to work on an issue:

`git checkout -b my-amazing-feature`

- Get started working!

- Once you're all done coding, it's time to open a PR :)
  Run the following commands from the root of the project directory:

`git add .`

`git commit -m "A short description about the feature."`

`git push origin <my-amazing-feature>`

Open your forked repo in your browser and then raise a PR to the `main` branch of this repository!

## Contributing

To start contributing, check out [CONTRIBUTING.md](https://github.com/IEEE-VIT/ToDo-iOS/blob/master/contributing.md). New contributors are always welcome to support this project. If you want something gentle to start with, check out issues labelled `easy` or `good-first-issue`. Check out issues labelled `hacktoberfest` if you are up for some fun hacktoberfest goodies! :)

## License

See the [LICENSE](https://github.com/kitrak-rev/pykitzoid/blob/main/LICENSE) file for license rights and limitations (MIT).

# test_README.md

import unittest

def binary_sort(arr):
    # Not a standard sort; assuming binary insertion sort
    for i in range(1, len(arr)):
        key = arr[i]
        left, right = 0, i - 1
        while left <= right:
            mid = (left + right) // 2
            if arr[mid] < key:
                left = mid + 1
            else:
                right = mid - 1
        for j in range(i, left, -1):
            arr[j] = arr[j-1]
        arr[left] = key
    return arr

def bubble_sort(arr):
    n = len(arr)
    arr = arr[:]
    for i in range(n):
        for j in range(0, n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
    return arr

def heapify(arr, n, i):
    largest = i
    l = 2*i + 1
    r = 2*i + 2
    if l < n and arr[l] > arr[largest]:
        largest = l
    if r < n and arr[r] > arr[largest]:
        largest = r
    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)

def heap_sort(arr):
    arr = arr[:]
    n = len(arr)
    for i in range(n//2-1, -1, -1):
        heapify(arr, n, i)
    for i in range(n-1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapify(arr, i, 0)
    return arr

def merge_sort(arr):
    if len(arr) <= 1:
        return arr
    mid = len(arr)//2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])
    result = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result

def insertion_sort(arr):
    arr = arr[:]
    for i in range(1, len(arr)):
        key = arr[i]
        j = i-1
        while j >= 0 and arr[j] > key:
            arr[j+1] = arr[j]
            j -= 1
        arr[j+1] = key
    return arr

def quick_sort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr)//2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quick_sort(left) + middle + quick_sort(right)

def pigeonhole_sort(arr):
    if not arr:
        return arr
    arr = arr[:]
    mn, mx = min(arr), max(arr)
    size = mx - mn + 1
    holes = [0] * size
    for x in arr:
        holes[x - mn] += 1
    i = 0
    for count in range(size):
        while holes[count] > 0:
            arr[i] = count + mn
            i += 1
            holes[count] -= 1
    return arr

def counting_sort_for_radix(arr, exp):
    n = len(arr)
    output = [0] * n
    count = [0] * 10
    for i in range(n):
        index = arr[i] // exp
        count[index % 10] += 1
    for i in range(1, 10):
        count[i] += count[i-1]
    for i in range(n-1, -1, -1):
        index = arr[i] // exp
        output[count[index % 10] - 1] = arr[i]
        count[index % 10] -= 1
    for i in range(n):
        arr[i] = output[i]

def radix_sort(arr):
    arr = arr[:]
    if not arr:
        return arr
    mx = max(arr)
    exp = 1
    while mx // exp > 0:
        counting_sort_for_radix(arr, exp)
        exp *= 10
    return arr

def shell_sort(arr):
    arr = arr[:]
    n = len(arr)
    gap = n // 2
    while gap > 0:
        for i in range(gap, n):
            temp = arr[i]
            j = i
            while j >= gap and arr[j-gap] > temp:
                arr[j] = arr[j-gap]
                j -= gap
            arr[j] = temp
        gap //= 2
    return arr

def cocktail_sort(arr):
    arr = arr[:]
    n = len(arr)
    swapped = True
    start = 0
    end = n-1
    while swapped:
        swapped = False
        for i in range(start, end):
            if arr[i] > arr[i+1]:
                arr[i], arr[i+1] = arr[i+1], arr[i]
                swapped = True
        if not swapped:
            break
        swapped = False
        end -= 1
        for i in range(end-1, start-1, -1):
            if arr[i] > arr[i+1]:
                arr[i], arr[i+1] = arr[i+1], arr[i]
                swapped = True
        start += 1
    return arr

class TestSortingAlgorithms(unittest.TestCase):
    def setUp(self):
        self.test_cases = [
            [],
            [1],
            [1, 2, 3, 4, 5],
            [5, 4, 3, 2, 1],
            [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]
        ]
        self.expected = [
            [],
            [1],
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
            [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]
        ]

    def test_binary_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(binary_sort(arr[:]), exp)

    def test_bubble_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(bubble_sort(arr[:]), exp)

    def test_heap_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(heap_sort(arr[:]), exp)

    def test_merge_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(merge_sort(arr[:]), exp)

    def test_insertion_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(insertion_sort(arr[:]), exp)

    def test_quick_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(quick_sort(arr[:]), exp)

    def test_pigeonhole_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(pigeonhole_sort(arr[:]), exp)

    def test_radix_sort(self):
        # Only non-negative integers for radix sort
        for arr, exp in zip(self.test_cases, self.expected):
            if all(x >= 0 for x in arr):
                self.assertEqual(radix_sort(arr[:]), exp)

    def test_shell_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(shell_sort(arr[:]), exp)

    def test_cocktail_sort(self):
        for arr, exp in zip(self.test_cases, self.expected):
            self.assertEqual(cocktail_sort(arr[:]), exp)

if __name__ == "__main__":
    unittest.main()