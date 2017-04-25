package strand

const testVersion = 3

var nucleotides = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(input string) string {
	output := make([]rune, len(input))
	for i, r := range input {
		output[i] = nucleotides[r]
	}
	return string(output)
}
