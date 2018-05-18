package xspinners

type Name int

const (
	DOTS1 Name = iota
	DOTS2
	CIRCLE1
	ARROWS1
)

var (
	spinners = make(map[string]*Spinner)
	lookup   = make(map[int]string)
)

func init() {
	spinners["DOTS1"] = &Spinner{Name: "Dots1", Delay: 160, Frames: []string{".", ":"}}
	spinners["DOTS2"] = &Spinner{Name: "Dots1", Delay: 120, Frames: []string{`⠋`, `⠙`, `⠹`, `⠸`, `⠼`, `⠴`, `⠦`, `⠧`, `⠇`, `⠏`}}
	spinners["CIRCLE1"] = &Spinner{Name: "Circle1", Delay: 120, Frames: []string{`◜`, `◠`, `◝`, `◞`, `◡`, `◟`}}
	spinners["ARROWS1"] = &Spinner{Name: "Arrows1", Delay: 100, Frames: []string{`▹▹▹▹▹`, `▸▹▹▹▹`, `▹▸▹▹▹`, `▹▹▸▹▹`, `▹▹▹▸▹`, `▹▹▹▹▸`}}
	buildSpinnerList()
}

func get(number Name) *Spinner {
	name := getSpinnerName(number)
	if inSpinners(name) {
		return spinners[name]
	}
	return nil
}

func inSpinners(name string) bool {
	if _, ok := spinners[name]; ok {
		return true
	}
	return false
}

func buildSpinnerList() {
	var tmp []string
	for k := range spinners {
		tmp = append(tmp, k)
	}

	for i := 0; i < len(spinners); i++ {
		lookup[i] = tmp[i]
	}
}

func getSpinnerName(name Name) string {
	return lookup[int(name)]
}
