package utils

import (
	"math"
)

// VectorMagnitude 计算向量的模
func VectorMagnitude(vector []float64) float64 {
	sum := 0.0
	for _, v := range vector {
		sum += v * v
	}
	return math.Sqrt(float64(sum))
}

// DotProduct 计算两个向量的点积
func DotProduct(vector1, vector2 []float64) float64 {
	product := 0.0
	for i := 0; i < len(vector1); i++ {
		product += vector1[i] * vector2[i]
	}
	return product
}

// CosineSimilarity 计算余弦相似度
func CosineSimilarity(vector1, vector2 []float64) float64 {
	dotProd := DotProduct(vector1, vector2)
	mag1 := VectorMagnitude(vector1)
	mag2 := VectorMagnitude(vector2)
	if mag1 == 0 || mag2 == 0 {
		return 0
	}
	return float64(dotProd) / (mag1 * mag2)
}
