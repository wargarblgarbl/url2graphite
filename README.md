# url2graphite
A very VERY dumb and insecure "server" to send arbitrary stats to graphite.

Parses the ouput of a curl request, and then forwards it to a graphite instance.  

This abuses go tutorials, fmtprint, unix dates,  and will probably cause your production systems to catch fire if you ever EVER decide to use this in production. 

```
  -gport string
    	graphite port (default "2003")
  -gurl string
    	url of graphite server (default "192.168.1.138")
  -laddress string
    	local server addressy
  -lport string
    	local server listen port (default "9090")
```

There is no security. None. If you wanted to run this somewhere on the public-facing internet, may the gods help you. 

To send a request: 

```curl -s http://host:9090/requestname/requestnamestat/$RANDOM > /dev/null```

You have to edit the graphite_ip and listen_port consts to change your settings.

This can't do requests to passworded graphite.  
