# learn go

- go知识归纳
 - go基本概念
   - GOPATH 配置工作目录
   - 工作目录下结构
       - src 存放源代码
       - pkg 编译后生产的文件
       - bin 编译后生产的可执行文件

- [go代码规范](http://colobu.com/2017/02/07/write-idiomatic-golang-codes/?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)
   - 注释
       - /** .....  **/ 注释片段、包名解释等，以句点结束
       - // .某段说明 开头留一个空格，以句点结束，
   - 字符串
       - 缩写词保持一致例如URL
       - 首字母小写：findById
       - 常量首字母大写,驼峰式： MaxLength
       - 错误字符串开头小写 fmt.Errorf("failed to write data")
       - 声明空的slice
       <pre><code>
       正确:
            var t []string
       长度为0的slice:
            t := []string{}
       </code></pre>
