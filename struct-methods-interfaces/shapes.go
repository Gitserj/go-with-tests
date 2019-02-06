package shapes

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

func Area(r Rectangle) float64 {
	return r.width * r.heigth
}
