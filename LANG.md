# Lang LK
## 第一行代码
```js
print("Hello World!")  // 我是注释
```
以上内容在屏幕上打印出 `Hello World!`。  


## 基本类型
```js
a = 1      // num
b = '1'    // str
a = false  // bool
a = {}     // table
c = nil    // nil
/* 我是多行注释 */
```
`lk` 中的基本类型有`num` `str` `bool` `nil` `table`。

```js
print(a == b) // 因为 num != str，所以 false
```
如上所示，虽然 `lk` 不需要明确指出变量的类型，但不意味着是弱类型语言。  
不同类型间比较，不会先转化为同类型。

```js
shy a = `
😊`
print(a)  // 😊
```
`shy` 关键字 表示 这是一个局部变量，与 `Lua` 中的 `local` 作用一致。
`str` 除了可以用 `'` `"` 包裹，还可以用 `` ` `` 包裹（ 表示这是个 `Raw String` ），这样可以避免被转义。   
⚠️ 如果使用 `Raw String` 构造字符，且第一个字符为换行 ( `\n` )，**这第一个**换行会被忽略（如上的变量 `a` 声明）。

```js
shy tb = {
    'a': 1,
    'b': 2,
}

// `:=` 表示该变量为私有，等同于 `shy a = 'b'`
a := 'b'

print(tb[a])    // 2
print(tb.a)     // 1
print(tb['a'])  // 1
```
请注意，TableAccess 有两种方式，一种是 `tb[a]`，另一种是 `tb.a`。
- `tb[a]` 会先计算 `a` 的值，然后再去访问 `tb` 中的值。
- `tb.a` 会直接去访问 `tb` 中 `key` 为 字符`a` 的值，不会计算变量 `a` 的值。


## 变量
```js
a = 1
if true {
    a := 3
    print(a)  // 3
}
print(a)  // 1
```
`shy` & `:=` 表示该变量为局部变量，只在当前作用域内有效。  
`a := 1` 等同于 `shy a = 1`。  
请**尽量**声明为私有变量，这会提高程序运行速度（因为不需要全局寻找变量）。  
`lk` 中已经声明的变量可以再次声明，这时会在其作用域内覆盖原来的值。  
 

```js
a, b = 1
print(a, b)  // 1 nil
```
`lk` 中可以使用 `,` 分隔多个变量，这时会将多个变量的值赋值给左边的变量。  
如果右边的变量个数少于左边的变量个数，那么多余的变量会被赋值为 `nil`。  


## 函数
```js
shy fn add2(a, b) {
    rt a + b
}
add2 := fn(a, b) {
    rt a + b
}
print(add2(1, 2))  // 3
```
和变量一致，`shy` 表示局部函数。  
以上两种声明作用一致，支持以变量方式声明函数。  

```js
fn addN(...) {
    sum = 0
    for _, i in {...} {
        sum += i
    }
    rt sum
}
addN(1, 2, 3, 4, 5)  // 15
```
`...` 为变长参数，表明0个或更多个参数。  
可以使用 `{...}` 来构造参数列表，再使用 `for in` 获取每一个参数。  

```js
a := fn(b) => 3 ^ b, 2 ^ b
print(a(2))

shy a = fn(b) {rt 3 ^ b, 2 ^ b}
print(a(2))
```
两个 `a` 函数声明的作用一致。  
`=>` 后返回值只能有一行。

## 循环
以下循环都支持 `break` 关键字。
```js
while condition {
    // ...
}
```
```js
for i = 0, 10 {
    // ...
}
```
等同于 `for i = 0; i <= 10; i++ {}`
```js
for i = 0, 10, 2 {
    // ...
}
```
等同于 `for i = 0; i <= 10; i += 2 {}`

## 流程控制
```js
if condition {
    // ...
} elif condition {
    // ...
} else {
    // ...
}

// 在 `if` 判断时，只有 `nil` 和 `false` 会被判断为 `false`
if '' and {} and 0 {
    if nil or false {
        print('never print')
    } else {
        print('only `nil` and `false` is false')
    }
}
```

## 运算
### 算术运算符
```js
a = 21
b = 10
c = a + b
print("Line 1 - c 的值为 ", c)
c = a - b
print("Line 2 - c 的值为 ", c)
c = a * b
print("Line 3 - c 的值为 ", c)
c = a / b
print("Line 4 - c 的值为 ", c)
c = a % b
print("Line 5 - c 的值为 ", c)
c = a ^ 2
print("Line 6 - c 的值为 ", c)
c = -a
print("Line 7 - c 的值为 ", c)
```
输出：
```
Line 1 - c 的值为     31
Line 2 - c 的值为     11
Line 3 - c 的值为     210
Line 4 - c 的值为     2.1
Line 5 - c 的值为     1
Line 6 - c 的值为     441
Line 7 - c 的值为     -21
```
同时，也支持 `a++` `a+=1` 等
### 关系运算符
```js
a = 21
b = 10

if (a == b) {
   print("Line 1 - a 等于 b")
else {
   print("Line 1 - a 不等于 b")
}

if (a != b) {
   print("Line 2 - a 不等于 b")
else {
   print("Line 2 - a 等于 b")
}

if (a < b) {
   print("Line 3 - a 小于 b")
else {
   print("Line 3 - a 大于等于 b")
}

if (a > b) {
   print("Line 4 - a 大于 b")
else {
   print("Line 5 - a 小于等于 b")
}

// 修改 a 和 b 的值
a = 5
b = 20
if (a <= b) {
   print("Line 5 - a 小于等于 b")
}

if (b >= a) {
   print("Line 6 - b 大于等于 a")
}
```
输出：
```
Line 1 - a 不等于 b
Line 2 - a 不等于 b
Line 3 - a 大于等于 b
Line 4 - a 大于 b
Line 5 - a 小于等于 b
Line 6 - b 大于等于 a
```

### 逻辑运算符
```js
print(false and true)  // false
print(false or true)   // true
print(not false)       // true
// `not false` 等于 `!false`
```

### 其他运算符
```js
a = "Hello "
b = "World"

// 连接字符串 a 和 b
print(a + b)  // Hello World
// b 字符串长度
print(#b)   // 5

// `#`获取长度，`...`为变长参数，`{}`构造Table
// `#{...}`即获取变长参数的长度（有多少个参数）
fn varagsLen(...) => print(#{...})

varagsLen(1, 2, 3, 4, 5)  // 5

// 三元操作符
print(true ? 'support ternary exp' : 'unreachable')
```

### 运算符优先级
```js
a = 20
b = 10
c = 15
d = 5

e = (a + b) * c / d  // (30 * 15) / 5
print(e)  // 90.0

e = ((a + b) * c) / d  // (30 * 15) / 5
print(e)  // 90.0

e = (a + b) * (c / d)  // (30) * (15 / 5)
print(e)  // 90.0

e = a + (b * c) / d  // 20 + (150 / 5)
print(e)  // 50.0
```

## 迭代器
### 默认迭代
```js
a := {'num': 1, 'str': '1', 'bool': false, 'nil': nil}
for k, v in a {
    print(k, v)
}
```
其中 `for k, v in a` 就创建了一个迭代器， 
当 `a` 是 `table` 时，编译器会使用内置的迭代器，在每次迭代时为 `k` 和 `v` 分别赋值为 `a` 的键和值。 

### 自定义迭代器
#### 无状态迭代器
```js
fn square(iteratorMaxCount, currentNumber) {
    if currentNumber < iteratorMaxCount {
        currentNumber = currentNumber + 1
        rt currentNumber, currentNumber * currentNumber
    }
}

for i, n in square, 3, 0 {
   print(i, n)
}
```
这样就实现了一个简单的平方迭代器，输出：
```
1    1
2    4
3    9
```

#### 有状态迭代器
```js
fn iter(a, i) {
    i++
    v := a[i]
    if v {
       rt i, v
    }
}
 
fn ipairs(a) {
    // lk 的 起始index 为 0，所以此处为 -1
    rt iter, a, -1
}
```
如上，实现了虚拟机内置的默认迭代器


## 面向对象 & 元表
```js
// 定义一个类，包含其默认属性值（x = 0, y = 0）
class Vector { 'x': 0, 'y': 0 }

// 创建一个 Vector 对象，调用这个方法可以在初始化对象时，为内部属性赋值
// 如果使用`new(Vector)`，则会使用默认值（x = 0, y = 0）
fn Vector.new(x, y) {
    shy v = new(Vector)
    v.x = x
    v.y = y
    rt v
}

// 为 `Vector` 设置 `__add` 元方法，实现 `+` 运算符
fn Vector.__add(v1, v2) {
    shy v = new(Vector)
    v.x = v1.x + v2.x
    v.y = v1.y + v2.y
    rt v
}

// `Object:function(...)` = `Object.function(self, ...)`
// 这里：`Vector:set(x, y)` = `Vector.set(self, x, y)`
fn Vector:set(x, y) {
    self.x = x
    self.y = y
}

// 为 `Vector` 设置 `__str` 元方法，`print` `Vector` 对象时会调用此方法
// 如果不实现此方法，会使用内置的转换为 `str` 的方法
fn Vector:__str() {
    rt 'Vector(' + str(self.x) + ', ' + str(self.y) + ')'
}

// 使用的`new(Object)`，所以使用的默认属性值
// 此时 x = 0, y = 0
shy v1 = new(Vector)
// 带值的初始化对象
// 此时 x = 3, y = 4
shy v2 = Vector.new(3, 4)
// 调用 `Vector:set(x, y)` 方法，修改v1的值
v1:set(1, 2)
shy v3 = v1 + v2
print(v3.x, v3.y)  // 4       6

// 上面实现了 `Vector:__str()` 方法，此处会调用
printf('%s + %s = %s', v1, v2, v3)  // Vector(1, 2) + Vector(3, 4) = Vector(4, 6)
``` 

以下是部分可以拓展的元方法表：  

|操作符/作用|metatable|
|-|-|
|`+`|`__add`|
|`-`|`__sub`|
|`*`|`__mul`|
|`/`|`__div`|
|`%`|`__mod`|
|`^`|`__pow`|
|`-`|`__unm`|
|`~/`|`__idiv`|
|`#`|`__len`|
|`==`|`__eq`|
|`<`|`__lt`|
|`<=`|`__le`|
|索引|`__index`|
|新索引|`__newindex`|
|转为`str`|`__str`|
|调用方法|`__call`|
|获取名称|`__name`|
|迭代器|`__iter`|


## 包
```js
// 文件名为 mod.lk
class module {}
 
// 定义一个常量
module.constant = "这是一个常量"
 
// 定义一个函数
fn module.func1() {
    print("这是一个公有函数！\n")
}
 
shy fn func2() {
    print("这是一个私有函数！")
}
 
fn module.func3() {
    func2()
}

rt module
```
如上定义了一个包，然后在另一个文件中导入：
```js
import "mod"
```
可以通过 `import` 关键字导入包，导入的包会在当前文件作用域中有效。  
导入路径 `mod` 为当前文件的相对路径。  
例如`import "a/b/c"`，会尝试导入：`./a/b/c.lk` `./a/b/c/init.lk`。  

导入后如下使用：
```js
module.func1()
// module.func2() 不可直接使用，因为是局部函数，但可以通过 module.func3() 调用
module.func3()
```

```js
// 设置别名，方便调用
m := import('mod')
m.func1()
```
需要注意，`class module` 在最后 `rt module`，如果不 `rt`，则导入时无法设置别名。

## 协程
```js
fn foo(a) {
    print("foo 函数输出", a)
    rt sync.yield(2 * a) // 返回 2*a 的值
}
 
co = sync.create(fn (a , b) {
    print("第一次协同程序执行输出", a, b) // co-body 1 10
    shy r = foo(a + 1)
     
    print("第二次协同程序执行输出", r)
    shy r, s = sync.yield(a + b, a - b)  // a，b的值为第一次调用协同程序时传入
     
    print("第三次协同程序执行输出", r, s)
    rt b, "结束协同程序"  // b的值为第二次调用协同程序时传入
})
       
print("main", sync.resume(co, 1, 10)) // true, 4
print()
print("main", sync.resume(co, "r")) // true 11 -9
print()
print("main", sync.resume(co, "x", "y")) // true 10 end
print()
print("main", sync.resume(co, "x", "y")) // cannot resume dead sync
print()
```


## 标准库
`[]`代表可选参数，`...`代表可变参数。
### `os`
- `os.exit([code])`
退出程序，`code` 为退出码，默认为 0。无返回值。
- `os.exec(script)`
执行脚本，`script` 为脚本内容。
返回两个值：`output` `err`，分别为输出和错误信息。  
- `os.getenv(name)`
获取环境变量，`name` 为环境变量名。
返回值一个：`value`，为环境变量值。  
如果当前环境不存在该环境变量，则返回 `nil`。
- `os.setenv(key, value)`
设置环境变量，`key` 为环境变量名，`value` 为环境变量值。
返回值一个：`err`，如果没有错误则为 `nil`。
- `os.tmp()`
获取临时文件夹路径。无返回值。
- `os.mv(src, dst)`
移动文件或文件夹，`src` 为源文件或文件夹路径，`dst` 为目标文件或文件夹路径。
返回值一个：`err`，为错误信息。
- `os.link(src, dst)`
创建文件或文件夹的硬链接，`src` 为源文件或文件夹路径，`dst` 为目标文件或文件夹路径。
返回值一个：`err`，为错误信息。
- `os.ls(path)`
获取文件夹下的文件列表，`path` 为文件夹路径。
返回值两个：`files` `err`，为文件列表、错误信息。
- `os.mkdir(path, [rescursive])`
创建文件夹，`path` 为文件夹路径，`rescursive` 为可选参数，为是否递归创建，默认为 `false`。
返回值一个：`err`，为错误信息。
- `os.rm(path, [rescursive])`
删除文件或文件夹，`path` 为文件或文件夹路径，`rescursive` 为可选参数，为是否递归删除，默认为 `false`。
返回值一个：`err`，为错误信息。
- `os.sleep(ms)`
休眠，`ms` 为休眠时间，单位为毫秒。
- `os.time([time, isUTC])`
获取当前时间戳（毫秒），`time` 为可选参数，为时间戳，`isUTC` 为可选参数，为是否使用 UTC 时间，默认为 `false`。
返回值一个：`time`，为时间戳。
- `os.date([format, time])`
获取当前时间，`format` 为可选参数，为时间格式，`time` 为可选参数，为时间戳。
返回值一个：`date`，为时间字符串。
- `os.read(file)`
读取文件内容，`file` 为文件路径。
返回值两个：`content` `err`，为文件内容(`str`)、错误信息。
- `os.write(file, content)`
写入文件内容，`file` 为文件路径，`content` 为文件内容。
返回值一个：`err`，为错误信息。

### `re`
- `re.have(pattern, str)`
判断字符串是否匹配正则表达式，`pattern` 为正则表达式，`str` 为字符串。
返回值一个：`have`，为是否匹配。
- `re.find(pattern, str)`
匹配字符串，`pattern` 为正则表达式，`str` 为字符串。
返回值一个：`list`，为匹配结果列表。如果没有匹配结果，则返回`nil`。

### `http`
- `http.get(url, [headers])`
发送 GET 请求，`url` 为请求地址，`headers` 为可选参数，为请求头。
返回值两个：`body` `err`，为响应内容(`table`)、错误信息。
- `http.post(url, data, [headers])`
发送 POST 请求，`url` 为请求地址，`data` 为请求数据，`headers` 为可选参数，为请求头。
返回值两个：`body` `err`，为响应内容(`table`)、错误信息。
- `http.req(method, url, [headers, body])`
发送请求，`method` 为请求方法，`url` 为请求地址，`headers` 为可选参数，为请求头，`body` 为可选参数，为请求数据。
返回值两个：`body` `err`，为响应内容(`table`)、错误信息。
- `http.listen(addr, [ fn (req) ])`
监听 HTTP 请求，`addr` 为监听地址，`fn` 为可选参数，为回调函数，`req` 为请求对象。
`req`包含属性`method`、`url`、`headers`、`body`，分别为请求方法、请求地址、请求头、请求数据。
返回值一个：`err`，为错误信息。
### `json`
- `json.get(source, path)`
获取 JSON 数据，`source` 为 JSON 数据，`path` 为路径。
`path`遵循`gjson`规则。详情请查看[gjson](https://github.com/tidwall/gjson)。

### 其他
- `string` https://www.runoob.com/lua/lua-strings.html
- `utf8` https://www.jianshu.com/p/dcbb6b47bb32