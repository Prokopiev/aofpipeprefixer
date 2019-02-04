# aofpipeprefixer

aofpipeprefixer is console utility used for filtering keys for redis-cluster. It passes only keys from selected db prefixes.

Based on modified library [aof](http://github.com/gato/aof)

## Usage
```bash
tail -f -n +1 appendonly.aof | ./aofpipeprefixer -base=2,12 | redis-cli --pipe
```

## Todo
- [x] tests