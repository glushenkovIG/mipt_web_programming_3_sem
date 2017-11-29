
package main
import "fmt"

var ans int = 1
/*
func main(){
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result) // Должно напечататься [3 5]
	gen := PowerGenerator(2)
	fmt.Println(gen()) // Должно напечатать 2
	fmt.Println(gen()) // Должно напечатать 4
	fmt.Println(gen()) // Должно напечатать 8

	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
	// Должно напечатать 2

}
*/

func RemoveEven(input []int) []int{
	result := []int{}
	for i := 0; i < len(input); i++ {
		if(input[i] % 2 != 0){
			result = append(result, input[i]);
		}
	}
	return result;
}

func PowerGenerator(a int) func() int {
	return func() int {
		ans = ans * a
		return ans
	}
}

func DifferentWordsCount(text string) int{
	new_word := ""
	dict := make(map[string]int)
	for i := 0; i < len(text); i++{
		if((int('A') <= int(text[i])) && (int(text[i]) <= int('Z'))){
			new_word += string(int(text[i]) - int('A') + int('a'))

		}else if((int('a') <= int(text[i])) && (int(text[i]) <= int('z'))){
			new_word += string(text[i])

		}else if(new_word != ""){
			dict[new_word] = 1
			new_word = ""
		}
	}
	if(new_word != ""){
		dict[new_word] = 1
	}
	return len(dict)
}
