package libs

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRect(x, y, width, height float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r Rect) MaxX() float64 {
	return r.X + r.Width
}

func (r Rect) MaxY() float64 {
	return r.Y + r.Height
}

// Check collision of two rectangles
func (r Rect) Intersects(other Rect) bool {
	return r.X <= other.MaxX() &&
		other.X <= r.MaxX() &&
		r.Y <= other.MaxY() &&
		other.Y <= r.MaxY()
}

// Check collision with an fixed position
func (r Rect) IntersectsAt(other Rect, offsetX, offsetY float64) bool {
	adjustedRect := Rect{
		X:      r.X + offsetX,
		Y:      r.Y + offsetY,
		Width:  r.Width,
		Height: r.Height,
	}
	return adjustedRect.Intersects(other)
}
