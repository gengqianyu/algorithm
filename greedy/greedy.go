package greedy

import (
	"log"

	mapSet "github.com/deckarep/golang-set"
)

type coverFun func(map[interface{}][]interface{}) []interface{}

//利用贪心算法解决集合覆盖问题
func SetCover(m map[interface{}][]interface{}) []interface{} {
	//存放选择子集合在map中对应 key
	var c []interface{}

	//将切片转换为集合
	mSet := make(map[interface{}]mapSet.Set)

	for k, v := range m {
		mSet[k] = mapSet.NewSetFromSlice(v)
	}

	//获取map集合所有不重复选项
	a := mapSet.NewSet()

	for _, set := range mSet {
		a = a.Union(set)
	}
	//迭代集合
	for e := range a.Iter() {
		log.Println(e)
	}

	//定义maxKey 保存再一轮遍历中，能够覆盖最大未覆盖的元素对应map 子集合的key
	var maxKey interface{}

	//如果包含所有元素的a集合，元素个数不为零说明还有没有被用覆盖的元素
	for len(a.ToSlice()) != 0 {
		maxKey = nil
		//遍历map获取每个子集合，选出一个最大的未覆盖子集合，加入到c中
		for key, set := range mSet {
			//	求出set子集合和a总集合的交集，然后再求出当前子集合set最大未覆盖元素个数 l
			l := len(set.Intersect(a).ToSlice())
			//如果l==0说明子集合set，已经把总集合a已经全覆盖了
			//如果l>0就说明还有未被覆盖的元素，如果当前这个集合包含的未覆盖元素的数量，比maxKey指向的未覆盖元素还多，重置maxKey
			if l > 0 && (maxKey == nil || l > len(mSet[maxKey].ToSlice())) { //贪婪
				maxKey = key
			}
		}
		//将最大覆盖子集对应的key放入c中
		if maxKey != nil {
			c = append(c, maxKey)
		}
		//从a总集合删除maxKey对应子集合覆盖的元素，也就是求差集
		a = a.Difference(mSet[maxKey])
	}

	return c
}
