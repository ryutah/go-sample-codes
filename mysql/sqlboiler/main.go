package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryutah/go-sample-codes/mysql/sqlboiler/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type myInt int

type Foo struct {
	models.Foo `boil:"foo,bind"`
	Child      models.FooChild `boil:"child,bind"`
}

func (f Foo) String() string {
	b, _ := json.MarshalIndent(f, "", "  ")
	return string(b)
}

func main() {
	boil.DebugMode = true

	ctx := context.Background()

	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/foo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	foo, err := models.Foos(qm.Where("id=?", 1)).One(ctx, db)
	if err != nil {
		log.Fatalf("%#v", err)
	}
	fmt.Printf("%v\n", foo)

	foos := make([]Foo, 0)
	if err := models.NewQuery(
		qm.Select(
			"foo.id", "foo.name",
			"child.id", "child.name",
		),
		qm.From("foo"),
		qm.InnerJoin("foo_child as child on foo.id = child.foo_id"),
	).Bind(ctx, db, &foos); err != nil {
		log.Fatalf("%#v", err)
	}
	fmt.Printf("%v\n", foos)

	children, err := models.FooChildren(
		qm.WhereIn("f.id in ?", 1, 2),
		qm.InnerJoin("foo as f on f.id = foo_child.id"),
	).All(ctx, db)
	if err != nil {
		log.Fatalf("%#v", err)
	}
	b, _ := json.MarshalIndent(children, "", "  ")
	fmt.Println(string(b))

	type barChild struct {
		FooID           myInt  `boil:"foo_id"`
		Bar             string `boil:"bar"`
		models.BarChild `boil:",bind"`
	}
	barChildren := make([]barChild, 0)
	if err := models.BarChildren(
		qm.Select("b.foo_id as foo_id, b.name as bar, bar_child.*"),
		qm.WhereIn("b.foo_id in ?", 1, 2),
		qm.InnerJoin("bar as b on b.id = bar_child.bar_id"),
	).Bind(ctx, db, &barChildren); err != nil {
		log.Fatalf("%#v", err)
	}
	b, _ = json.MarshalIndent(barChildren, "", "  ")
	fmt.Println(string(b))
}
