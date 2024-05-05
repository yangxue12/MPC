package main

import (
	"fmt"
	"math"
	"math/rand"
)

const p = 257

func main() {
	secrets := []int{23, 55, 17}
	users := getSegment(secrets, 7, 3)
	fmt.Println(users)
	data := getSecret(users[0], users[1], users[2])
	fmt.Println(data)
}

func getSegment(secrets []int, n int, t int) [][]int {
	users := make([][]int, n)
	for i := 0; i < len(secrets); i++ {
		result := segmentCreate(n, t, secrets[i])
		for j := 0; j < n; j++ {
			//如果users[j][]为空，则往里添加index
			if len(users[j]) == 0 {
				users[j] = append(users[j], j+1)
			}
			users[j] = append(users[j], result[j])
		}
	}
	return users
}

func segmentCreate(n int, t int, m int) []int {
	//判断p是否为质数
	if !isPrime(p) {
		fmt.Println("p is not a prime number")
		return nil
	}
	//随机生成小于p的权重（系数）,并用数组存储
	weight := make([]int, t)
	weight[0] = m
	for i := 1; i < t; i++ {
		weight[i] = rand.Intn(p)
	}
	results := make([]int, n)
	for i := 0; i < n; i++ {
		result := 0
		for j := 0; j < t; j++ {
			result += weight[j] * int(math.Pow(float64(i+1), float64(j)))

			/*
				用tmp存储*index
			*/
		}
		results[i] = result % p
	}
	return results
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getSecret(users ...[]int) []int { //输入对应的
	secrets := []int{}
	data := make([][][]int, len(users[0])-1)
	for _, user := range users {
		for j := 0; j < len(user)-1; j++ {
			singledata := []int{user[0], user[j+1]}
			data[j] = append(data[j], singledata)
		}
	}
	for i := 0; i < len(data); i++ {
		//secret := getSingleSeg(data[i])
		secret := getSingleSeg(data[i]) % p
		secrets = append(secrets, secret)
	}
	return secrets
}

func getSingleSeg(data [][]int) int {
	n := len(data)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, n)
		for j := 0; j < n; j++ {
			A[i][j] = int(math.Pow(float64(data[i][0]), float64(j)))
		}
	}

	B := make([]int, n)
	for i := 0; i < n; i++ {
		B[i] = data[i][1]
	}
	x := gaussElimination(A, B)
	if x[0] < 0 {
		x[0] = x[0] + p

	}
	return x[0]
}

func gaussElimination(A [][]int, b []int) []int {
	n := len(A)

	// 增广矩阵
	AB := make([][]int, n)
	for i := range AB {
		AB[i] = make([]int, n+1)
		copy(AB[i], A[i])
		AB[i][n] = b[i]
	}
	// 消元
	for k := 0; k < n-1; k++ {
		for i := k + 1; i < n; i++ {
			factor := AB[i][k] / AB[k][k]
			for j := k; j < n+1; j++ {
				AB[i][j] -= factor * AB[k][j]
			}
		}
	}
	fmt.Println(AB)
	// 回代求解
	x := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0
		for j := i + 1; j < n; j++ {
			sum += (AB[i][j] * x[j])
		}
		x[i] = ((AB[i][n] - sum) / AB[i][i]) % p
	}
	return x
}
