package main

import (
"fmt"
"log"
"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello From Go Service!")
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/app" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, mainAppPage)
}

func main() {
	http.HandleFunc("/app", appHandler) // Update this line of code
	http.HandleFunc("/", helloHandler) // Update this line of code
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

var mainAppPage = `
<!doctype html>
<html lang=en>
<head>
<meta charset=utf-8>
<title>blah</title>
<style type="text/css">
body {
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
	line-height: 1.6;
	color: #222;
	max-width: 40rem;
	padding: 2rem;
	margin: auto;
	background: #fafafa;
  }
  img {
	max-width: 100%;
  }
  a {
	color: #2ECC40;
  }
  h1, h2, strong {
	color: #111;
  }
</style>
</head>
<body>
<h1>Schooner</h1>
<img src="https://upload.wikimedia.org/wikipedia/commons/b/ba/The_America_Schooner_Yacht_-_New_York_Yacht_Club.jpg" width="500" />
<p><strong>Schooner</strong>, a vessel rigged with fore and aft sails, properly with two masts, 
but now often with three, four 
and sometimes more masts.
They are much used in the coasting trade, and require a smaller crew in proportion to their size than 
square-rigged vessels.
They also tend to be used where sea-breezes and land-breezes alternate, using such wind for following the coast. 
</p>
<h3>Name Origin</h3>
<p>According to the story, which is probably true, 
the name arose from a chance spectator's exclamation “there she scoons,” i.e. glides, 
slips free, at the launch of the first vessel of this type at Gloucester, Massachusetts, in 1713, 
her builder being one Andrew Robinson. 
</p>
<p>The spelling “schooner” is due to a supposed derivation from the 
Dutch <em>schooner</em>, but that and the other European equivalents, Ger. Schoner, Dan. skonnert, Span. 
and Portuguese escuna, &c., are all from English. “ To scoon,” according to Skeat, is a 
Scottish (Clydesdale) dialect word, meaning to skip over water like a flat stone, and is ultimately 
connected with the root, implying quick motion, seen in shoot, scud, &c. In American colloquial 
usage “schooner” is applied to the covered prairie-wagons used by the 
emigrants moving westward before the construction of railways, and to a tall, narrow, lager-beer glass.</p>
</body>
</html>
`