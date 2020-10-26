package math

type Vector2di struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (vec *Vector2di)Plus(vec2 Vector2di) {
	vec.X += vec2.X
	vec.Y += vec2.Y
}

func (vec *Vector2di)Minus(vec2 Vector2di) {
	vec.X -= vec2.X
	vec.Y -= vec2.Y
}

func (vec *Vector2di)Mult(vec2 Vector2di) {
	vec.X *= vec2.X
	vec.Y *= vec2.Y
}

func (vec *Vector2di)Div(vec2 Vector2di) {
	vec.X /= vec2.X
	vec.Y /= vec2.Y
}
