package poledata

import calc2 "github.com/GrigoryKrasnochub/go-simple-chart-project/calc"

type PoleData struct {
	Name    string
	PointsX []float64
	PointsY []float64
}

func GetResource1ToSpentResourceGraph(calcStep []calc2.CalculationStep, Resource1Label string) PoleData {
	var C1 []float64
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

func GetResource2ToSpentResourceGraph(calcStep []calc2.CalculationStep, Resource2Label string) PoleData {
	var C2 []float64
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

func GetSumQualityToSpentResourceGraph(calcStep []calc2.CalculationStep, SumQualityLabel string) PoleData {
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