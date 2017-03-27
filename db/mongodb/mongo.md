# mongodb在golang的使用方式

## 查询 find
    - 查找一个
    selecter := bson.M{"_id":"test"}
    s.DB("test").C("counters").Find(selecter).One(&result)
    
    - 查找多个
    selecter = bson.M{"user_name":"ggg"}
    iter := s.DB("test").C("users").Find(selecter).Sort("_id").Iter()
    for iter.Next(&users) {
        fmt.Printf("users name: %v age: %v\n",users.Username,users.Age)
    }
      
## 新增 insert   
    selecter := bson.M{"user_name":"ggg","age":20}
    err = s.DB("test").C("users").Insert(selecter)
    
## 更新 update 
    selecter = bson.M{"user_name":"ggg"}
    update := bson.M{"$set":bson.M{"age":21}}
    err = s.DB("test").C("users").Update(selecter, update)
      
## 删除 delete
    selecter = bson.M{"user_name":"ggg1"}
    info, err := s.DB("test").C("users").RemoveAll(selecter)
    
