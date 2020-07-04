package charts

import (
	"bytes"
	"image"

	"github.com/GrigoryKrasnochub/go-simple-chart-project/calc/poledata"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func GetChartFromPoleData(polesData []poledata.PoleData, XAxisLabel string, YAxisLabel string) image.Image {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = ""
	p.X.Label.Text = XAxisLabel
	p.Y.Label.Text = YAxisLabel
	//So, i should do it by foreach, but in this way, the styling of line will be the same. So it looks like additional work for Homework))
	err = plotutil.AddLinePoints(p,
		polesData[0].Name, poleDataToPoints(polesData[0]),
		polesData[1].Name, poleDataToPoints(polesData[1]),
		polesData[2].Name, poleDataToPoints(polesData[2]),
	)
	if err != nil {
		panic(err)
	}

	writer, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	_, _ = writer.WriteTo(buf)
	img, _, err := image.Decode(buf)
	if err != nil {
		panic(err)
	}

	return img
}

func poleDataToPoints(data poledata.PoleData) plotter.XYs {
	pts := make(plotter.XYs, len(data.PointsX))
	for i := range pts {
		pts[i].X = data.PointsX[i]
		pts[i].Y = data.PointsY[i]
	}
	return pts
}
