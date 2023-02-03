# switch
> vela 中自定义的switch 语句结构

## 内置方法
- [vela.switch(prefix , method)](#switch)&emsp;switch构造

## switch
> switch = vela.switch(prefix , method) <br />
> prefix: 默认的字段前缀 如:field  method:匹配方法 默认:=

内置方法:
- switch._              &emsp;初始化方法，如果switch 不为空的话 那么会自动拼接方法
- switch.once(bool)     &emsp;是否只匹配中一个就结束
- switch.case(cnd)      &emsp;添加复杂条件和处理对象
- switch.default(v...)  &emsp;添加默认方法
- switch.match(value)   &emsp;开始匹配对象

```lua
    local s = vela.switch("name")
    local s = vela.switch("age" , "=")
    
    s._{
        ["name = vela.com"] = print
    }

    s.once(true)
    s.case("name eq baidu.com"  , "age > 188").pipe()
    s.case("name eq google.com" , "age > 188").pipe()
    s.case("name eq baidu.com" , "age > 188").pipe()
    s.case("name eq baidu.com" , "age > 188").pipe()
    s.case("name eq baidu.com" , "age > 188").pipe()
    s.case("name eq baidu.com" , "age > 188").pipe()
    s.default(pipe1 , pipe2 , pipe3)

    local app = {
        name = "baidu.com",
        age  = 15,
    }
    s.match(app)
```

## case
> case = switch.case(cnd) 返回的是case的对象 提供的是一种匹配场景

内置方法:
- pipe(v...) &emsp;命中后如何操作
- over() &emsp;命中这条结束switch 类似break


```lua
    local switch = vela.switch()
    
    local case = switch.case("name = 123")
    case.pipe("123123")
    case.over()

    --简写
    switch.case("name = 456").pipe(function(v , case_id , cnd) end)        -- v match 的对象; case_id:条件编号 cnd:条件内容
    switch.case("name = 123").pipe(function(v , case_id , cnd) end).over() -- 如果匹配跳出switch
    switch.case("name = 789").pipe(function(v , case_id , cnd) end)

```
