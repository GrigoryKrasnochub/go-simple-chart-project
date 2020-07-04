package main

import (
	"fmt"
	"image"
	"log"
	"strconv"

	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	_calc "github.com/GrigoryKrasnochub/go-simple-chart-project/calc"
	"github.com/GrigoryKrasnochub/go-simple-chart-project/calc/poledata"
	"github.com/GrigoryKrasnochub/go-simple-chart-project/charts"
	"github.com/GrigoryKrasnochub/go-simple-chart-project/fyne_utils"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

type applicationConstantsKeeper struct {
	resultSpentResourceLabel string
	resultResource1Label     string
	resultResource2Label     string
	resultQuality1Label      string
	resultQuality2Label      string
	resultQualitySumLabel    string
}

var applicationConstantsMap = map[string]applicationConstantsKeeper{
	_calc.Maximize.String(): applicationConstantsKeeper{
		resultSpentResourceLabel: "Csum",
		resultResource1Label:     "С1",
		resultResource2Label:     "С2",
		resultQuality1Label:      "W2",
		resultQuality2Label:      "W1",
		resultQualitySumLabel:    "Wsum",
	},
	_calc.Minimize.String(): applicationConstantsKeeper{
		resultSpentResourceLabel: "Lsum",
		resultResource1Label:     "L1",
		resultResource2Label:     "L2",
		resultQuality1Label:      "t1",
		resultQuality2Label:      "t1",
		resultQualitySumLabel:    "tsum",
	},
}

func main() {
	a := app.New()
	w := a.NewWindow("Task1")
	a.Settings().SetTheme(theme.LightTheme())

	//Block graph

	imgCanvas := &canvas.Image{}
	imgCanvas.FillMode = canvas.ImageFillContain

	graph := fyne.NewContainerWithLayout(
		layout.NewFixedGridLayout(fyne.NewSize(800, 800)),
		imgCanvas,
	)

	//Block result

	resultSpentResource := widget.NewLabel("")
	resultSpentResourceLabel := widget.NewLabel("")
	resultResource1Label := widget.NewLabel("")
	resultResource2Label := widget.NewLabel("")
	resultQuality1Label := widget.NewLabel("")
	resultQuality2Label := widget.NewLabel("")
	resultQualitySumLabel := widget.NewLabel("")
	resultResource1 := widget.NewLabel("")
	resultResource2 := widget.NewLabel("")
	resultQuality1 := widget.NewLabel("")
	resultQuality2 := widget.NewLabel("")
	resultQualitySum := widget.NewLabel("")
	resultSpentResourceContainer := widget.NewHBox(resultSpentResourceLabel, resultSpentResource)
	resultResource1Container := widget.NewHBox(resultResource1Label, resultResource1)
	resultResource2Container := widget.NewHBox(resultResource2Label, resultResource2)
	resultQuality1Container := widget.NewHBox(resultQuality1Label, resultQuality1)
	resultQuality2Container := widget.NewHBox(resultQuality2Label, resultQuality2)
	resultQualitySumContainer := widget.NewHBox(resultQualitySumLabel, resultQualitySum)

	resultContainer := widget.NewVBox(
		widget.NewLabel("Results"),
		resultSpentResourceContainer,
		resultResource1Container,
		resultResource2Container,
		resultQuality1Container,
		resultQuality2Container,
		resultQualitySumContainer,
	)

	//Block input form

	variantNumber := widget.NewEntry()
	variantNumber.Text = "0"
	variantNumber.OnChanged = func(value string) {
		variantNumber.SetText(fyne_utils.FilterNumeric(value))
	}

	calculationStep := widget.NewEntry()
	calculationStep.Text = "0"
	calculationStep.OnChanged = func(value string) {
		calculationStep.SetText(fyne_utils.FilterNumeric(value))
	}

	resourceVolume := widget.NewEntry()
	resourceVolume.Text = "0"
	resourceVolume.OnChanged = func(value string) {
		resourceVolume.SetText(fyne_utils.FilterNumeric(value))
	}

	operationType := widget.NewSelect([]string{_calc.Maximize.String(), _calc.Minimize.String()}, func(value string) {
		if settings, found := applicationConstantsMap[value]; found {
			resultSpentResourceLabel.Text = settings.resultSpentResourceLabel
			resultSpentResourceLabel.Text = settings.resultSpentResourceLabel
			resultResource1Label.Text = settings.resultResource1Label
			resultResource2Label.Text = settings.resultResource2Label
			resultQuality1Label.Text = settings.resultQuality1Label
			resultQuality2Label.Text = settings.resultQuality2Label
			resultQualitySumLabel.Text = settings.resultQualitySumLabel
			resultContainer.Refresh()
		}
	})

	operationType.SetSelected("max")

	type inputFieldsKeeper struct {
		Widget       *widget.Entry
		From         float64
		To           float64
		DefaultValue float64
		ValueLabel   string
	}

	inputFields := []inputFieldsKeeper{
		{variantNumber, 1, 7, 1, "variant"},
		{calculationStep, 0, 0.05, 0.05, "step"},
		{resourceVolume, 0, 2.0, 0, "resource"},
	}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{"Operation", operationType},
			{"Variant", variantNumber},
			{"Step", calculationStep},
			{"Max size", resourceVolume},
		},
		OnSubmit: func() {
			submitted := true
			for _, inputField := range inputFields {
				value := inputField.Widget.Text
				value, err := fyne_utils.FilterNumericInDiapason(value, inputField.From, inputField.To, inputField.DefaultValue)
				if err != nil {
					err = fmt.Errorf("error in %s input %v", inputField.ValueLabel, err)
					inputField.Widget.SetText(value)
					dialog.ShowError(err, w)
					submitted = false
					break
				}
				log.Printf("Form submitted with value %s in %s input \n", value, inputField.ValueLabel)
			}

			if !submitted {
				return
			}
			variant, _ := strconv.Atoi(variantNumber.Text)
			step, _ := strconv.ParseFloat(calculationStep.Text, 64)
			volume, _ := strconv.ParseFloat(resourceVolume.Text, 64)
			calc := _calc.Calc{
				VariantNumber:  variant,
				Step:           step,
				Type:           _calc.Type(operationType.Selected),
				ResourceVolume: volume,
				CalcStep: []_calc.CalculationStep{
					{
						Resource1:     0,
						Resource2:     0,
						Quality1:      0,
						Quality2:      0,
						SpentResource: 0,
					},
				},
			}

			log.Println("Form onSubmitClick finish start another goroutine ")

			resultChanel := make(chan _calc.CalculationStep, 1)
			poleChanel := make(chan image.Image, 1)
			go func(resultChanel chan _calc.CalculationStep, poleChanel chan image.Image) {
				calc.DoCalc()
				log.Println("Finish Calculation")
				log.Println(calc.CalcStep)
				resultChanel <- calc.CalcStep[len(calc.CalcStep)-1]

				res := imgCanvas.Image
				res = charts.GetChartFromPoleData([]poledata.PoleData{
					poledata.GetResource1ToSpentResourceGraph(calc.CalcStep, applicationConstantsMap[operationType.Selected].resultResource1Label),
					poledata.GetResource2ToSpentResourceGraph(calc.CalcStep, applicationConstantsMap[operationType.Selected].resultResource2Label),
					poledata.GetSumQualityToSpentResourceGraph(calc.CalcStep, applicationConstantsMap[operationType.Selected].resultQualitySumLabel),
				}, applicationConstantsMap[operationType.Selected].resultSpentResourceLabel, fmt.Sprintf("%s %s %s", applicationConstantsMap[operationType.Selected].resultResource1Label, applicationConstantsMap[operationType.Selected].resultResource2Label, applicationConstantsMap[operationType.Selected].resultQualitySumLabel))
				poleChanel <- res
				log.Println("Form onSubmitClick  another goroutine finish")
			}(resultChanel, poleChanel)

			lastCalcStep := <-resultChanel
			resultSpentResource.Text = fmt.Sprint(lastCalcStep.SpentResource)
			resultResource1.Text = fmt.Sprint(lastCalcStep.Resource1)
			resultResource2.Text = fmt.Sprint(lastCalcStep.Resource2)
			resultQuality1.Text = fmt.Sprint(lastCalcStep.Quality1)
			resultQuality2.Text = fmt.Sprint(lastCalcStep.Quality2)
			resultQualitySum.Text = fmt.Sprint(lastCalcStep.Quality1 + lastCalcStep.Quality2)
			resultContainer.Refresh()

			poleRes := <-poleChanel
			imgCanvas.Image = poleRes
			imgCanvas.Refresh()
			log.Println("Form onSubmitClick main finish")

		},
	}

	w.SetContent(widget.NewHBox(
		widget.NewVBox(
			form,
			resultContainer,
		),
		graph,
	))

	w.CenterOnScreen()
	w.ShowAndRun()
}
