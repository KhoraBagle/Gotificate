# Gotificate
----- ABOUT THIS PROJECT---- Gotificate is "GO" and "Certificate" as one word.
This is an exploratory project created as a sample submission to Digicert with GO and Javascript. 
It creates a structure with some information of different dogs that I am friends with and creates an API using Mux router in GO. 
Javascript accesses the above api on either http or https localhost and shows dog information in browser. 
PKI: Security authority can be created and security certificate can be issued via included code. 
------------------
PKI using sha 256 rsa 2048
GO Struct ---API (Mux router)--->|PKI| ---> html and jss front end
--------------------

DogPage.png in root directory is what final result should look like.

-----BUILT WITH------

    // PKI
	"crypto/rand"
	"crypto/rsa"
    "crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"

    //API
    "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"


----GETTING STARTED-----

Clonse repository from:  https://github.com/"your username"/Gotificate.git
or with: 
	git clone https://github.com/your_username_/Gotificate.git
Install GO if not installed already. 


-----USAGE-----

1. PKI 

CREATE CERTIFICATE TRUSTED AUTHORITY
    go run \Gotificate\src\authority
creates 'ca' certificate authority certificate. 
below fields in create_authority.go set up basic authority info.

	Organization:  []string{"DOG AUTHORITY"},
	Country:       []string{"USA"},
	Province:      []string{"OHIO"},
	Locality:      []string{"COLUMBUS"},
	StreetAddress: []string{"DOG PARK"},
	PostalCode:    []string{"43221"},

below lines set up valid date range. 

    NotBefore:             time.Now(),
	NotAfter:              time.Now().AddDate(1000, 0, 0),

    go run \Gotificate\src\certify  
certify creates local certificate that is signed by the above created certificate authority
below fields in create_and_sign.go set up basic authority info.

	Organization:  []string{"Dogwebsite.com"},
	Country:       []string{"USA"},
	Province:      []string{"OHIO"},
	Locality:      []string{"COLUMBUS"},
	StreetAddress: []string{"DOG PARK"},
	PostalCode:    []string{"43221"},

below lines set up valid date range. 
    NotBefore:             time.Now(),
	NotAfter:              time.Now().AddDate(1000, 0, 0),


certificates are saved in root folder as ca.crt | key  and certificate.crt | key pairs


API:
go run \Gotificate\src\api     
Starts API on port 8080. 
api end points:

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/dog", createDog).Methods("POST")
	router.HandleFunc("/dogs", getAllDogs).Methods("GET")
	router.HandleFunc("/dogs/{name}", getDogInfo).Methods("GET")
	router.HandleFunc("/dogs/{name}", updateDog).Methods("PATCH")
	router.HandleFunc("/dogs/{name}", deleteDog).Methods("DELETE")

The server runs as http because I cant issue security ceritifacte to local host/ip and have it recognized. 
So an https handshake will fail. 

However if not hosted locally (same ip) and you are using multiple machines the last 2 lines in main.go  can be comment swappped with previous line and the server will run as https 
and will use previously created security certificates. 

	log.Fatal(http.ListenAndServe(":8080", router))

	//	log.Fatal(http.ListenAndServeTLS(":8080", "certificate.crt", "certificate.key", router))
	//	http.Handle("/", router)

test API through your browser on: (default) http://api.localhost:8080/dogs or https://api.localhost:8080/dogs depending on which setting you used. 
or using *curl or an API testing app. 
Once it is running.


FRONT END:

open \Gotificate\src\webjs\index.html in any browser
if you are running the server as https you will have to install the security certificate from the root folder. "ca.crt". install it to tursted authority folder. 
if you are running it as http on local, most browsers will block api fetch, this is browser specific and will need to be disabled for front end to properly work. 
App called CORS Unblock will work. 

Now you will see all of the nice dogs in my neightborhood. 
You can add, modify, query, and delete dogs with curl commands and refresh browser window to display changes. 
(Dont delete dogs though as its not nice)