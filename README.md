    
BLOCK DIAGRAM:
----------------
				   
           
		           -------------------
	   HTTP/POST      |                   |         ------------
CLIENT  -------------->   |     WEB SERVER    |        |  SQLITE    | 
		          |                   |------->|     	    |
	                  |		      |         ------------
			   -------------------
                                   |
				   | 
	                           |
	                    -------------------	
		           |                   | 
		           |   BASE-36 ENCODER |
			    -------------------
						   
               
Steps for URL Shortening:
------------------------

1. Client sends HTTP/POST request to Web Server with "LongURL" 
2. Web Server inserts "LongURL" to sqlite database and gets autoIncrement ID for that row.
3. Web Server encodes the ID into base-36, prefix "http://" to the result,  and returns the result back to Client as Short-URL.


Steps for getting Original URL:
-----------------------

1. Client sends HTTP/POST request to Web Server with "ShortURL"
2. Web Server removes the "http://" prefix from the "ShortURL", and decodes it using Base-36 decoding. Result is ID ( integer value)
3. Web Server fetches the "LongURL" corresponding to the ID ( from step 2) from sqlite database, and returns the result back to client as Original-URL.

Deployment Steps:
-----------------



