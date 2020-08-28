package greedy

import (
	"testing"

	mapset "github.com/deckarep/golang-set"
)

func TestSetCover(t *testing.T) {

	marshalTests := []struct {
		name   string
		new    func() coverFun
		m      map[string][]string
		golden []interface{}
	}{
		{
			name: "SetCover",
			new: func() coverFun {
				return SetCover
			},
			m: map[string][]string{
				"f1": []string{"北京", "上海", "天津"},
				"f2": []string{"广州", "北京", "深圳"},
				"f3": []string{"成都", "上海", "杭州"},
				"f4": []string{"上海", "天津"},
				"f5": []string{"杭州", "大连"},
			},
			golden: []interface{}{"f1", "f5", "f2", "f3"},
		},
	}

	//电台
	//radios := map[string][]string{
	//	"f1": []string{"北京", "上海", "天津"},
	//	"f2": []string{"广州", "北京", "深圳"},
	//	"f3": []string{"成都", "上海", "杭州"},
	//	"f4": []string{"上海", "天津"},
	//	"f5": []string{"杭州", "大连"},
	//}

	//radios["f1"] = []string{"北京", "上海", "天津"}
	//radios["f2"] = []string{"广州", "北京", "深圳"}
	//radios["f3"] = []string{"成都", "上海", "杭州"}
	//radios["f4"] = []string{"上海", "天津"}
	//radios["f5"] = []string{"杭州", "大连"}

	for _, tt := range marshalTests {

		t.Run(tt.name, func(t *testing.T) {

			m := make(map[interface{}][]interface{})
			for key, value := range tt.m {
				m[key] = func(s []string) []interface{} {
					t := make([]interface{}, len(value))
					for i, e := range s {
						t[i] = e
					}
					return t
				}(value)
			}

			c := (tt.new())(m)
			cSet := mapset.NewSetFromSlice(c)
			gSet := mapset.NewSetFromSlice(tt.golden)

			if !cSet.Equal(gSet) {
				t.Errorf("expected:%v,actual:%v", tt.golden, c)
			}

			//if !reflect.DeepEqual(tt.golden, c) {
			//	t.Errorf("expected:%s,actual:%s", func() (str string) {
			//		for i, e := range tt.golden {
			//			if i == 0 {
			//				str += e
			//			} else {
			//				str = str + "," + e
			//			}
			//		}
			//		return
			//	}(), c)
			//}
		})

	}

}
