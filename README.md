# url2graphite
A very VERY dumb and insecure "server" to send arbitrary stats to graphite.

Parses the ouput of a curl request, and then using bash (!!) forwards it on to a graphite (or a statsd) instance. 

This abuses unix pipes, go tutorials, and will probably cause your production systems to catch fire if you ever EVER decide to use this in production. 

There is no security. None. If you wanted to run this somewhere on the public-facing internet, nginx with .htpasswd would probably help. You have been warned though. 

To send a request: 

```curl -s http://host:9090/requestname/requestnamestat/$RANDOM > /dev/null```

To change the port and IP of your graphite instance, set the PORT and SERVER variables in forward.sh 
