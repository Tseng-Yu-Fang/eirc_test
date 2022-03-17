package main

import(
	"fmt"
	"encoding/json"
) 

var x,y,z int = 1,2,3
var(
	a = 4
	b ="hay"
	c = 5
)

type students struct{
	Name string `json:name`
	Number string `json:number`
}

func main(){
	// d := "wow" 
	amy := students{
		Name: "amy",
		Number:"C108118221",
	}
	vin := students{}
	vin.Name ="vin"
	vin.Number="C108118217"
	
	// num :=[]int{1,2,3,4}
	class := []students{vin,amy}
	// class2 :=[]students{}

	// for index,value := range class{
	// 	class2 = append(class2, value) 
	// 	fmt.Println(index,class2)
	// }

	//序列化
	marshal, err := json.Marshal(class)
	if err != nil {
		fmt.Println(err)
	}
	// []byte
	fmt.Println(string(marshal))


	class3 := []students{}
	err = json.Unmarshal(marshal, &class3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(class3)
}