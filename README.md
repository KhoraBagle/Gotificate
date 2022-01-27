# Gotificate
----- ABOUT THIS PROJECT---- 
--Gotificate is "GO" and "Certificate" as one word--
This is an exploratory project created as a sample submission to Digicert with GO and Javascript. 
It creates a structure with some information of different dogs that I am friends with and creates an API using Mux router in GO. 
Javascript calls the above api on either http or https localhost and shows dog information in browser. 
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

Install GO if not installed already.

Clonse repository from:  https://github.com/KhoraBagle/Gotificate
or with: 

	git clone https://github.com/KhoraBagle/Gotificate

Open cloned repository. 'Gotificate'
Open terminal on repository location 'Gotificate' as root folder. 

-----USAGE-----

1. PKI 

1.1 CREATE CERTIFICATE TRUSTED AUTHORITY -- creates 'ca' certificate authority certificate. 

To create 'ca' trusted authority certificate in your local repository folder run the below in terminal:

	 go run src\authority\create_authority.go


To modify trusted authority settings: navigate to src\authority. type in terminal:

	cd src/authority
open create_authority.go in editor:
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

If you made any changes save the file and back out to parent folder to rerun with the above commnad line. 
And/or run from current folder with:

	go run .

Once the authority certificate is created it will show up as ca.crt in root folder. 
Open it and install it in trusted authority folder 
for your browsers to recognize it. 

If you are in the authority folder in terminal, back out to root folder with:

	 cd ../..

1.2 CREATE A CERTIFICATE SIGNED BY THE TRUSTED AUTHORITY 

To create a local certificate that is signed by the above created certificate authority.
From root folder run:

	go run src\certify\create_and_sign.go

To edit certificate settings navigate to src\certify\ or from root folder type:

	cd src\certify\

open create_and_sign.go in editor.
Below fields in create_and_sign.go set up basic certificate  info.

	Organization:  []string{"Dogwebsite.com"},
	Country:       []string{"USA"},
	Province:      []string{"OHIO"},
	Locality:      []string{"COLUMBUS"},
	StreetAddress: []string{"DOG PARK"},
	PostalCode:    []string{"43221"},

below lines set up valid date range. Note: as it is it will be valid for 1k years (as any dog related certificate should be)

    NotBefore:             time.Now(),
	NotAfter:              time.Now().AddDate(1000, 0, 0),


certificates are saved in root folder as ca.crt | key  and certificate.crt | key pair

If you made changes to the certificate and are in the authority folder in terminal, back out to root folder with:

	 cd ../..


2. API:

2.1 ---- START API -----

Make sure that you are in the root folder in terminal '\Gotificate' and start API with:

	go run \Gotificate\src\api\main.go     

This will start API on port 8080. 

2.2 ----  MODIFY API (skip this step if you want to use default settings)-----

Navigate to API folder cd src\api or via terminal with:

	 cd src\api

Open main.go in editor
The Mux server is set to run as  http because it is troublesome to issue security ceritifacte to local host/ip and have it authenticated. 
Since, as opposed to a unique address, such as: https://www.ilikedogsalottheyarenice123.com ... seeing "localhost" on a certificate is not unique 
and gets rejected by the authenticator and browsers. 

However, the code is included, so if you want to test the server on a local network where a certificate can be issued and certified properly:
Swap commenting out on the following lines of code in main.go (they are at the end of the file)

	log.Fatal(http.ListenAndServe(":8080", router))

	//	log.Fatal(http.ListenAndServeTLS(":8080", "certificate.crt", "certificate.key", router))
	//	http.Handle("/", router)

To look like this

	//	log.Fatal(http.ListenAndServe(":8080", router))

	log.Fatal(http.ListenAndServeTLS(":8080", "certificate.crt", "certificate.key", router))
	http.Handle("/", router)

Save main.go and rerun src\api\main.go The server will now use the previously created certificates and will run as https.
You will also need to go back to step 1.2 and reissue the certificate to your local server address. 

2.3 --- TEST API ---

If you are running the server with its default settings. Open any browser to:

	http://api.localhost:8080/dogs 

You can try fetching individual dogs, such as:

	http://api.localhost:8080/dogs/Waf

or 

	http://api.localhost:8080/Yozhik

If you modified the code to run the server with PKI then open the following:	
	https://api.localhost:8080/dogs

Additionally you can use curl, or any API testing software.
The end points are as follows:	

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/dog", createDog).Methods("POST")
	router.HandleFunc("/dogs", getAllDogs).Methods("GET")
	router.HandleFunc("/dogs/{name}", getDogInfo).Methods("GET")
	router.HandleFunc("/dogs/{name}", updateDog).Methods("PATCH")
	router.HandleFunc("/dogs/{name}", deleteDog).Methods("DELETE")

3. FRONT END:

In this step we will run front end that calls one of our API endpoints "GET" all dogs and displays the list of all the dogs in memory. 
Furthermore you can use other endpoints to modify the dog struct in memory and refresh the site to reflect it.

If you are running the server with default settings without PKI go to step 3.1
Else: step 3.2

3.1 SERVER RUNNING WITHOUT PKI

open \Gotificate\src\webjs\index.html in any browser

The result should look like: DogPage.png in the root folder.
If it does not it means that the browser is rejecting the jss call to our API. 
The reason being that it is on local and running on http so the browser is blocking CORS (Cross original resource sharing) for security reasons.
Basically because the server is not using PKI, it is on local, and it just looks really uncool to the browser. 
You can install an app called CORS unblock and restart the index.html file for it to work. 
Though I wouldnt instlal 3rd party apps on a work machine :/ 

if you are running the server as https you will have to install the security certificate from the root folder. "ca.crt". install it to tursted authority folder. 
if you are running it as http on local, most browsers will block api fetch, this is browser specific and will need to be disabled for front end to properly work. 
App called CORS Unblock will work. 


3.2 SERVER RUNNING WITH PKI

Navigate to \Gotificate\src\webjs\ and open scripts.js
Change line 17 to:

	request.open('GET', 'https://api.localhost:8080/dogs', true)

This will call the server on https instead of http. Save the file. 
open \Gotificate\src\webjs\index.html in any browser. 

The result should once again look like: DogPage.png in the root folder.

If the page does not open at all, it means that you need to go to step 2.2, uncomment the code and rerun src/api/main.go to start the mux server on https.

If the page loads and all you see is a really cute dog but no dog descriptions, it means that the ceritifcate is invalid and you are failing the 
handshakle with the sevrer. 
Check security warning info on your browser, if it shows certificate as invalid then
you need to go to step 1.1 and make sure that you add the trusted authority certificate to your trusted folder. 
If the certificate is valid but the browser is telling you that it still cant be authenticated,
Go to step 1.2 and change the Organiation name to reflect the ip or the address of the server that runs the main.go code. 

4. CONCLUSION 

Now you will see all of the nice dogs in my neightborhood. 
You can add, modify, query, and delete dogs with curl commands and refresh browser window to display changes. 
(Dont delete dogs though as its not nice)