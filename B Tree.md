# B Tree

B树是一个多路平衡查找树

T.root 性质
1. x 为一个节点
	a. x.n 为关键字个数
	b. x.key1 <= x.key2 <= ... <= x.key(x.n)
	c. x.leaf 是否是页节点
2. x.ci 为孩子的指针 x.c1 x.c2 ... x.c(x.n+1) 叶节点没有孩子，x.ci 没有定义 null
3. 对任意的ki属于x.ci, 满足
	k1 <= x.key1 <= k2 <= x.key2 <= ... <= k(x.n) <= x.key(x.n) <= k(x.n+1)
4. 每个页节点具有相同的高度 h
5. t
	a. 最少有 t 个孩子，t-1 个关键字
	b. 最多有 2t 个孩子，2t-1 个关键字
