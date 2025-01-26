package utils

import "image"

func RectIntersectFloat(s image.Rectangle, x, y float64, width, height int) (dx, dy float64) {
	xMin, xMax := x, x+float64(width)
	yMin, yMax := y, y+float64(height)
	if xMin < float64(s.Min.X) {
		xMin = float64(s.Min.X)
	}
	if yMin < float64(s.Min.Y) {
		yMin = float64(s.Min.Y)
	}
	if xMax > float64(s.Max.X) {
		xMax = float64(s.Max.X)
	}
	if yMax > float64(s.Max.Y) {
		yMax = float64(s.Max.Y)
	}
	return xMax - xMin, yMax - yMin
}
