package kata

func Tickets(peopleInLine []int) string {
	sum := 0

	for _, item := range peopleInLine {
		switch item {
		case 25:
			{
				sum += 25
			}
		case 50:
			{
				sum += 25
				sum -= 25
			}
		case 100:
			{
				sum += 25
				sum -= 75
			}

		}

		if sum < 0 {
			return "NO"
		}
	}

	return "YES"
}
