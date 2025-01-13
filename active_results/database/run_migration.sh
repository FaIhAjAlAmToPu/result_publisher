# Step 1: Install Charybdis Migrator
cargo install charybdis-migrate
# Step 2: Run the Migration
cd twitch-sentinel-rs
migrate --keyspace=twitch --host=localhost:9042
# Step 3: Verify the Migration
docker exec -it node1 cqlsh
DESC KEYSPACE twitch;