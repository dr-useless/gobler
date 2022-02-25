gobler
-----

A CLI for gobkv

# bind
```
./gobler bind $NETWORK $ADDRESS --a $AUTH
```
Bind writes the connection details to the local (pwd) config file.

### $NETWORK
This can be any value supported by Go stdlib's `net.Conn` [Dial](https://pkg.go.dev/net?utm_source=gopls#Dial) method.

For example `unix`, or `tcp`.

### $ADDRESS
As per [$NETWORK](#$NETWORK), can be any value supported by the Dial method, depending on your choice of network.

### --a $AUTH
The `-a` auth flag is for the (optional) auth secret, as set in your gobkv config.

# get
```
./gobler get $KEY
```
Fairly self-explanatory; get the value for a given key.

# set
```
./gobler set $KEY $VALUE --ttl $TTL
```
Write a value for a given key.

### $TTL
Optional. This is the number of seconds after which the key will expire, and be cleaned up.

# del
```
./gobler del $KEY
```
Removes the given key.

# list
```
./gobler list $KEY_PREFIX
```
This will stream (individually) every key that matches the prefix.