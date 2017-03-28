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
    
