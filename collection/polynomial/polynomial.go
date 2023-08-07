package polynomial

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Polynomial struct {
	degree      int
	coefficient []int
}

func New() *Polynomial {
	result := new(Polynomial)
	result.coefficient = make([]int, 0)
	result.degree = 0
	return result
}

func (p *Polynomial) Read() {

	scanner := bufio.NewScanner(os.Stdin)

	var degree int
	fmt.Printf("다항식의 최고 차수를 입력하시오: ")
	fmt.Scanln(&degree)

	var numbers = make([]int, 0, degree)
	fmt.Printf("각 항의 계수를 입력하시오 (총 %d개): ", degree+1)

	if scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, " ")
		for _, char := range result {
			num, _ := strconv.Atoi(char)
			numbers = append(numbers, num)
		}
	}
	p.degree = degree
	p.coefficient = numbers
}

func (p *Polynomial) Add(p1, p2 Polynomial) {

}

func (p *Polynomial) Display() {
	length := len(p.coefficient)
	for idx, value := range p.coefficient {
		if idx == length-1 {
			fmt.Printf("%d", value)
		} else {
			fmt.Printf("%d X %d + \t", value, (length-idx)-1)
		}
	}
	println()
}
