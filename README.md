# Struct Interface (Day 12)

##Struct

- Tugas dari struct adalah membangun/membuat beberapa object atau properties, kuleb mirip dengan Class pada materi
Javascript sebelumnya 

contoh code dari struct :

```go
type Name struct {
    FirstName string
    LastName string
}
```
contoh code untuk penginputan data
```go
var data = []Name{
    {
    FirstName : "ucok ",
    LastName : "markocop",
    },
    {
    FirstName : "ucok2 ",
    LastName : "markocop2",
    },
}
```

## Interface

- Point dari interface adalah membangun/membuat beberapa method yang kurang lebihnya mirip seperti konsep
enkapsulasi pada javascript, yaitu membangun dari formatter atau returning

Contoh code dari Interface :

```go
type person interface {
    getFirstName()string
}

func () getFirstName() string {
    return FirstName
}
```

# PostGre (Day 13)

package use to connect golang and postgre : 
```bash
go get github.com/jackc/pgx/v4 
```

URL connection database :

`postgres://user:password@host:port/dbname`
- user = user dari database
- password = password dari database
- host = host database, secara local akan menggunakan localhost
- port = port database
- dbname = nama database