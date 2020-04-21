package main

import (
	"bytes"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/GrigoryKrasnochub/go-simple-chart-project/fyne_utils"
	"image"
	"log"
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	a := app.New()
	w := a.NewWindow("Task1")

	imgCanvas := canvas.NewImageFromImage(getChart())
	imgCanvas.FillMode = canvas.ImageFillContain

	graph := fyne.NewContainerWithLayout(
		layout.NewFixedGridLayout(fyne.NewSize(800, 800)),
		imgCanvas,
	)

	variantNumber := widget.NewEntry()
	variantNumber.Text = "0"
	variantNumber.OnChanged = func(value string) {
		fyne_utils.Numeric(&value)
		variantNumber.SetText(value)
	}

	calculationStep := widget.NewEntry()
	calculationStep.Text = "0"
	calculationStep.OnChanged = func(value string) {
		fyne_utils.Numeric(&value)
		calculationStep.SetText(value)
	}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{"Номер варианта", variantNumber},
			{"Шаг отклонения", calculationStep},
		},
		OnSubmit: func() {
			type inputFieldsKeeper struct {
				Widget       *widget.Entry
				From         float64
				To           float64
				DefaultValue float64
			}

			inputFields := [2]inputFieldsKeeper{{variantNumber, 1, 7, 1}, {calculationStep, 0, 0.05, 0.05}}
			for _, inputField := range inputFields {
				value := inputField.Widget.Text
				err := fyne_utils.NumericInDiapason(&value, inputField.From, inputField.To, inputField.DefaultValue)
				if err != nil {
					dialog.ShowError(err, w)
				}
				inputField.Widget.SetText(value)
			}

			//log
			submittedVariantNumber := variantNumber.Text
			submittedCalculationStep := calculationStep.Text
			log.Printf("Form submitted with values: variant %s, step %s \n", submittedVariantNumber, submittedCalculationStep)
		},
	}

	w.SetContent(widget.NewHBox(
		form,
		graph,
	))

	w.CenterOnScreen()
	w.ShowAndRun()

}

func getChart() image.Image {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = ""
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		panic(err)
	}

	p.X.Min = 0
	p.X.Max = 15
	p.Y.Min = 0
	p.Y.Max = 15

	writer, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	writer.WriteTo(buf)
	img, _, err := image.Decode(buf)
	if err != nil {
		panic(err)
	}

	return img
}

func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
