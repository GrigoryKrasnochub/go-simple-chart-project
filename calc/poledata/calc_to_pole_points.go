package poledata

import _calc "github.com/GrigoryKrasnochub/go-simple-chart-project/calc"

type PoleData struct {
	Name    string
	PointsX []float64
	PointsY []float64
}

func GetResource1ToSpentResourceGraph(calcStep []_calc.CalculationStep, Resource1Label string) PoleData {
	C1 := make([]float64, 0, len(calcStep))
	var SpentResource []float64
	for _, step := range calcStep {
		C1 = append(C1, step.Resource1)
		SpentResource = append(SpentResource, step.SpentResource)
	}

	return PoleData{
		Name:    Resource1Label,
		PointsX: SpentResource,
		PointsY: C1,
	}
}

func GetResource2ToSpentResourceGraph(calcStep []_calc.CalculationStep, Resource2Label string) PoleData {
	C2 := make([]float64, 0, len(calcStep))
	var SpentResource []float64
	for _, step := range calcStep {
		C2 = append(C2, step.Resource2)
		SpentResource = append(SpentResource, step.SpentResource)
	}

	return PoleData{
		Name:    Resource2Label,
		PointsX: SpentResource,
		PointsY: C2,
	}
}

func GetSumQualityToSpentResourceGraph(calcStep []_calc.CalculationStep, SumQualityLabel string) PoleData {
	var RecyclingQuality []float64
	var SpentResource []float64
	for _, step := range calcStep {
		RecyclingQuality = append(RecyclingQuality, step.Quality1+step.Quality2)
		SpentResource = append(SpentResource, step.SpentResource)
	}

	return PoleData{
		Name:    SumQualityLabel,
		PointsX: SpentResource,
		PointsY: RecyclingQuality,
	}
}
