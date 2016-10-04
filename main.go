package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

var m = make(map[string]int)

// A data structure to hold key/value pairs
type Pair struct {
	Key   string
	Value int
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortedValues(m map[string]int) {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
		fmt.Printf("%v\n", i)
	}
	for _, k := range p {
		fmt.Printf("%s\n", k.Key)
	}
	fmt.Println("")

	sort.Sort(p)

	fmt.Printf("Post-sorted: ")
	for _, k := range p {
		fmt.Printf("%s: %v\n", k.Key, k.Value)
	}
	fmt.Println("")

}
func dumpobj(prefix string, x interface{}) {

	switch t := x.(type) {
	case map[string]interface{}:
		for k, v := range t {
			if len(prefix) > 0 {
				dumpobj(prefix+"."+k, v)
			} else {
				dumpobj(k, v)
			}
		}
	case []interface{}:
		for i, v := range t {
			dumpobj(prefix+"["+strconv.Itoa(i)+"]", v)
		}
	case string:
		//fmt.Printf("%s = %d\n", prefix, len(t))
		m[prefix] = len(t)
	default:
		fmt.Printf("Unhandled: %T\n", t)
	}
}

func main() {
	data, err := ioutil.ReadFile("./export.json")
	if err != nil {
		fmt.Printf("can not read file %v", err)
		os.Exit(1)
	}
	var etcd map[string]interface{}
	if err = json.Unmarshal(data, &etcd); err != nil {
		fmt.Printf("Could not unmarshal json %v", err)
	}
	dumpobj("", etcd)
	for i, v := range m {
		fmt.Printf("%v: %v\n", i, v)
	}
	fmt.Printf("--------")
	sortedValues(m)

}
