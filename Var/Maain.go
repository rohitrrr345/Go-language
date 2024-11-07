package main
import "fmt"

func main(){


	fmt.Println("HI there ")
	fmt.Println(1+2+3)
	var name="rohit is giving it alll to the people"
	fmt.Println(name);
	var na int=13;
	fmt.Println(na);
	const remaninder=34;
	fmt.Println(remaninder);
	ans:="string";
	fmt.Println(ans);
	const(
		port=500;
		rem="localhost"
	)
	fmt.Println(port,rem);
	i:=1
	for i<=100{
		fmt.Println(i);
		i++
	}
	var nums[4]int;
	nums[0]=123
	var value=nums[0];
	fmt.Println(value);
	args:=[67]int{1,23,4,5,6,6};
	for i:=0;i<67;i++{
		fmt.Println(args[i])
	}
	arr:=[2][2]int{{1,3},{2,4}};
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Println(arr[i][j])
		}
	}
	 var arena[]int;
	 arena=append(arena,142);
	 arena=append(arena,142);
	 arena=append(arena,142);
	 arena=append(arena,142);
	 arena=append(arena,142);
	 fmt.Println(len(arena));
	 fmt.Println(arena);
	 n:=[4]int{};
	 fmt.Println(n);
	 mpp:=make(map[string]int);
	 mpp["arr"]=34;
	 mpp["rem"]=32535;
	 clear(mpp);
	 fmt.Println(mpp);
	 remea:=[34]int{14,325,25,235,234,12,4,124}
	 fmt.Println(remea)
    maps:=map[string]int{"jbesgj":13};
	fmt.Println(maps)
    

}