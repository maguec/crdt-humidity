# crdt-humidity


Using Redis CRDT to simulate what would happen if you placed a humidifier and a dehumidifier in the same room.


## 
![alt text](https://raw.githubusercontent.com/maguec/crdt-humidity/master/docs/diagram.png)


## 
The humidifier *increments* the net_humidity by 1 when it runs
The dehumidifier *decrements* the net_humidity by 1 when it runs

The key is of type HASH with the following values:

```
| humidity           |
|--------------------|
| net_humidity       |
| humidifier_count   |
| dehumidifier_count |
| winner             |
```

## Running

### Start the humidifier and dehumidifier

if wait is set they will not start writing until the watcher is started

on CRDB1
```
 ./humidity -port 10000 -password 12345 -wait -role humidifier
```

on CRDB2
```
 ./humidity -port 10000 -password 12345 -wait -role dehumidifier
```

### Start the watcher in a watch

on CRDB1
```
watch -n 1 -d "./humidity-watcher -port 10000 -password 12345"
```

### Resetting

stop the watcher process

on either node run 
```
redis-cli -a 12345 -p 10000 del humidity
```



