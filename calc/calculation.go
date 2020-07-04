package calc

import (
	"log"
	"math"
)

type CalculationStep struct {
	Resource1     float64
	Resource2     float64
	Quality1      float64
	Quality2      float64
	SpentResource float64
}

type Calc struct {
	VariantNumber  int
	Type           Type
	Step           float64
	ResourceVolume float64
	CalcStep       []CalculationStep
}

func (calc *Calc) DoCalc() {
	if calc.Type == Maximize {
		log.Printf("Maximize!\n")
		calc.maximize()
	}
	if calc.Type == Minimize {
		log.Printf("Minimize!\n")
		calc.minimize()
	}
}

func (calc *Calc) maximize() {
	for calc.ResourceVolume > 0 {
		currentStep := calc.Step
		log.Printf("New iteration! Step:%.2f", calc.Step)
		if calc.ResourceVolume < currentStep {
			currentStep = calc.ResourceVolume
		}
		calc.ResourceVolume -= currentStep
		currentC1 := calc.CalcStep[len(calc.CalcStep)-1].Resource1
		currentC2 := calc.CalcStep[len(calc.CalcStep)-1].Resource2
		w1Derivative := calc.calcW1Derivative(currentC1)
		w2Derivative := calc.calcW2Derivative(currentC2)

		if w1Derivative > w2Derivative {
			currentC1 += currentStep
		}

		if w1Derivative < w2Derivative {
			currentC2 += currentStep
		}

		if w1Derivative == w2Derivative {
			currentStep = currentStep / 2
			currentC1 += currentStep
			currentC2 += currentStep
		}

		calc.CalcStep = append(calc.CalcStep, CalculationStep{
			Resource1:     currentC1,
			Resource2:     currentC2,
			Quality1:      calc.calcW1(currentC1),
			Quality2:      calc.calcW2(currentC2),
			SpentResource: currentC1 + currentC2,
		})
	}
}

func (calc *Calc) calcW1Derivative(C1 float64) float64 {
	return 3 * math.Exp(-3*C1)
}

func (calc *Calc) calcW2Derivative(C2 float64) float64 {
	return (2 * math.Exp(-C2/float64(calc.VariantNumber))) / float64(calc.VariantNumber)
}

func (calc *Calc) calcW1(C1 float64) float64 {
	return 1 - math.Exp(-3*C1)
}

func (calc *Calc) calcW2(C2 float64) float64 {
	return 2 - 2*math.Exp(-C2/float64(calc.VariantNumber))
}

func (calc *Calc) minimize() {
	for calc.ResourceVolume > 0 {
		currentStep := calc.Step
		log.Printf("New iteration! Step:%.2f", calc.Step)
		if calc.ResourceVolume < currentStep {
			currentStep = calc.ResourceVolume
		}
		calc.ResourceVolume -= currentStep
		currentL1 := calc.CalcStep[len(calc.CalcStep)-1].Resource1
		currentL2 := calc.CalcStep[len(calc.CalcStep)-1].Resource2
		t1Derivative := calc.calcW1Derivative(currentL1)
		t2Derivative := calc.calcW2Derivative(currentL2)

		if t1Derivative > t2Derivative {
			currentL1 += currentStep
		}

		if t1Derivative < t2Derivative {
			currentL2 += currentStep
		}

		if t1Derivative == t2Derivative {
			currentStep = currentStep / 2
			currentL1 += currentStep
			currentL2 += currentStep
		}

		calc.CalcStep = append(calc.CalcStep, CalculationStep{
			Resource1:     currentL1,
			Resource2:     currentL2,
			Quality1:      calc.calcT1(currentL1),
			Quality2:      calc.calcT2(currentL2),
			SpentResource: currentL1 + currentL2,
		})
	}
}

func (calc *Calc) calcT1(L1 float64) float64 {
	return 1 / ((1 + 1/float64(calc.VariantNumber)) - L1)
}

func (calc *Calc) calcT2(L2 float64) float64 {
	return 1 / (3 - L2)
}
