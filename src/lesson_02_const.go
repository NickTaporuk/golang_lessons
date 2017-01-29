package main

import "fmt"

const SingleLineString = "Single line constant";

const (
	MultiLineFirst = 1
	MultiLineSecond
	MultiLineThird = "third"
	MultiLineFourth
)


const (
	IotaUseCaseFirst = 1 << iota
	IotaUseCaseSecond
	IotaUseCaseThird
	IotaUseCaseFourth
)

const (
	IotaUseCase2First = iota * 3 - 2
	IotaUseCase2Second
	IotaUseCase2Third
	IotaUseCase2Fourth
)

func main() {
	fmt.Println(SingleLineString, "\n")

	fmt.Println("MultiLineFirst " , MultiLineFirst)
	fmt.Println("MultiLineSecond " , MultiLineSecond )
	fmt.Println("MultiLineThird " , MultiLineThird)
	fmt.Println("MultiLineFourth " , MultiLineFourth, "\n")

	fmt.Println("IotaUseCaseFirst " , IotaUseCaseFirst)
	fmt.Println("IotaUseCaseSecond " , IotaUseCaseSecond )
	fmt.Println("IotaUseCaseThird " , IotaUseCaseThird)
	fmt.Println("IotaUseCaseFourth " , IotaUseCaseFourth, "\n")

	fmt.Println("IotaUseCase2First " , IotaUseCase2First)
	fmt.Println("IotaUseCase2Second " , IotaUseCase2Second )
	fmt.Println("IotaUseCase2Third " , IotaUseCase2Third)
	fmt.Println("IotaUseCase2Fourth " , IotaUseCase2Fourth, "\n")

}
