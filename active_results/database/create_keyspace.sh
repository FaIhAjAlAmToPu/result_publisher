# Step 1: Enter CQLSH
docker exec -it node1 cqlsh
# Step 2: Create the Keyspace
CREATE KEYSPACE IF NOT EXISTS twitch
  WITH replication = {
    'class': 'NetworkTopologyStrategy',
    'replication_factor': '3'
  } AND tablets = {'enabled': false};

# Step 3: Exit CQLSH
exit