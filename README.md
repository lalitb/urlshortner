
# BLOCK DIAGRAM:
 
				   
   		                           -------------------
	             HTTP/POST        |                   |         ------------
     CLIENT  -------------->    |     WEB SERVER    |        |  SQLITE    | 
		                            |                   |------->|        	  |
	                              |	            	    |         ------------
			                           -------------------
                                         |
				                                 | 
	                                       |
	                                -------------------	
		                             |                   | 
		                             |   BASE-36 ENCODER |
			                            -------------------
					   
               
# Steps for URL Shortening:

1. Client sends HTTP/POST request to Web Server with "LongURL" 
2. Web Server inserts "LongURL" to sqlite database and gets autoIncrement ID for that row.
3. Web Server encodes the ID into base-36, prefix "http://" to the result,  and returns the result back to Client as Short-URL.


# Steps for getting Original URL:


1. Client sends HTTP/POST request to Web Server with "ShortURL"
2. Web Server removes the "http://" prefix from the "ShortURL", and decodes it using Base-36 decoding. Result is ID ( integer value)
3. Web Server fetches the "LongURL" corresponding to the ID ( from step 2) from sqlite database, and returns the result back to client as Original-URL.

# Deployment Steps for urlshortner Docker image:

1. This has been tested in below linux distro, but it should work in any linux distro:

$ cat /etc/lsb-release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=16.04
DISTRIB_CODENAME=xenial
DISTRIB_DESCRIPTION="Ubuntu 16.04.1 LTS"


2. The Docker image for application is published in docker hub ( lalitbhasin123/urlshortner ). The image was created using golang:onbuild image (https://hub.docker.com/_/golang/) .

3. Steps to perform in bare linux machine:
     -> Install Docker ( refer to link  https://docs.docker.com/engine/installation/ for installation instructions. )
     
     -> Pull the image: 
     
		$ sudo docker pull lalitbhasin123/urlshortner
		
      -> Validate that the docker image is downloaded, and available locally:
      
      		$  sudo docker images
		
      -> Run the Web application through docker image:
      
      		$  sudo docker run --publish 8080:8080 --name test --rm lalitbhasin123/urlshortner
		
	The urlshortner web application is now running locally on port 8080
	

# Testing using curl:

 1. Execute below command locally on the machine to create the short url:
 
         $ curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/shorten' -d '{"url":"http://a.very.long.url"}'
         {"short":"http://5yc1u"}
	 
 2. Execute below command locally on the machine to get the original url:
 
         $ curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/original' -d '{"short":"http://5yc1u"}'
         {"original":"http://a.very.long.url"}
	 
3. To run the unit test cases:
	-> Login to the running docker image, and run the test suite:
	
	      $ sudo docker exec -i -t test '/bin/bash'
	      root@fa83f7337833:/go/src/app# go test
              PASS
              ok      app     0.057s
              root@fa83f7337833:/go/src/app# exit
	      $
	   
4. To stop the running docker image:

	$  sudo docker stop test
	
