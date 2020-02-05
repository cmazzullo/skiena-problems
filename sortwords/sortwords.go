// 4-37. Implement versions of several different sorting algorithms, such
// as selection sort, insertion sort, heapsort, mergesort, and
// quicksort. Conduct experiments to assess the relative performance of
// these algorithms in a simple application that reads a large text file
// and reports exactly one instance of each word that appears within
// it. This application can be efficiently implemented by sorting all the
// words that occur in the text and then passing through the sorted
// sequence to identify one instance of each distinct word. Write a brief
// report with your conclusions.

package sortwords

func Swap(a *[5]string, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func InsertionSort(a *[5]string) {
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				Swap(a, i, j)
			}
		}
	}
}
