import React, { useEffect } from 'react';
// import Chart from 'chart.js/auto'; // Import Chart.js library

const LineChart = () => {
  // 2023-10-13T09:43:48.650000Z  INFO avail_light: Using config: RuntimeConfig { http_server_host: "127.0.0.2", http_server_port: (7001, 0), secret_key: Some(Seed { seed: "avail" }), port: 37001, tcp_port_reuse: false, autonat_only_global_ips: false, autonat_throttle: 1, autonat_retry_interval: 10, autonat_refresh_interval: 30, autonat_boot_delay: 5, identify_protocol: "/avail_kad/id/1.0.0", identify_agent: "avail-light-client/rust-client", bootstraps: [(PeerId("12D3KooWS9BZnKansQfo6AHJ54gZbSPTPENHEdr3JWEFJ5NELKAw"), "/ip4/192.168.1.125/udp/37000/quic-v1")], bootstrap_period: 300, relays: [], full_node_ws: ["ws://127.0.0.1:9944"], app_id: Some(0), confidence: 92.0, avail_path: "avail_path", log_level: "info", log_format_json: false, ot_collector_endpoint: "http://otelcollector.avail.tools:4317", disable_rpc: false, disable_proof_verification: false, dht_parallelization_limit: 20, put_batch_size: 1000, query_proof_rpc_parallel_tasks: 8, block_processing_delay: None, block_matrix_partition: None, sync_start_block: None, max_cells_per_rpc: Some(30), threshold: 5000, kad_record_ttl: 86400, publication_interval: 43200, replication_interval: 10800, replication_factor: 20, connection_idle_timeout: 30, query_timeout: 60, query_parallelism: 3, caching_max_peers: 1, disjoint_query_paths: false, max_kad_record_number: 2400000, max_kad_record_size: 8192, max_kad_provided_keys: 1024, avail_secret_key: None }

// Local peer id: PeerId("12D3KooWBmiYs4mfEDBqRQ8itZ3Ex8nZVFLwJeMbp3pED65VRRxM"). Public key: PublicKey { publickey: Ed25519(PublicKey(compressed): 1d9ad3d40b799dd5b3574e57142aff0f137aedc47bf6707433ec9cc339d02a) }.
  useEffect(() => {
    // Get the canvas element
    const ctx = document.getElementById('lineChart').getContext('2d');

    // Define the data for the chart
    const data = {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
      datasets: [{
        label: 'Sales',
        backgroundColor: 'rgba(255, 99, 132, 0.2)',
        borderColor: 'rgba(255, 99, 132, 1)',
        borderWidth: 2,
        data: [65, 59, 80, 81, 56, 55, 40]
      }]
    };

    // Define the options for the chart
    const options = {
      scales: {
        yAxes: [{
          ticks: {
            beginAtZero: true
          }
        }]
      }
    };

    // Create the Line Chart
    new Chart(ctx, {
      type: 'line',
      data: data,
      options: options
    });
  }, []);

  return (
    <div className='flex justify-center mt-[21rem]'>
      <canvas id="lineChart" width="800" height="400"></canvas>
    </div>
  );
};

export default LineChart;
