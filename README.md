# Lang LK
改编自Lua5.3，[luago](https://github.com/zxh0/luago-book)

## 速览
**详细语法**，可以查看[test](test)文件夹的内容

```js
// 发送请求
shy _, resp = http.post(
    'http://httpbin.org/post', 
    {'accept': 'application/json'}, 
    '{"foo": "bar"}'
)
print(resp)

// json解析
if json.get(resp, 'json.foo') != 'bar' {
    error('mismatch result')
}

// 设置metatable
shy headers = {}
headers.__str = fn(a) {
    shy s = ''
    for k, v in a {
        shy ss = ''
        for _, vv in v {
            ss = ss .. vv .. ';'
        }
        s = s .. k .. ': ' .. ss .. '\n'
    }
    rt s
}


/*
处理监听事件
`req`包含属性`method`, `url`, `body`, `headers`
*/
shy fn handle(req) {
    setmetatable(req.headers, headers)
    rt 200, fmt('%s %s\n\n%s\n%s', req.method, req.url, req.headers, req.body)
}

// 监听
if http.listen(':8080', handle) != nil {
    error(err)
}
```

## CLI
```bash
# 编译test/basic.lk，输出到test/basic.lkc
./go-lang-lk -c test/basic.lk
# 运行test/basic.lkc
./go-lang-lk test/basic.lkc
# 也可以运行test/basic.lk（内部会先进行编译）
./go-lang-lk test/basic.lk
```

## TODO
- 语法
  - [x] 注释`//` `/* */`
  - [x] 去除`repeat` `until`
  - [x] Raw String, `\``
  - [x] 支持任意对象 Concat
- 编译器
  - [x] 自动添加`range` ( `paris` )
- Table
  - [ ] 索引从0开始 
  - [x] key为StringExp，而不是NameExp
  - [x] `=` -> `:`, eg: `{a = 'a'}` -> `{a: 'a'}`
- CLI
  - [x] 利用HASH，如果文件内容没变化，就不需要重新编译