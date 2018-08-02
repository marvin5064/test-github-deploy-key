a side note to myself

If you stucking at input passphrase

e.g.
```
Enter passphrase for key '/Users/ABC/.ssh/id_rsa':
```

here is a proper solution to set up ssh-key without passphrase (never ask passphrase input during process)

```
ssh-keygen -b 2048 -t rsa -f /tmp/sshkey -q -N ""
```

also, all content matters for private key
```
-----BEGIN RSA PRIVATE KEY-----
abc
def
hij
-----END RSA PRIVATE KEY-----
```

hence dont missing any from above