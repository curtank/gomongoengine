# gomongoengine
ODM model like mongoengine for go 
## proposal
basic operation
### save
```go
type A struct{
  Name string
}
database.Save(&A{})
```
### fill
```go
type A struct{
  Name string
  ID int
}
a:=A{Name:"bill"}
database.Fill(&a)
//will find one A with name bill and use it to fill a
```
### find
```go
type A struct{
  Name string
  ID int
}
a:=A{Name:"bill"}
as:=[]*A{}
database.Find(&a,as)
//will find all A with name bill and use it to fill as
```
### declare unique with struct tags
```go
type A struct{
  Name string "index:"unique"
  ID int
}
database.Find(&A{Name:"bill"})
//will find all A with name bill, zero or one
//also make sure a unique index in mongodb
```
### use struct name generate collection name
```go
type A struct{
  Name string
}
database.Save(&A{})

func collectionname(name string){
  return string.LowerCase(name)
}
```
the collection in database will be a,and the colletionname func is not visable for the beginer,we provide a basic func 
which transform "UserName" to "user_name" collection name 
### fuzzy search with struct tag
```go
type A struct{
  Name string `index:"unique",search:"fuzzy"`
  ID int
}
database.Find(&A{Name:"bill"})
//will find all A with name contain bill, zero or one
```
