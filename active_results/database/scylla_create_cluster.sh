# creating 3 nodes
docker run --name result_node1 --network ws-scylla -p "9042:9042" -d scylladb/scylla:6.1.1 \
  --overprovisioned 1 \
  --smp 1

docker run --name result_node2 --network ws-scylla -d scylladb/scylla:6.1.1 \
  --overprovisioned 1 --smp 1 \
  --seeds="result_node1"

docker run --name result_node3 --network ws-scylla -d scylladb/scylla:6.1.1 \
  --overprovisioned 1 \
  --smp 1 \
  --seeds="result_node1"

  # check status
  docker exec -it result_node1 nodetool status