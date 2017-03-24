# learn go

- go知识归纳
 - go基本概念
   - GOPATH 配置工作目录
   - 常用基本命令
       - go build 编译包
            - 参数 -o output指定输出名称，代替默认包名
            - 参数 -i 安装相应的包，编译+go install
       - go install 安装并生成pkg、bin目录下的文件
       - go run  执行程序
       - go get 动态获取远程代码包
            - -d 只下载不安装
            - -u 强制使用网络去更新包和它的依赖包
            - -v 显示执行的命令
       - go 
   - 工作目录下结构
       - src 存放源代码，多项目可以分多个子目录
       - pkg 编译后生产的文件 
       - bin 编译后生产的可执行文件

- [go代码规范](http://colobu.com/2017/02/07/write-idiomatic-golang-codes/?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)
   - 注释
       - /** .....  **/ 注释片段、包名解释等
       - // 某段说明. 开头留一个空格
   - 字符串
       - 缩写词保持一致,例如URL
       - 首字母小写: findById
       - 常量首字母大写,驼峰式: MaxLength
       - 错误字符串开头小写: fmt.Errorf("failed to write data")
       - 声明空的slice
       <pre><code>
       正确:
            var t []string
       长度为0的slice:
            t := []string{}
       </code></pre>
       - 报名称使用单数 例如model而不是models
   - 代码片段
     - 包含 分割字符串
     <pre><code>
     str := "223,344,"
	 fmt.Println(strings.Contains(str,","))
	 s := strings.Split(str,",")
	 for j := 0; j< len(s) ; j++  {
		fmt.Println(s[j])
	 }
     </code></pre>