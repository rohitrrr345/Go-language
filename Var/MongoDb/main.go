package mongodb
import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
	"net/http"
	"github.com/rohit/mongo-gplang/Controllers"
)
  func main() {
	
      r:=httprouter.New();
	  uc:=controllers.NewUserController(getSession());
	  r.GET("user/:id ",uc.GetUser)
	  r.POST("/user ", uc.CreateUser)
	  r.DELETE("/user/:id ",uc.DeleteUser)
 }
 func getSession() *mgo.Session{
	   s,err=mgo.Dial("mongodb://localhost:27107")
	   if err!=nil{
		panic(err)
	   }
	   return s
 }