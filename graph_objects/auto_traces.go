package grob

import (
	"encoding/json"
	"errors"
)

// Area <no value>
type Area struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo AreaHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *AreaHoverlabel `json:"hoverlabel,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *AreaMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// R data_array Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the radial coordinates for legacy polar chart only.
	R interface{} `json:"r,omitempty"`

	// Rsrc string Sets the source reference on Chart Studio Cloud for  r .
	Rsrc String `json:"rsrc,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *AreaStream `json:"stream,omitempty"`

	// T data_array Area traces are deprecated! Please switch to the *barpolar* trace type. Sets the angular coordinates for legacy polar chart only.
	T interface{} `json:"t,omitempty"`

	// Transforms <no value> <no value>
	Transforms AreaTransforms `json:"transforms,omitempty"`

	// Tsrc string Sets the source reference on Chart Studio Cloud for  t .
	Tsrc String `json:"tsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible AreaVisible `json:"visible,omitempty"`
}

func NewArea() *Area {
	return &Area{
		Type: TraceTypeArea,
	}
}

func (this *Area) GetType() TraceType {
	return this.Type
}

// Bar The data visualized by the span of the bars is set in `y` if `orientation` is set th *v* (the default) and the labels are set in `x`. By setting `orientation` to *h*, the roles are interchanged.
type Bar struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// _deprecated <no value> <no value>
	// Pending... _deprecated No valTyp <no value> `json:"_deprecated,omitempty"`

	// Alignmentgroup string Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Base any Sets where the bar base is drawn (in position axis units). In *stack* or *relative* barmode, traces that set *base* will be excluded and drawn in *overlay* mode instead.
	Base interface{} `json:"base,omitempty"`

	// Basesrc string Sets the source reference on Chart Studio Cloud for  base .
	Basesrc String `json:"basesrc,omitempty"`

	// Cliponaxis boolean Determines whether the text nodes are clipped about the subplot axes. To show the text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*.
	Cliponaxis Bool `json:"cliponaxis,omitempty"`

	// Constraintext enumerated Constrain the size of text inside or outside a bar to be no larger than the bar itself.
	Constraintext BarConstraintext `json:"constraintext,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// Error_x <no value> <no value>
	Error_x *BarErrorX `json:"error_x,omitempty"`

	// Error_y <no value> <no value>
	Error_y *BarErrorY `json:"error_y,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo BarHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *BarHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `value` and `label`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Insidetextanchor enumerated Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
	Insidetextanchor BarInsidetextanchor `json:"insidetextanchor,omitempty"`

	// Insidetextfont <no value> Sets the font used for `text` lying inside the bar.
	Insidetextfont *BarInsidetextfont `json:"insidetextfont,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *BarMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Offset number Shifts the position where the bar is drawn (in position axis units). In *group* barmode, traces that set *offset* will be excluded and drawn in *overlay* mode instead.
	Offset interface{} `json:"offset,omitempty"`

	// Offsetgroup string Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up.
	Offsetgroup String `json:"offsetgroup,omitempty"`

	// Offsetsrc string Sets the source reference on Chart Studio Cloud for  offset .
	Offsetsrc String `json:"offsetsrc,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation enumerated Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
	Orientation BarOrientation `json:"orientation,omitempty"`

	// Outsidetextfont <no value> Sets the font used for `text` lying outside the bar.
	Outsidetextfont *BarOutsidetextfont `json:"outsidetextfont,omitempty"`

	// R data_array r coordinates in scatter traces are deprecated!Please switch to the *scatterpolar* trace type.Sets the radial coordinatesfor legacy polar chart only.
	R interface{} `json:"r,omitempty"`

	// Rsrc string Sets the source reference on Chart Studio Cloud for  r .
	Rsrc String `json:"rsrc,omitempty"`

	// Selected <no value> <no value>
	Selected *BarSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *BarStream `json:"stream,omitempty"`

	// T data_array t coordinates in scatter traces are deprecated!Please switch to the *scatterpolar* trace type.Sets the angular coordinatesfor legacy polar chart only.
	T interface{} `json:"t,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textangle angle Sets the angle of the tick labels with respect to the bar. For example, a `tickangle` of -90 draws the tick labels vertically. With *auto* the texts may automatically be rotated to fit with the maximum size in bars.
	Textangle float64 `json:"textangle,omitempty"`

	// Textfont <no value> Sets the font used for `text`.
	Textfont *BarTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
	Textposition BarTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `value` and `label`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms BarTransforms `json:"transforms,omitempty"`

	// Tsrc string Sets the source reference on Chart Studio Cloud for  t .
	Tsrc String `json:"tsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *BarUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible BarVisible `json:"visible,omitempty"`

	// Width number Sets the bar width (in position axis units).
	Width interface{} `json:"width,omitempty"`

	// Widthsrc string Sets the source reference on Chart Studio Cloud for  width .
	Widthsrc String `json:"widthsrc,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar BarXcalendar `json:"xcalendar,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment BarXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar BarYcalendar `json:"ycalendar,omitempty"`

	// Yperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment BarYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewBar() *Bar {
	return &Bar{
		Type: TraceTypeBar,
	}
}

func (this *Bar) GetType() TraceType {
	return this.Type
}

// Barpolar The data visualized by the radial span of the bars is set in `r`
type Barpolar struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Base any Sets where the bar base is drawn (in radial axis units). In *stack* barmode, traces that set *base* will be excluded and drawn in *overlay* mode instead.
	Base interface{} `json:"base,omitempty"`

	// Basesrc string Sets the source reference on Chart Studio Cloud for  base .
	Basesrc String `json:"basesrc,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dr number Sets the r coordinate step.
	Dr float64 `json:"dr,omitempty"`

	// Dtheta number Sets the theta coordinate step. By default, the `dtheta` step equals the subplot's period divided by the length of the `r` coordinates.
	Dtheta float64 `json:"dtheta,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo BarpolarHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *BarpolarHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *BarpolarMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Offset number Shifts the angular position where the bar is drawn (in *thetatunit* units).
	Offset interface{} `json:"offset,omitempty"`

	// Offsetsrc string Sets the source reference on Chart Studio Cloud for  offset .
	Offsetsrc String `json:"offsetsrc,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// R data_array Sets the radial coordinates
	R interface{} `json:"r,omitempty"`

	// R0 any Alternate to `r`. Builds a linear space of r coordinates. Use with `dr` where `r0` is the starting coordinate and `dr` the step.
	R0 interface{} `json:"r0,omitempty"`

	// Rsrc string Sets the source reference on Chart Studio Cloud for  r .
	Rsrc String `json:"rsrc,omitempty"`

	// Selected <no value> <no value>
	Selected *BarpolarSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *BarpolarStream `json:"stream,omitempty"`

	// Subplot subplotid Sets a reference between this trace's data coordinates and a polar subplot. If *polar* (the default value), the data refer to `layout.polar`. If *polar2*, the data refer to `layout.polar2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Text string Sets hover text elements associated with each bar. If a single string, the same string appears over all bars. If an array of string, the items are mapped in order to the this trace's coordinates.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Theta data_array Sets the angular coordinates
	Theta interface{} `json:"theta,omitempty"`

	// Theta0 any Alternate to `theta`. Builds a linear space of theta coordinates. Use with `dtheta` where `theta0` is the starting coordinate and `dtheta` the step.
	Theta0 interface{} `json:"theta0,omitempty"`

	// Thetasrc string Sets the source reference on Chart Studio Cloud for  theta .
	Thetasrc String `json:"thetasrc,omitempty"`

	// Thetaunit enumerated Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
	Thetaunit BarpolarThetaunit `json:"thetaunit,omitempty"`

	// Transforms <no value> <no value>
	Transforms BarpolarTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *BarpolarUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible BarpolarVisible `json:"visible,omitempty"`

	// Width number Sets the bar angular width (in *thetaunit* units).
	Width interface{} `json:"width,omitempty"`

	// Widthsrc string Sets the source reference on Chart Studio Cloud for  width .
	Widthsrc String `json:"widthsrc,omitempty"`
}

func NewBarpolar() *Barpolar {
	return &Barpolar{
		Type: TraceTypeBarpolar,
	}
}

func (this *Barpolar) GetType() TraceType {
	return this.Type
}

// Box Each box spans from quartile 1 (Q1) to quartile 3 (Q3). The second quartile (Q2, i.e. the median) is marked by a line inside the box. The fences grow outward from the boxes' edges, by default they span +/- 1.5 times the interquartile range (IQR: Q3-Q1), The sample mean and standard deviation as well as notches and the sample, outlier and suspected outliers points can be optionally added to the box plot. The values and positions corresponding to each boxes can be input using two signatures. The first signature expects users to supply the sample values in the `y` data array for vertical boxes (`x` for horizontal boxes). By supplying an `x` (`y`) array, one box per distinct `x` (`y`) value is drawn If no `x` (`y`) {array} is provided, a single box is drawn. In this case, the box is positioned with the trace `name` or with `x0` (`y0`) if provided. The second signature expects users to supply the boxes corresponding Q1, median and Q3 statistics in the `q1`, `median` and `q3` data arrays respectively. Other box features relying on statistics namely `lowerfence`, `upperfence`, `notchspan` can be set directly by the users. To have plotly compute them or to show sample points besides the boxes, users can set the `y` data array for vertical boxes (`x` for horizontal boxes) to a 2D array with the outer length corresponding to the number of boxes in the traces and the inner length corresponding the sample size.
type Box struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Alignmentgroup string Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Boxmean enumerated If *true*, the mean of the box(es)' underlying distribution is drawn as a dashed line inside the box(es). If *sd* the standard deviation is also drawn. Defaults to *true* when `mean` is set. Defaults to *sd* when `sd` is set Otherwise defaults to *false*.
	Boxmean BoxBoxmean `json:"boxmean,omitempty"`

	// Boxpoints enumerated If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the box(es) are shown with no sample points Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set. Defaults to *all* under the q1/median/q3 signature. Otherwise defaults to *outliers*.
	Boxpoints BoxBoxpoints `json:"boxpoints,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Sets the x coordinate step for multi-box traces set using q1/median/q3.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step for multi-box traces set using q1/median/q3.
	Dy float64 `json:"dy,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo BoxHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *BoxHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron flaglist Do the hover effects highlight individual boxes  or sample points or both?
	Hoveron BoxHoveron `json:"hoveron,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Jitter number Sets the amount of jitter in the sample points drawn. If *0*, the sample points align along the distribution axis. If *1*, the sample points are drawn in a random jitter of width equal to the width of the box(es).
	Jitter float64 `json:"jitter,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *BoxLine `json:"line,omitempty"`

	// Lowerfence data_array Sets the lower fence values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `lowerfence` is not provided but a sample (in `y` or `x`) is set, we compute the lower as the last sample point below 1.5 times the IQR.
	Lowerfence interface{} `json:"lowerfence,omitempty"`

	// Lowerfencesrc string Sets the source reference on Chart Studio Cloud for  lowerfence .
	Lowerfencesrc String `json:"lowerfencesrc,omitempty"`

	// Marker <no value> <no value>
	Marker *BoxMarker `json:"marker,omitempty"`

	// Mean data_array Sets the mean values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `mean` is not provided but a sample (in `y` or `x`) is set, we compute the mean for each box using the sample values.
	Mean interface{} `json:"mean,omitempty"`

	// Meansrc string Sets the source reference on Chart Studio Cloud for  mean .
	Meansrc String `json:"meansrc,omitempty"`

	// Median data_array Sets the median values. There should be as many items as the number of boxes desired.
	Median interface{} `json:"median,omitempty"`

	// Mediansrc string Sets the source reference on Chart Studio Cloud for  median .
	Mediansrc String `json:"mediansrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover. For box traces, the name will also be used for the position coordinate, if `x` and `x0` (`y` and `y0` if horizontal) are missing and the position axis is categorical
	Name String `json:"name,omitempty"`

	// Notched boolean Determines whether or not notches are drawn. Notches displays a confidence interval around the median. We compute the confidence interval as median +/- 1.57 * IQR / sqrt(N), where IQR is the interquartile range and N is the sample size. If two boxes' notches do not overlap there is 95% confidence their medians differ. See https://sites.google.com/site/davidsstatistics/home/notched-box-plots for more info. Defaults to *false* unless `notchwidth` or `notchspan` is set.
	Notched Bool `json:"notched,omitempty"`

	// Notchspan data_array Sets the notch span from the boxes' `median` values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `notchspan` is not provided but a sample (in `y` or `x`) is set, we compute it as 1.57 * IQR / sqrt(N), where N is the sample size.
	Notchspan interface{} `json:"notchspan,omitempty"`

	// Notchspansrc string Sets the source reference on Chart Studio Cloud for  notchspan .
	Notchspansrc String `json:"notchspansrc,omitempty"`

	// Notchwidth number Sets the width of the notches relative to the box' width. For example, with 0, the notches are as wide as the box(es).
	Notchwidth float64 `json:"notchwidth,omitempty"`

	// Offsetgroup string Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up.
	Offsetgroup String `json:"offsetgroup,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation enumerated Sets the orientation of the box(es). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
	Orientation BoxOrientation `json:"orientation,omitempty"`

	// Pointpos number Sets the position of the sample points in relation to the box(es). If *0*, the sample points are places over the center of the box(es). Positive (negative) values correspond to positions to the right (left) for vertical boxes and above (below) for horizontal boxes
	Pointpos float64 `json:"pointpos,omitempty"`

	// Q1 data_array Sets the Quartile 1 values. There should be as many items as the number of boxes desired.
	Q1 interface{} `json:"q1,omitempty"`

	// Q1src string Sets the source reference on Chart Studio Cloud for  q1 .
	Q1src String `json:"q1src,omitempty"`

	// Q3 data_array Sets the Quartile 3 values. There should be as many items as the number of boxes desired.
	Q3 interface{} `json:"q3,omitempty"`

	// Q3src string Sets the source reference on Chart Studio Cloud for  q3 .
	Q3src String `json:"q3src,omitempty"`

	// Quartilemethod enumerated Sets the method used to compute the sample's Q1 and Q3 quartiles. The *linear* method uses the 25th percentile for Q1 and 75th percentile for Q3 as computed using method #10 (listed on http://www.amstat.org/publications/jse/v14n3/langford.html). The *exclusive* method uses the median to divide the ordered dataset into two halves if the sample is odd, it does not include the median in either half - Q1 is then the median of the lower half and Q3 the median of the upper half. The *inclusive* method also uses the median to divide the ordered dataset into two halves but if the sample is odd, it includes the median in both halves - Q1 is then the median of the lower half and Q3 the median of the upper half.
	Quartilemethod BoxQuartilemethod `json:"quartilemethod,omitempty"`

	// Sd data_array Sets the standard deviation values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `sd` is not provided but a sample (in `y` or `x`) is set, we compute the standard deviation for each box using the sample values.
	Sd interface{} `json:"sd,omitempty"`

	// Sdsrc string Sets the source reference on Chart Studio Cloud for  sd .
	Sdsrc String `json:"sdsrc,omitempty"`

	// Selected <no value> <no value>
	Selected *BoxSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *BoxStream `json:"stream,omitempty"`

	// Text string Sets the text elements associated with each sample value. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms BoxTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *BoxUnselected `json:"unselected,omitempty"`

	// Upperfence data_array Sets the upper fence values. There should be as many items as the number of boxes desired. This attribute has effect only under the q1/median/q3 signature. If `upperfence` is not provided but a sample (in `y` or `x`) is set, we compute the lower as the last sample point above 1.5 times the IQR.
	Upperfence interface{} `json:"upperfence,omitempty"`

	// Upperfencesrc string Sets the source reference on Chart Studio Cloud for  upperfence .
	Upperfencesrc String `json:"upperfencesrc,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible BoxVisible `json:"visible,omitempty"`

	// Whiskerwidth number Sets the width of the whiskers relative to the box' width. For example, with 1, the whiskers are as wide as the box(es).
	Whiskerwidth float64 `json:"whiskerwidth,omitempty"`

	// Width number Sets the width of the box in data coordinate If *0* (default value) the width is automatically selected based on the positions of other box traces in the same subplot.
	Width float64 `json:"width,omitempty"`

	// X data_array Sets the x sample data or coordinates. See overview for more info.
	X interface{} `json:"x,omitempty"`

	// X0 any Sets the x coordinate for single-box traces or the starting coordinate for multi-box traces set using q1/median/q3. See overview for more info.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar BoxXcalendar `json:"xcalendar,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment BoxXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y sample data or coordinates. See overview for more info.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Sets the y coordinate for single-box traces or the starting coordinate for multi-box traces set using q1/median/q3. See overview for more info.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar BoxYcalendar `json:"ycalendar,omitempty"`

	// Yperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment BoxYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewBox() *Box {
	return &Box{
		Type: TraceTypeBox,
	}
}

func (this *Box) GetType() TraceType {
	return this.Type
}

// Candlestick The candlestick is a style of financial chart describing open, high, low and close for a given `x` coordinate (most likely time). The boxes represent the spread between the `open` and `close` values and the lines represent the spread between the `low` and `high` values Sample points where the close value is higher (lower) then the open value are called increasing (decreasing). By default, increasing candles are drawn in green whereas decreasing are drawn in red.
type Candlestick struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Close data_array Sets the close values.
	Close interface{} `json:"close,omitempty"`

	// Closesrc string Sets the source reference on Chart Studio Cloud for  close .
	Closesrc String `json:"closesrc,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Decreasing <no value> <no value>
	Decreasing *CandlestickDecreasing `json:"decreasing,omitempty"`

	// High data_array Sets the high values.
	High interface{} `json:"high,omitempty"`

	// Highsrc string Sets the source reference on Chart Studio Cloud for  high .
	Highsrc String `json:"highsrc,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo CandlestickHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *CandlestickHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Increasing <no value> <no value>
	Increasing *CandlestickIncreasing `json:"increasing,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *CandlestickLine `json:"line,omitempty"`

	// Low data_array Sets the low values.
	Low interface{} `json:"low,omitempty"`

	// Lowsrc string Sets the source reference on Chart Studio Cloud for  low .
	Lowsrc String `json:"lowsrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Open data_array Sets the open values.
	Open interface{} `json:"open,omitempty"`

	// Opensrc string Sets the source reference on Chart Studio Cloud for  open .
	Opensrc String `json:"opensrc,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *CandlestickStream `json:"stream,omitempty"`

	// Text string Sets hover text elements associated with each sample point. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to this trace's sample points.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms CandlestickTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible CandlestickVisible `json:"visible,omitempty"`

	// Whiskerwidth number Sets the width of the whiskers relative to the box' width. For example, with 1, the whiskers are as wide as the box(es).
	Whiskerwidth float64 `json:"whiskerwidth,omitempty"`

	// X data_array Sets the x coordinates. If absent, linear coordinate will be generated.
	X interface{} `json:"x,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar CandlestickXcalendar `json:"xcalendar,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment CandlestickXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`
}

func NewCandlestick() *Candlestick {
	return &Candlestick{
		Type: TraceTypeCandlestick,
	}
}

func (this *Candlestick) GetType() TraceType {
	return this.Type
}

// Carpet The data describing carpet axis layout is set in `y` and (optionally) also `x`. If only `y` is present, `x` the plot is interpreted as a cheater plot and is filled in using the `y` values. `x` and `y` may either be 2D arrays matching with each dimension matching that of `a` and `b`, or they may be 1D arrays with total length equal to that of `a` and `b`.
type Carpet struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// A data_array An array containing values of the first parameter value
	A interface{} `json:"a,omitempty"`

	// A0 number Alternate to `a`. Builds a linear space of a coordinates. Use with `da` where `a0` is the starting coordinate and `da` the step.
	A0 float64 `json:"a0,omitempty"`

	// Aaxis <no value> <no value>
	Aaxis *CarpetAaxis `json:"aaxis,omitempty"`

	// Asrc string Sets the source reference on Chart Studio Cloud for  a .
	Asrc String `json:"asrc,omitempty"`

	// B data_array A two dimensional array of y coordinates at each carpet point.
	B interface{} `json:"b,omitempty"`

	// B0 number Alternate to `b`. Builds a linear space of a coordinates. Use with `db` where `b0` is the starting coordinate and `db` the step.
	B0 float64 `json:"b0,omitempty"`

	// Baxis <no value> <no value>
	Baxis *CarpetBaxis `json:"baxis,omitempty"`

	// Bsrc string Sets the source reference on Chart Studio Cloud for  b .
	Bsrc String `json:"bsrc,omitempty"`

	// Carpet string An identifier for this carpet, so that `scattercarpet` and `contourcarpet` traces can specify a carpet plot on which they lie
	Carpet String `json:"carpet,omitempty"`

	// Cheaterslope number The shift applied to each successive row of data in creating a cheater plot. Only used if `x` is been omitted.
	Cheaterslope float64 `json:"cheaterslope,omitempty"`

	// Color color Sets default for all colors associated with this axis all at once: line, font, tick, and grid colors. Grid color is lightened by blending this with the plot background Individual pieces can override this.
	Color String `json:"color,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Da number Sets the a coordinate step. See `a0` for more info.
	Da float64 `json:"da,omitempty"`

	// Db number Sets the b coordinate step. See `b0` for more info.
	Db float64 `json:"db,omitempty"`

	// Font <no value> The default font used for axis & tick labels on this carpet
	Font *CarpetFont `json:"font,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Stream <no value> <no value>
	Stream *CarpetStream `json:"stream,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible CarpetVisible `json:"visible,omitempty"`

	// X data_array A two dimensional array of x coordinates at each carpet point. If omitted, the plot is a cheater plot and the xaxis is hidden by default.
	X interface{} `json:"x,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array A two dimensional array of y coordinates at each carpet point.
	Y interface{} `json:"y,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewCarpet() *Carpet {
	return &Carpet{
		Type: TraceTypeCarpet,
	}
}

func (this *Carpet) GetType() TraceType {
	return this.Type
}

// Choropleth The data that describes the choropleth value-to-color mapping is set in `z`. The geographic locations corresponding to each value in `z` are set in `locations`.
type Choropleth struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *ChoroplethColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Featureidkey string Sets the key in GeoJSON features which is used as id to match the items included in the `locations` array. Only has an effect when `geojson` is set. Support nested property, for example *properties.name*.
	Featureidkey String `json:"featureidkey,omitempty"`

	// Geo subplotid Sets a reference between this trace's geospatial coordinates and a geographic map. If *geo* (the default value), the geospatial coordinates refer to `layout.geo`. If *geo2*, the geospatial coordinates refer to `layout.geo2`, and so on.
	Geo String `json:"geo,omitempty"`

	// Geojson any Sets optional GeoJSON data associated with this trace. If not given, the features on the base map are used. It can be set as a valid GeoJSON object or as a URL string. Note that we only accept GeoJSONs of type *FeatureCollection* or *Feature* with geometries of type *Polygon* or *MultiPolygon*.
	Geojson interface{} `json:"geojson,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ChoroplethHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ChoroplethHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Locationmode enumerated Determines the set of locations used to match entries in `locations` to regions on the map. Values *ISO-3*, *USA-states*, *country names* correspond to features on the base map and value *geojson-id* corresponds to features from a custom GeoJSON linked to the `geojson` attribute.
	Locationmode ChoroplethLocationmode `json:"locationmode,omitempty"`

	// Locations data_array Sets the coordinates via location IDs or names. See `locationmode` for more info.
	Locations interface{} `json:"locations,omitempty"`

	// Locationssrc string Sets the source reference on Chart Studio Cloud for  locations .
	Locationssrc String `json:"locationssrc,omitempty"`

	// Marker <no value> <no value>
	Marker *ChoroplethMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Selected <no value> <no value>
	Selected *ChoroplethSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *ChoroplethStream `json:"stream,omitempty"`

	// Text string Sets the text elements associated with each location.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ChoroplethTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ChoroplethUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ChoroplethVisible `json:"visible,omitempty"`

	// Z data_array Sets the color values.
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewChoropleth() *Choropleth {
	return &Choropleth{
		Type: TraceTypeChoropleth,
	}
}

func (this *Choropleth) GetType() TraceType {
	return this.Type
}

// Choroplethmapbox GeoJSON features to be filled are set in `geojson` The data that describes the choropleth value-to-color mapping is set in `locations` and `z`.
type Choroplethmapbox struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Below string Determines if the choropleth polygons will be inserted before the layer with the specified ID. By default, choroplethmapbox traces are placed above the water layers. If set to '', the layer will be inserted above every existing layer.
	Below String `json:"below,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *ChoroplethmapboxColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Featureidkey string Sets the key in GeoJSON features which is used as id to match the items included in the `locations` array. Support nested property, for example *properties.name*.
	Featureidkey String `json:"featureidkey,omitempty"`

	// Geojson any Sets the GeoJSON data associated with this trace. It can be set as a valid GeoJSON object or as a URL string. Note that we only accept GeoJSONs of type *FeatureCollection* or *Feature* with geometries of type *Polygon* or *MultiPolygon*.
	Geojson interface{} `json:"geojson,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ChoroplethmapboxHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ChoroplethmapboxHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variable `properties` Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Locations data_array Sets which features found in *geojson* to plot using their feature `id` field.
	Locations interface{} `json:"locations,omitempty"`

	// Locationssrc string Sets the source reference on Chart Studio Cloud for  locations .
	Locationssrc String `json:"locationssrc,omitempty"`

	// Marker <no value> <no value>
	Marker *ChoroplethmapboxMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Selected <no value> <no value>
	Selected *ChoroplethmapboxSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *ChoroplethmapboxStream `json:"stream,omitempty"`

	// Subplot subplotid Sets a reference between this trace's data coordinates and a mapbox subplot. If *mapbox* (the default value), the data refer to `layout.mapbox`. If *mapbox2*, the data refer to `layout.mapbox2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Text string Sets the text elements associated with each location.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ChoroplethmapboxTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ChoroplethmapboxUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ChoroplethmapboxVisible `json:"visible,omitempty"`

	// Z data_array Sets the color values.
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewChoroplethmapbox() *Choroplethmapbox {
	return &Choroplethmapbox{
		Type: TraceTypeChoroplethmapbox,
	}
}

func (this *Choroplethmapbox) GetType() TraceType {
	return this.Type
}

// Cone Use cone traces to visualize vector fields.  Specify a vector field using 6 1D arrays, 3 position arrays `x`, `y` and `z` and 3 vector component arrays `u`, `v`, `w`. The cones are drawn exactly at the positions given by `x`, `y` and `z`.
type Cone struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Anchor enumerated Sets the cones' anchor with respect to their x/y/z positions. Note that *cm* denote the cone's center of mass which corresponds to 1/4 from the tail to tip.
	Anchor ConeAnchor `json:"anchor,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto boolean Determines whether or not the color domain is computed with respect to the input data (here u/v/w norm) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax number Sets the upper bound of the color domain. Value should have the same units as u/v/w norm and if set, `cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid number Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as u/v/w norm. Has no effect when `cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin number Sets the lower bound of the color domain. Value should have the same units as u/v/w norm and if set, `cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *ConeColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ConeHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ConeHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variable `norm` Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Lighting <no value> <no value>
	Lighting *ConeLighting `json:"lighting,omitempty"`

	// Lightposition <no value> <no value>
	Lightposition *ConeLightposition `json:"lightposition,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the surface. Please note that in the case of using high `opacity` values for example a value greater than or equal to 0.5 on two surfaces (and 0.25 with four surfaces), an overlay of multiple transparent surfaces may not perfectly be sorted in depth by the webgl API. This behavior may be improved in the near future and is subject to change.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Scene subplotid Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Sizemode enumerated Determines whether `sizeref` is set as a *scaled* (i.e unitless) scalar (normalized by the max u/v/w norm in the vector field) or as *absolute* value (in the same units as the vector field).
	Sizemode ConeSizemode `json:"sizemode,omitempty"`

	// Sizeref number Adjusts the cone size scaling. The size of the cones is determined by their u/v/w norm multiplied a factor and `sizeref`. This factor (computed internally) corresponds to the minimum "time" to travel across two successive x/y/z positions at the average velocity of those two successive positions. All cones in a given trace use the same factor. With `sizemode` set to *scaled*, `sizeref` is unitless, its default value is *0.5* With `sizemode` set to *absolute*, `sizeref` has the same units as the u/v/w vector field, its the default value is half the sample's maximum vector norm.
	Sizeref float64 `json:"sizeref,omitempty"`

	// Stream <no value> <no value>
	Stream *ConeStream `json:"stream,omitempty"`

	// Text string Sets the text elements associated with the cones. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// U data_array Sets the x components of the vector field.
	U interface{} `json:"u,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Usrc string Sets the source reference on Chart Studio Cloud for  u .
	Usrc String `json:"usrc,omitempty"`

	// V data_array Sets the y components of the vector field.
	V interface{} `json:"v,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ConeVisible `json:"visible,omitempty"`

	// Vsrc string Sets the source reference on Chart Studio Cloud for  v .
	Vsrc String `json:"vsrc,omitempty"`

	// W data_array Sets the z components of the vector field.
	W interface{} `json:"w,omitempty"`

	// Wsrc string Sets the source reference on Chart Studio Cloud for  w .
	Wsrc String `json:"wsrc,omitempty"`

	// X data_array Sets the x coordinates of the vector field and of the displayed cones.
	X interface{} `json:"x,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates of the vector field and of the displayed cones.
	Y interface{} `json:"y,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the z coordinates of the vector field and of the displayed cones.
	Z interface{} `json:"z,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewCone() *Cone {
	return &Cone{
		Type: TraceTypeCone,
	}
}

func (this *Cone) GetType() TraceType {
	return this.Type
}

// Contour The data from which contour lines are computed is set in `z`. Data in `z` must be a {2D array} of numbers. Say that `z` has N rows and M columns, then by default, these N rows correspond to N y coordinates (set in `y` or auto-generated) and the M columns correspond to M x coordinates (set in `x` or auto-generated). By setting `transpose` to *true*, the above behavior is flipped.
type Contour struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Autocontour boolean Determines whether or not the contour level attributes are picked by an algorithm. If *true*, the number of contour levels can be set in `ncontours`. If *false*, set the contour level attributes in `contours`.
	Autocontour Bool `json:"autocontour,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *ContourColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the `z` data are filled in. It is defaulted to true if `z` is a one dimensional array otherwise it is defaulted to false.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Contours <no value> <no value>
	Contours *ContourContours `json:"contours,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// Fillcolor color Sets the fill color if `contours.type` is *constraint*. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ContourHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ContourHoverlabel `json:"hoverlabel,omitempty"`

	// Hoverongaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the `z` data have hover labels associated with them.
	Hoverongaps Bool `json:"hoverongaps,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext data_array Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ContourLine `json:"line,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Ncontours integer Sets the maximum number of contour levels. The actual number of contours will be chosen automatically to be less than or equal to the value of `ncontours`. Has an effect only if `autocontour` is *true* or if `contours.size` is missing.
	Ncontours int64 `json:"ncontours,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *ContourStream `json:"stream,omitempty"`

	// Text data_array Sets the text elements associated with each z value.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ContourTransforms `json:"transforms,omitempty"`

	// Transpose boolean Transposes the z data.
	Transpose Bool `json:"transpose,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ContourVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar ContourXcalendar `json:"xcalendar,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment ContourXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Xtype enumerated If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
	Xtype ContourXtype `json:"xtype,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar ContourYcalendar `json:"ycalendar,omitempty"`

	// Yperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment ContourYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Ytype enumerated If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
	Ytype ContourYtype `json:"ytype,omitempty"`

	// Z data_array Sets the z data.
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zhoverformat string Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. See: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format
	Zhoverformat String `json:"zhoverformat,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewContour() *Contour {
	return &Contour{
		Type: TraceTypeContour,
	}
}

func (this *Contour) GetType() TraceType {
	return this.Type
}

// Contourcarpet Plots contours on either the first carpet axis or the carpet axis with a matching `carpet` attribute. Data `z` is interpreted as matching that of the corresponding carpet axis.
type Contourcarpet struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// A data_array Sets the x coordinates.
	A interface{} `json:"a,omitempty"`

	// A0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	A0 interface{} `json:"a0,omitempty"`

	// Asrc string Sets the source reference on Chart Studio Cloud for  a .
	Asrc String `json:"asrc,omitempty"`

	// Atype enumerated If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
	Atype ContourcarpetAtype `json:"atype,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Autocontour boolean Determines whether or not the contour level attributes are picked by an algorithm. If *true*, the number of contour levels can be set in `ncontours`. If *false*, set the contour level attributes in `contours`.
	Autocontour Bool `json:"autocontour,omitempty"`

	// B data_array Sets the y coordinates.
	B interface{} `json:"b,omitempty"`

	// B0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	B0 interface{} `json:"b0,omitempty"`

	// Bsrc string Sets the source reference on Chart Studio Cloud for  b .
	Bsrc String `json:"bsrc,omitempty"`

	// Btype enumerated If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
	Btype ContourcarpetBtype `json:"btype,omitempty"`

	// Carpet string The `carpet` of the carpet axes on which this contour trace lies
	Carpet String `json:"carpet,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *ContourcarpetColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Contours <no value> <no value>
	Contours *ContourcarpetContours `json:"contours,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Da number Sets the x coordinate step. See `x0` for more info.
	Da float64 `json:"da,omitempty"`

	// Db number Sets the y coordinate step. See `y0` for more info.
	Db float64 `json:"db,omitempty"`

	// Fillcolor color Sets the fill color if `contours.type` is *constraint*. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hovertext data_array Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ContourcarpetLine `json:"line,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Ncontours integer Sets the maximum number of contour levels. The actual number of contours will be chosen automatically to be less than or equal to the value of `ncontours`. Has an effect only if `autocontour` is *true* or if `contours.size` is missing.
	Ncontours int64 `json:"ncontours,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *ContourcarpetStream `json:"stream,omitempty"`

	// Text data_array Sets the text elements associated with each z value.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transpose boolean Transposes the z data.
	Transpose Bool `json:"transpose,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ContourcarpetVisible `json:"visible,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Z data_array Sets the z data.
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewContourcarpet() *Contourcarpet {
	return &Contourcarpet{
		Type: TraceTypeContourcarpet,
	}
}

func (this *Contourcarpet) GetType() TraceType {
	return this.Type
}

// Densitymapbox Draws a bivariate kernel density estimation with a Gaussian kernel from `lon` and `lat` coordinates and optional `z` values using a colorscale.
type Densitymapbox struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Below string Determines if the densitymapbox trace will be inserted before the layer with the specified ID. By default, densitymapbox traces are placed below the first layer of type symbol If set to '', the layer will be inserted above every existing layer.
	Below String `json:"below,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *DensitymapboxColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo DensitymapboxHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *DensitymapboxHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (lon,lat) pair If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (lon,lat) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Lat data_array Sets the latitude coordinates (in degrees North).
	Lat interface{} `json:"lat,omitempty"`

	// Latsrc string Sets the source reference on Chart Studio Cloud for  lat .
	Latsrc String `json:"latsrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Lon data_array Sets the longitude coordinates (in degrees East).
	Lon interface{} `json:"lon,omitempty"`

	// Lonsrc string Sets the source reference on Chart Studio Cloud for  lon .
	Lonsrc String `json:"lonsrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Radius number Sets the radius of influence of one `lon` / `lat` point in pixels. Increasing the value makes the densitymapbox trace smoother, but less detailed.
	Radius interface{} `json:"radius,omitempty"`

	// Radiussrc string Sets the source reference on Chart Studio Cloud for  radius .
	Radiussrc String `json:"radiussrc,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *DensitymapboxStream `json:"stream,omitempty"`

	// Subplot subplotid Sets a reference between this trace's data coordinates and a mapbox subplot. If *mapbox* (the default value), the data refer to `layout.mapbox`. If *mapbox2*, the data refer to `layout.mapbox2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Text string Sets text elements associated with each (lon,lat) pair If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (lon,lat) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms DensitymapboxTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible DensitymapboxVisible `json:"visible,omitempty"`

	// Z data_array Sets the points' weight. For example, a value of 10 would be equivalent to having 10 points of weight 1 in the same spot
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewDensitymapbox() *Densitymapbox {
	return &Densitymapbox{
		Type: TraceTypeDensitymapbox,
	}
}

func (this *Densitymapbox) GetType() TraceType {
	return this.Type
}

// Funnel Visualize stages in a process using length-encoded bars. This trace can be used to show data in either a part-to-whole representation wherein each item appears in a single stage, or in a "drop-off" representation wherein each item appears in each stage it traversed. See also the "funnelarea" trace type for a different approach to visualizing funnel data.
type Funnel struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Alignmentgroup string Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Cliponaxis boolean Determines whether the text nodes are clipped about the subplot axes. To show the text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*.
	Cliponaxis Bool `json:"cliponaxis,omitempty"`

	// Connector <no value> <no value>
	Connector *FunnelConnector `json:"connector,omitempty"`

	// Constraintext enumerated Constrain the size of text inside or outside a bar to be no larger than the bar itself.
	Constraintext FunnelConstraintext `json:"constraintext,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo FunnelHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *FunnelHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `percentInitial`, `percentPrevious` and `percentTotal`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Insidetextanchor enumerated Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
	Insidetextanchor FunnelInsidetextanchor `json:"insidetextanchor,omitempty"`

	// Insidetextfont <no value> Sets the font used for `text` lying inside the bar.
	Insidetextfont *FunnelInsidetextfont `json:"insidetextfont,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *FunnelMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Offset number Shifts the position where the bar is drawn (in position axis units). In *group* barmode, traces that set *offset* will be excluded and drawn in *overlay* mode instead.
	Offset float64 `json:"offset,omitempty"`

	// Offsetgroup string Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up.
	Offsetgroup String `json:"offsetgroup,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation enumerated Sets the orientation of the funnels. With *v* (*h*), the value of the each bar spans along the vertical (horizontal). By default funnels are tend to be oriented horizontally; unless only *y* array is presented or orientation is set to *v*. Also regarding graphs including only 'horizontal' funnels, *autorange* on the *y-axis* are set to *reversed*.
	Orientation FunnelOrientation `json:"orientation,omitempty"`

	// Outsidetextfont <no value> Sets the font used for `text` lying outside the bar.
	Outsidetextfont *FunnelOutsidetextfont `json:"outsidetextfont,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *FunnelStream `json:"stream,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textangle angle Sets the angle of the tick labels with respect to the bar. For example, a `tickangle` of -90 draws the tick labels vertically. With *auto* the texts may automatically be rotated to fit with the maximum size in bars.
	Textangle float64 `json:"textangle,omitempty"`

	// Textfont <no value> Sets the font used for `text`.
	Textfont *FunnelTextfont `json:"textfont,omitempty"`

	// Textinfo flaglist Determines which trace information appear on the graph. In the case of having multiple funnels, percentages & totals are computed separately (per trace).
	Textinfo FunnelTextinfo `json:"textinfo,omitempty"`

	// Textposition enumerated Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
	Textposition FunnelTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `percentInitial`, `percentPrevious`, `percentTotal`, `label` and `value`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms FunnelTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible FunnelVisible `json:"visible,omitempty"`

	// Width number Sets the bar width (in position axis units).
	Width float64 `json:"width,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment FunnelXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Yperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment FunnelYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewFunnel() *Funnel {
	return &Funnel{
		Type: TraceTypeFunnel,
	}
}

func (this *Funnel) GetType() TraceType {
	return this.Type
}

// Funnelarea Visualize stages in a process using area-encoded trapezoids. This trace can be used to show data in a part-to-whole representation similar to a "pie" trace, wherein each item appears in a single stage. See also the "funnel" trace type for a different approach to visualizing funnel data.
type Funnelarea struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Aspectratio number Sets the ratio between height and width
	Aspectratio float64 `json:"aspectratio,omitempty"`

	// Baseratio number Sets the ratio between bottom length and maximum top length.
	Baseratio float64 `json:"baseratio,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dlabel number Sets the label step. See `label0` for more info.
	Dlabel float64 `json:"dlabel,omitempty"`

	// Domain <no value> <no value>
	Domain *FunnelareaDomain `json:"domain,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo FunnelareaHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *FunnelareaHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `label`, `color`, `value`, `text` and `percent`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each sector. If a single string, the same string appears for all data points. If an array of string, the items are mapped in order of this trace's sectors. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Insidetextfont <no value> Sets the font used for `textinfo` lying inside the sector.
	Insidetextfont *FunnelareaInsidetextfont `json:"insidetextfont,omitempty"`

	// Label0 number Alternate to `labels`. Builds a numeric set of labels. Use with `dlabel` where `label0` is the starting label and `dlabel` the step.
	Label0 float64 `json:"label0,omitempty"`

	// Labels data_array Sets the sector labels. If `labels` entries are duplicated, we sum associated `values` or simply count occurrences if `values` is not provided. For other array attributes (including color) we use the first non-empty entry among all occurrences of the label.
	Labels interface{} `json:"labels,omitempty"`

	// Labelssrc string Sets the source reference on Chart Studio Cloud for  labels .
	Labelssrc String `json:"labelssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *FunnelareaMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Scalegroup string If there are multiple funnelareas that should be sized according to their totals, link them by providing a non-empty group id here shared by every trace in the same group.
	Scalegroup String `json:"scalegroup,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *FunnelareaStream `json:"stream,omitempty"`

	// Text data_array Sets text elements associated with each sector. If trace `textinfo` contains a *text* flag, these elements will be seen on the chart. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the font used for `textinfo`.
	Textfont *FunnelareaTextfont `json:"textfont,omitempty"`

	// Textinfo flaglist Determines which trace information appear on the graph.
	Textinfo FunnelareaTextinfo `json:"textinfo,omitempty"`

	// Textposition enumerated Specifies the location of the `textinfo`.
	Textposition FunnelareaTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `label`, `color`, `value`, `text` and `percent`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Title <no value> <no value>
	Title *FunnelareaTitle `json:"title,omitempty"`

	// Transforms <no value> <no value>
	Transforms FunnelareaTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Values data_array Sets the values of the sectors. If omitted, we count occurrences of each label.
	Values interface{} `json:"values,omitempty"`

	// Valuessrc string Sets the source reference on Chart Studio Cloud for  values .
	Valuessrc String `json:"valuessrc,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible FunnelareaVisible `json:"visible,omitempty"`
}

func NewFunnelarea() *Funnelarea {
	return &Funnelarea{
		Type: TraceTypeFunnelarea,
	}
}

func (this *Funnelarea) GetType() TraceType {
	return this.Type
}

// Heatmap The data that describes the heatmap value-to-color mapping is set in `z`. Data in `z` can either be a {2D array} of values (ragged or not) or a 1D array of values. In the case where `z` is a {2D array}, say that `z` has N rows and M columns. Then, by default, the resulting heatmap will have N partitions along the y axis and M partitions along the x axis. In other words, the i-th row/ j-th column cell in `z` is mapped to the i-th partition of the y axis (starting from the bottom of the plot) and the j-th partition of the x-axis (starting from the left of the plot). This behavior can be flipped by using `transpose`. Moreover, `x` (`y`) can be provided with M or M+1 (N or N+1) elements. If M (N), then the coordinates correspond to the center of the heatmap cells and the cells have equal width. If M+1 (N+1), then the coordinates correspond to the edges of the heatmap cells. In the case where `z` is a 1D {array}, the x and y coordinates must be provided in `x` and `y` respectively to form data triplets.
type Heatmap struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *HeatmapColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the `z` data are filled in. It is defaulted to true if `z` is a one dimensional array and `zsmooth` is not false; otherwise it is defaulted to false.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo HeatmapHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *HeatmapHoverlabel `json:"hoverlabel,omitempty"`

	// Hoverongaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the `z` data have hover labels associated with them.
	Hoverongaps Bool `json:"hoverongaps,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext data_array Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *HeatmapStream `json:"stream,omitempty"`

	// Text data_array Sets the text elements associated with each z value.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms HeatmapTransforms `json:"transforms,omitempty"`

	// Transpose boolean Transposes the z data.
	Transpose Bool `json:"transpose,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible HeatmapVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar HeatmapXcalendar `json:"xcalendar,omitempty"`

	// Xgap number Sets the horizontal gap (in pixels) between bricks.
	Xgap float64 `json:"xgap,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment HeatmapXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Xtype enumerated If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
	Xtype HeatmapXtype `json:"xtype,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar HeatmapYcalendar `json:"ycalendar,omitempty"`

	// Ygap number Sets the vertical gap (in pixels) between bricks.
	Ygap float64 `json:"ygap,omitempty"`

	// Yperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment HeatmapYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Ytype enumerated If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
	Ytype HeatmapYtype `json:"ytype,omitempty"`

	// Z data_array Sets the z data.
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zhoverformat string Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. See: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format
	Zhoverformat String `json:"zhoverformat,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsmooth enumerated Picks a smoothing algorithm use to smooth `z` data.
	Zsmooth HeatmapZsmooth `json:"zsmooth,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewHeatmap() *Heatmap {
	return &Heatmap{
		Type: TraceTypeHeatmap,
	}
}

func (this *Heatmap) GetType() TraceType {
	return this.Type
}

// Heatmapgl WebGL version of the heatmap trace type.
type Heatmapgl struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *HeatmapglColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo HeatmapglHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *HeatmapglHoverlabel `json:"hoverlabel,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *HeatmapglStream `json:"stream,omitempty"`

	// Text data_array Sets the text elements associated with each z value.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms HeatmapglTransforms `json:"transforms,omitempty"`

	// Transpose boolean Transposes the z data.
	Transpose Bool `json:"transpose,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible HeatmapglVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Xtype enumerated If *array*, the heatmap's x coordinates are given by *x* (the default behavior when `x` is provided). If *scaled*, the heatmap's x coordinates are given by *x0* and *dx* (the default behavior when `x` is not provided).
	Xtype HeatmapglXtype `json:"xtype,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Ytype enumerated If *array*, the heatmap's y coordinates are given by *y* (the default behavior when `y` is provided) If *scaled*, the heatmap's y coordinates are given by *y0* and *dy* (the default behavior when `y` is not provided)
	Ytype HeatmapglYtype `json:"ytype,omitempty"`

	// Z data_array Sets the z data.
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsmooth enumerated Picks a smoothing algorithm use to smooth `z` data.
	Zsmooth HeatmapglZsmooth `json:"zsmooth,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewHeatmapgl() *Heatmapgl {
	return &Heatmapgl{
		Type: TraceTypeHeatmapgl,
	}
}

func (this *Heatmapgl) GetType() TraceType {
	return this.Type
}

// Histogram The sample data from which statistics are computed is set in `x` for vertically spanning histograms and in `y` for horizontally spanning histograms. Binning options are set `xbins` and `ybins` respectively if no aggregation data is provided.
type Histogram struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// _deprecated <no value> <no value>
	// Pending... _deprecated No valTyp <no value> `json:"_deprecated,omitempty"`

	// Alignmentgroup string Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Autobinx boolean Obsolete: since v1.42 each bin attribute is auto-determined separately and `autobinx` is not needed. However, we accept `autobinx: true` or `false` and will update `xbins` accordingly before deleting `autobinx` from the trace.
	Autobinx Bool `json:"autobinx,omitempty"`

	// Autobiny boolean Obsolete: since v1.42 each bin attribute is auto-determined separately and `autobiny` is not needed. However, we accept `autobiny: true` or `false` and will update `ybins` accordingly before deleting `autobiny` from the trace.
	Autobiny Bool `json:"autobiny,omitempty"`

	// Bingroup string Set a group of histogram traces which will have compatible bin settings. Note that traces on the same subplot and with the same *orientation* under `barmode` *stack*, *relative* and *group* are forced into the same bingroup, Using `bingroup`, traces under `barmode` *overlay* and on different axes (of the same axis type) can have compatible bin settings. Note that histogram and histogram2d* trace can share the same `bingroup`
	Bingroup String `json:"bingroup,omitempty"`

	// Cumulative <no value> <no value>
	Cumulative *HistogramCumulative `json:"cumulative,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Error_x <no value> <no value>
	Error_x *HistogramErrorX `json:"error_x,omitempty"`

	// Error_y <no value> <no value>
	Error_y *HistogramErrorY `json:"error_y,omitempty"`

	// Histfunc enumerated Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
	Histfunc HistogramHistfunc `json:"histfunc,omitempty"`

	// Histnorm enumerated Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
	Histnorm HistogramHistnorm `json:"histnorm,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo HistogramHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *HistogramHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variable `binNumber` Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *HistogramMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Nbinsx integer Specifies the maximum number of desired bins. This value will be used in an algorithm that will decide the optimal bin size such that the histogram best visualizes the distribution of the data. Ignored if `xbins.size` is provided.
	Nbinsx int64 `json:"nbinsx,omitempty"`

	// Nbinsy integer Specifies the maximum number of desired bins. This value will be used in an algorithm that will decide the optimal bin size such that the histogram best visualizes the distribution of the data. Ignored if `ybins.size` is provided.
	Nbinsy int64 `json:"nbinsy,omitempty"`

	// Offsetgroup string Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up.
	Offsetgroup String `json:"offsetgroup,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation enumerated Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
	Orientation HistogramOrientation `json:"orientation,omitempty"`

	// Selected <no value> <no value>
	Selected *HistogramSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *HistogramStream `json:"stream,omitempty"`

	// Text string Sets hover text elements associated with each bar. If a single string, the same string appears over all bars. If an array of string, the items are mapped in order to the this trace's coordinates.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms HistogramTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *HistogramUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible HistogramVisible `json:"visible,omitempty"`

	// X data_array Sets the sample data to be binned on the x axis.
	X interface{} `json:"x,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xbins <no value> <no value>
	Xbins *HistogramXbins `json:"xbins,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar HistogramXcalendar `json:"xcalendar,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the sample data to be binned on the y axis.
	Y interface{} `json:"y,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ybins <no value> <no value>
	Ybins *HistogramYbins `json:"ybins,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar HistogramYcalendar `json:"ycalendar,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewHistogram() *Histogram {
	return &Histogram{
		Type: TraceTypeHistogram,
	}
}

func (this *Histogram) GetType() TraceType {
	return this.Type
}

// Histogram2d The sample data from which statistics are computed is set in `x` and `y` (where `x` and `y` represent marginal distributions, binning is set in `xbins` and `ybins` in this case) or `z` (where `z` represent the 2D distribution and binning set, binning is set by `x` and `y` in this case). The resulting distribution is visualized as a heatmap.
type Histogram2d struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autobinx boolean Obsolete: since v1.42 each bin attribute is auto-determined separately and `autobinx` is not needed. However, we accept `autobinx: true` or `false` and will update `xbins` accordingly before deleting `autobinx` from the trace.
	Autobinx Bool `json:"autobinx,omitempty"`

	// Autobiny boolean Obsolete: since v1.42 each bin attribute is auto-determined separately and `autobiny` is not needed. However, we accept `autobiny: true` or `false` and will update `ybins` accordingly before deleting `autobiny` from the trace.
	Autobiny Bool `json:"autobiny,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Bingroup string Set the `xbingroup` and `ybingroup` default prefix For example, setting a `bingroup` of *1* on two histogram2d traces will make them their x-bins and y-bins match separately.
	Bingroup String `json:"bingroup,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *Histogram2dColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Histfunc enumerated Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
	Histfunc Histogram2dHistfunc `json:"histfunc,omitempty"`

	// Histnorm enumerated Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
	Histnorm Histogram2dHistnorm `json:"histnorm,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo Histogram2dHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *Histogram2dHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variable `z` Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *Histogram2dMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Nbinsx integer Specifies the maximum number of desired bins. This value will be used in an algorithm that will decide the optimal bin size such that the histogram best visualizes the distribution of the data. Ignored if `xbins.size` is provided.
	Nbinsx int64 `json:"nbinsx,omitempty"`

	// Nbinsy integer Specifies the maximum number of desired bins. This value will be used in an algorithm that will decide the optimal bin size such that the histogram best visualizes the distribution of the data. Ignored if `ybins.size` is provided.
	Nbinsy int64 `json:"nbinsy,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *Histogram2dStream `json:"stream,omitempty"`

	// Transforms <no value> <no value>
	Transforms Histogram2dTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible Histogram2dVisible `json:"visible,omitempty"`

	// X data_array Sets the sample data to be binned on the x axis.
	X interface{} `json:"x,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xbingroup string Set a group of histogram traces which will have compatible x-bin settings. Using `xbingroup`, histogram2d and histogram2dcontour traces  (on axes of the same axis type) can have compatible x-bin settings. Note that the same `xbingroup` value can be used to set (1D) histogram `bingroup`
	Xbingroup String `json:"xbingroup,omitempty"`

	// Xbins <no value> <no value>
	Xbins *Histogram2dXbins `json:"xbins,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar Histogram2dXcalendar `json:"xcalendar,omitempty"`

	// Xgap number Sets the horizontal gap (in pixels) between bricks.
	Xgap float64 `json:"xgap,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the sample data to be binned on the y axis.
	Y interface{} `json:"y,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ybingroup string Set a group of histogram traces which will have compatible y-bin settings. Using `ybingroup`, histogram2d and histogram2dcontour traces  (on axes of the same axis type) can have compatible y-bin settings. Note that the same `ybingroup` value can be used to set (1D) histogram `bingroup`
	Ybingroup String `json:"ybingroup,omitempty"`

	// Ybins <no value> <no value>
	Ybins *Histogram2dYbins `json:"ybins,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar Histogram2dYcalendar `json:"ycalendar,omitempty"`

	// Ygap number Sets the vertical gap (in pixels) between bricks.
	Ygap float64 `json:"ygap,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the aggregation data.
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zhoverformat string Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. See: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format
	Zhoverformat String `json:"zhoverformat,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsmooth enumerated Picks a smoothing algorithm use to smooth `z` data.
	Zsmooth Histogram2dZsmooth `json:"zsmooth,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewHistogram2d() *Histogram2d {
	return &Histogram2d{
		Type: TraceTypeHistogram2d,
	}
}

func (this *Histogram2d) GetType() TraceType {
	return this.Type
}

// Histogram2dcontour The sample data from which statistics are computed is set in `x` and `y` (where `x` and `y` represent marginal distributions, binning is set in `xbins` and `ybins` in this case) or `z` (where `z` represent the 2D distribution and binning set, binning is set by `x` and `y` in this case). The resulting distribution is visualized as a contour plot.
type Histogram2dcontour struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autobinx boolean Obsolete: since v1.42 each bin attribute is auto-determined separately and `autobinx` is not needed. However, we accept `autobinx: true` or `false` and will update `xbins` accordingly before deleting `autobinx` from the trace.
	Autobinx Bool `json:"autobinx,omitempty"`

	// Autobiny boolean Obsolete: since v1.42 each bin attribute is auto-determined separately and `autobiny` is not needed. However, we accept `autobiny: true` or `false` and will update `ybins` accordingly before deleting `autobiny` from the trace.
	Autobiny Bool `json:"autobiny,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Autocontour boolean Determines whether or not the contour level attributes are picked by an algorithm. If *true*, the number of contour levels can be set in `ncontours`. If *false*, set the contour level attributes in `contours`.
	Autocontour Bool `json:"autocontour,omitempty"`

	// Bingroup string Set the `xbingroup` and `ybingroup` default prefix For example, setting a `bingroup` of *1* on two histogram2d traces will make them their x-bins and y-bins match separately.
	Bingroup String `json:"bingroup,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *Histogram2dcontourColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`zmin` and `zmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Contours <no value> <no value>
	Contours *Histogram2dcontourContours `json:"contours,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Histfunc enumerated Specifies the binning function used for this histogram trace. If *count*, the histogram values are computed by counting the number of values lying inside each bin. If *sum*, *avg*, *min*, *max*, the histogram values are computed using the sum, the average, the minimum or the maximum of the values lying inside each bin respectively.
	Histfunc Histogram2dcontourHistfunc `json:"histfunc,omitempty"`

	// Histnorm enumerated Specifies the type of normalization used for this histogram trace. If **, the span of each bar corresponds to the number of occurrences (i.e. the number of data points lying inside the bins). If *percent* / *probability*, the span of each bar corresponds to the percentage / fraction of occurrences with respect to the total number of sample points (here, the sum of all bin HEIGHTS equals 100% / 1). If *density*, the span of each bar corresponds to the number of occurrences in a bin divided by the size of the bin interval (here, the sum of all bin AREAS equals the total number of sample points). If *probability density*, the area of each bar corresponds to the probability that an event will fall into the corresponding bin (here, the sum of all bin AREAS equals 1).
	Histnorm Histogram2dcontourHistnorm `json:"histnorm,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo Histogram2dcontourHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *Histogram2dcontourHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variable `z` Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *Histogram2dcontourLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *Histogram2dcontourMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Nbinsx integer Specifies the maximum number of desired bins. This value will be used in an algorithm that will decide the optimal bin size such that the histogram best visualizes the distribution of the data. Ignored if `xbins.size` is provided.
	Nbinsx int64 `json:"nbinsx,omitempty"`

	// Nbinsy integer Specifies the maximum number of desired bins. This value will be used in an algorithm that will decide the optimal bin size such that the histogram best visualizes the distribution of the data. Ignored if `ybins.size` is provided.
	Nbinsy int64 `json:"nbinsy,omitempty"`

	// Ncontours integer Sets the maximum number of contour levels. The actual number of contours will be chosen automatically to be less than or equal to the value of `ncontours`. Has an effect only if `autocontour` is *true* or if `contours.size` is missing.
	Ncontours int64 `json:"ncontours,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `zmin` will correspond to the last color in the array and `zmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *Histogram2dcontourStream `json:"stream,omitempty"`

	// Transforms <no value> <no value>
	Transforms Histogram2dcontourTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible Histogram2dcontourVisible `json:"visible,omitempty"`

	// X data_array Sets the sample data to be binned on the x axis.
	X interface{} `json:"x,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xbingroup string Set a group of histogram traces which will have compatible x-bin settings. Using `xbingroup`, histogram2d and histogram2dcontour traces  (on axes of the same axis type) can have compatible x-bin settings. Note that the same `xbingroup` value can be used to set (1D) histogram `bingroup`
	Xbingroup String `json:"xbingroup,omitempty"`

	// Xbins <no value> <no value>
	Xbins *Histogram2dcontourXbins `json:"xbins,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar Histogram2dcontourXcalendar `json:"xcalendar,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the sample data to be binned on the y axis.
	Y interface{} `json:"y,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ybingroup string Set a group of histogram traces which will have compatible y-bin settings. Using `ybingroup`, histogram2d and histogram2dcontour traces  (on axes of the same axis type) can have compatible y-bin settings. Note that the same `ybingroup` value can be used to set (1D) histogram `bingroup`
	Ybingroup String `json:"ybingroup,omitempty"`

	// Ybins <no value> <no value>
	Ybins *Histogram2dcontourYbins `json:"ybins,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar Histogram2dcontourYcalendar `json:"ycalendar,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the aggregation data.
	Z interface{} `json:"z,omitempty"`

	// Zauto boolean Determines whether or not the color domain is computed with respect to the input data (here in `z`) or the bounds set in `zmin` and `zmax`  Defaults to `false` when `zmin` and `zmax` are set by the user.
	Zauto Bool `json:"zauto,omitempty"`

	// Zhoverformat string Sets the hover text formatting rule using d3 formatting mini-languages which are very similar to those in Python. See: https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format
	Zhoverformat String `json:"zhoverformat,omitempty"`

	// Zmax number Sets the upper bound of the color domain. Value should have the same units as in `z` and if set, `zmin` must be set as well.
	Zmax float64 `json:"zmax,omitempty"`

	// Zmid number Sets the mid-point of the color domain by scaling `zmin` and/or `zmax` to be equidistant to this point. Value should have the same units as in `z`. Has no effect when `zauto` is `false`.
	Zmid float64 `json:"zmid,omitempty"`

	// Zmin number Sets the lower bound of the color domain. Value should have the same units as in `z` and if set, `zmax` must be set as well.
	Zmin float64 `json:"zmin,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewHistogram2dcontour() *Histogram2dcontour {
	return &Histogram2dcontour{
		Type: TraceTypeHistogram2dcontour,
	}
}

func (this *Histogram2dcontour) GetType() TraceType {
	return this.Type
}

// Image Display an image, i.e. data on a 2D regular raster. By default, when an image is displayed in a subplot, its y axis will be reversed (ie. `autorange: 'reversed'`), constrained to the domain (ie. `constrain: 'domain'`) and it will have the same scale as its x axis (ie. `scaleanchor: 'x,`) in order for pixels to be rendered as squares.
type Image struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Colormodel enumerated Color model used to map the numerical color components described in `z` into colors. If `source` is specified, this attribute will be set to `rgba256` otherwise it defaults to `rgb`.
	Colormodel ImageColormodel `json:"colormodel,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Set the pixel's horizontal size.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Set the pixel's vertical size
	Dy float64 `json:"dy,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ImageHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ImageHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `z`, `color` and `colormodel`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext data_array Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Source string Specifies the data URI of the image to be visualized. The URI consists of "data:image/[<media subtype>][;base64],<data>"
	Source String `json:"source,omitempty"`

	// Stream <no value> <no value>
	Stream *ImageStream `json:"stream,omitempty"`

	// Text data_array Sets the text elements associated with each z value.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ImageVisible `json:"visible,omitempty"`

	// X0 any Set the image's x position.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Y0 any Set the image's y position.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Z data_array A 2-dimensional array in which each element is an array of 3 or 4 numbers representing a color.
	Z interface{} `json:"z,omitempty"`

	// Zmax info_array Array defining the higher bound for each color component. Note that the default value will depend on the colormodel. For the `rgb` colormodel, it is [255, 255, 255]. For the `rgba` colormodel, it is [255, 255, 255, 1]. For the `rgba256` colormodel, it is [255, 255, 255, 255]. For the `hsl` colormodel, it is [360, 100, 100]. For the `hsla` colormodel, it is [360, 100, 100, 1].
	Zmax interface{} `json:"zmax,omitempty"`

	// Zmin info_array Array defining the lower bound for each color component. Note that the default value will depend on the colormodel. For the `rgb` colormodel, it is [0, 0, 0]. For the `rgba` colormodel, it is [0, 0, 0, 0]. For the `rgba256` colormodel, it is [0, 0, 0, 0]. For the `hsl` colormodel, it is [0, 0, 0]. For the `hsla` colormodel, it is [0, 0, 0, 0].
	Zmin interface{} `json:"zmin,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewImage() *Image {
	return &Image{
		Type: TraceTypeImage,
	}
}

func (this *Image) GetType() TraceType {
	return this.Type
}

// Indicator An indicator is used to visualize a single `value` along with some contextual information such as `steps` or a `threshold`, using a combination of three visual elements: a number, a delta, and/or a gauge. Deltas are taken with respect to a `reference`. Gauges can be either angular or bullet (aka linear) gauges.
type Indicator struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Align enumerated Sets the horizontal alignment of the `text` within the box. Note that this attribute has no effect if an angular gauge is displayed: in this case, it is always centered
	Align IndicatorAlign `json:"align,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Delta <no value> <no value>
	Delta *IndicatorDelta `json:"delta,omitempty"`

	// Domain <no value> <no value>
	Domain *IndicatorDomain `json:"domain,omitempty"`

	// Gauge <no value> The gauge of the Indicator plot.
	Gauge *IndicatorGauge `json:"gauge,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines how the value is displayed on the graph. `number` displays the value numerically in text. `delta` displays the difference to a reference value in text. Finally, `gauge` displays the value graphically on an axis.
	Mode IndicatorMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Number <no value> <no value>
	Number *IndicatorNumber `json:"number,omitempty"`

	// Stream <no value> <no value>
	Stream *IndicatorStream `json:"stream,omitempty"`

	// Title <no value> <no value>
	Title *IndicatorTitle `json:"title,omitempty"`

	// Transforms <no value> <no value>
	Transforms IndicatorTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Value number Sets the number to be displayed.
	Value float64 `json:"value,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible IndicatorVisible `json:"visible,omitempty"`
}

func NewIndicator() *Indicator {
	return &Indicator{
		Type: TraceTypeIndicator,
	}
}

func (this *Indicator) GetType() TraceType {
	return this.Type
}

// Isosurface Draws isosurfaces between iso-min and iso-max values with coordinates given by four 1-dimensional arrays containing the `value`, `x`, `y` and `z` of every vertex of a uniform or non-uniform 3-D grid. Horizontal or vertical slices, caps as well as spaceframe between iso-min and iso-max values could also be drawn using this trace.
type Isosurface struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Caps <no value> <no value>
	Caps *IsosurfaceCaps `json:"caps,omitempty"`

	// Cauto boolean Determines whether or not the color domain is computed with respect to the input data (here `value`) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax number Sets the upper bound of the color domain. Value should have the same units as `value` and if set, `cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid number Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as `value`. Has no effect when `cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin number Sets the lower bound of the color domain. Value should have the same units as `value` and if set, `cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *IsosurfaceColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Contour <no value> <no value>
	Contour *IsosurfaceContour `json:"contour,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Flatshading boolean Determines whether or not normal smoothing is applied to the meshes, creating meshes with an angular, low-poly look via flat reflections.
	Flatshading Bool `json:"flatshading,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo IsosurfaceHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *IsosurfaceHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Isomax number Sets the maximum boundary for iso-surface plot.
	Isomax float64 `json:"isomax,omitempty"`

	// Isomin number Sets the minimum boundary for iso-surface plot.
	Isomin float64 `json:"isomin,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Lighting <no value> <no value>
	Lighting *IsosurfaceLighting `json:"lighting,omitempty"`

	// Lightposition <no value> <no value>
	Lightposition *IsosurfaceLightposition `json:"lightposition,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the surface. Please note that in the case of using high `opacity` values for example a value greater than or equal to 0.5 on two surfaces (and 0.25 with four surfaces), an overlay of multiple transparent surfaces may not perfectly be sorted in depth by the webgl API. This behavior may be improved in the near future and is subject to change.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Scene subplotid Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Slices <no value> <no value>
	Slices *IsosurfaceSlices `json:"slices,omitempty"`

	// Spaceframe <no value> <no value>
	Spaceframe *IsosurfaceSpaceframe `json:"spaceframe,omitempty"`

	// Stream <no value> <no value>
	Stream *IsosurfaceStream `json:"stream,omitempty"`

	// Surface <no value> <no value>
	Surface *IsosurfaceSurface `json:"surface,omitempty"`

	// Text string Sets the text elements associated with the vertices. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Value data_array Sets the 4th dimension (value) of the vertices.
	Value interface{} `json:"value,omitempty"`

	// Valuesrc string Sets the source reference on Chart Studio Cloud for  value .
	Valuesrc String `json:"valuesrc,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible IsosurfaceVisible `json:"visible,omitempty"`

	// X data_array Sets the X coordinates of the vertices on X axis.
	X interface{} `json:"x,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the Y coordinates of the vertices on Y axis.
	Y interface{} `json:"y,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the Z coordinates of the vertices on Z axis.
	Z interface{} `json:"z,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewIsosurface() *Isosurface {
	return &Isosurface{
		Type: TraceTypeIsosurface,
	}
}

func (this *Isosurface) GetType() TraceType {
	return this.Type
}

// Mesh3d Draws sets of triangles with coordinates given by three 1-dimensional arrays in `x`, `y`, `z` and (1) a sets of `i`, `j`, `k` indices (2) Delaunay triangulation or (3) the Alpha-shape algorithm or (4) the Convex-hull algorithm
type Mesh3d struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Alphahull number Determines how the mesh surface triangles are derived from the set of vertices (points) represented by the `x`, `y` and `z` arrays, if the `i`, `j`, `k` arrays are not supplied. For general use of `mesh3d` it is preferred that `i`, `j`, `k` are supplied. If *-1*, Delaunay triangulation is used, which is mainly suitable if the mesh is a single, more or less layer surface that is perpendicular to `delaunayaxis`. In case the `delaunayaxis` intersects the mesh surface at more than one point it will result triangles that are very long in the dimension of `delaunayaxis`. If *>0*, the alpha-shape algorithm is used. In this case, the positive `alphahull` value signals the use of the alpha-shape algorithm, _and_ its value acts as the parameter for the mesh fitting. If *0*,  the convex-hull algorithm is used. It is suitable for convex bodies or if the intention is to enclose the `x`, `y` and `z` point set into a convex hull.
	Alphahull float64 `json:"alphahull,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto boolean Determines whether or not the color domain is computed with respect to the input data (here `intensity`) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax number Sets the upper bound of the color domain. Value should have the same units as `intensity` and if set, `cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid number Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as `intensity`. Has no effect when `cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin number Sets the lower bound of the color domain. Value should have the same units as `intensity` and if set, `cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Color color Sets the color of the whole mesh
	Color String `json:"color,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *Mesh3dColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Contour <no value> <no value>
	Contour *Mesh3dContour `json:"contour,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Delaunayaxis enumerated Sets the Delaunay axis, which is the axis that is perpendicular to the surface of the Delaunay triangulation. It has an effect if `i`, `j`, `k` are not provided and `alphahull` is set to indicate Delaunay triangulation.
	Delaunayaxis Mesh3dDelaunayaxis `json:"delaunayaxis,omitempty"`

	// Facecolor data_array Sets the color of each face Overrides *color* and *vertexcolor*.
	Facecolor interface{} `json:"facecolor,omitempty"`

	// Facecolorsrc string Sets the source reference on Chart Studio Cloud for  facecolor .
	Facecolorsrc String `json:"facecolorsrc,omitempty"`

	// Flatshading boolean Determines whether or not normal smoothing is applied to the meshes, creating meshes with an angular, low-poly look via flat reflections.
	Flatshading Bool `json:"flatshading,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo Mesh3dHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *Mesh3dHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// I data_array A vector of vertex indices, i.e. integer values between 0 and the length of the vertex vectors, representing the *first* vertex of a triangle. For example, `{i[m], j[m], k[m]}` together represent face m (triangle m) in the mesh, where `i[m] = n` points to the triplet `{x[n], y[n], z[n]}` in the vertex arrays. Therefore, each element in `i` represents a point in space, which is the first vertex of a triangle.
	I interface{} `json:"i,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Intensity data_array Sets the intensity values for vertices or cells as defined by `intensitymode`. It can be used for plotting fields on meshes.
	Intensity interface{} `json:"intensity,omitempty"`

	// Intensitymode enumerated Determines the source of `intensity` values.
	Intensitymode Mesh3dIntensitymode `json:"intensitymode,omitempty"`

	// Intensitysrc string Sets the source reference on Chart Studio Cloud for  intensity .
	Intensitysrc String `json:"intensitysrc,omitempty"`

	// Isrc string Sets the source reference on Chart Studio Cloud for  i .
	Isrc String `json:"isrc,omitempty"`

	// J data_array A vector of vertex indices, i.e. integer values between 0 and the length of the vertex vectors, representing the *second* vertex of a triangle. For example, `{i[m], j[m], k[m]}`  together represent face m (triangle m) in the mesh, where `j[m] = n` points to the triplet `{x[n], y[n], z[n]}` in the vertex arrays. Therefore, each element in `j` represents a point in space, which is the second vertex of a triangle.
	J interface{} `json:"j,omitempty"`

	// Jsrc string Sets the source reference on Chart Studio Cloud for  j .
	Jsrc String `json:"jsrc,omitempty"`

	// K data_array A vector of vertex indices, i.e. integer values between 0 and the length of the vertex vectors, representing the *third* vertex of a triangle. For example, `{i[m], j[m], k[m]}` together represent face m (triangle m) in the mesh, where `k[m] = n` points to the triplet  `{x[n], y[n], z[n]}` in the vertex arrays. Therefore, each element in `k` represents a point in space, which is the third vertex of a triangle.
	K interface{} `json:"k,omitempty"`

	// Ksrc string Sets the source reference on Chart Studio Cloud for  k .
	Ksrc String `json:"ksrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Lighting <no value> <no value>
	Lighting *Mesh3dLighting `json:"lighting,omitempty"`

	// Lightposition <no value> <no value>
	Lightposition *Mesh3dLightposition `json:"lightposition,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the surface. Please note that in the case of using high `opacity` values for example a value greater than or equal to 0.5 on two surfaces (and 0.25 with four surfaces), an overlay of multiple transparent surfaces may not perfectly be sorted in depth by the webgl API. This behavior may be improved in the near future and is subject to change.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Scene subplotid Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *Mesh3dStream `json:"stream,omitempty"`

	// Text string Sets the text elements associated with the vertices. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Vertexcolor data_array Sets the color of each vertex Overrides *color*. While Red, green and blue colors are in the range of 0 and 255; in the case of having vertex color data in RGBA format, the alpha color should be normalized to be between 0 and 1.
	Vertexcolor interface{} `json:"vertexcolor,omitempty"`

	// Vertexcolorsrc string Sets the source reference on Chart Studio Cloud for  vertexcolor .
	Vertexcolorsrc String `json:"vertexcolorsrc,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible Mesh3dVisible `json:"visible,omitempty"`

	// X data_array Sets the X coordinates of the vertices. The nth element of vectors `x`, `y` and `z` jointly represent the X, Y and Z coordinates of the nth vertex.
	X interface{} `json:"x,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar Mesh3dXcalendar `json:"xcalendar,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the Y coordinates of the vertices. The nth element of vectors `x`, `y` and `z` jointly represent the X, Y and Z coordinates of the nth vertex.
	Y interface{} `json:"y,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar Mesh3dYcalendar `json:"ycalendar,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the Z coordinates of the vertices. The nth element of vectors `x`, `y` and `z` jointly represent the X, Y and Z coordinates of the nth vertex.
	Z interface{} `json:"z,omitempty"`

	// Zcalendar enumerated Sets the calendar system to use with `z` date data.
	Zcalendar Mesh3dZcalendar `json:"zcalendar,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewMesh3d() *Mesh3d {
	return &Mesh3d{
		Type: TraceTypeMesh3d,
	}
}

func (this *Mesh3d) GetType() TraceType {
	return this.Type
}

// Ohlc The ohlc (short for Open-High-Low-Close) is a style of financial chart describing open, high, low and close for a given `x` coordinate (most likely time). The tip of the lines represent the `low` and `high` values and the horizontal segments represent the `open` and `close` values. Sample points where the close value is higher (lower) then the open value are called increasing (decreasing). By default, increasing items are drawn in green whereas decreasing are drawn in red.
type Ohlc struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Close data_array Sets the close values.
	Close interface{} `json:"close,omitempty"`

	// Closesrc string Sets the source reference on Chart Studio Cloud for  close .
	Closesrc String `json:"closesrc,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Decreasing <no value> <no value>
	Decreasing *OhlcDecreasing `json:"decreasing,omitempty"`

	// High data_array Sets the high values.
	High interface{} `json:"high,omitempty"`

	// Highsrc string Sets the source reference on Chart Studio Cloud for  high .
	Highsrc String `json:"highsrc,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo OhlcHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *OhlcHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Increasing <no value> <no value>
	Increasing *OhlcIncreasing `json:"increasing,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *OhlcLine `json:"line,omitempty"`

	// Low data_array Sets the low values.
	Low interface{} `json:"low,omitempty"`

	// Lowsrc string Sets the source reference on Chart Studio Cloud for  low .
	Lowsrc String `json:"lowsrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Open data_array Sets the open values.
	Open interface{} `json:"open,omitempty"`

	// Opensrc string Sets the source reference on Chart Studio Cloud for  open .
	Opensrc String `json:"opensrc,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *OhlcStream `json:"stream,omitempty"`

	// Text string Sets hover text elements associated with each sample point. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to this trace's sample points.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Tickwidth number Sets the width of the open/close tick marks relative to the *x* minimal interval.
	Tickwidth float64 `json:"tickwidth,omitempty"`

	// Transforms <no value> <no value>
	Transforms OhlcTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible OhlcVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates. If absent, linear coordinate will be generated.
	X interface{} `json:"x,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar OhlcXcalendar `json:"xcalendar,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment OhlcXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`
}

func NewOhlc() *Ohlc {
	return &Ohlc{
		Type: TraceTypeOhlc,
	}
}

func (this *Ohlc) GetType() TraceType {
	return this.Type
}

// Parcats Parallel categories diagram for multidimensional categorical data.
type Parcats struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Arrangement enumerated Sets the drag interaction mode for categories and dimensions. If `perpendicular`, the categories can only move along a line perpendicular to the paths. If `freeform`, the categories can freely move on the plane. If `fixed`, the categories and dimensions are stationary.
	Arrangement ParcatsArrangement `json:"arrangement,omitempty"`

	// Bundlecolors boolean Sort paths so that like colors are bundled together within each category.
	Bundlecolors Bool `json:"bundlecolors,omitempty"`

	// Counts number The number of observations represented by each state. Defaults to 1 so that each state represents one observation
	Counts interface{} `json:"counts,omitempty"`

	// Countssrc string Sets the source reference on Chart Studio Cloud for  counts .
	Countssrc String `json:"countssrc,omitempty"`

	// Dimensions <no value> <no value>
	Dimensions ParcatsDimensions `json:"dimensions,omitempty"`

	// Domain <no value> <no value>
	Domain *ParcatsDomain `json:"domain,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ParcatsHoverinfo `json:"hoverinfo,omitempty"`

	// Hoveron enumerated Sets the hover interaction mode for the parcats diagram. If `category`, hover interaction take place per category. If `color`, hover interactions take place per color per category. If `dimension`, hover interactions take place across all categories per dimension.
	Hoveron ParcatsHoveron `json:"hoveron,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `count`, `probability`, `category`, `categorycount`, `colorcount` and `bandcolorcount`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate String `json:"hovertemplate,omitempty"`

	// Labelfont <no value> Sets the font for the `dimension` labels.
	Labelfont *ParcatsLabelfont `json:"labelfont,omitempty"`

	// Line <no value> <no value>
	Line *ParcatsLine `json:"line,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Sortpaths enumerated Sets the path sorting algorithm. If `forward`, sort paths based on dimension categories from left to right. If `backward`, sort paths based on dimensions categories from right to left.
	Sortpaths ParcatsSortpaths `json:"sortpaths,omitempty"`

	// Stream <no value> <no value>
	Stream *ParcatsStream `json:"stream,omitempty"`

	// Tickfont <no value> Sets the font for the `category` labels.
	Tickfont *ParcatsTickfont `json:"tickfont,omitempty"`

	// Transforms <no value> <no value>
	Transforms ParcatsTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ParcatsVisible `json:"visible,omitempty"`
}

func NewParcats() *Parcats {
	return &Parcats{
		Type: TraceTypeParcats,
	}
}

func (this *Parcats) GetType() TraceType {
	return this.Type
}

// Parcoords Parallel coordinates for multidimensional exploratory data analysis. The samples are specified in `dimensions`. The colors are set in `line.color`.
type Parcoords struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dimensions <no value> <no value>
	Dimensions ParcoordsDimensions `json:"dimensions,omitempty"`

	// Domain <no value> <no value>
	Domain *ParcoordsDomain `json:"domain,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Labelangle angle Sets the angle of the labels with respect to the horizontal. For example, a `tickangle` of -90 draws the labels vertically. Tilted labels with *labelangle* may be positioned better inside margins when `labelposition` is set to *bottom*.
	Labelangle float64 `json:"labelangle,omitempty"`

	// Labelfont <no value> Sets the font for the `dimension` labels.
	Labelfont *ParcoordsLabelfont `json:"labelfont,omitempty"`

	// Labelside enumerated Specifies the location of the `label`. *top* positions labels above, next to the title *bottom* positions labels below the graph Tilted labels with *labelangle* may be positioned better inside margins when `labelposition` is set to *bottom*.
	Labelside ParcoordsLabelside `json:"labelside,omitempty"`

	// Line <no value> <no value>
	Line *ParcoordsLine `json:"line,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Rangefont <no value> Sets the font for the `dimension` range values.
	Rangefont *ParcoordsRangefont `json:"rangefont,omitempty"`

	// Stream <no value> <no value>
	Stream *ParcoordsStream `json:"stream,omitempty"`

	// Tickfont <no value> Sets the font for the `dimension` tick values.
	Tickfont *ParcoordsTickfont `json:"tickfont,omitempty"`

	// Transforms <no value> <no value>
	Transforms ParcoordsTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ParcoordsVisible `json:"visible,omitempty"`
}

func NewParcoords() *Parcoords {
	return &Parcoords{
		Type: TraceTypeParcoords,
	}
}

func (this *Parcoords) GetType() TraceType {
	return this.Type
}

// Pie A data visualized by the sectors of the pie is set in `values`. The sector labels are set in `labels`. The sector colors are set in `marker.colors`
type Pie struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// _deprecated <no value> <no value>
	// Pending... _deprecated No valTyp <no value> `json:"_deprecated,omitempty"`

	// Automargin boolean Determines whether outside text labels can push the margins.
	Automargin Bool `json:"automargin,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Direction enumerated Specifies the direction at which succeeding sectors follow one another.
	Direction PieDirection `json:"direction,omitempty"`

	// Dlabel number Sets the label step. See `label0` for more info.
	Dlabel float64 `json:"dlabel,omitempty"`

	// Domain <no value> <no value>
	Domain *PieDomain `json:"domain,omitempty"`

	// Hole number Sets the fraction of the radius to cut out of the pie. Use this to make a donut chart.
	Hole float64 `json:"hole,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo PieHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *PieHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `label`, `color`, `value`, `percent` and `text`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each sector. If a single string, the same string appears for all data points. If an array of string, the items are mapped in order of this trace's sectors. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Insidetextfont <no value> Sets the font used for `textinfo` lying inside the sector.
	Insidetextfont *PieInsidetextfont `json:"insidetextfont,omitempty"`

	// Insidetextorientation enumerated Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
	Insidetextorientation PieInsidetextorientation `json:"insidetextorientation,omitempty"`

	// Label0 number Alternate to `labels`. Builds a numeric set of labels. Use with `dlabel` where `label0` is the starting label and `dlabel` the step.
	Label0 float64 `json:"label0,omitempty"`

	// Labels data_array Sets the sector labels. If `labels` entries are duplicated, we sum associated `values` or simply count occurrences if `values` is not provided. For other array attributes (including color) we use the first non-empty entry among all occurrences of the label.
	Labels interface{} `json:"labels,omitempty"`

	// Labelssrc string Sets the source reference on Chart Studio Cloud for  labels .
	Labelssrc String `json:"labelssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *PieMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Outsidetextfont <no value> Sets the font used for `textinfo` lying outside the sector.
	Outsidetextfont *PieOutsidetextfont `json:"outsidetextfont,omitempty"`

	// Pull number Sets the fraction of larger radius to pull the sectors out from the center. This can be a constant to pull all slices apart from each other equally or an array to highlight one or more slices.
	Pull interface{} `json:"pull,omitempty"`

	// Pullsrc string Sets the source reference on Chart Studio Cloud for  pull .
	Pullsrc String `json:"pullsrc,omitempty"`

	// Rotation number Instead of the first slice starting at 12 o'clock, rotate to some other angle.
	Rotation float64 `json:"rotation,omitempty"`

	// Scalegroup string If there are multiple pie charts that should be sized according to their totals, link them by providing a non-empty group id here shared by every trace in the same group.
	Scalegroup String `json:"scalegroup,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Sort boolean Determines whether or not the sectors are reordered from largest to smallest.
	Sort Bool `json:"sort,omitempty"`

	// Stream <no value> <no value>
	Stream *PieStream `json:"stream,omitempty"`

	// Text data_array Sets text elements associated with each sector. If trace `textinfo` contains a *text* flag, these elements will be seen on the chart. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the font used for `textinfo`.
	Textfont *PieTextfont `json:"textfont,omitempty"`

	// Textinfo flaglist Determines which trace information appear on the graph.
	Textinfo PieTextinfo `json:"textinfo,omitempty"`

	// Textposition enumerated Specifies the location of the `textinfo`.
	Textposition PieTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `label`, `color`, `value`, `percent` and `text`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Title <no value> <no value>
	Title *PieTitle `json:"title,omitempty"`

	// Transforms <no value> <no value>
	Transforms PieTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Values data_array Sets the values of the sectors. If omitted, we count occurrences of each label.
	Values interface{} `json:"values,omitempty"`

	// Valuessrc string Sets the source reference on Chart Studio Cloud for  values .
	Valuessrc String `json:"valuessrc,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible PieVisible `json:"visible,omitempty"`
}

func NewPie() *Pie {
	return &Pie{
		Type: TraceTypePie,
	}
}

func (this *Pie) GetType() TraceType {
	return this.Type
}

// Pointcloud The data visualized as a point cloud set in `x` and `y` using the WebGl plotting engine.
type Pointcloud struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo PointcloudHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *PointcloudHoverlabel `json:"hoverlabel,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Indices data_array A sequential value, 0..n, supply it to avoid creating this array inside plotting. If specified, it must be a typed `Int32Array` array. Its length must be equal to or greater than the number of points. For the best performance and memory use, create one large `indices` typed array that is guaranteed to be at least as long as the largest number of points during use, and reuse it on each `Plotly.restyle()` call.
	Indices interface{} `json:"indices,omitempty"`

	// Indicessrc string Sets the source reference on Chart Studio Cloud for  indices .
	Indicessrc String `json:"indicessrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *PointcloudMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *PointcloudStream `json:"stream,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible PointcloudVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xbounds data_array Specify `xbounds` in the shape of `[xMin, xMax] to avoid looping through the `xy` typed array. Use it in conjunction with `xy` and `ybounds` for the performance benefits.
	Xbounds interface{} `json:"xbounds,omitempty"`

	// Xboundssrc string Sets the source reference on Chart Studio Cloud for  xbounds .
	Xboundssrc String `json:"xboundssrc,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Xy data_array Faster alternative to specifying `x` and `y` separately. If supplied, it must be a typed `Float32Array` array that represents points such that `xy[i * 2] = x[i]` and `xy[i * 2 + 1] = y[i]`
	Xy interface{} `json:"xy,omitempty"`

	// Xysrc string Sets the source reference on Chart Studio Cloud for  xy .
	Xysrc String `json:"xysrc,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ybounds data_array Specify `ybounds` in the shape of `[yMin, yMax] to avoid looping through the `xy` typed array. Use it in conjunction with `xy` and `xbounds` for the performance benefits.
	Ybounds interface{} `json:"ybounds,omitempty"`

	// Yboundssrc string Sets the source reference on Chart Studio Cloud for  ybounds .
	Yboundssrc String `json:"yboundssrc,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewPointcloud() *Pointcloud {
	return &Pointcloud{
		Type: TraceTypePointcloud,
	}
}

func (this *Pointcloud) GetType() TraceType {
	return this.Type
}

// Sankey Sankey plots for network flow data analysis. The nodes are specified in `nodes` and the links between sources and targets in `links`. The colors are set in `nodes[i].color` and `links[i].color`, otherwise defaults are used.
type Sankey struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Arrangement enumerated If value is `snap` (the default), the node arrangement is assisted by automatic snapping of elements to preserve space between nodes specified via `nodepad`. If value is `perpendicular`, the nodes can only move along a line perpendicular to the flow. If value is `freeform`, the nodes can freely move on the plane. If value is `fixed`, the nodes are stationary.
	Arrangement SankeyArrangement `json:"arrangement,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Domain <no value> <no value>
	Domain *SankeyDomain `json:"domain,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired. Note that this attribute is superseded by `node.hoverinfo` and `node.hoverinfo` for nodes and links respectively.
	Hoverinfo SankeyHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *SankeyHoverlabel `json:"hoverlabel,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Link <no value> The links of the Sankey plot.
	Link *SankeyLink `json:"link,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Node <no value> The nodes of the Sankey plot.
	Node *SankeyNode `json:"node,omitempty"`

	// Orientation enumerated Sets the orientation of the Sankey diagram.
	Orientation SankeyOrientation `json:"orientation,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Stream <no value> <no value>
	Stream *SankeyStream `json:"stream,omitempty"`

	// Textfont <no value> Sets the font for node labels
	Textfont *SankeyTextfont `json:"textfont,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Valueformat string Sets the value formatting rule using d3 formatting mini-language which is similar to those of Python. See https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format
	Valueformat String `json:"valueformat,omitempty"`

	// Valuesuffix string Adds a unit to follow the value in the hover tooltip. Add a space if a separation is necessary from the value.
	Valuesuffix String `json:"valuesuffix,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible SankeyVisible `json:"visible,omitempty"`
}

func NewSankey() *Sankey {
	return &Sankey{
		Type: TraceTypeSankey,
	}
}

func (this *Sankey) GetType() TraceType {
	return this.Type
}

// Scatter The scatter trace type encompasses line charts, scatter charts, text charts, and bubble charts. The data visualized as scatter point or lines is set in `x` and `y`. Text (appearing either on the chart or on hover only) is via `text`. Bubble charts are achieved by setting `marker.size` and/or `marker.color` to numerical arrays.
type Scatter struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Cliponaxis boolean Determines whether or not markers and text nodes are clipped about the subplot axes. To show markers and text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*.
	Cliponaxis Bool `json:"cliponaxis,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// Error_x <no value> <no value>
	Error_x *ScatterErrorX `json:"error_x,omitempty"`

	// Error_y <no value> <no value>
	Error_y *ScatterErrorY `json:"error_y,omitempty"`

	// Fill enumerated Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
	Fill ScatterFill `json:"fill,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Groupnorm enumerated Only relevant when `stackgroup` is used, and only the first `groupnorm` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the normalization for the sum of this `stackgroup`. With *fraction*, the value of each trace at each location is divided by the sum of all trace values at that location. *percent* is the same but multiplied by 100 to show percentages. If there are multiple subplots, or multiple `stackgroup`s on one subplot, each will be normalized within its own set.
	Groupnorm ScatterGroupnorm `json:"groupnorm,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ScatterHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron flaglist Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
	Hoveron ScatterHoveron `json:"hoveron,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ScatterLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *ScatterMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode ScatterMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation enumerated Only relevant when `stackgroup` is used, and only the first `orientation` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Sets the stacking direction. With *v* (*h*), the y (x) values of subsequent traces are added. Also affects the default value of `fill`.
	Orientation ScatterOrientation `json:"orientation,omitempty"`

	// R data_array r coordinates in scatter traces are deprecated!Please switch to the *scatterpolar* trace type.Sets the radial coordinatesfor legacy polar chart only.
	R interface{} `json:"r,omitempty"`

	// Rsrc string Sets the source reference on Chart Studio Cloud for  r .
	Rsrc String `json:"rsrc,omitempty"`

	// Selected <no value> <no value>
	Selected *ScatterSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stackgaps enumerated Only relevant when `stackgroup` is used, and only the first `stackgaps` found in the `stackgroup` will be used - including if `visible` is *legendonly* but not if it is `false`. Determines how we handle locations at which other traces in this group have data but this one does not. With *infer zero* we insert a zero at these locations. With *interpolate* we linearly interpolate between existing values, and extrapolate a constant beyond the existing values.
	Stackgaps ScatterStackgaps `json:"stackgaps,omitempty"`

	// Stackgroup string Set several scatter traces (on the same subplot) to the same stackgroup in order to add their y values (or their x values if `orientation` is *h*). If blank or omitted this trace will not be stacked. Stacking also turns `fill` on by default, using *tonexty* (*tonextx*) if `orientation` is *h* (*v*) and sets the default `mode` to *lines* irrespective of point count. You can only stack on a numeric (linear or log) axis. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
	Stackgroup String `json:"stackgroup,omitempty"`

	// Stream <no value> <no value>
	Stream *ScatterStream `json:"stream,omitempty"`

	// T data_array t coordinates in scatter traces are deprecated!Please switch to the *scatterpolar* trace type.Sets the angular coordinatesfor legacy polar chart only.
	T interface{} `json:"t,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the text font.
	Textfont *ScatterTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ScatterTransforms `json:"transforms,omitempty"`

	// Tsrc string Sets the source reference on Chart Studio Cloud for  t .
	Tsrc String `json:"tsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ScatterUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar ScatterXcalendar `json:"xcalendar,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment ScatterXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar ScatterYcalendar `json:"ycalendar,omitempty"`

	// Yperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment ScatterYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewScatter() *Scatter {
	return &Scatter{
		Type: TraceTypeScatter,
	}
}

func (this *Scatter) GetType() TraceType {
	return this.Type
}

// Scatter3d The data visualized as scatter point or lines in 3D dimension is set in `x`, `y`, `z`. Text (appearing either on the chart or on hover only) is via `text`. Bubble charts are achieved by setting `marker.size` and/or `marker.color` Projections are achieved via `projection`. Surface fills are achieved via `surfaceaxis`.
type Scatter3d struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Error_x <no value> <no value>
	Error_x *Scatter3dErrorX `json:"error_x,omitempty"`

	// Error_y <no value> <no value>
	Error_y *Scatter3dErrorY `json:"error_y,omitempty"`

	// Error_z <no value> <no value>
	Error_z *Scatter3dErrorZ `json:"error_z,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo Scatter3dHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *Scatter3dHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets text elements associated with each (x,y,z) triplet. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y,z) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *Scatter3dLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *Scatter3dMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode Scatter3dMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Projection <no value> <no value>
	Projection *Scatter3dProjection `json:"projection,omitempty"`

	// Scene subplotid Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *Scatter3dStream `json:"stream,omitempty"`

	// Surfaceaxis enumerated If *-1*, the scatter points are not fill with a surface If *0*, *1*, *2*, the scatter points are filled with a Delaunay surface about the x, y, z respectively.
	Surfaceaxis Scatter3dSurfaceaxis `json:"surfaceaxis,omitempty"`

	// Surfacecolor color Sets the surface fill color.
	Surfacecolor String `json:"surfacecolor,omitempty"`

	// Text string Sets text elements associated with each (x,y,z) triplet. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y,z) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> <no value>
	Textfont *Scatter3dTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition Scatter3dTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms Scatter3dTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible Scatter3dVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar Scatter3dXcalendar `json:"xcalendar,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar Scatter3dYcalendar `json:"ycalendar,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the z coordinates.
	Z interface{} `json:"z,omitempty"`

	// Zcalendar enumerated Sets the calendar system to use with `z` date data.
	Zcalendar Scatter3dZcalendar `json:"zcalendar,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewScatter3d() *Scatter3d {
	return &Scatter3d{
		Type: TraceTypeScatter3d,
	}
}

func (this *Scatter3d) GetType() TraceType {
	return this.Type
}

// Scattercarpet Plots a scatter trace on either the first carpet axis or the carpet axis with a matching `carpet` attribute.
type Scattercarpet struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// A data_array Sets the a-axis coordinates.
	A interface{} `json:"a,omitempty"`

	// Asrc string Sets the source reference on Chart Studio Cloud for  a .
	Asrc String `json:"asrc,omitempty"`

	// B data_array Sets the b-axis coordinates.
	B interface{} `json:"b,omitempty"`

	// Bsrc string Sets the source reference on Chart Studio Cloud for  b .
	Bsrc String `json:"bsrc,omitempty"`

	// Carpet string An identifier for this carpet, so that `scattercarpet` and `contourcarpet` traces can specify a carpet plot on which they lie
	Carpet String `json:"carpet,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Fill enumerated Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
	Fill ScattercarpetFill `json:"fill,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScattercarpetHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ScattercarpetHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron flaglist Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
	Hoveron ScattercarpetHoveron `json:"hoveron,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (a,b) point. If a single string, the same string appears over all the data points. If an array of strings, the items are mapped in order to the the data points in (a,b). To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ScattercarpetLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *ScattercarpetMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode ScattercarpetMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Selected <no value> <no value>
	Selected *ScattercarpetSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *ScattercarpetStream `json:"stream,omitempty"`

	// Text string Sets text elements associated with each (a,b) point. If a single string, the same string appears over all the data points. If an array of strings, the items are mapped in order to the the data points in (a,b). If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the text font.
	Textfont *ScattercarpetTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScattercarpetTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `a`, `b` and `text`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ScattercarpetTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ScattercarpetUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScattercarpetVisible `json:"visible,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`
}

func NewScattercarpet() *Scattercarpet {
	return &Scattercarpet{
		Type: TraceTypeScattercarpet,
	}
}

func (this *Scattercarpet) GetType() TraceType {
	return this.Type
}

// Scattergeo The data visualized as scatter point or lines on a geographic map is provided either by longitude/latitude pairs in `lon` and `lat` respectively or by geographic location IDs or names in `locations`.
type Scattergeo struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Featureidkey string Sets the key in GeoJSON features which is used as id to match the items included in the `locations` array. Only has an effect when `geojson` is set. Support nested property, for example *properties.name*.
	Featureidkey String `json:"featureidkey,omitempty"`

	// Fill enumerated Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape.
	Fill ScattergeoFill `json:"fill,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Geo subplotid Sets a reference between this trace's geospatial coordinates and a geographic map. If *geo* (the default value), the geospatial coordinates refer to `layout.geo`. If *geo2*, the geospatial coordinates refer to `layout.geo2`, and so on.
	Geo String `json:"geo,omitempty"`

	// Geojson any Sets optional GeoJSON data associated with this trace. If not given, the features on the base map are used when `locations` is set. It can be set as a valid GeoJSON object or as a URL string. Note that we only accept GeoJSONs of type *FeatureCollection* or *Feature* with geometries of type *Polygon* or *MultiPolygon*.
	Geojson interface{} `json:"geojson,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScattergeoHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ScattergeoHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (lon,lat) pair or item in `locations`. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (lon,lat) or `locations` coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Lat data_array Sets the latitude coordinates (in degrees North).
	Lat interface{} `json:"lat,omitempty"`

	// Latsrc string Sets the source reference on Chart Studio Cloud for  lat .
	Latsrc String `json:"latsrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ScattergeoLine `json:"line,omitempty"`

	// Locationmode enumerated Determines the set of locations used to match entries in `locations` to regions on the map. Values *ISO-3*, *USA-states*, *country names* correspond to features on the base map and value *geojson-id* corresponds to features from a custom GeoJSON linked to the `geojson` attribute.
	Locationmode ScattergeoLocationmode `json:"locationmode,omitempty"`

	// Locations data_array Sets the coordinates via location IDs or names. Coordinates correspond to the centroid of each location given. See `locationmode` for more info.
	Locations interface{} `json:"locations,omitempty"`

	// Locationssrc string Sets the source reference on Chart Studio Cloud for  locations .
	Locationssrc String `json:"locationssrc,omitempty"`

	// Lon data_array Sets the longitude coordinates (in degrees East).
	Lon interface{} `json:"lon,omitempty"`

	// Lonsrc string Sets the source reference on Chart Studio Cloud for  lon .
	Lonsrc String `json:"lonsrc,omitempty"`

	// Marker <no value> <no value>
	Marker *ScattergeoMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode ScattergeoMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Selected <no value> <no value>
	Selected *ScattergeoSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *ScattergeoStream `json:"stream,omitempty"`

	// Text string Sets text elements associated with each (lon,lat) pair or item in `locations`. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (lon,lat) or `locations` coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the text font.
	Textfont *ScattergeoTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScattergeoTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `lat`, `lon`, `location` and `text`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ScattergeoTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ScattergeoUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScattergeoVisible `json:"visible,omitempty"`
}

func NewScattergeo() *Scattergeo {
	return &Scattergeo{
		Type: TraceTypeScattergeo,
	}
}

func (this *Scattergeo) GetType() TraceType {
	return this.Type
}

// Scattergl The data visualized as scatter point or lines is set in `x` and `y` using the WebGL plotting engine. Bubble charts are achieved by setting `marker.size` and/or `marker.color` to a numerical arrays.
type Scattergl struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dx number Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// Error_x <no value> <no value>
	Error_x *ScatterglErrorX `json:"error_x,omitempty"`

	// Error_y <no value> <no value>
	Error_y *ScatterglErrorY `json:"error_y,omitempty"`

	// Fill enumerated Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
	Fill ScatterglFill `json:"fill,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterglHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ScatterglHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ScatterglLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *ScatterglMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace.
	Mode ScatterglMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Selected <no value> <no value>
	Selected *ScatterglSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *ScatterglStream `json:"stream,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the text font.
	Textfont *ScatterglTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterglTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ScatterglTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ScatterglUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterglVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar ScatterglXcalendar `json:"xcalendar,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment ScatterglXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar ScatterglYcalendar `json:"ycalendar,omitempty"`

	// Yperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment ScatterglYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewScattergl() *Scattergl {
	return &Scattergl{
		Type: TraceTypeScattergl,
	}
}

func (this *Scattergl) GetType() TraceType {
	return this.Type
}

// Scattermapbox The data visualized as scatter point, lines or marker symbols on a Mapbox GL geographic map is provided by longitude/latitude pairs in `lon` and `lat`.
type Scattermapbox struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Below string Determines if this scattermapbox trace's layers are to be inserted before the layer with the specified ID. By default, scattermapbox layers are inserted above all the base layers. To place the scattermapbox layers above every other layer, set `below` to *''*.
	Below String `json:"below,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Fill enumerated Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape.
	Fill ScattermapboxFill `json:"fill,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScattermapboxHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ScattermapboxHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (lon,lat) pair If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (lon,lat) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Lat data_array Sets the latitude coordinates (in degrees North).
	Lat interface{} `json:"lat,omitempty"`

	// Latsrc string Sets the source reference on Chart Studio Cloud for  lat .
	Latsrc String `json:"latsrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ScattermapboxLine `json:"line,omitempty"`

	// Lon data_array Sets the longitude coordinates (in degrees East).
	Lon interface{} `json:"lon,omitempty"`

	// Lonsrc string Sets the source reference on Chart Studio Cloud for  lon .
	Lonsrc String `json:"lonsrc,omitempty"`

	// Marker <no value> <no value>
	Marker *ScattermapboxMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover.
	Mode ScattermapboxMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Selected <no value> <no value>
	Selected *ScattermapboxSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *ScattermapboxStream `json:"stream,omitempty"`

	// Subplot subplotid Sets a reference between this trace's data coordinates and a mapbox subplot. If *mapbox* (the default value), the data refer to `layout.mapbox`. If *mapbox2*, the data refer to `layout.mapbox2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Text string Sets text elements associated with each (lon,lat) pair If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (lon,lat) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the icon text font (color=mapbox.layer.paint.text-color, size=mapbox.layer.layout.text-size). Has an effect only when `type` is set to *symbol*.
	Textfont *ScattermapboxTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScattermapboxTextposition `json:"textposition,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `lat`, `lon` and `text`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ScattermapboxTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ScattermapboxUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScattermapboxVisible `json:"visible,omitempty"`
}

func NewScattermapbox() *Scattermapbox {
	return &Scattermapbox{
		Type: TraceTypeScattermapbox,
	}
}

func (this *Scattermapbox) GetType() TraceType {
	return this.Type
}

// Scatterpolar The scatterpolar trace type encompasses line charts, scatter charts, text charts, and bubble charts in polar coordinates. The data visualized as scatter point or lines is set in `r` (radial) and `theta` (angular) coordinates Text (appearing either on the chart or on hover only) is via `text`. Bubble charts are achieved by setting `marker.size` and/or `marker.color` to numerical arrays.
type Scatterpolar struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Cliponaxis boolean Determines whether or not markers and text nodes are clipped about the subplot axes. To show markers and text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*.
	Cliponaxis Bool `json:"cliponaxis,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dr number Sets the r coordinate step.
	Dr float64 `json:"dr,omitempty"`

	// Dtheta number Sets the theta coordinate step. By default, the `dtheta` step equals the subplot's period divided by the length of the `r` coordinates.
	Dtheta float64 `json:"dtheta,omitempty"`

	// Fill enumerated Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterpolar has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
	Fill ScatterpolarFill `json:"fill,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterpolarHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ScatterpolarHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron flaglist Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
	Hoveron ScatterpolarHoveron `json:"hoveron,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ScatterpolarLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *ScatterpolarMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode ScatterpolarMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// R data_array Sets the radial coordinates
	R interface{} `json:"r,omitempty"`

	// R0 any Alternate to `r`. Builds a linear space of r coordinates. Use with `dr` where `r0` is the starting coordinate and `dr` the step.
	R0 interface{} `json:"r0,omitempty"`

	// Rsrc string Sets the source reference on Chart Studio Cloud for  r .
	Rsrc String `json:"rsrc,omitempty"`

	// Selected <no value> <no value>
	Selected *ScatterpolarSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *ScatterpolarStream `json:"stream,omitempty"`

	// Subplot subplotid Sets a reference between this trace's data coordinates and a polar subplot. If *polar* (the default value), the data refer to `layout.polar`. If *polar2*, the data refer to `layout.polar2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the text font.
	Textfont *ScatterpolarTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterpolarTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `r`, `theta` and `text`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Theta data_array Sets the angular coordinates
	Theta interface{} `json:"theta,omitempty"`

	// Theta0 any Alternate to `theta`. Builds a linear space of theta coordinates. Use with `dtheta` where `theta0` is the starting coordinate and `dtheta` the step.
	Theta0 interface{} `json:"theta0,omitempty"`

	// Thetasrc string Sets the source reference on Chart Studio Cloud for  theta .
	Thetasrc String `json:"thetasrc,omitempty"`

	// Thetaunit enumerated Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
	Thetaunit ScatterpolarThetaunit `json:"thetaunit,omitempty"`

	// Transforms <no value> <no value>
	Transforms ScatterpolarTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ScatterpolarUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterpolarVisible `json:"visible,omitempty"`
}

func NewScatterpolar() *Scatterpolar {
	return &Scatterpolar{
		Type: TraceTypeScatterpolar,
	}
}

func (this *Scatterpolar) GetType() TraceType {
	return this.Type
}

// Scatterpolargl The scatterpolargl trace type encompasses line charts, scatter charts, and bubble charts in polar coordinates using the WebGL plotting engine. The data visualized as scatter point or lines is set in `r` (radial) and `theta` (angular) coordinates Bubble charts are achieved by setting `marker.size` and/or `marker.color` to numerical arrays.
type Scatterpolargl struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Dr number Sets the r coordinate step.
	Dr float64 `json:"dr,omitempty"`

	// Dtheta number Sets the theta coordinate step. By default, the `dtheta` step equals the subplot's period divided by the length of the `r` coordinates.
	Dtheta float64 `json:"dtheta,omitempty"`

	// Fill enumerated Sets the area to fill with a solid color. Defaults to *none* unless this trace is stacked, then it gets *tonexty* (*tonextx*) if `orientation` is *v* (*h*) Use with `fillcolor` if not *none*. *tozerox* and *tozeroy* fill to x=0 and y=0 respectively. *tonextx* and *tonexty* fill between the endpoints of this trace and the endpoints of the trace before it, connecting those endpoints with straight lines (to make a stacked area graph); if there is no trace before it, they behave like *tozerox* and *tozeroy*. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other. Traces in a `stackgroup` will only fill to (or be filled to) other traces in the same group. With multiple `stackgroup`s or some traces stacked and some not, if fill-linked traces are not already consecutive, the later ones will be pushed down in the drawing order.
	Fill ScatterpolarglFill `json:"fill,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterpolarglHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ScatterpolarglHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ScatterpolarglLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *ScatterpolarglMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode ScatterpolarglMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// R data_array Sets the radial coordinates
	R interface{} `json:"r,omitempty"`

	// R0 any Alternate to `r`. Builds a linear space of r coordinates. Use with `dr` where `r0` is the starting coordinate and `dr` the step.
	R0 interface{} `json:"r0,omitempty"`

	// Rsrc string Sets the source reference on Chart Studio Cloud for  r .
	Rsrc String `json:"rsrc,omitempty"`

	// Selected <no value> <no value>
	Selected *ScatterpolarglSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *ScatterpolarglStream `json:"stream,omitempty"`

	// Subplot subplotid Sets a reference between this trace's data coordinates and a polar subplot. If *polar* (the default value), the data refer to `layout.polar`. If *polar2*, the data refer to `layout.polar2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the text font.
	Textfont *ScatterpolarglTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterpolarglTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `r`, `theta` and `text`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Theta data_array Sets the angular coordinates
	Theta interface{} `json:"theta,omitempty"`

	// Theta0 any Alternate to `theta`. Builds a linear space of theta coordinates. Use with `dtheta` where `theta0` is the starting coordinate and `dtheta` the step.
	Theta0 interface{} `json:"theta0,omitempty"`

	// Thetasrc string Sets the source reference on Chart Studio Cloud for  theta .
	Thetasrc String `json:"thetasrc,omitempty"`

	// Thetaunit enumerated Sets the unit of input *theta* values. Has an effect only when on *linear* angular axes.
	Thetaunit ScatterpolarglThetaunit `json:"thetaunit,omitempty"`

	// Transforms <no value> <no value>
	Transforms ScatterpolarglTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ScatterpolarglUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterpolarglVisible `json:"visible,omitempty"`
}

func NewScatterpolargl() *Scatterpolargl {
	return &Scatterpolargl{
		Type: TraceTypeScatterpolargl,
	}
}

func (this *Scatterpolargl) GetType() TraceType {
	return this.Type
}

// Scatterternary Provides similar functionality to the *scatter* type but on a ternary phase diagram. The data is provided by at least two arrays out of `a`, `b`, `c` triplets.
type Scatterternary struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// A data_array Sets the quantity of component `a` in each data point. If `a`, `b`, and `c` are all provided, they need not be normalized, only the relative values matter. If only two arrays are provided they must be normalized to match `ternary<i>.sum`.
	A interface{} `json:"a,omitempty"`

	// Asrc string Sets the source reference on Chart Studio Cloud for  a .
	Asrc String `json:"asrc,omitempty"`

	// B data_array Sets the quantity of component `a` in each data point. If `a`, `b`, and `c` are all provided, they need not be normalized, only the relative values matter. If only two arrays are provided they must be normalized to match `ternary<i>.sum`.
	B interface{} `json:"b,omitempty"`

	// Bsrc string Sets the source reference on Chart Studio Cloud for  b .
	Bsrc String `json:"bsrc,omitempty"`

	// C data_array Sets the quantity of component `a` in each data point. If `a`, `b`, and `c` are all provided, they need not be normalized, only the relative values matter. If only two arrays are provided they must be normalized to match `ternary<i>.sum`.
	C interface{} `json:"c,omitempty"`

	// Cliponaxis boolean Determines whether or not markers and text nodes are clipped about the subplot axes. To show markers and text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*.
	Cliponaxis Bool `json:"cliponaxis,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the provided data arrays are connected.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Csrc string Sets the source reference on Chart Studio Cloud for  c .
	Csrc String `json:"csrc,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Fill enumerated Sets the area to fill with a solid color. Use with `fillcolor` if not *none*. scatterternary has a subset of the options available to scatter. *toself* connects the endpoints of the trace (or each segment of the trace if it has gaps) into a closed shape. *tonext* fills the space between two traces if one completely encloses the other (eg consecutive contour lines), and behaves like *toself* if there is no trace before it. *tonext* should not be used if one trace does not enclose the other.
	Fill ScatterternaryFill `json:"fill,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ScatterternaryHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ScatterternaryHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron flaglist Do the hover effects highlight individual points (markers or line points) or do they highlight filled regions? If the fill is *toself* or *tonext* and there are no markers or text, then the default is *fills*, otherwise it is *points*.
	Hoveron ScatterternaryHoveron `json:"hoveron,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (a,b,c) point. If a single string, the same string appears over all the data points. If an array of strings, the items are mapped in order to the the data points in (a,b,c). To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ScatterternaryLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *ScatterternaryMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Mode flaglist Determines the drawing mode for this scatter trace. If the provided `mode` includes *text* then the `text` elements appear at the coordinates. Otherwise, the `text` elements appear on hover. If there are less than 20 points and the trace is not stacked then the default is *lines+markers*. Otherwise, *lines*.
	Mode ScatterternaryMode `json:"mode,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Selected <no value> <no value>
	Selected *ScatterternarySelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *ScatterternaryStream `json:"stream,omitempty"`

	// Subplot subplotid Sets a reference between this trace's data coordinates and a ternary subplot. If *ternary* (the default value), the data refer to `layout.ternary`. If *ternary2*, the data refer to `layout.ternary2`, and so on.
	Subplot String `json:"subplot,omitempty"`

	// Sum number The number each triplet should sum to, if only two of `a`, `b`, and `c` are provided. This overrides `ternary<i>.sum` to normalize this specific trace, but does not affect the values displayed on the axes. 0 (or missing) means to use ternary<i>.sum
	Sum float64 `json:"sum,omitempty"`

	// Text string Sets text elements associated with each (a,b,c) point. If a single string, the same string appears over all the data points. If an array of strings, the items are mapped in order to the the data points in (a,b,c). If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the text font.
	Textfont *ScatterternaryTextfont `json:"textfont,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements with respects to the (x,y) coordinates.
	Textposition ScatterternaryTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `a`, `b`, `c` and `text`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ScatterternaryTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ScatterternaryUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ScatterternaryVisible `json:"visible,omitempty"`
}

func NewScatterternary() *Scatterternary {
	return &Scatterternary{
		Type: TraceTypeScatterternary,
	}
}

func (this *Scatterternary) GetType() TraceType {
	return this.Type
}

// Splom Splom traces generate scatter plot matrix visualizations. Each splom `dimensions` items correspond to a generated axis. Values for each of those dimensions are set in `dimensions[i].values`. Splom traces support all `scattergl` marker style attributes. Specify `layout.grid` attributes and/or layout x-axis and y-axis attributes for more control over the axis positioning and style.
type Splom struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Diagonal <no value> <no value>
	Diagonal *SplomDiagonal `json:"diagonal,omitempty"`

	// Dimensions <no value> <no value>
	Dimensions SplomDimensions `json:"dimensions,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo SplomHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *SplomHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Marker <no value> <no value>
	Marker *SplomMarker `json:"marker,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Selected <no value> <no value>
	Selected *SplomSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showlowerhalf boolean Determines whether or not subplots on the lower half from the diagonal are displayed.
	Showlowerhalf Bool `json:"showlowerhalf,omitempty"`

	// Showupperhalf boolean Determines whether or not subplots on the upper half from the diagonal are displayed.
	Showupperhalf Bool `json:"showupperhalf,omitempty"`

	// Stream <no value> <no value>
	Stream *SplomStream `json:"stream,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair to appear on hover. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms SplomTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *SplomUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible SplomVisible `json:"visible,omitempty"`

	// Xaxes info_array Sets the list of x axes corresponding to dimensions of this splom trace. By default, a splom will match the first N xaxes where N is the number of input dimensions. Note that, in case where `diagonal.visible` is false and `showupperhalf` or `showlowerhalf` is false, this splom trace will generate one less x-axis and one less y-axis.
	Xaxes interface{} `json:"xaxes,omitempty"`

	// Yaxes info_array Sets the list of y axes corresponding to dimensions of this splom trace. By default, a splom will match the first N yaxes where N is the number of input dimensions. Note that, in case where `diagonal.visible` is false and `showupperhalf` or `showlowerhalf` is false, this splom trace will generate one less x-axis and one less y-axis.
	Yaxes interface{} `json:"yaxes,omitempty"`
}

func NewSplom() *Splom {
	return &Splom{
		Type: TraceTypeSplom,
	}
}

func (this *Splom) GetType() TraceType {
	return this.Type
}

// Streamtube Use a streamtube trace to visualize flow in a vector field.  Specify a vector field using 6 1D arrays of equal length, 3 position arrays `x`, `y` and `z` and 3 vector component arrays `u`, `v`, and `w`.  By default, the tubes' starting positions will be cut from the vector field's x-z plane at its minimum y value. To specify your own starting position, use attributes `starts.x`, `starts.y` and `starts.z`. The color is encoded by the norm of (u, v, w), and the local radius by the divergence of (u, v, w).
type Streamtube struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto boolean Determines whether or not the color domain is computed with respect to the input data (here u/v/w norm) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax number Sets the upper bound of the color domain. Value should have the same units as u/v/w norm and if set, `cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid number Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as u/v/w norm. Has no effect when `cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin number Sets the lower bound of the color domain. Value should have the same units as u/v/w norm and if set, `cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *StreamtubeColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo StreamtubeHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *StreamtubeHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `tubex`, `tubey`, `tubez`, `tubeu`, `tubev`, `tubew`, `norm` and `divergence`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext String `json:"hovertext,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Lighting <no value> <no value>
	Lighting *StreamtubeLighting `json:"lighting,omitempty"`

	// Lightposition <no value> <no value>
	Lightposition *StreamtubeLightposition `json:"lightposition,omitempty"`

	// Maxdisplayed integer The maximum number of displayed segments in a streamtube.
	Maxdisplayed int64 `json:"maxdisplayed,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the surface. Please note that in the case of using high `opacity` values for example a value greater than or equal to 0.5 on two surfaces (and 0.25 with four surfaces), an overlay of multiple transparent surfaces may not perfectly be sorted in depth by the webgl API. This behavior may be improved in the near future and is subject to change.
	Opacity float64 `json:"opacity,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Scene subplotid Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Sizeref number The scaling factor for the streamtubes. The default is 1, which avoids two max divergence tubes from touching at adjacent starting positions.
	Sizeref float64 `json:"sizeref,omitempty"`

	// Starts <no value> <no value>
	Starts *StreamtubeStarts `json:"starts,omitempty"`

	// Stream <no value> <no value>
	Stream *StreamtubeStream `json:"stream,omitempty"`

	// Text string Sets a text element associated with this trace. If trace `hoverinfo` contains a *text* flag, this text element will be seen in all hover labels. Note that streamtube traces do not support array `text` values.
	Text String `json:"text,omitempty"`

	// U data_array Sets the x components of the vector field.
	U interface{} `json:"u,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Usrc string Sets the source reference on Chart Studio Cloud for  u .
	Usrc String `json:"usrc,omitempty"`

	// V data_array Sets the y components of the vector field.
	V interface{} `json:"v,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible StreamtubeVisible `json:"visible,omitempty"`

	// Vsrc string Sets the source reference on Chart Studio Cloud for  v .
	Vsrc String `json:"vsrc,omitempty"`

	// W data_array Sets the z components of the vector field.
	W interface{} `json:"w,omitempty"`

	// Wsrc string Sets the source reference on Chart Studio Cloud for  w .
	Wsrc String `json:"wsrc,omitempty"`

	// X data_array Sets the x coordinates of the vector field.
	X interface{} `json:"x,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates of the vector field.
	Y interface{} `json:"y,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the z coordinates of the vector field.
	Z interface{} `json:"z,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewStreamtube() *Streamtube {
	return &Streamtube{
		Type: TraceTypeStreamtube,
	}
}

func (this *Streamtube) GetType() TraceType {
	return this.Type
}

// Sunburst Visualize hierarchal data spanning outward radially from root to leaves. The sunburst sectors are determined by the entries in *labels* or *ids* and in *parents*.
type Sunburst struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Branchvalues enumerated Determines how the items in `values` are summed. When set to *total*, items in `values` are taken to be value of all its descendants. When set to *remainder*, items in `values` corresponding to the root and the branches sectors are taken to be the extra part not part of the sum of the values at their leaves.
	Branchvalues SunburstBranchvalues `json:"branchvalues,omitempty"`

	// Count flaglist Determines default for `values` when it is not provided, by inferring a 1 for each of the *leaves* and/or *branches*, otherwise 0.
	Count SunburstCount `json:"count,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Domain <no value> <no value>
	Domain *SunburstDomain `json:"domain,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo SunburstHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *SunburstHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `currentPath`, `root`, `entry`, `percentRoot`, `percentEntry` and `percentParent`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each sector. If a single string, the same string appears for all data points. If an array of string, the items are mapped in order of this trace's sectors. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Insidetextfont <no value> Sets the font used for `textinfo` lying inside the sector.
	Insidetextfont *SunburstInsidetextfont `json:"insidetextfont,omitempty"`

	// Insidetextorientation enumerated Controls the orientation of the text inside chart sectors. When set to *auto*, text may be oriented in any direction in order to be as big as possible in the middle of a sector. The *horizontal* option orients text to be parallel with the bottom of the chart, and may make text smaller in order to achieve that goal. The *radial* option orients text along the radius of the sector. The *tangential* option orients text perpendicular to the radius of the sector.
	Insidetextorientation SunburstInsidetextorientation `json:"insidetextorientation,omitempty"`

	// Labels data_array Sets the labels of each of the sectors.
	Labels interface{} `json:"labels,omitempty"`

	// Labelssrc string Sets the source reference on Chart Studio Cloud for  labels .
	Labelssrc String `json:"labelssrc,omitempty"`

	// Leaf <no value> <no value>
	Leaf *SunburstLeaf `json:"leaf,omitempty"`

	// Level any Sets the level from which this trace hierarchy is rendered. Set `level` to `''` to start from the root node in the hierarchy. Must be an "id" if `ids` is filled in, otherwise plotly attempts to find a matching item in `labels`.
	Level interface{} `json:"level,omitempty"`

	// Marker <no value> <no value>
	Marker *SunburstMarker `json:"marker,omitempty"`

	// Maxdepth integer Sets the number of rendered sectors from any given `level`. Set `maxdepth` to *-1* to render all the levels in the hierarchy.
	Maxdepth int64 `json:"maxdepth,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Outsidetextfont <no value> Sets the font used for `textinfo` lying outside the sector. This option refers to the root of the hierarchy presented at the center of a sunburst graph. Please note that if a hierarchy has multiple root nodes, this option won't have any effect and `insidetextfont` would be used.
	Outsidetextfont *SunburstOutsidetextfont `json:"outsidetextfont,omitempty"`

	// Parents data_array Sets the parent sectors for each of the sectors. Empty string items '' are understood to reference the root node in the hierarchy. If `ids` is filled, `parents` items are understood to be "ids" themselves. When `ids` is not set, plotly attempts to find matching items in `labels`, but beware they must be unique.
	Parents interface{} `json:"parents,omitempty"`

	// Parentssrc string Sets the source reference on Chart Studio Cloud for  parents .
	Parentssrc String `json:"parentssrc,omitempty"`

	// Root <no value> <no value>
	Root *SunburstRoot `json:"root,omitempty"`

	// Rotation angle Rotates the whole diagram counterclockwise by some angle. By default the first slice starts at 3 o'clock.
	Rotation float64 `json:"rotation,omitempty"`

	// Sort boolean Determines whether or not the sectors are reordered from largest to smallest.
	Sort Bool `json:"sort,omitempty"`

	// Stream <no value> <no value>
	Stream *SunburstStream `json:"stream,omitempty"`

	// Text data_array Sets text elements associated with each sector. If trace `textinfo` contains a *text* flag, these elements will be seen on the chart. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the font used for `textinfo`.
	Textfont *SunburstTextfont `json:"textfont,omitempty"`

	// Textinfo flaglist Determines which trace information appear on the graph.
	Textinfo SunburstTextinfo `json:"textinfo,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `currentPath`, `root`, `entry`, `percentRoot`, `percentEntry`, `percentParent`, `label` and `value`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms SunburstTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Values data_array Sets the values associated with each of the sectors. Use with `branchvalues` to determine how the values are summed.
	Values interface{} `json:"values,omitempty"`

	// Valuessrc string Sets the source reference on Chart Studio Cloud for  values .
	Valuessrc String `json:"valuessrc,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible SunburstVisible `json:"visible,omitempty"`
}

func NewSunburst() *Sunburst {
	return &Sunburst{
		Type: TraceTypeSunburst,
	}
}

func (this *Sunburst) GetType() TraceType {
	return this.Type
}

// Surface The data the describes the coordinates of the surface is set in `z`. Data in `z` should be a {2D array}. Coordinates in `x` and `y` can either be 1D {arrays} or {2D arrays} (e.g. to graph parametric surfaces). If not provided in `x` and `y`, the x and y coordinates are assumed to be linear starting at 0 with a unit step. The color scale corresponds to the `z` values by default. For custom color scales, use `surfacecolor` which should be a {2D array}, where its bounds can be controlled using `cmin` and `cmax`.
type Surface struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// _deprecated <no value> <no value>
	// Pending... _deprecated No valTyp <no value> `json:"_deprecated,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Cauto boolean Determines whether or not the color domain is computed with respect to the input data (here z or surfacecolor) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax number Sets the upper bound of the color domain. Value should have the same units as z or surfacecolor and if set, `cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid number Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as z or surfacecolor. Has no effect when `cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin number Sets the lower bound of the color domain. Value should have the same units as z or surfacecolor and if set, `cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *SurfaceColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Connectgaps boolean Determines whether or not gaps (i.e. {nan} or missing values) in the `z` data are filled in.
	Connectgaps Bool `json:"connectgaps,omitempty"`

	// Contours <no value> <no value>
	Contours *SurfaceContours `json:"contours,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Hidesurface boolean Determines whether or not a surface is drawn. For example, set `hidesurface` to *false* `contours.x.show` to *true* and `contours.y.show` to *true* to draw a wire frame plot.
	Hidesurface Bool `json:"hidesurface,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo SurfaceHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *SurfaceHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Lighting <no value> <no value>
	Lighting *SurfaceLighting `json:"lighting,omitempty"`

	// Lightposition <no value> <no value>
	Lightposition *SurfaceLightposition `json:"lightposition,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the surface. Please note that in the case of using high `opacity` values for example a value greater than or equal to 0.5 on two surfaces (and 0.25 with four surfaces), an overlay of multiple transparent surfaces may not perfectly be sorted in depth by the webgl API. This behavior may be improved in the near future and is subject to change.
	Opacity float64 `json:"opacity,omitempty"`

	// Opacityscale any Sets the opacityscale. The opacityscale must be an array containing arrays mapping a normalized value to an opacity value. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 1], [0.5, 0.2], [1, 1]]` means that higher/lower values would have higher opacity values and those in the middle would be more transparent Alternatively, `opacityscale` may be a palette name string of the following list: 'min', 'max', 'extremes' and 'uniform'. The default is 'uniform'.
	Opacityscale interface{} `json:"opacityscale,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Scene subplotid Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Stream <no value> <no value>
	Stream *SurfaceStream `json:"stream,omitempty"`

	// Surfacecolor data_array Sets the surface color values, used for setting a color scale independent of `z`.
	Surfacecolor interface{} `json:"surfacecolor,omitempty"`

	// Surfacecolorsrc string Sets the source reference on Chart Studio Cloud for  surfacecolor .
	Surfacecolorsrc String `json:"surfacecolorsrc,omitempty"`

	// Text string Sets the text elements associated with each z value. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible SurfaceVisible `json:"visible,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// Xcalendar enumerated Sets the calendar system to use with `x` date data.
	Xcalendar SurfaceXcalendar `json:"xcalendar,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Ycalendar enumerated Sets the calendar system to use with `y` date data.
	Ycalendar SurfaceYcalendar `json:"ycalendar,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the z coordinates.
	Z interface{} `json:"z,omitempty"`

	// Zcalendar enumerated Sets the calendar system to use with `z` date data.
	Zcalendar SurfaceZcalendar `json:"zcalendar,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewSurface() *Surface {
	return &Surface{
		Type: TraceTypeSurface,
	}
}

func (this *Surface) GetType() TraceType {
	return this.Type
}

// Table Table view for detailed data viewing. The data are arranged in a grid of rows and columns. Most styling can be specified for columns, rows or individual cells. Table is using a column-major order, ie. the grid is represented as a vector of column vectors.
type Table struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Cells <no value> <no value>
	Cells *TableCells `json:"cells,omitempty"`

	// Columnorder data_array Specifies the rendered order of the data columns; for example, a value `2` at position `0` means that column index `0` in the data will be rendered as the third column, as columns have an index base of zero.
	Columnorder interface{} `json:"columnorder,omitempty"`

	// Columnordersrc string Sets the source reference on Chart Studio Cloud for  columnorder .
	Columnordersrc String `json:"columnordersrc,omitempty"`

	// Columnwidth number The width of columns expressed as a ratio. Columns fill the available width in proportion of their specified column widths.
	Columnwidth interface{} `json:"columnwidth,omitempty"`

	// Columnwidthsrc string Sets the source reference on Chart Studio Cloud for  columnwidth .
	Columnwidthsrc String `json:"columnwidthsrc,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Domain <no value> <no value>
	Domain *TableDomain `json:"domain,omitempty"`

	// Header <no value> <no value>
	Header *TableHeader `json:"header,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo TableHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *TableHoverlabel `json:"hoverlabel,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Stream <no value> <no value>
	Stream *TableStream `json:"stream,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible TableVisible `json:"visible,omitempty"`
}

func NewTable() *Table {
	return &Table{
		Type: TraceTypeTable,
	}
}

func (this *Table) GetType() TraceType {
	return this.Type
}

// Treemap Visualize hierarchal data from leaves (and/or outer branches) towards root with rectangles. The treemap sectors are determined by the entries in *labels* or *ids* and in *parents*.
type Treemap struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Branchvalues enumerated Determines how the items in `values` are summed. When set to *total*, items in `values` are taken to be value of all its descendants. When set to *remainder*, items in `values` corresponding to the root and the branches sectors are taken to be the extra part not part of the sum of the values at their leaves.
	Branchvalues TreemapBranchvalues `json:"branchvalues,omitempty"`

	// Count flaglist Determines default for `values` when it is not provided, by inferring a 1 for each of the *leaves* and/or *branches*, otherwise 0.
	Count TreemapCount `json:"count,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Domain <no value> <no value>
	Domain *TreemapDomain `json:"domain,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo TreemapHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *TreemapHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `currentPath`, `root`, `entry`, `percentRoot`, `percentEntry` and `percentParent`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each sector. If a single string, the same string appears for all data points. If an array of string, the items are mapped in order of this trace's sectors. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Insidetextfont <no value> Sets the font used for `textinfo` lying inside the sector.
	Insidetextfont *TreemapInsidetextfont `json:"insidetextfont,omitempty"`

	// Labels data_array Sets the labels of each of the sectors.
	Labels interface{} `json:"labels,omitempty"`

	// Labelssrc string Sets the source reference on Chart Studio Cloud for  labels .
	Labelssrc String `json:"labelssrc,omitempty"`

	// Level any Sets the level from which this trace hierarchy is rendered. Set `level` to `''` to start from the root node in the hierarchy. Must be an "id" if `ids` is filled in, otherwise plotly attempts to find a matching item in `labels`.
	Level interface{} `json:"level,omitempty"`

	// Marker <no value> <no value>
	Marker *TreemapMarker `json:"marker,omitempty"`

	// Maxdepth integer Sets the number of rendered sectors from any given `level`. Set `maxdepth` to *-1* to render all the levels in the hierarchy.
	Maxdepth int64 `json:"maxdepth,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Outsidetextfont <no value> Sets the font used for `textinfo` lying outside the sector. This option refers to the root of the hierarchy presented on top left corner of a treemap graph. Please note that if a hierarchy has multiple root nodes, this option won't have any effect and `insidetextfont` would be used.
	Outsidetextfont *TreemapOutsidetextfont `json:"outsidetextfont,omitempty"`

	// Parents data_array Sets the parent sectors for each of the sectors. Empty string items '' are understood to reference the root node in the hierarchy. If `ids` is filled, `parents` items are understood to be "ids" themselves. When `ids` is not set, plotly attempts to find matching items in `labels`, but beware they must be unique.
	Parents interface{} `json:"parents,omitempty"`

	// Parentssrc string Sets the source reference on Chart Studio Cloud for  parents .
	Parentssrc String `json:"parentssrc,omitempty"`

	// Pathbar <no value> <no value>
	Pathbar *TreemapPathbar `json:"pathbar,omitempty"`

	// Root <no value> <no value>
	Root *TreemapRoot `json:"root,omitempty"`

	// Sort boolean Determines whether or not the sectors are reordered from largest to smallest.
	Sort Bool `json:"sort,omitempty"`

	// Stream <no value> <no value>
	Stream *TreemapStream `json:"stream,omitempty"`

	// Text data_array Sets text elements associated with each sector. If trace `textinfo` contains a *text* flag, these elements will be seen on the chart. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textfont <no value> Sets the font used for `textinfo`.
	Textfont *TreemapTextfont `json:"textfont,omitempty"`

	// Textinfo flaglist Determines which trace information appear on the graph.
	Textinfo TreemapTextinfo `json:"textinfo,omitempty"`

	// Textposition enumerated Sets the positions of the `text` elements.
	Textposition TreemapTextposition `json:"textposition,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `currentPath`, `root`, `entry`, `percentRoot`, `percentEntry`, `percentParent`, `label` and `value`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Tiling <no value> <no value>
	Tiling *TreemapTiling `json:"tiling,omitempty"`

	// Transforms <no value> <no value>
	Transforms TreemapTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Values data_array Sets the values associated with each of the sectors. Use with `branchvalues` to determine how the values are summed.
	Values interface{} `json:"values,omitempty"`

	// Valuessrc string Sets the source reference on Chart Studio Cloud for  values .
	Valuessrc String `json:"valuessrc,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible TreemapVisible `json:"visible,omitempty"`
}

func NewTreemap() *Treemap {
	return &Treemap{
		Type: TraceTypeTreemap,
	}
}

func (this *Treemap) GetType() TraceType {
	return this.Type
}

// Violin In vertical (horizontal) violin plots, statistics are computed using `y` (`x`) values. By supplying an `x` (`y`) array, one violin per distinct x (y) value is drawn If no `x` (`y`) {array} is provided, a single violin is drawn. That violin position is then positioned with with `name` or with `x0` (`y0`) if provided.
type Violin struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Alignmentgroup string Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Bandwidth number Sets the bandwidth used to compute the kernel density estimate. By default, the bandwidth is determined by Silverman's rule of thumb.
	Bandwidth float64 `json:"bandwidth,omitempty"`

	// Box <no value> <no value>
	Box *ViolinBox `json:"box,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Fillcolor color Sets the fill color. Defaults to a half-transparent variant of the line color, marker color, or marker line color, whichever is available.
	Fillcolor String `json:"fillcolor,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo ViolinHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *ViolinHoverlabel `json:"hoverlabel,omitempty"`

	// Hoveron flaglist Do the hover effects highlight individual violins or sample points or the kernel density estimate or any combination of them?
	Hoveron ViolinHoveron `json:"hoveron,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Jitter number Sets the amount of jitter in the sample points drawn. If *0*, the sample points align along the distribution axis. If *1*, the sample points are drawn in a random jitter of width equal to the width of the violins.
	Jitter float64 `json:"jitter,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Line <no value> <no value>
	Line *ViolinLine `json:"line,omitempty"`

	// Marker <no value> <no value>
	Marker *ViolinMarker `json:"marker,omitempty"`

	// Meanline <no value> <no value>
	Meanline *ViolinMeanline `json:"meanline,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover. For violin traces, the name will also be used for the position coordinate, if `x` and `x0` (`y` and `y0` if horizontal) are missing and the position axis is categorical. Note that the trace name is also used as a default value for attribute `scalegroup` (please see its description for details).
	Name String `json:"name,omitempty"`

	// Offsetgroup string Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up.
	Offsetgroup String `json:"offsetgroup,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation enumerated Sets the orientation of the violin(s). If *v* (*h*), the distribution is visualized along the vertical (horizontal).
	Orientation ViolinOrientation `json:"orientation,omitempty"`

	// Pointpos number Sets the position of the sample points in relation to the violins. If *0*, the sample points are places over the center of the violins. Positive (negative) values correspond to positions to the right (left) for vertical violins and above (below) for horizontal violins.
	Pointpos float64 `json:"pointpos,omitempty"`

	// Points enumerated If *outliers*, only the sample points lying outside the whiskers are shown If *suspectedoutliers*, the outlier points are shown and points either less than 4*Q1-3*Q3 or greater than 4*Q3-3*Q1 are highlighted (see `outliercolor`) If *all*, all sample points are shown If *false*, only the violins are shown with no sample points. Defaults to *suspectedoutliers* when `marker.outliercolor` or `marker.line.outliercolor` is set, otherwise defaults to *outliers*.
	Points ViolinPoints `json:"points,omitempty"`

	// Scalegroup string If there are multiple violins that should be sized according to to some metric (see `scalemode`), link them by providing a non-empty group id here shared by every trace in the same group. If a violin's `width` is undefined, `scalegroup` will default to the trace's name. In this case, violins with the same names will be linked together
	Scalegroup String `json:"scalegroup,omitempty"`

	// Scalemode enumerated Sets the metric by which the width of each violin is determined.*width* means each violin has the same (max) width*count* means the violins are scaled by the number of sample points makingup each violin.
	Scalemode ViolinScalemode `json:"scalemode,omitempty"`

	// Selected <no value> <no value>
	Selected *ViolinSelected `json:"selected,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Side enumerated Determines on which side of the position value the density function making up one half of a violin is plotted. Useful when comparing two violin traces under *overlay* mode, where one trace has `side` set to *positive* and the other to *negative*.
	Side ViolinSide `json:"side,omitempty"`

	// Span info_array Sets the span in data space for which the density function will be computed. Has an effect only when `spanmode` is set to *manual*.
	Span interface{} `json:"span,omitempty"`

	// Spanmode enumerated Sets the method by which the span in data space where the density function will be computed. *soft* means the span goes from the sample's minimum value minus two bandwidths to the sample's maximum value plus two bandwidths. *hard* means the span goes from the sample's minimum to its maximum value. For custom span settings, use mode *manual* and fill in the `span` attribute.
	Spanmode ViolinSpanmode `json:"spanmode,omitempty"`

	// Stream <no value> <no value>
	Stream *ViolinStream `json:"stream,omitempty"`

	// Text string Sets the text elements associated with each sample value. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Transforms <no value> <no value>
	Transforms ViolinTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Unselected <no value> <no value>
	Unselected *ViolinUnselected `json:"unselected,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible ViolinVisible `json:"visible,omitempty"`

	// Width number Sets the width of the violin in data coordinates. If *0* (default value) the width is automatically selected based on the positions of other violin traces in the same subplot.
	Width float64 `json:"width,omitempty"`

	// X data_array Sets the x sample data or coordinates. See overview for more info.
	X interface{} `json:"x,omitempty"`

	// X0 any Sets the x coordinate for single-box traces or the starting coordinate for multi-box traces set using q1/median/q3. See overview for more info.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y sample data or coordinates. See overview for more info.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Sets the y coordinate for single-box traces or the starting coordinate for multi-box traces set using q1/median/q3. See overview for more info.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewViolin() *Violin {
	return &Violin{
		Type: TraceTypeViolin,
	}
}

func (this *Violin) GetType() TraceType {
	return this.Type
}

// Volume Draws volume trace between iso-min and iso-max values with coordinates given by four 1-dimensional arrays containing the `value`, `x`, `y` and `z` of every vertex of a uniform or non-uniform 3-D grid. Horizontal or vertical slices, caps as well as spaceframe between iso-min and iso-max values could also be drawn using this trace.
type Volume struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Autocolorscale boolean Determines whether the colorscale is a default palette (`autocolorscale: true`) or the palette determined by `colorscale`. In case `colorscale` is unspecified or `autocolorscale` is true, the default  palette will be chosen according to whether numbers in the `color` array are all positive, all negative or mixed.
	Autocolorscale Bool `json:"autocolorscale,omitempty"`

	// Caps <no value> <no value>
	Caps *VolumeCaps `json:"caps,omitempty"`

	// Cauto boolean Determines whether or not the color domain is computed with respect to the input data (here `value`) or the bounds set in `cmin` and `cmax`  Defaults to `false` when `cmin` and `cmax` are set by the user.
	Cauto Bool `json:"cauto,omitempty"`

	// Cmax number Sets the upper bound of the color domain. Value should have the same units as `value` and if set, `cmin` must be set as well.
	Cmax float64 `json:"cmax,omitempty"`

	// Cmid number Sets the mid-point of the color domain by scaling `cmin` and/or `cmax` to be equidistant to this point. Value should have the same units as `value`. Has no effect when `cauto` is `false`.
	Cmid float64 `json:"cmid,omitempty"`

	// Cmin number Sets the lower bound of the color domain. Value should have the same units as `value` and if set, `cmax` must be set as well.
	Cmin float64 `json:"cmin,omitempty"`

	// Coloraxis subplotid Sets a reference to a shared color axis. References to these shared color axes are *coloraxis*, *coloraxis2*, *coloraxis3*, etc. Settings for these shared color axes are set in the layout, under `layout.coloraxis`, `layout.coloraxis2`, etc. Note that multiple color scales can be linked to the same color axis.
	Coloraxis String `json:"coloraxis,omitempty"`

	// Colorbar <no value> <no value>
	Colorbar *VolumeColorbar `json:"colorbar,omitempty"`

	// Colorscale colorscale Sets the colorscale. The colorscale must be an array containing arrays mapping a normalized value to an rgb, rgba, hex, hsl, hsv, or named color string. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 'rgb(0,0,255)'], [1, 'rgb(255,0,0)']]`. To control the bounds of the colorscale in color space, use`cmin` and `cmax`. Alternatively, `colorscale` may be a palette name string of the following list: Greys,YlGnBu,Greens,YlOrRd,Bluered,RdBu,Reds,Blues,Picnic,Rainbow,Portland,Jet,Hot,Blackbody,Earth,Electric,Viridis,Cividis.
	// Pending of type "colorscale"Colorscale  `json:"colorscale,omitempty"`

	// Contour <no value> <no value>
	Contour *VolumeContour `json:"contour,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Flatshading boolean Determines whether or not normal smoothing is applied to the meshes, creating meshes with an angular, low-poly look via flat reflections.
	Flatshading Bool `json:"flatshading,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo VolumeHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *VolumeHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available.  Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Same as `text`.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Isomax number Sets the maximum boundary for iso-surface plot.
	Isomax float64 `json:"isomax,omitempty"`

	// Isomin number Sets the minimum boundary for iso-surface plot.
	Isomin float64 `json:"isomin,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Lighting <no value> <no value>
	Lighting *VolumeLighting `json:"lighting,omitempty"`

	// Lightposition <no value> <no value>
	Lightposition *VolumeLightposition `json:"lightposition,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Opacity number Sets the opacity of the surface. Please note that in the case of using high `opacity` values for example a value greater than or equal to 0.5 on two surfaces (and 0.25 with four surfaces), an overlay of multiple transparent surfaces may not perfectly be sorted in depth by the webgl API. This behavior may be improved in the near future and is subject to change.
	Opacity float64 `json:"opacity,omitempty"`

	// Opacityscale any Sets the opacityscale. The opacityscale must be an array containing arrays mapping a normalized value to an opacity value. At minimum, a mapping for the lowest (0) and highest (1) values are required. For example, `[[0, 1], [0.5, 0.2], [1, 1]]` means that higher/lower values would have higher opacity values and those in the middle would be more transparent Alternatively, `opacityscale` may be a palette name string of the following list: 'min', 'max', 'extremes' and 'uniform'. The default is 'uniform'.
	Opacityscale interface{} `json:"opacityscale,omitempty"`

	// Reversescale boolean Reverses the color mapping if true. If true, `cmin` will correspond to the last color in the array and `cmax` will correspond to the first color.
	Reversescale Bool `json:"reversescale,omitempty"`

	// Scene subplotid Sets a reference between this trace's 3D coordinate system and a 3D scene. If *scene* (the default value), the (x,y,z) coordinates refer to `layout.scene`. If *scene2*, the (x,y,z) coordinates refer to `layout.scene2`, and so on.
	Scene String `json:"scene,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Showscale boolean Determines whether or not a colorbar is displayed for this trace.
	Showscale Bool `json:"showscale,omitempty"`

	// Slices <no value> <no value>
	Slices *VolumeSlices `json:"slices,omitempty"`

	// Spaceframe <no value> <no value>
	Spaceframe *VolumeSpaceframe `json:"spaceframe,omitempty"`

	// Stream <no value> <no value>
	Stream *VolumeStream `json:"stream,omitempty"`

	// Surface <no value> <no value>
	Surface *VolumeSurface `json:"surface,omitempty"`

	// Text string Sets the text elements associated with the vertices. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Value data_array Sets the 4th dimension (value) of the vertices.
	Value interface{} `json:"value,omitempty"`

	// Valuesrc string Sets the source reference on Chart Studio Cloud for  value .
	Valuesrc String `json:"valuesrc,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible VolumeVisible `json:"visible,omitempty"`

	// X data_array Sets the X coordinates of the vertices on X axis.
	X interface{} `json:"x,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the Y coordinates of the vertices on Y axis.
	Y interface{} `json:"y,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`

	// Z data_array Sets the Z coordinates of the vertices on Z axis.
	Z interface{} `json:"z,omitempty"`

	// Zsrc string Sets the source reference on Chart Studio Cloud for  z .
	Zsrc String `json:"zsrc,omitempty"`
}

func NewVolume() *Volume {
	return &Volume{
		Type: TraceTypeVolume,
	}
}

func (this *Volume) GetType() TraceType {
	return this.Type
}

// Waterfall Draws waterfall trace which is useful graph to displays the contribution of various elements (either positive or negative) in a bar chart. The data visualized by the span of the bars is set in `y` if `orientation` is set th *v* (the default) and the labels are set in `x`. By setting `orientation` to *h*, the roles are interchanged.
type Waterfall struct {

	// Type is the type of the plot
	Type TraceType `json:"type,omitempty"`

	// Alignmentgroup string Set several traces linked to the same position axis or matching axes to the same alignmentgroup. This controls whether bars compute their positional range dependently or independently.
	Alignmentgroup String `json:"alignmentgroup,omitempty"`

	// Base number Sets where the bar base is drawn (in position axis units).
	Base float64 `json:"base,omitempty"`

	// Cliponaxis boolean Determines whether the text nodes are clipped about the subplot axes. To show the text nodes above axis lines and tick labels, make sure to set `xaxis.layer` and `yaxis.layer` to *below traces*.
	Cliponaxis Bool `json:"cliponaxis,omitempty"`

	// Connector <no value> <no value>
	Connector *WaterfallConnector `json:"connector,omitempty"`

	// Constraintext enumerated Constrain the size of text inside or outside a bar to be no larger than the bar itself.
	Constraintext WaterfallConstraintext `json:"constraintext,omitempty"`

	// Customdata data_array Assigns extra data each datum. This may be useful when listening to hover, click and selection events. Note that, *scatter* traces also appends customdata items in the markers DOM elements
	Customdata interface{} `json:"customdata,omitempty"`

	// Customdatasrc string Sets the source reference on Chart Studio Cloud for  customdata .
	Customdatasrc String `json:"customdatasrc,omitempty"`

	// Decreasing <no value> <no value>
	Decreasing *WaterfallDecreasing `json:"decreasing,omitempty"`

	// Dx number Sets the x coordinate step. See `x0` for more info.
	Dx float64 `json:"dx,omitempty"`

	// Dy number Sets the y coordinate step. See `y0` for more info.
	Dy float64 `json:"dy,omitempty"`

	// Hoverinfo flaglist Determines which trace information appear on hover. If `none` or `skip` are set, no information is displayed upon hovering. But, if `none` is set, click and hover events are still fired.
	Hoverinfo WaterfallHoverinfo `json:"hoverinfo,omitempty"`

	// Hoverinfosrc string Sets the source reference on Chart Studio Cloud for  hoverinfo .
	Hoverinfosrc String `json:"hoverinfosrc,omitempty"`

	// Hoverlabel <no value> <no value>
	Hoverlabel *WaterfallHoverlabel `json:"hoverlabel,omitempty"`

	// Hovertemplate string Template string used for rendering the information that appear on hover box. Note that this will override `hoverinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. The variables available in `hovertemplate` are the ones emitted as event data described at this link https://plotly.com/javascript/plotlyjs-events/#event-data. Additionally, every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `initial`, `delta` and `final`. Anything contained in tag `<extra>` is displayed in the secondary box, for example "<extra>{fullData.name}</extra>". To hide the secondary box completely, use an empty tag `<extra></extra>`.
	Hovertemplate interface{} `json:"hovertemplate,omitempty"`

	// Hovertemplatesrc string Sets the source reference on Chart Studio Cloud for  hovertemplate .
	Hovertemplatesrc String `json:"hovertemplatesrc,omitempty"`

	// Hovertext string Sets hover text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. To be seen, trace `hoverinfo` must contain a *text* flag.
	Hovertext interface{} `json:"hovertext,omitempty"`

	// Hovertextsrc string Sets the source reference on Chart Studio Cloud for  hovertext .
	Hovertextsrc String `json:"hovertextsrc,omitempty"`

	// Ids data_array Assigns id labels to each datum. These ids for object constancy of data points during animation. Should be an array of strings, not numbers or any other type.
	Ids interface{} `json:"ids,omitempty"`

	// Idssrc string Sets the source reference on Chart Studio Cloud for  ids .
	Idssrc String `json:"idssrc,omitempty"`

	// Increasing <no value> <no value>
	Increasing *WaterfallIncreasing `json:"increasing,omitempty"`

	// Insidetextanchor enumerated Determines if texts are kept at center or start/end points in `textposition` *inside* mode.
	Insidetextanchor WaterfallInsidetextanchor `json:"insidetextanchor,omitempty"`

	// Insidetextfont <no value> Sets the font used for `text` lying inside the bar.
	Insidetextfont *WaterfallInsidetextfont `json:"insidetextfont,omitempty"`

	// Legendgroup string Sets the legend group for this trace. Traces part of the same legend group hide/show at the same time when toggling legend items.
	Legendgroup String `json:"legendgroup,omitempty"`

	// Measure data_array An array containing types of values. By default the values are considered as 'relative'. However; it is possible to use 'total' to compute the sums. Also 'absolute' could be applied to reset the computed total or to declare an initial value where needed.
	Measure interface{} `json:"measure,omitempty"`

	// Measuresrc string Sets the source reference on Chart Studio Cloud for  measure .
	Measuresrc String `json:"measuresrc,omitempty"`

	// Meta any Assigns extra meta information associated with this trace that can be used in various text attributes. Attributes such as trace `name`, graph, axis and colorbar `title.text`, annotation `text` `rangeselector`, `updatemenues` and `sliders` `label` text all support `meta`. To access the trace `meta` values in an attribute in the same trace, simply use `%{meta[i]}` where `i` is the index or key of the `meta` item in question. To access trace `meta` in layout attributes, use `%{data[n[.meta[i]}` where `i` is the index or key of the `meta` and `n` is the trace index.
	Meta interface{} `json:"meta,omitempty"`

	// Metasrc string Sets the source reference on Chart Studio Cloud for  meta .
	Metasrc String `json:"metasrc,omitempty"`

	// Name string Sets the trace name. The trace name appear as the legend item and on hover.
	Name String `json:"name,omitempty"`

	// Offset number Shifts the position where the bar is drawn (in position axis units). In *group* barmode, traces that set *offset* will be excluded and drawn in *overlay* mode instead.
	Offset interface{} `json:"offset,omitempty"`

	// Offsetgroup string Set several traces linked to the same position axis or matching axes to the same offsetgroup where bars of the same position coordinate will line up.
	Offsetgroup String `json:"offsetgroup,omitempty"`

	// Offsetsrc string Sets the source reference on Chart Studio Cloud for  offset .
	Offsetsrc String `json:"offsetsrc,omitempty"`

	// Opacity number Sets the opacity of the trace.
	Opacity float64 `json:"opacity,omitempty"`

	// Orientation enumerated Sets the orientation of the bars. With *v* (*h*), the value of the each bar spans along the vertical (horizontal).
	Orientation WaterfallOrientation `json:"orientation,omitempty"`

	// Outsidetextfont <no value> Sets the font used for `text` lying outside the bar.
	Outsidetextfont *WaterfallOutsidetextfont `json:"outsidetextfont,omitempty"`

	// Selectedpoints any Array containing integer indices of selected points. Has an effect only for traces that support selections. Note that an empty array means an empty selection where the `unselected` are turned on for all points, whereas, any other non-array values means no selection all where the `selected` and `unselected` styles have no effect.
	Selectedpoints interface{} `json:"selectedpoints,omitempty"`

	// Showlegend boolean Determines whether or not an item corresponding to this trace is shown in the legend.
	Showlegend Bool `json:"showlegend,omitempty"`

	// Stream <no value> <no value>
	Stream *WaterfallStream `json:"stream,omitempty"`

	// Text string Sets text elements associated with each (x,y) pair. If a single string, the same string appears over all the data points. If an array of string, the items are mapped in order to the this trace's (x,y) coordinates. If trace `hoverinfo` contains a *text* flag and *hovertext* is not set, these elements will be seen in the hover labels.
	Text interface{} `json:"text,omitempty"`

	// Textangle angle Sets the angle of the tick labels with respect to the bar. For example, a `tickangle` of -90 draws the tick labels vertically. With *auto* the texts may automatically be rotated to fit with the maximum size in bars.
	Textangle float64 `json:"textangle,omitempty"`

	// Textfont <no value> Sets the font used for `text`.
	Textfont *WaterfallTextfont `json:"textfont,omitempty"`

	// Textinfo flaglist Determines which trace information appear on the graph. In the case of having multiple waterfalls, totals are computed separately (per trace).
	Textinfo WaterfallTextinfo `json:"textinfo,omitempty"`

	// Textposition enumerated Specifies the location of the `text`. *inside* positions `text` inside, next to the bar end (rotated and scaled if needed). *outside* positions `text` outside, next to the bar end (scaled if needed), unless there is another bar stacked on this one, then the text gets pushed inside. *auto* tries to position `text` inside the bar, but if the bar is too small and no bar is stacked on this one the text is moved outside.
	Textposition WaterfallTextposition `json:"textposition,omitempty"`

	// Textpositionsrc string Sets the source reference on Chart Studio Cloud for  textposition .
	Textpositionsrc String `json:"textpositionsrc,omitempty"`

	// Textsrc string Sets the source reference on Chart Studio Cloud for  text .
	Textsrc String `json:"textsrc,omitempty"`

	// Texttemplate string Template string used for rendering the information text that appear on points. Note that this will override `textinfo`. Variables are inserted using %{variable}, for example "y: %{y}". Numbers are formatted using d3-format's syntax %{variable:d3-format}, for example "Price: %{y:$.2f}". https://github.com/d3/d3-3.x-api-reference/blob/master/Formatting.md#d3_format for details on the formatting syntax. Dates are formatted using d3-time-format's syntax %{variable|d3-time-format}, for example "Day: %{2019-01-01|%A}". https://github.com/d3/d3-time-format#locale_format for details on the date formatting syntax. Every attributes that can be specified per-point (the ones that are `arrayOk: true`) are available. variables `initial`, `delta`, `final` and `label`.
	Texttemplate interface{} `json:"texttemplate,omitempty"`

	// Texttemplatesrc string Sets the source reference on Chart Studio Cloud for  texttemplate .
	Texttemplatesrc String `json:"texttemplatesrc,omitempty"`

	// Totals <no value> <no value>
	Totals *WaterfallTotals `json:"totals,omitempty"`

	// Transforms <no value> <no value>
	Transforms WaterfallTransforms `json:"transforms,omitempty"`

	// Uid string Assign an id to this trace, Use this to provide object constancy between traces during animations and transitions.
	Uid String `json:"uid,omitempty"`

	// Uirevision any Controls persistence of some user-driven changes to the trace: `constraintrange` in `parcoords` traces, as well as some `editable: true` modifications such as `name` and `colorbar.title`. Defaults to `layout.uirevision`. Note that other user-driven trace attribute changes are controlled by `layout` attributes: `trace.visible` is controlled by `layout.legend.uirevision`, `selectedpoints` is controlled by `layout.selectionrevision`, and `colorbar.(x|y)` (accessible with `config: {editable: true}`) is controlled by `layout.editrevision`. Trace changes are tracked by `uid`, which only falls back on trace index if no `uid` is provided. So if your app can add/remove traces before the end of the `data` array, such that the same trace has a different index, you can still preserve user-driven changes if you give each trace a `uid` that stays with it as it moves.
	Uirevision interface{} `json:"uirevision,omitempty"`

	// Visible enumerated Determines whether or not this trace is visible. If *legendonly*, the trace is not drawn, but can appear as a legend item (provided that the legend itself is visible).
	Visible WaterfallVisible `json:"visible,omitempty"`

	// Width number Sets the bar width (in position axis units).
	Width interface{} `json:"width,omitempty"`

	// Widthsrc string Sets the source reference on Chart Studio Cloud for  width .
	Widthsrc String `json:"widthsrc,omitempty"`

	// X data_array Sets the x coordinates.
	X interface{} `json:"x,omitempty"`

	// X0 any Alternate to `x`. Builds a linear space of x coordinates. Use with `dx` where `x0` is the starting coordinate and `dx` the step.
	X0 interface{} `json:"x0,omitempty"`

	// Xaxis subplotid Sets a reference between this trace's x coordinates and a 2D cartesian x axis. If *x* (the default value), the x coordinates refer to `layout.xaxis`. If *x2*, the x coordinates refer to `layout.xaxis2`, and so on.
	Xaxis String `json:"xaxis,omitempty"`

	// Xperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the x axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Xperiod interface{} `json:"xperiod,omitempty"`

	// Xperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the x0 axis. When `x0period` is round number of weeks, the `x0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Xperiod0 interface{} `json:"xperiod0,omitempty"`

	// Xperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the x axis.
	Xperiodalignment WaterfallXperiodalignment `json:"xperiodalignment,omitempty"`

	// Xsrc string Sets the source reference on Chart Studio Cloud for  x .
	Xsrc String `json:"xsrc,omitempty"`

	// Y data_array Sets the y coordinates.
	Y interface{} `json:"y,omitempty"`

	// Y0 any Alternate to `y`. Builds a linear space of y coordinates. Use with `dy` where `y0` is the starting coordinate and `dy` the step.
	Y0 interface{} `json:"y0,omitempty"`

	// Yaxis subplotid Sets a reference between this trace's y coordinates and a 2D cartesian y axis. If *y* (the default value), the y coordinates refer to `layout.yaxis`. If *y2*, the y coordinates refer to `layout.yaxis2`, and so on.
	Yaxis String `json:"yaxis,omitempty"`

	// Yperiod any Only relevant when the axis `type` is *date*. Sets the period positioning in milliseconds or *M<n>* on the y axis. Special values in the form of *M<n>* could be used to declare the number of months. In this case `n` must be a positive integer.
	Yperiod interface{} `json:"yperiod,omitempty"`

	// Yperiod0 any Only relevant when the axis `type` is *date*. Sets the base for period positioning in milliseconds or date string on the y0 axis. When `y0period` is round number of weeks, the `y0period0` by default would be on a Sunday i.e. 2000-01-02, otherwise it would be at 2000-01-01.
	Yperiod0 interface{} `json:"yperiod0,omitempty"`

	// Yperiodalignment enumerated Only relevant when the axis `type` is *date*. Sets the alignment of data points on the y axis.
	Yperiodalignment WaterfallYperiodalignment `json:"yperiodalignment,omitempty"`

	// Ysrc string Sets the source reference on Chart Studio Cloud for  y .
	Ysrc String `json:"ysrc,omitempty"`
}

func NewWaterfall() *Waterfall {
	return &Waterfall{
		Type: TraceTypeWaterfall,
	}
}

func (this *Waterfall) GetType() TraceType {
	return this.Type
}

// Trace is an interface to identify all available traces. The main purpose is to provide autocompletion.
type Trace interface {
	GetType() TraceType
}

// TraceType represents an specific trace type.
type TraceType string

const (
	TraceTypeArea               TraceType = "area"
	TraceTypeBar                TraceType = "bar"
	TraceTypeBarpolar           TraceType = "barpolar"
	TraceTypeBox                TraceType = "box"
	TraceTypeCandlestick        TraceType = "candlestick"
	TraceTypeCarpet             TraceType = "carpet"
	TraceTypeChoropleth         TraceType = "choropleth"
	TraceTypeChoroplethmapbox   TraceType = "choroplethmapbox"
	TraceTypeCone               TraceType = "cone"
	TraceTypeContour            TraceType = "contour"
	TraceTypeContourcarpet      TraceType = "contourcarpet"
	TraceTypeDensitymapbox      TraceType = "densitymapbox"
	TraceTypeFunnel             TraceType = "funnel"
	TraceTypeFunnelarea         TraceType = "funnelarea"
	TraceTypeHeatmap            TraceType = "heatmap"
	TraceTypeHeatmapgl          TraceType = "heatmapgl"
	TraceTypeHistogram          TraceType = "histogram"
	TraceTypeHistogram2d        TraceType = "histogram2d"
	TraceTypeHistogram2dcontour TraceType = "histogram2dcontour"
	TraceTypeImage              TraceType = "image"
	TraceTypeIndicator          TraceType = "indicator"
	TraceTypeIsosurface         TraceType = "isosurface"
	TraceTypeMesh3d             TraceType = "mesh3d"
	TraceTypeOhlc               TraceType = "ohlc"
	TraceTypeParcats            TraceType = "parcats"
	TraceTypeParcoords          TraceType = "parcoords"
	TraceTypePie                TraceType = "pie"
	TraceTypePointcloud         TraceType = "pointcloud"
	TraceTypeSankey             TraceType = "sankey"
	TraceTypeScatter            TraceType = "scatter"
	TraceTypeScatter3d          TraceType = "scatter3d"
	TraceTypeScattercarpet      TraceType = "scattercarpet"
	TraceTypeScattergeo         TraceType = "scattergeo"
	TraceTypeScattergl          TraceType = "scattergl"
	TraceTypeScattermapbox      TraceType = "scattermapbox"
	TraceTypeScatterpolar       TraceType = "scatterpolar"
	TraceTypeScatterpolargl     TraceType = "scatterpolargl"
	TraceTypeScatterternary     TraceType = "scatterternary"
	TraceTypeSplom              TraceType = "splom"
	TraceTypeStreamtube         TraceType = "streamtube"
	TraceTypeSunburst           TraceType = "sunburst"
	TraceTypeSurface            TraceType = "surface"
	TraceTypeTable              TraceType = "table"
	TraceTypeTreemap            TraceType = "treemap"
	TraceTypeViolin             TraceType = "violin"
	TraceTypeVolume             TraceType = "volume"
	TraceTypeWaterfall          TraceType = "waterfall"
)

type unmarshalType struct {
	Type TraceType `json:"type,omitempty"`
}

// UnmarshallTrace decodes an array of bytes into a Trace interface.
func UnmarshallTrace(data []byte) (Trace, error) {
	traceType := unmarshalType{}
	err := json.Unmarshal(data, &traceType)
	if err != nil {
		return nil, err
	}
	switch traceType.Type {
	case TraceTypeArea:
		trace := &Area{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeBar:
		trace := &Bar{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeBarpolar:
		trace := &Barpolar{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeBox:
		trace := &Box{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeCandlestick:
		trace := &Candlestick{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeCarpet:
		trace := &Carpet{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeChoropleth:
		trace := &Choropleth{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeChoroplethmapbox:
		trace := &Choroplethmapbox{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeCone:
		trace := &Cone{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeContour:
		trace := &Contour{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeContourcarpet:
		trace := &Contourcarpet{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeDensitymapbox:
		trace := &Densitymapbox{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeFunnel:
		trace := &Funnel{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeFunnelarea:
		trace := &Funnelarea{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeHeatmap:
		trace := &Heatmap{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeHeatmapgl:
		trace := &Heatmapgl{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeHistogram:
		trace := &Histogram{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeHistogram2d:
		trace := &Histogram2d{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeHistogram2dcontour:
		trace := &Histogram2dcontour{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeImage:
		trace := &Image{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeIndicator:
		trace := &Indicator{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeIsosurface:
		trace := &Isosurface{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeMesh3d:
		trace := &Mesh3d{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeOhlc:
		trace := &Ohlc{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeParcats:
		trace := &Parcats{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeParcoords:
		trace := &Parcoords{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypePie:
		trace := &Pie{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypePointcloud:
		trace := &Pointcloud{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeSankey:
		trace := &Sankey{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScatter:
		trace := &Scatter{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScatter3d:
		trace := &Scatter3d{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScattercarpet:
		trace := &Scattercarpet{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScattergeo:
		trace := &Scattergeo{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScattergl:
		trace := &Scattergl{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScattermapbox:
		trace := &Scattermapbox{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScatterpolar:
		trace := &Scatterpolar{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScatterpolargl:
		trace := &Scatterpolargl{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeScatterternary:
		trace := &Scatterternary{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeSplom:
		trace := &Splom{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeStreamtube:
		trace := &Streamtube{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeSunburst:
		trace := &Sunburst{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeSurface:
		trace := &Surface{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeTable:
		trace := &Table{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeTreemap:
		trace := &Treemap{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeViolin:
		trace := &Violin{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeVolume:
		trace := &Volume{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	case TraceTypeWaterfall:
		trace := &Waterfall{}
		err = json.Unmarshal(data, trace)
		if err != nil {
			return nil, err
		}
		return trace, nil
	default:
		return nil, errors.New("Trace Type is not registered")
	}
}
