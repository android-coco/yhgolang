package main

import "fmt"

func main() {
	var a int = 8;
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
       fmt.Println("a 小于 20\n" );
	}
	fmt.Println("a 的值为 : %d\n", a);

	var marks int = 90;
	var grder = "B";
	switch marks {
		case 90 : grder = "A";
		case 80 : grder = "B";
		case 50,60,70 : grder = "C";
		default : grder = "D";
	}
	switch {
      case grder == "A" :
         fmt.Println("优秀!\n" );     
      case grder == "B", grder == "C" :
         fmt.Println("良好\n" );      
      case grder == "D" :
         fmt.Println("及格\n" );      
      case grder == "F":
         fmt.Println("不及格\n" );
      default:
         fmt.Println("差\n" );
   }
   fmt.Println("你的等级是 %s\n", grder );

   var x interface{}
   switch i := x.(type) {
      case nil:	  
         fmt.Println(" x 的类型 :%T",i);                
      case int:	  
         fmt.Println("x 是 int 型");                       
      case float64:
         fmt.Println("x 是 float64 型");           
      case func(int) float64:
         fmt.Println("x 是 func(int) 型");                      
      case bool, string:
         fmt.Println("x 是 bool 或 string 型" );       
      default:
         fmt.Println("未知型");     
   }

   var c1, c2, c3 chan int;
   var i1, i2 int;
   select {
      case i1 = <-c1:
         fmt.Println("received ", i1, " from c1");
      case c2 <- i2:
         fmt.Println("sent ", i2, " to c2");
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Println("received ", i3, " from c3");
         } else {
            fmt.Println("c3 is closed");
         }
      default:
         fmt.Println("no communication");
   }        
}