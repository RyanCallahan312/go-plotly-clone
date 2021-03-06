package main

import (
	grob "github.com/RyanCallahan312/go-plotly-clone/graph_objects"
	"github.com/RyanCallahan312/go-plotly-clone/offline"
)

func main() {
	/*
	   fig = dict({
	       "data": [{"type": "bar",
	                 "x": [1, 2, 3],
	                 "y": [1, 3, 2]}],
	       "layout": {"title": {"text": "A Figure Specified By Python Dictionary"}}
	   })
	*/
	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Bar{
				Type: grob.TraceTypeBar,
				X:    []float64{1, 2, 3},
				Y:    []float64{1, 2, 3},
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "A Figure Specified By Go Struct",
			},
		},
		Config: &grob.Config{
			Responsive: grob.True,
		},
	}

	offline.ToHtml(fig, "bar.html")
	offline.Show(fig)
}
