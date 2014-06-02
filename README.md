asterisk-redis
====

## How to install

```
git clone https://github.com/t-k/asterisk-redis

cd asterisk-redis

go get github.com/mattn/gom
gom install
gom build
```

##Example

```
; /etc/asterisk/extensions.conf
; System(path_to_redis_script command value})

exten => _1000,1,Answer()
 same => n,System(/etc/asterisk/redis publish foobar ${CALLERID(num)})
 same => n,System(/etc/asterisk/redis setex foo 20 ${CALLERID(num)})
 same => n,Playback(beep)
 same => n,Echo()
 same => n,Hangup()

```

## Setting config values

```
# bash_profile

export REDIS_HOST='foo.com'
export REDIS_PORT='6389'
```