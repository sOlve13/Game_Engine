package objects

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func orientation(p, q, r Point2D) int {
	qX, qY := q.GetCoords()
	pX, pY := p.GetCoords()
	rX, rY := r.GetCoords()
	val := (qY-pY)*(rX-qX) - (qX-pX)*(rY-qY)
	if val == 0 {
		return 0 // коллинеарны
	} else if val > 0 {
		return 1 // по часовой стрелке
	} else {
		return 2 // против часовой стрелки
	}
}
func onSegment(p, q, r Point2D) bool {
	qX, qY := q.GetCoords()
	pX, pY := p.GetCoords()
	rX, rY := r.GetCoords()
	return qX <= max(pX, rX) && qX >= min(pX, rX) &&
		qY <= max(pY, rY) && qY >= min(pY, rY)
}
func segmentsIntersect(p1, q1, p2, q2 Point2D) bool {
	// Находим ориентации для четырех пар точек
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	// Общий случай
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Специальные случаи:
	// p1, q1 и p2 коллинеарны, и p2 лежит на отрезке p1q1
	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}
	// p1, q1 и q2 коллинеарны, и q2 лежит на отрезке p1q1
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}
	// p2, q2 и p1 коллинеарны, и p1 лежит на отрезке p2q2
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}
	// p2, q2 и q1 коллинеарны, и q1 лежит на отрезке p2q2
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}

	return false // Не пересекаются
}

func isPointInPolygon(p Point2D, polygon []Point2D, screen *ebiten.Image, backgroundColor color.Color) bool {
	n := len(polygon)
	if n < 3 {
		return false
	}
	_, py := p.GetCoords()

	// Создаем точку, которая явно вне полигона
	extreme := NewPoint2D(screen, backgroundColor, 1e10, py, backgroundColor)

	count := 0
	for i := 0; i < n; i++ {
		next := (i + 1) % n
		// Проверяем пересекается ли линия многоугольника с лучом
		if segmentsIntersect(polygon[i], polygon[next], p, extreme) {
			// Проверяем, лежит ли точка на ребре многоугольника
			if orientation(polygon[i], p, polygon[next]) == 0 && onSegment(polygon[i], p, polygon[next]) {
				return false // Если точка лежит на линии, возвращаем false
			}
			count++
		}
	}
	// Точка внутри полигона, если количество пересечений нечетное
	return count%2 == 1
}

func rotatePoint(cx, cy, x, y int, angle float64) (int, int) {
	// Переводим в float64 для вычислений
	xf, yf := float64(x-cx), float64(y-cy)

	// Применяем формулы поворота
	newX := xf*math.Cos(angle) - yf*math.Sin(angle)
	newY := xf*math.Sin(angle) + yf*math.Cos(angle)

	// Переводим обратно к целым числам
	return int(newX) + cx, int(newY) + cy
}