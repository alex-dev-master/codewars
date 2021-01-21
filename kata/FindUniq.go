package kata


func FindUniq(arr []float32) float32 {
	for i := 0; i < len(arr); i++ {
		current := arr[i]
		if (i + 1) == len(arr) {
			return current
		}
		next := arr[i+1]
		if current != next {
			if (i + 2) >= len(arr) {
				return next
			}
			if current != arr[i+2] {
				return current
			} else {
				return next
			}
		}
	}
	return 0
}
