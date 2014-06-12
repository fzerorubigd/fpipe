Friend Pipe
===========

A simple friend to friend pipe written in Go language to learn that language itself. 
At first you need to add your friend configuration like this : 

```bash
fpipe -mode add -friend friend_name address 1.2.3.4 -port 9876
```

Your friend should add you the same way, with your address but the same port.
then you can send data to each other like this : 

```bash
# Sender 
command | fpipe -mode server -friend receiver_name
```

```bash
# Receiver 
fpipe -mode client -friend sender_name
```

WARNING : Its just a fun project, do not take it serious.