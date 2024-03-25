package main

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func InsertSort(arr []int) {
	for i := 1; i <= len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
