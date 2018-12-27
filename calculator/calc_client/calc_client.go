package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jarrocha/go_grpc/calculator/calcpb"
	"google.golang.org/grpc"
)

var server_result = make(chan float64)

func parseOperator(in string) calcpb.Operand {

	switch in {
	case "+":
		return calcpb.Operand_SUM
	case "-":
		return calcpb.Operand_SUB
	case "*":
		return calcpb.Operand_MUL
	case "/":
		return calcpb.Operand_DIV
	default:
		return calcpb.Operand_UNKNOWN
	}
}

func printResult() {
	for res := range server_result {
		fmt.Printf("Result: %g\n", res)
	}
}

func doRequest(client calcpb.OperServiceClient) {

	var num1, num2 float64
	var oper string
	var operand calcpb.Operand

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		raw := scanner.Text()
		nraw := strings.Fields(raw)

		if len(nraw) != 3 {
			fmt.Println("Invalid Operation. Expects: NUM1 [=-/] NUM2")
			return
		}

		num1, _ = strconv.ParseFloat(nraw[0], 64)
		num2, _ = strconv.ParseFloat(nraw[2], 64)
		oper = nraw[1]
		operand = parseOperator(oper)
	}

	req := &calcpb.OperRequest{
		Operation: &calcpb.Operation{
			Operator: operand,
			Number1:  num1,
			Number2:  num2,
		},
	}

	response, err := client.Calculate(context.Background(), req)
	if err != nil {
		log.Println("Response error", err)
	}

	server_result <- response.GetResult()
}

func main() {

	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("gRPC Dial error. ", err)
	}
	defer conn.Close()

	client := calcpb.NewOperServiceClient(conn)

	go printResult()

	fmt.Printf("Enter operation with spaces in between.")
	for {
		doRequest(client)
	}

}
