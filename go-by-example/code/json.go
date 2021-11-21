package gobyex

import (
	"encoding/json"
	"fmt"
	"os"
)

type res1 struct {
	Page   int
	Fruits []string
}

//fields needs to be exported to encode to json
type res2 struct {
	Page   int      `json:page`
	Fruits []string `json:fruits`
}

func JSONFn() {

	b, _ := json.Marshal(true)
	fmt.Println(string(b))

	i, _ := json.Marshal(1)
	fmt.Println(string(i))

	f, _ := json.Marshal(3.4)
	fmt.Println(string(f))

	s, _ := json.Marshal("ad")
	fmt.Println(string(s))

	sarr := []string{"aple", "banan"}
	sarrb, _ := json.Marshal(sarr)
	fmt.Println(string(sarrb))

	marr := map[string]int{"aple": 1, "banan": 2}
	marrb, _ := json.Marshal(marr)
	fmt.Println(string(marrb))

	res1d := res1{
		Page:   1,
		Fruits: []string{"apple", "orange"},
	}
	res1b, _ := json.Marshal(res1d)
	fmt.Println(string(res1b))

	res2d := res1{
		Page:   1,
		Fruits: []string{"apple", "orange"},
	}
	res2b, _ := json.Marshal(res2d)
	fmt.Println(string(res2b))

	//decode
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(dat["num"])

	//acccess nested entry after unmarshal. needs conversion

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res2ob := res2{}
	json.Unmarshal([]byte(str), &res2ob)
	fmt.Println(res2ob)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
