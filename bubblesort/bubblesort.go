/**
The program prompts the user to type in a sequence of up to 10 integers,
sort them using bubble sort algorithm and prints the result in sorted order.
*/
package bubblesort

func BubbleSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
			}
		}
	}
}

func Swap(sli []int, index int) {
	if index < 0 || index >= len(sli) {
		return
	}
	aux := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = aux
}
