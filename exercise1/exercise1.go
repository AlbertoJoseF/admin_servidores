package main

import "fmt"
import "math/rand"
import "time"

func main(){
	//HELLOWORLD
	//helloWorld()	

	//VECTOR OF ZEROES
	/*vector := vectorOfZeroes()
	fmt.Println(vector)*/

	//VECTOR OF RANDOM VALUES
	/*vector := vectorOfRandom()
	fmt.Println(vector)*/

	//MATRIX OF ZEROES
	/*matrix := matrixOfZeroes()
	fmt.Println(matrix)*/

	//MATRIX OF RANDOM VALUES
	/*matrix := matrixOfRandom()
	fmt.Println(matrix)*/

	//MULTIPLICATION OF VECTOR AND MATRIX
	/*n := 5
	m := 6
	threshold := 123
	vector := createRandomVector(n, threshold)
	time.Sleep(1)
	matrix := createNonSquareRandomMatrix(m, n, threshold)
	result := vectorMatrixMult(vector, matrix)
	fmt.Println(vector)
	fmt.Println(matrix)
	fmt.Println(result)*/

	//SOLVE SYSTEM OF EQUATIONS GIVEN TRIANGULAR MATRIX AND AN OBJECTIVE VECTOR
	n := 6
	threshold := 123
	matrix := createRandomMatrix(n, threshold)
	matrix = makeIntoUpperTriangularMatrix(matrix)
	//vector := createRandomVector(n, threshold)
	fmt.Println(matrix)
}

/*
func solveSystemOfEquations(system [][]int, objective []int) [] int {
	var values []int
	for i := (len(system) - 1); i > -1; i--{
		if len(values) == 0{
			if 
		} else{

		}
	}
		col := 
	for i := (len(system) - 1); i > -1; i--{
		for j := ()
	}
	return
}*/

func makeIntoUpperTriangularMatrix(matrix [][]int) [][]int {
	col := 0
	for i := 0; i < len(matrix); i++{
		for j := 0; j < col; j++{
			matrix[i][j] = 0
		}
		col ++
	}
	 
	return matrix
}

func vectorMatrixMult(vector []int, matrix [][]int) []int {
	var result []int
	temp := 0
	for i := 0; i < len(matrix); i++{
		for j := 0; j < len(vector); j++{
			temp += vector[j] * matrix[i][j]
		}
		result = append(result, temp)
		temp = 0 
	}
	return result
}

func matrixOfRandom() [][]int {
	var n int
	var threshold int
	fmt.Println("Give me the size of the n x n matrix of random numbers:")
	fmt.Scanf("%d\n", &n)
	fmt.Println("Give me the maximum number as threshold for random numbers:")
	fmt.Scanf("%d\n", &threshold)
	matrix := createRandomMatrix(n, threshold)
	return matrix
}

func createRandomMatrix(n, threshold int) [][]int {
	var matrix [][]int
	for i := 0; i < n; i++{
		matrix = append(matrix, createRandomVector(n, threshold))
		time.Sleep(1)
	}
	return matrix
}

func createNonSquareRandomMatrix(n, m, threshold int) [][]int {
	var matrix [][]int
	for i := 0; i < n; i++{
		matrix = append(matrix, createRandomVector(m, threshold))
		time.Sleep(1)
	}
	return matrix
}


func matrixOfZeroes() [][]int {
	var n int
	var matrix [][]int
	fmt.Println("Give me the size of the n x n matrix of 0's:")
	fmt.Printf("n: ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++{
		matrix = append(matrix, createZeroesVector(n))
	}
	return matrix
}

func vectorOfRandom() []int{
	var size int
	var threshold int
	fmt.Println("Give me the size of the vector of random numbers:")
	fmt.Scanf("%d\n", &size)
	fmt.Println("Give me the maximum number as threshold for random numbers:")
	fmt.Scanf("%d\n", &threshold)
	vector := createRandomVector(size, threshold)
	return vector
}

func createRandomVector(size, threshold int) []int{
	var vector []int
	var number int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++{
		number = rand.Intn(threshold)
		vector = append(vector, number)
	}
	return vector
}

func vectorOfZeroes() []int{
	var size int
	fmt.Println("Give size of the vector to initialize in 0's:")
	fmt.Scanf("%d\n", &size)
	vector := createZeroesVector(size)
	//var x [size]int -> NO FUNCIONA porque size se define en tiempo de ejecución
	//fmt.Println(vector) -> TEST
	return vector
}

func createZeroesVector(size int) []int{
	vector := make([]int, size)
	return vector
}

func helloWorld(){
	name := ""
	fmt.Println("What's your name?")
	fmt.Scanf("%s\n", &name)
	fmt.Println("hello", name)
}