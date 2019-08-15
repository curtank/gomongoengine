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
### find
```go
type A struct{
  Name string
  ID int
}
database.Find(&A{Name:"bill"})
//will find all A with name bill
```
### declare unique with struct tags
```go
type A struct{
  Name string "index:"unique"
  ID int
}
database.Find(&A{Name:"bill"})
//will find all A with name bill, zero or one
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
the collection in database will be a
