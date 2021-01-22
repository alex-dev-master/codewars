package kata

func Tickets(peopleInLine []int) string {
	if len(peopleInLine) == 0 {
		return "NO"
	}

	var cashier = make(map[int]int)
	cashier[25] = 0
	cashier[50] = 0
	cashier[100] = 0

	for _, cash := range peopleInLine {
		cashier[cash]++
		if cash == 25 {
			continue
		}

		switch cash {
		case 50:
			{
				cashier[25]--
			}
		case 100:
			{
				if cashier[25] >= 1 && cashier[50] >= 1 {
					cashier[25]--
					cashier[50]--
				} else if cashier[25] >= 3 {
					cashier[25] = cashier[25] - 3
				} else {
					return "NO"
				}
			}
		}

		for _, cashierValue := range cashier {
			if cashierValue < 0 {
				return "NO"
			}
		}

	}

	return "YES"
}
