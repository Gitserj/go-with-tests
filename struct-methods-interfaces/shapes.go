package shapes

import "math"

// Функции и методы очень похожи
// Функция может принимать на вход тип "обьекта" func(obj)
// Метод может быть вызван для экзепляра "обьекта" obj.func()

type (
	Rectangle struct {
		width  float64
		heigth float64
	}
	Circle struct {
		radius float64
	}
)

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.heigth)
}

// Функция
// func Area(r Rectangle) float64 {
// 	return r.width * r.heigth
// }

// Её аналог в виде методов для 2х типов данных
func (r Rectangle) Area() float64 {
	return r.width * r.heigth
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
