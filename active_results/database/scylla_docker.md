Using multiple cores requires storing a proper value to `/proc/sys/fs/aio-max-nr`.
While the default value for `aio-max-nr` on many non-production systems is 64K,
this may not be optimal for high-performance workloads. The ideal value depends
on the current value of `/proc/sys/fs/aio-nr` and also on the number of cores to
be used by ScyllaDB:

    Available AIO on the system >= (AIO requests per-cpu) * ncpus

Expanding the definitions on both sides, we get:

    aio_max_nr - aio_nr >= (storage_iocbs + preempt_iocbs + network_iocbs) * ncpus
                                     1024               2           50000

Which yields, for `/proc/sys/fs/aio-max-nr`:

    aio_max_nr >= aio_nr + 51026 * ncpus


#### options
1. `--overprovisioned`: for environments where resources are shared
2. `--smp`: number of CPU cores ScyllaDB uses
### getting started
```sh
docker network create scylla-net
docker run --name result-node1 --network scylla-net -p 9042:9042 -d scylladb/scylla:6.1.1 \
  --overprovisioned 1 --smp 1
```

#### volume for persistency??

```sh
docker stop result-node1
docker rm result-node1
docker volume ls
docker inspect result-node1
docker volume rm <volume_name>
docker network rm scylla-net # can also remove the network
# verify removals
docker ps -a
docker volume ls
docker network ls
docker image rm scylladb/scylla:6.1.1 # if want to remove scyllaDB image
```

```sh
docker volume create result-node1-data
docker volume create result-node2-data
docker volume create result-node3-data

docker run --name result-node1 --network scylla-net -p 9042:9042 -d -v result-node1-data:/var/lib/scylla scylladb/scylla:6.1.1 --overprovisioned 1 --smp 1
docker run --name result-node2 --network scylla-net -d -v result-node2-data:/var/lib/scylla scylladb/scylla:6.1.1 --overprovisioned 1 --smp 1 --seeds="result-node1"

docker run --name result-node3 --network scylla-net -d -v result-node3-data:/var/lib/scylla scylladb/scylla:6.1.1 --overprovisioned 1 --smp 1 --seeds="result-node1"
```

#### mounted path for volume
search for ``mount path``
```sh
docker volume ls
docker volume inspect result-node1-data
```

By using ``-v result-node1-data:/var/lib/scylla``, you're telling Docker to map the ```host path <mount path> to /var/lib/scylla inside the container```.

#### XFS file system
The default filesystem in Docker is inadequate for anything else than just testing out ScyllaDB.
``for production``: data volumes, ensure first that it’s on a ScyllaDB-supported filesystem like XFS, then create a ScyllaDB data directory ``/var/lib/scylla`` on the host
