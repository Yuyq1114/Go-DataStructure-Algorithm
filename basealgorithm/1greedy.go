package basealgorithm

// 贪心算法（Greedy Algorithm）：
// 每步选择当前状态下最优的解，从而得到全局最优解的一种算法。

// 1定义最优解的特征，并确定可行解的范围。首先需要明确问题的最优解具备哪些特征，比如最大值、最小值等，然后需要确定可行解的范围，即哪些解是可行的。

// 2根据贪心策略，选取当前最优解。在每一步中，根据贪心策略选取当前最优解。贪心策略可以是基于局部最优、全局最优或者其他特定条件的。

// 3判断当前解是否满足可行性条件。每选取一个解，都需要判断其是否满足可行性条件，即是否符合问题的限制条件。

// 4若当前解满足可行性条件，则将其加入到最优解集合中。如果当前解满足可行性条件，即符合问题的限制条件，则将其加入到最优解集合中。

// 5重复上述步骤，直至问题得到解决。在完成一次选择后，需要重新进入第二步，选取下一个最优解，直至问题得到解决。
import (
	"sort"
)

// 贪心算法求解背包问题
func KnapsackG(w []int, v []int, c int) int {
	n := len(w)
	items := make([]item, n)

	//求单位价值
	for i := 0; i < n; i++ {
		items[i] = item{w[i], v[i], float64(v[i]) / float64(w[i])}
	}
	sort.Sort(sort.Reverse(byRatio(items))) // 按单位价值从大到小排序

	var (
		totalValue int // 总价值
		remain     int = c
	)
	//装入背包
	for i := 0; i < n; i++ {
		if remain < items[i].weight { // 物品装不下，装满背包
			totalValue += int(float64(remain) * items[i].ratio)
			break
		}
		remain -= items[i].weight
		totalValue += items[i].value
	}
	return totalValue
}

// 物品结构体
type item struct {
	weight int     // 重量
	value  int     // 价值
	ratio  float64 // 单位价值
}

// 按单位价值从大到小排序
type byRatio []item

func (r byRatio) Len() int           { return len(r) }
func (r byRatio) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byRatio) Less(i, j int) bool { return r[i].ratio < r[j].ratio }

// func main() {
// 	w := []int{2, 2, 6, 5, 4} // 物品重量
// 	v := []int{6, 3, 5, 4, 6} // 物品价值
// 	c := 10                   // 背包容量

// 	totalValue := KnapsackG(w, v, c)

// 	fmt.Printf("背包容量：%d\n", c)
// 	fmt.Printf("物品重量：%v\n", w)
// 	fmt.Printf("物品价值：%v\n", v)
// 	fmt.Printf("最大总价值：%d\n", totalValue)
// }
