//					ASSIGNMENT-2
// CREATE A BILLING SYSTEM THAT A USER SENDS NUMBER OF MESSAGES AND WHICH COSTS 0.02 DOLLARS
// CALCULATE THE TOTAL COST, IF DORIS SEND 72 MESSAGES IN A SINGLE DAY

package main

import "fmt"

func main() {
	numMessagesFromDoris := 72
	costPerMessage := 0.02
	totalCost := costPerMessage * float64(numMessagesFromDoris)
	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}
