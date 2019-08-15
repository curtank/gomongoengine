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
### unique with struct tags
```go
type A struct{
  Name string "index:"unique"
  ID int
}
database.Find(&A{Name:"bill"})
//will find all A with name bill, zero or one
```
