package main
import "fmt"
func getLang() (string, string, string) {
	return "Hindi", "Sanskrit", "English"
}
func getSum(nums...int)int{
	result:=0;
	for _,num:=range nums{
		result+=num;
	}
	return result

}
  func main(){
	nums:=[]int{1,2,4,5,5};
	result:=0;
	for _,num:=range nums{
       result+=num;
	}
	fmt.Println(result);
	m:=map[string]string {"Rohit":"singh","Singularian":"Plural"};
	fmt.Println(m)
	
	
	
		lang1, lang2, lang3 := getLang()
		fmt.Println(lang1, lang2, lang3)

       nums:=[]int{1,3,5,6,7};
	   response:=getSum(nums...);
	   fmt.Println(response)

  }
