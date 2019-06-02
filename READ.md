
## page and node

```
const (
	branchPageFlag   = 0x01
	leafPageFlag     = 0x02
	metaPageFlag     = 0x04
	freelistPageFlag = 0x10
)

branchPageElement
leafPageElement
meta
func (f *freelist) write(p *page) error


type page struct {
	id       pgid
	flags    uint16
	count    uint16
	overflow uint32
	ptr      uintptr
}

  pgid flags count overflow ptr
|--------|--|--|----|-------------


// leaf

type leafPageElement struct {
	flags uint32
	pos   uint32
	ksize uint32
	vsize uint32
}

flags pos ksize vsize
|----|----|----|----|
|----|----|----|----|
|----|----|----|----|

// branch

type branchPageElement struct {
	pos   uint32
	ksize uint32
	pgid  pgid
}

 pos  ksize  pgid
|----|----|--------|
|----|----|--------|
|----|----|--------|
|----|----|--------|

// meta

type meta struct {
	magic    uint32
	version  uint32
	pageSize uint32
	flags    uint32
	root     bucket
	freelist pgid
	pgid     pgid
	txid     txid
	checksum uint64
}

// freelist

pgid,pgid,pgid

```






## old

事务和并发性没有关系

db.Batch


UBC: https://www.usenix.org/legacy/publications/library/proceedings/usenix2000/freenix/full_papers/silvers/silvers_html/index.html
BSD


## Question

- bolt_linux.go 为什么linux自定义实现了 fdatasync，不能使用 db.file.Sync() ？
    - 另外，怎么发现上面问题。意思是，怎么在其它地方能够了解到这么细节的不同。linux 和其它操作系统
    - 细致了解系统实现
    - 测试。但是一般只有知道差异，才会去测试
    - 提高警惕。在写这种基础组件的时候，当涉及系统调用的时候，要记得各个系统的差异性
- 现象解释：随着 bench count 的提高 从 1000 - 100000，write op/s 一直在降低。
    - 规律是什么？画曲线
    - 原因是什么？
    - 得到哪些收获？

## conclusion

- 写存储，一定要保证测试完整。（这里可以采用TDD）因为，不写测试很难保证正确性
- 错误日志，要留存。异常情况特别多。包括系统之间的差异


## read

page 序列化后的结构，内存对应文件

node 序列化前的结构，对应实际上层接口（逻辑实现）


## link

https://lrita.github.io/categories/#bolt

http://www.nextblockchain.top/categories/boltdb

