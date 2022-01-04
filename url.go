package main

//法一：
//type entry struct {
//	key ,value int
//}
//
//type LRUCache struct {
//	cap int
//	cache map[int]*list.Element
//	lst *list.List
//}
//
//
//func Constructor1 (capacity int) LRUCache{
//	return LRUCache{capacity, map[int]*list.Element{},list.New()}
//}
//
//func (c *LRUCache) Get(key int) int{
//	//先考察key在不在map里面
//	e :=c.cache[key]
//	//不在的情况返回-1
//	if e ==nil {
//		return -1
//	}
//	//MoveToFront将元素e移动到链表的第一个位置，如果e不是l的元素，l不会被修改。
//	c.lst.MoveToFront(e)
//	return e.Value.(entry).value
//}
//
//func (c *LRUCache) Put(key,value int) {
//	if e :=c.cache[key];e!=nil{
//		e.Value =entry{key,value}
//		c.lst.MoveToFront(e)
//		return
//	}
//	c.cache[key] =c.lst.PushFront(entry{key,value})
//	if len(c.cache) >c.cap{
//		delete(c.cache,c.lst.Remove(c.lst.Back()).(entry).key)
//	}
//
//}


//法二：
//定义双链表

//https://leetcode-cn.com/problems/lru-cache/

/*我们定义一个 LinkNode ，用以存储元素。因为是双向链表，
自然我们要定义 pre 和 next。同时，我们需要存储下元素的 key 和 value 。
val大家应该都能理解，关键是为什么需要存储 key？举个例子，比如当整个cache 的元素满了，此时我们需要删除 map 中的数据，
需要通过 LinkNode 中的 key 来进行查询，否则无法获取到 key。
*/
type  LinkNode struct {
	key,val int
	pre,next *LinkNode
}
/*自然需要一个 Cache 来存储所有的 Node。我们定义 cap 为 cache 的长度，m 用来存储元素。head 和 tail 作为 Cache 的首尾。*/

type LRUCache struct {
	m  map[int]*LinkNode
	cap  int
	//记录连表的头和尾
	head,tail *LinkNode
}

func Construct1(capacity int) LRUCache{
  head :=&LinkNode{0,0,nil,nil}
  tail :=&LinkNode{0,0,nil,nil}
  head.next =tail
  tail.pre =head
  return LRUCache{
  	make(map[int]*LinkNode),
  	capacity,
  	head,
  	tail,
  }
}

//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
func (this *LRUCache) Get(key int) int {

	cache :=this.m
	//查找，如果存在就就将其放在最前面
	if v,exist :=cache[key];exist{
		this.MoveToHead(v)
		return v.val
	}else{
		//找不到就返回-1
		return -1
	}

}


//void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
//当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

func (this *LRUCache) Put(key int,value int){
	tail :=this.tail
	cache :=this.m
	//判断存的这个key在map里面有没有
	if v,exist :=cache[key];exist{
		//如果有，变更其数据值
		v.val =value
		//如果有，就把他放到双向链表最前面
		this.MoveToHead(v)
	}else{
		//如果关键字不存在，则插入数组
		//v就相当于初始化了一些节点，将key，value存进去
		v :=&LinkNode{key,value,nil,nil}
		//当缓存池长度达到上限
		if len(cache) == this.cap{
			//删除缓存池的最久未使用的数据末尾的（可以这么理解，head节点，tail节点都是不存数据的，所以删或者加的时候都是在head之后，tail之前）
			//缓存删的是key注意
			delete(cache,tail.pre.key)
			//删除节点
			this.RemoveNode(tail.pre)
		}
		this.AddNode(v)
		cache[key] =v
	}
}

//双链表的删除节点(删的是尾节点)
func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//双链表的添加节点（添加节点也是添加到head节点之后）

func (this *LRUCache) AddNode(node *LinkNode) {
	//这个顺序也要注意
	head := this.head
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}

//将元素移动到双链表的最前面节点
func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}





/*
146. LRU 缓存机制

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v, exist := cache[key]; exist {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) AddNode(node *LinkNode) {
	head := this.head
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) Put(key int, value int) {
	tail := this.tail
	cache := this.m
	if v, exist := cache[key]; exist {
		v.val = value
		this.MoveToHead(v)
	} else {
		v := &LinkNode{key, value, nil, nil}
		if len(cache) == this.cap {
			delete(cache, tail.pre.key)
			this.RemoveNode(tail.pre)
		}
		this.AddNode(v)
		cache[key] = v
	}
}

作者：ivan1
链接：https://leetcode-cn.com/problems/lru-cache/solution/hashmap-shuang-xiang-lian-biao-chao-ji-xiang-xi-by/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。













*/