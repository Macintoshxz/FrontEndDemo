package utils

import "sort"

// 类型约束
// type Number interface {
// 	int | float32 | float64 | int32
// }

func Map1(s []int, f func(int) int) []int {
	result := make([]int, len(s))

	for i, v := range s {
		result[i] = f(v)
	}

	return result
}

// 泛型工具切片
type GSice[T any] []T

// 根据函数创建新数组
func (s GSice[T]) Map(f func(T) T) GSice[T] {
	result := make(GSice[T], len(s))

	for i, v := range s {
		result[i] = f(v)
	}

	return result
}

// 聚合函数
//  prv 前一个值, 当前值
func (s GSice[T]) Reduce(f func(prv T, curent T) T) T {
	var result T

	for _, v := range s {
		result = f(result, v)
	}

	return result
}

// 过滤
func (s GSice[T]) Filter(f func(T) bool) GSice[T] {
	result := GSice[T]{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}

	}
	return result
}

// 找到符合条件的第一个元素
func (s GSice[T]) Find(f func(T) bool) *T {
	for _, v := range s {
		if f(v) {
			return &v
		}
	}

	return nil
}

// 排序
// 使用
// Sort(func(a, b int) bool {
//     return a < b
// })
func (s GSice[T]) Sort(less func(T, T) bool) {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
}

// 颠倒数组
func (s GSice[T]) Reverse() {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
