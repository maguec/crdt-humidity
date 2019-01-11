# crdt-humidity


Using Redis CRDT to simulate what would happen if you placed a humidifier and a dehumidifier in the same room.


## 
![alt text](https://raw.githubusercontent.com/maguec/crdt-humidity/master/docs/img.png)


## 
The humidifier *increments* the net_humidity by 1 when it runs
The dehumidifier *decrements* the net_humidity by 1 when it runs

The key is of type HASH with the following values:

|--------------------|
| net_humidity       |
| humidifier_count   |
| dehumidifier_count |
| winner             |
