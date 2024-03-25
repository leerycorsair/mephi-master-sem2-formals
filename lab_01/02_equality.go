package main

func InterpolationSearch(nums []int, elem int) int {
	l, r := 0, len(nums)-1
	for l <= r && elem >= nums[l] && elem <= nums[r] {
		if l == r {
			if nums[l] == elem {
				return l
			}
			return -1
		}
		pos := l + ((elem-nums[l])*(r-l))/(nums[r]-nums[l]+(nums[r]-nums[l]+1)%2)
		if nums[pos] == elem {
			return pos
		}
		if nums[pos] < elem {
			l = pos + 1
		} else {
			r = pos - 1
		}
	}

	return -1
}

func BinarySearch(nums []int, elem int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == elem {
			return m
		} else if nums[m] < elem {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func LinearSearch(nums []int, elem int) int {
	for i, num := range nums {
		if num == elem {
			return i
		}
	}
	return -1
}
