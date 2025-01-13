# go to ubuntu
wsl
# check aio-max-nr
cat /proc/sys/fs/aio-max-nr
# logout
exit
# create volume
docker volume create result-node1-data
docker volume create result-node2-data
docker volume create result-node3-data
# create cluster
docker run --name result-node1 --network scylla-net -p 9042:9042 -d -v result-node1-data:/var/lib/scylla scylladb/scylla:6.1.1 --overprovisioned 1 --smp 1
docker run --name result-node2 --network scylla-net -d -v result-node2-data:/var/lib/scylla scylladb/scylla:6.1.1 --overprovisioned 1 --smp 1 --seeds="result-node1"

docker run --name result-node3 --network scylla-net -d -v result-node3-data:/var/lib/scylla scylladb/scylla:6.1.1 --overprovisioned 1 --smp 1 --seeds="result-node1"