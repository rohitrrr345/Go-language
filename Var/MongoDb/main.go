package mongodb
import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
	"net/http"
)
  func main() {
	
      r:=httprouter.New();
	  uc:=controllers.NewUserController(getSession());
	  r.GET("")

  













  }