package hge

import "math"

func InvSqrt(x float64) float64 {
	return 1 / math.Sqrt(x)
}

type Vector struct {
	X, Y float64
}

func NewVector(x, y float64) Vector {
	var v Vector
	v.X = x
	v.Y = y
	return v
}

func (v Vector) Negate() Vector {
	return NewVector(-v.X, -v.Y)
}

func (v Vector) Subtract(v2 Vector) Vector {
	return NewVector(v.X-v2.X, v.Y-v2.Y)
}

func (v Vector) Add(v2 Vector) Vector {
	return NewVector(v.X+v2.X, v.Y+v2.Y)
}

func (v *Vector) SubtractEqual(v2 Vector) *Vector {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v *Vector) AddEqual(v2 Vector) *Vector {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v Vector) EQ(v2 Vector) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vector) NEQ(v2 Vector) bool {
	return v.X != v2.X && v.Y != v2.Y
}

func (v Vector) Divide(scalar float64) Vector {
	return NewVector(v.X/scalar, v.Y/scalar)
}

func (v Vector) Multiply(scalar float64) Vector {
	return NewVector(v.X*scalar, v.Y*scalar)
}

func (v *Vector) MultiplyEqual(scalar float64) *Vector {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vector) Angle(arg ...interface{}) float64 {
	if len(arg) == 1 {
		if vec, ok := arg[0].(Vector); ok {
			s := vec
			t := vec

			s.Normalize()
			t.Normalize()
			return math.Acos(s.Dot(t))
		}
	} else {
		return math.Atan2(v.Y, v.X)
	}

	return 0.0
}

func (v *Vector) Clamp(max float64) {
	if v.Length() > max {
		v.Normalize()
		v.X *= max
		v.Y *= max
	}
}

func (v *Vector) Normalize() *Vector {
	rc := InvSqrt(v.Dot(*v))
	v.X *= rc
	v.Y *= rc

	return v
}

func (v *Vector) Rotate(a float64) *Vector {
	var vec Vector

	vec.X = v.X*math.Cos(a) - v.Y*math.Sin(a)
	vec.Y = v.X*math.Sin(a) + v.Y*math.Cos(a)

	v.X = vec.X
	v.Y = vec.Y

	return v
}

func VectorAngle(v Vector, u Vector) float64 {
	return v.Angle(u)
}

func VectorDot(v Vector, u Vector) float64 {
	return v.Dot(u)
}
