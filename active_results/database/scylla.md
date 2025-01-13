```sh
docker exec -it scylla nodetool status
```
> Could not initialize seastar: std::runtime_error (Your system does not satisfy minimum AIO requirements. Set /proc/sys/fs/aio-max-nr to at least 66559 (minimum) or 76558 (recommended for networking performance).)

### Maximum number of allowable concurrent requests
```sh
cat /proc/sys/fs/aio-max-nr
```
> 65536

```sh
sudo sysctl -w fs.aio-max-nr=76558
```
> fs.aio-max-nr = 76558

#### make it permanent
**Not yet**
```sh
echo "fs.aio-max-nr=76558" | sudo tee -a /etc/sysctl.conf
```