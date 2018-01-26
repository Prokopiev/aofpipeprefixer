# aofpipeprefixer

aofpipeprefixer is console utility used for converting keys for redis-cluster. Redis cluster does not support databases, 
so keys from redis db 0 are converted to **_0:key_name_** and so on.

Based on modified library [aof](http://github.com/gato/aof)

## Usage
```bash
tail -f -n +1 appendonly.aof | ./aofpipeprefixer | redis-cli --pipe
```

## Todo
- [x] tests