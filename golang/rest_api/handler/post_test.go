package handler

//
//import (
//	"context"
//	"crud-api/db"
//	"crud-api/models"
//	"crud-api/repository"
//	"crud-api/repository/post"
//	"github.com/go-chi/chi"
//	"gorm.io/gorm"
//	"os"
//	"testing"
//)
//
//var (
//	d  *gorm.DB
//	h  *Handler
//	p  repository.PostRepo
//	c *chi.Router
//	p1 = models.Post{
//		ID:      1,
//		Title:   "One",
//		Content: "Two",
//	}
//)
//
//func TestMain(m *testing.M) {
//	setup()
//	code := m.Run()
//	db.TruncateTable()
//	os.Exit(code)
//
//}
//
//func setup() {
//	d = db.TestDB()
//	_ = d.AutoMigrate(d)
//	h = NewHandler(d)
//	p = post.NewSQLPostRepo(d)
//
//	loadFixtures()
//}
//
//func loadFixtures() error {
//	ctx := context.TODO()
//	if err := p.Create(ctx, &p1); err != nil {
//		return err
//	}
//	return nil
//}
//
//
