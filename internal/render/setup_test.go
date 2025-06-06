package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/markledger/bookings/internal/config"
	"github.com/markledger/bookings/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	//What am I going to put in session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (mw *myWriter) Header() http.Header {
	return http.Header{}
}

func (mw *myWriter) WriteHeader(statusCode int) {

}

func (mw *myWriter) Write(b []byte) (int, error) {
	return len(b), nil
}
