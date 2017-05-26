# mongodb在golang的应用

## 定义结构体
    type User struct {
		Username string "user_name"
		Age      int    "age"
	}
	users := []User{}

## 查询 find

   - 查找一个

    selecter := bson.M{"_id":"test"}
    err = s.DB("test").C("counters").Find(selecter).One(&result)
    
   - 根据objectId查找
   
    err = s.DB("test").C("users").FindId("58d9024569ad41655482e5df").One(&users)
    
   - 查找多个
   
    selecter = bson.M{"user_name":"ggg"}
    iter := s.DB("test").C("users").Find(selecter).Sort("_id").Iter()
    for iter.Next(&users) {
        fmt.Printf("users name: %v age: %v\n",users.Username,users.Age)
    }
    
   - 单条件查找 等于=
   
    selecter := bson.M{"user_name":"ggg"}
    err = s.DB("test").C("users").Find(selecter).All(&users)
    fmt.Printf("%v", users)
    
   - 单条件查找 不等于!= $ne
   
    selecter := bson.M{"user_name":bson.M{"$ne":"ggg"}}
    
   - 单条件查找 >($gt)  >=($gte)  <=($lte) <(lt)
   
    selecter := bson.M{"age":bson.M{"$gte":"10"}}    
    
   - 单条件查找 in
   
    selecter := bson.M{"user_name":bson.M{"$in":[]{"ggg","ggg1"}}}
    
   - 多条件查询 and
   
    selecter := bson.M{"user_name":"ggg","age":20}
    
   - 多条件查询 or
   
    selecter := bson.M{"$or":[]bson.M{bson.M{"user_name":"ggg"},bson.M{"age":20}}}
    
   - 分组查询
   
    match := bson.M{"user_name" : "ggg"}
    group := bson.M{
            "_id":  nil,
            "user_name" :bson.M{"$last":"$type"},
            "age" :bson.M{"$last":"$type"},
            "total_num" :bson.M{"$sum":1},
            "avg_age" :bson.M{"$avg":"$age"},
        }
    sort :=  bson.M{
            "age": -1,
        }
    user := User{}
    pipeline := []bson.M{{"$match":match},{"$group":group},{"$sort":sort}}
    iter := s.DB("test").C("users").Pipe(pipeline).Iter()
    for iter.Next(&user) {
        fmt.Printf("users name: %v age: %v\n",user.Username,user.AvgAge)
    }

    管道增加group可对分组后的数据统计
    
## 新增 insert   
    selecter := bson.M{"user_name":"ggg","age":20}
    err = s.DB("test").C("users").Insert(selecter)
    
## 更新 update 

   - 更新（$set）增加数量($inc) 数组增加一个元素($push) 数组删除一个元素($pull)
  
    selecter = bson.M{"user_name":"ggg"}
    update := bson.M{"$set":bson.M{"age":21}}
    err = s.DB("test").C("users").Update(selecter, update)
      
## 删除 delete
    selecter = bson.M{"user_name":"ggg1"}
    info, err := s.DB("test").C("users").RemoveAll(selecter)

## 简单查询操作符（非聚合操作）
   - (1）$eq : 用来等值条件过滤某一个key的值。
    用法示例（过滤某个key等于某个值，可以用 $eq)
    db.op_test.find({"name":{$eq:"steven"}})
   - 2）$gt : 用来判断某个key值大于某个指定的值。

     用法示例（过滤某个key值，大于某个指定的值）
    db.op_test.find({"age":{$gt:19}})
   - 3）$gt : 用来判断某个key值大于等于某个指定的值。

    用法示例（过滤某个key值，大于某个指定的值）
    db.op_test.find({"age":{$gte:19}})
   - 4）$lt : 用来判断某个key值小于某个指定的值。

    用法示例（过滤某个key值，小于某个指定的值）
    db.op_test.find({"age":{$lt:20}})
   - 5）$lte : 用来判断某个key值小于等于某个指定的值。

    用法示例（过滤某个key值，小于等于某个指定的值）
    db.op_test.find({"age":{$lte:20}})
   - 6）$ne : 用来不等值条件过滤某一个key的值。

    用法示例（过滤某个key不等于某个值，可以用 $ne)
    db.op_test.find({"name":{$ne:"steven"}})
   - 7）$in : 用来指定某个key的值在指定的离散的值域内

    用法示例（过滤某个key的值是否符合条件，和$or的区别就是，or可以组合不同key的不同条件,而$in则只的是单一条件）
    db.op_test.find({"name":{$in:["steven","jack"]}})
   - 8）$nin : 和$in相对，用来指定key值不存在某个指定的离散值域内。

    用法示例（和$in的用法相反）
    db.op_test.find({"name":{$nin:["steven","jack"]}})
    逻辑（Logic）运算相关
   - 1） $or : 任意组合不同的查询条件（可以针对任意key的限制条件），只要满足任意组合条件中的一个即可。
    用法示例（返回 name 为 steven 或者 age 等于 20 的文档）：
    db.op_test.find({"$or" : [{"name":"steven"},{"age":20}]})
   - 2） $and: 和$or操作符相对，任意组合不同的查询条件（可以针对任意key的限制条件），并且必须同时满足所有条件。

    用法示例（返回 name 为 steven 并且 age 等于 20 的文档）：
    db.op_test.find({"$and" : [{"name":"steven"},{"age":20}]})
   - 3） $not： 元条件语句，需要和其他条件语句组合使用。

    用法示例（$not 和 $lt 操作符组合使用，返回 age 大于等于20的文档）：
    db.op_test.find({"age":{"$not":{"$lt":20}}})
   - 4） $nor：和$or相反，表示所有条件均不能满足则返回。

    用法示例（凡是 name 为 steven 或者 age 为 20 的全部过滤掉）：
    db.op_test.find({"$nor" : [{"name":"steven"},{"age":20}]})
    元素(Element)运算相关
   - 1） $exists: 查询不包含某一个属性（key）的文档。

    用法实例（凡是包含name这个key的文档全部返回）
    db.op_test.find({"name":{"$exists":true}})

    用法实例（凡是不包含name这个key的文档全部返回）
    db.op_test.find({"name":{"$exists":false}})

    ps：true和false的区别就是判断是否包含这个key
   - 2） $type : 过滤某个字段是某一个BSON数据类型的数据。

    用法示例（返回所有name字段为String类型的所有文档）
    db.op_test.find({"name":{"$type":2}})
    ps：name后面的数字具体查询列表参见：http://docs.mongodb.org/manual/reference/operator/query/type/#op._S_type
    求值（Evaluation）操作相关
   - 1） $mod : 取余操作符，筛选经过区域操作后，结果符合条件的文档。

    用法示例(返回age的值和 4 求余后 结果为 0 的数据)
    db.op_test.find({"age" : {"$mod" : [4,0]}})
   - 2） $regex : 筛选值满足正则表达式的文档。

    用法示例（返回 name 符合指定正则的数据，option选项限定正则的形式）
    db.op_test.find({"name" : {$regex:"stev*",$options:"i"}})
    ps：options相关参见：http://docs.mongodb.org/manual/reference/operator/query/regex/#op._S_regex
   - 3） $text: 针对建立了全文索引的字段，实施全文检索匹配。

    用法示例（针对构建全文索引的字段进行搜索，默认为英文，不支持中文）
    db.op_test.find({"$text":{$search:"steven",$language:"en"}})
    ps：目前支持的语言以及缩写，参见：http://docs.mongodb.org/manual/reference/text-search-languages/
   - 4） $where: 强大的查询关键字，但性能较差，可以传入js表达式或js函数来筛选数据。

    用法示例（返回满足传入的js函数的文档，函数表示文档中只要任意字段的值为"steven"则返回）
    db.op_test.find({"$where":function(){
        for(var index in this) {
            if(this[index] == "steven") {
                return true;
            }
        }
        return false;
    }})
    数组（Array）相关操作
    $all : 数组查询操作符，查询条件是一个数组，被查询的字段也是一个数组，要求被查询的数组类型的字段要是查询条件数组的超集（即大数组要包含小数组）。

    用法示例：(查询key value对应的数组值，要同时包含"a","b","c"三个元素)
    db.op_test.find({"values":{$all:["a","b"]}})
    $elemMatch : 数组查询操作符，用来指定数组的每一个元素同时满足所罗列的条件，如果不指定，则条件会是或的关系

    用法示例：(要匹配 values数组中，至少有一个元素，满足所有的条件)
    用于指定嵌套文档操作，具体事例参见：http://docs.mongodb.org/manual/reference/operator/projection/elemMatch/
    $size: 用于某个数组类型的key对应值的数量满足要求。

    用法示例：筛选出来包含数组元素个数为3的文档。
    db.op_test.find({"values":{$size : 3}})
    评论（Comments）相关操作
    $comment: 在查询、更新或其他操作执行过程中，可以通过添加$comment操作符添加评论。改评论会被记录在日志中，用于后续分析。

    用法示例
    db.collection.find( { <query>, $comment: <comment> } )
    地理位置（Geospatial）相关操作
    $geoWithin: 这个操作符基于2d 空间索引，首先要针对文档的某个字段建立一个2d的空间索引，然后利用此操作符，可以在一个2d空间范围内指定一个多变形，$geoWithin操作符就是查询出包含在多变形范围内的点。

    详见：http://docs.mongodb.org/manual/reference/operator/query/geoWithin/#op._S_geoWithin
    $geoIntersects: 同样基于2d空间索引，计算当前的空间范围和指定的geo多变形取交集。

    详见：http://docs.mongodb.org/manual/reference/operator/query/geoIntersects/#op._S_geoIntersects
    $near ：基于2d空间索引，指定一个点，返回该点有近及远的所有的点。

    详见：http://docs.mongodb.org/manual/reference/operator/query/near/#op._S_near
    $nearSphere: 基于2d空间索引，指定一个点，由近及远的返回所有的点，和$near操作符不同的是计算距离的方式 $nearSphere计算的是球面距离。$near计算的是坐标距离。

    投影相关操作
    $ : 对你们看错，就只是一个$操作符，如果文档中某个value是数组类型，通过 $ 操作符可以指定数组字段的投影，返回数组字段中第一个匹配的那个元素，相当于截断了原来的整个数组，只返回第一个值。

    用法示例：(会返回values数组中，第一个和"a"相等的元素，也就是返回"a")
    db.op_test.find({"values":{$eq:"a"}},{"values.$":1})
    返回结果如下：
    { "_id" : ObjectId("551117417cbfa0a55db5c5b9"), "values" : [ "a" ] }
    $elemMatch : 这个操作符上面数组操作有涉及，其另外一个效果就是，在嵌套文档的应用中，返回数组中第一个符合条件的文档，可以限定多种组合条件。

    详见：http://docs.mongodb.org/manual/reference/operator/projection/elemMatch/#proj._S_elemMatch
    $meta ： 和全文索引 text index 组合使用，针对一个带有全文索引的元素，指定改操作符，可以返回和查询条件相似的分数，分数越高，匹配度越高。

    用法示例：
    db.op_test.find({"$text":{$search:"steven",$language:"en"}},{score:{$meta:"textScore"}})
    执行结果：
    { "_id" : ObjectId("550fdba3c118f1b20bd51a9f"), "name" : "steven", "age" : 20, "score" : 1.1 }
    $slice : 数组类型字段的投影操作，返回原来数据的一个子集.针对一个数组，其有如下几种返回子集的方式：

    用法示例： 返回博客的前10条评论
    db.blog.find({"comments":{"$slice":10}})

    用法示例： 返回博客的后10条评论
    db.blog.find({"comments":{"$slice":10}})

    用法示例： 返回博客跳过前10条，然后返回第11 ~ 15条
    db.blog.find({"comments":{"$slice":[10,5]}})

