# Length Extension Attack

This is an implementation for md5. You can read more information about this attack [here](https://en.wikipedia.org/wiki/Length_extension_attack).

**WARNING**: This repository has changed the standard workflow of the md5 Go package, in order to implement a succesfull attack. 


## One-pass HMAC length extension attack

Steps to recreate the attack:

```
echo -n "1234" > key
echo -n "hello" > message
cat key message | openssl dgst -md5 -binary > tag
```

Now we suppose that the attacker knows only the information of `message` and `tag`. The goal is to add information to the message and compute the new One-pass HMAC without knowing the `key`. 

The following example add the information "\x20world" to the hashed information.

```
go run cmd/main.go
// a1b18483023ea03a9c3e68c9bf00f650
```

This will print the new md5 without knowing the key.

To recreate the example with the `key` and openssl to compare with the Go output, can be done with the following commands:

First of all, we must create a padding block following the same structure as the md5 algorith will create

```
echo -n ' world' > toAdd
dd if=/dev/zero bs=1 count=1 | tr "\0" $'\x80' > padding2.bin
dd if=/dev/zero bs=1 count=46 >> padding2.bin
dd if=/dev/zero bs=1 count=1 | tr "\0" $'\x48' >> padding2.bin
dd if=/dev/zero bs=1 count=7 >> padding2.bin
xxd padding2.bin
```

Now we have all the necessary files to compute the hash that will match the one that Go forged:

```
cat key message padding2.bin toAdd | openssl dgst -md5 -binary | xdd -p
// a1b18483023ea03a9c3e68c9bf00f650
```

TODO => More documentation on md5 standard library modifications
