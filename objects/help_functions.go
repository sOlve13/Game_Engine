package objects

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"

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

func rotatePoint(x, y, cx, cy int, angle float64) (int, int) {
	xf, yf := float64(x-cx), float64(y-cy)
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)

	// Если угол близок к 90°, 180°, 270°, 360°, то округляем cosA и sinA
	if math.Abs(cosA) < 1e-10 {
		cosA = 0
	}
	if math.Abs(sinA) < 1e-10 {
		sinA = 0
	}

	newX := xf*cosA - yf*sinA
	newY := xf*sinA + yf*cosA

	return int(math.Round(newX)) + cx, int(math.Round(newY)) + cy
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getImageSize(filePath string) (int, int, error) {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// Декодируем изображение
	img, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	// Получаем размеры изображения
	return img.Bounds().Dx(), img.Bounds().Dy(), nil
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func indexOf(slice []int, item int) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1 // Если элемент не найден
}

func writeToFile(filename, text string) error {
	// Открываем файл для записи, который перезапишет существующее содержимое
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Записываем текст в файл
	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("Ошибка при записи в файл: %v", err)
	}

	return nil
}
