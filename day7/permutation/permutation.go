package permutation

//Request calculates all permutations if a given string of characters
type Request struct {
	available map[rune]int
	chars     []rune
	results   []string
	temp      []rune
}

//GenerateFor initializes and returns all permutations of the given string of characters
func (c *Request) GenerateFor(s string) []string {
	c.chars = make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		c.chars[i] = rune(s[i])
	}
	c.results = []string{}
	c.temp = make([]rune, len(s))

	c.available = make(map[rune]int)
	for i := 0; i < len(s); i++ {
		if _, ok := c.available[rune(s[i])]; ok {
			c.available[rune(s[i])]++
		} else {
			c.available[rune(s[i])] = 1
		}
	}

	c.step(0)
	return c.results
}

// GetResults return all the found permutations
func (c *Request) getResults() []string {
	results := make([]string, len(c.results))
	for i := 0; i < len(c.results); i++ {
		results[i] = c.results[i]
	}
	return results
}

//Generate Generates all possible permutations of a string
func (c *Request) step(depth int) {
	// only the underlying array of results changes, so return the new copy
	if depth == len(c.chars) {
		c.results = append(c.results, string(c.temp))
	}
	for k, v := range c.available {

		if v > 0 {
			c.available[k]--
			c.temp[depth] = k
			c.step(depth + 1)
			c.available[k]++
			c.temp[depth] = 0
		}
	}
}
