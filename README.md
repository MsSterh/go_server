# Server on Golang

Server is running on http://localhost:8080/ and has 4 end-points:

<ul>
<li> <b>GET /</b> - shows all blocks</li>
<li> <b>GET /all_data</b> - shows all data (all records)</li>
<li> <b>POST /add_data</b> - adds data in json <code>{ data: "string of data" }</code></li>
<li> <b>GET /last_blocks/{blockId}</b> - shows 'blockId' count of last blocks</li>
</ul>


## Data (record) structure

<pre><code>{
  data: "string of data"
}</code></pre>

## Block structure

Block is forming when server gets 5 records:

<pre><code>{
  // string of previous block hash or '0' for the first block
  previous_block_hash: '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08',
  // array of 5 records data
  rows: ['data1','data2','data3','data4','data5'],
  // current timestamp
  timestamp: 12123889,
  // string of SHA256 of previous 3 fields of this block
  block_hash: '1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014'
}</code></pre>

## Install and run

<a href="https://golang.org/doc/install#install" target="_blank">Install Go lang</a>.

Build project:
<pre><code>go build</code></pre>

Run on http://localhost:8080/:
<pre><code>./go_server</code></pre>

Add new data from console: 
<pre><code>curl -H "Content-Type: application/json" -d '{"data":"New data"}' http://localhost:8080/add_data</code></pre>

