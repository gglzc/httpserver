package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Table struct{
	 DiceNumber int `json:"dicenumber"`	
	 Val        int `json:"val"` 
	 Bet	    int `json:"bet"` 
}


func (t *Table) CalWin(c *fiber.Ctx)error{

	dicenumber:=t.DiceNumber
	bet:=t.Bet
	val:=t.Val
	
	//dicenumber must between 1 and 6
	if dicenumber<=6 && dicenumber>=1{
		win:=dicenumber*bet*val
		err:=c.JSON(fiber.Map{
			"Win":win,
		})
		if err!=nil{
			fmt.Println("some wrong")
		}
		return err
	}else{
		fmt.Println("Please input the right dicenumber")
	}
	return nil
 }

