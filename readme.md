# rkteer
-----

Rocketeer, a CLI for rocketkv.
Using the [client package](https://github.com/intob/rocketkv/tree/main/client) in rocketkv.

# bind
```
./rkteer bind $NETWORK $ADDRESS --a $AUTH
```
Bind writes the connection details to the local (pwd) config file.

### $NETWORK
This can be any value supported by Go stdlib's `net.Conn` [Dial](https://pkg.go.dev/net?utm_source=gopls#Dial) method.

For example `unix`, or `tcp`.

### $ADDRESS
As per [$NETWORK](#$NETWORK), can be any value supported by the Dial method, depending on your choice of network.

### --a $AUTH
The `-a` auth flag is for the (optional) auth secret, as set in your rkteer config.

# get
```
./rkteer get $KEY
```
Fairly self-explanatory; get the value for a given key.

# set
```
./rkteer set $KEY $VALUE --ttl $TTL
```
Write a value for a given key.

### $TTL
Optional. This is the number of seconds after which the key will expire, and be cleaned up.

# del
```
./rkteer del $KEY
```
Removes the given key.

# list
```
./rkteer list $KEY_PREFIX
```
This will stream (individually) every key that matches the prefix.

# count
```
./rkteer count $KEY_PREFIX
```
Returns the count of keys that match the given prefix. If no prefix is given, the count will be of all keys in the KV store.