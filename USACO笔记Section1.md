## Section 1

### P1203 [USACO1.1]坏掉的项链Broken Necklace

> 一串项链，r、b两种颜色，w可以表示任意一种颜色。假如你要在一些点打破项链,展开成一条直线，然后从一端开始收集同颜色的珠子直到你遇到一个不同的颜色珠子，在另一端做同样的事(颜色可能与在这之前收集的不同)。 确定应该在哪里打破项链来收集到最大数目的珠子。

+ 思路1：枚举断开点

+ 思路2：动归

    $left[2][i] $与$right[2][i]$表示向左和向右找到 红色r/蓝色b 最多分别为多少



### P1209 [USACO1.3]修理牛棚 Barn Repair

> 计算拦住所有有牛的牛棚所需木板的最小总长度

转为 【最小m段问题】

+ 思路1：动归，$d[i][j] $前$i$个分成$j$段 所需要木材的最小值  $O(n^2)$

    $$ d[i][j] = min{d[i-1][j-1]+1, d[i-1][j]+num[i]-num[i-1]}, i>m $$

    ​            $$  = d[i-1][j-1] + 1 , i<=m $$  即有足够的木板

+ 思路2：贪心，每次选取距离最小的连接  $O(nlgn)$



### P3650 [USACO1.3]滑雪课程设计Ski Course Design

> 因此,如果他改变山峰的高度（使最高与最低的山峰海拔高度差不超过17）,约翰可以避免支付税收。如果改变一座山x单位的高度成本是x^2单位,约翰最少需要付多少钱

思路1：枚举起点夹逼 $O(n^2)$

思路2：贪心，



### P1444 [USACO1.3]虫洞wormhole

> N个虫洞在农场上（2<=N<=12，n是偶数）,虫洞将形成 N/2 连接配对。例如，如果A和B的虫洞连接成一对，进入虫洞A的任何对象体将从虫洞B出去，朝着同一个方向，而且进入虫洞B的任何对象将同样从虫洞A出去，朝着相同的方向前进。计算不同的虫洞配对（情况），使贝茜可能被困在一个无限循环中，如果她从不幸的位置开始。

思路：虫洞配对（回溯） + 检查是否有环（dfs）

给虫洞配对，对每种配对方式调用checkRing，检查是否有环。检查是否环时，枚举每一点作为出发点，使用visited数组标记每一点是从哪里被访问的。注意：通过虫洞到达和+x到达 是不同的，要分两个数组标记

本题特殊在于，从虫洞走出后只能+x走，不能再走一遍虫洞，即虫洞后一定为+x，+x后一定进入虫洞



### P1215 [USACO1.4]母亲的牛奶 Mother's Milk

> 有三个容量分别是A,B,C升的桶，A,B,C分别是三个从1到20的整数， 最初，A和B桶都是空的，而C桶是装满牛奶的。有时，农民把牛奶从一个桶倒到另一个桶中，直到被灌桶装满或原桶空了。当然每一次灌注都是完全的。
>
> 找出当A桶是空的时候，C桶中牛奶所剩量的所有可能性。

思路：类似回溯

注意：**局部变量**！！





### P1217 [USACO1.5]回文质数 Prime Palindromes

> 因为151既是一个质数又是一个回文数(从左到右和从右到左是看一样的)，所以 151 是回文质数。
>
> 写一个程序来找出范围$[a,b](5 <= a < b <= 100,000,000)$( 一亿)间的所有回文质数;

思路：生成回文数，判断是否质数

**判断质数：**

+ 方法1：枚举判断n能否能被从2到sqrt(n)之间的质数x整除，若不能则为质数。 $O(\sqrt{n})$

    ```go
    prime []int   //所有<sqrt(n)的质数的集合
    func isPrime(n int) int {
    	if n == 1 {
    		return 0
    	}
    	n_sqrt := int(math.Sqrt(float64(n)))
      for _,x := range() {
    		if n%x == 0 {
    			return 0
    		}
    	}
    	return 1
    }
    ```

+ 方法二：Miller-Rabin 素性检验  $O(k*logn^2)$ , k 为重复次数，一般次数<20

    

生成回文数可优化-但复杂度没有变化



### P1218 [USACO1.5]特殊的质数肋骨 Superprime Rib

> 7 3 3 1 全部肋骨上的数字 7331是质数;三根肋骨 733是质数;二根肋骨 73 是质数;当然,最后一根肋骨 7 也是质数。 7331 被叫做长度 4 的特殊质数。写一个程序对给定的肋骨的数目 N (1<=N<=8),求出所有的特殊质数。数字1不被看作一个质数。

+ 思路1：从低位至高位判断是否质数——过不了 
+ 思路2：从高位至低位判断质数。$v=v*10+ 2,3,5,7$





## Section2

### P1457 [USACO2.1]城堡 The Castle

### P1458 [USACO2.1]顺序的分数 Ordered Fractions

> 输入一个自然数N,对于一个最简分数a/b（分子和分母互质的分数），满足1<=b<=N，0<=a/b<=1。给定一个自然数N，1<=n<=160，按分数值递增的顺序输出所有解。

+ 思路1：维护一个大小为N的堆，n个数分母分别为1-n，每次弹出堆中最小的数（x/y），加入和这个数相同分母、分子+1（ (x+1)/y ），继续维护这个堆。 复杂度：$O(N^2 logN)$

### P1459 [USACO2.1]三值的排序 Sorting a Three-Valued Sequence

> Consider the special sorting problem in which the records to be sorted have at most three different key values . Please computes the minimal number of exchange operations that are necessary to make the sequence sorted.

+ 思路1：计算3种值每一种各有多少个，找出位置不正确的数，如果可以直接交换，交换。交换完成后剩下的都是不可以直接被交换的，一组（3个）的操作数为2。如：2,3,1，需要交换2次

### P1460 [USACO2.1] 健康的荷斯坦奶牛 Healthy Holsteins

> Help Farmer John feed the cows so they stay healthy while minimizing the number of scoops that a cow is fed.
>
> Given the daily requirements of each kind of vitamin that a cow needs, identify the smallest combination of scoops of feed a cow can be fed in order to meet at least the minimum vitamin requirements.

+ 思路1：DFS，每次选或者不选。用一个数组记录选择的方案。$O(2^n)$
+ 思路2：$O(n!)$

### 2.2.2 Subset Sums

> For many sets of consecutive integers from 1 through N (1 <= N <= 39), one can partition the set into two sets whose sums are identical

+ 思路1：DFS，一般O(2^n) 的算法n<=25可以通过。此处不能过
+ 思路2：动态规划-背包问题。`F(i,j) = F(i-1, j) + F(i-1, j-i)`，**到F(i,j)的方法数**





### 2.3.1  Longest Prefix 最长前缀 

> 如果一个集合 P 中的元素可以通过串联（元素可以重复使用，相当于 Pascal 中的 “+” 运算符）组成一个序列 S ，那么我们认为序列 S 可以分解为 P 中的元素。元素不一定要全部出现（如下例中BBC就没有出现）
>
> 序列 S 的前面 K 个字符称作 S 中长度为 K 的前缀。设计一个程序，输入一个元素集合以及一个大写字母序列 S ，设S'是序列S的最长前缀，使其可以分解为给出的集合P中的元素，求S'的长度K。





### 2.3.2 Cow Pedigrees 奶牛家谱

> 这些二叉树总共有N个节点(3 <= N < 200)。每个节点有两个子孩子或者没有孩子。树的高度等于K(1 < K < 100)。高度是从根到最远的那个叶子所需要经过的结点数。
>
> 有多少不同的家谱结构? 如果一个家谱的树结构不同于另一个的, 那么这两个家谱就是不同的。输出可能的家谱树的个数除以9901的余数。



### 2.3.3 Zero Sum 

> 考虑一个由1到N（N属于[3,9]）的数字组成的递增数列：1 2 3 ... N。 在数列中插入`“+”` `“-”`或`“ ”`空白，来将每一对数字组合在一起（请不要在第一个数字前插入符号）。计算该表达式的结果并判断其值是否为0。



### 2.3.4 Money System 货币系统

> 货币系统中货币的种类数目是 V (1<=V<=25)。要构造的数量钱是 N (1<= N<=10,000)。

+ 思路1：动态规划-完全背包。难点在于**没想清楚状态**…… $F(i,j) = F(i-1,j) + F(i-1, j-h[i])$ 用前 $i$ 种面值的货币构成 $j$ 值的方法数



### 2.3.5 Controlling Companies 

> 如果至少满足了以下三个条件之一，公司A就可以控制公司B了：
>
> 1、公司A = 公司B。
>
> 2、公司A拥有大于50%的公司B的股票。
>
> 3、公司A控制K(K >= 1)个公司，记为C1, ..., CK，每个公司Ci拥有xi%的公司B的股票，并且x1+ .... + xK > 50%。
>
> 已知公司i享有公司j的p%的股票。计算所有的数对(h,s)，表明公司h控制公司s。至多有100个公司。

+ 思路：搜索。设 `a[i][j]`为 `i `公司控制 j 公司的股份,`con[i][j]`表示 `i `公司是否控制` j `公司。首先枚举每个 `i,j `公司,对于每个` i `能控制的公司 `j`，则要继续枚举`j`占有股份的公司 `k`,计算其股份和，更新 `a[i][k]`,然后添加` i` 控制` j`,而每次添加新的控制关系则要再更新一次所有关系。时间复杂度小于 `O(n^4)`