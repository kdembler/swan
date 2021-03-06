// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conf

import "fmt"

// CassandraAddress represents cassandra address flag.
var CassandraAddress = NewStringFlag("cassandra_address", "Address of Cassandra DB endpoint for Metadata and Snap Publishers.", "127.0.0.1")

// CassandraUsername holds the user name which will be presented when connecting to the cluster at CassandraAddress.
var CassandraUsername = NewStringFlag("cassandra_username", "The user name which will be presented when connecting to the cluster at 'cassandra_address'.", "")

// CassandraPassword holds the password which will be presented when connecting to the cluster at CassandraAddress.
var CassandraPassword = NewStringFlag("cassandra_password", "The password which will be presented when connecting to the cluster at 'cassandra_address'.", "")

// CassandraConnectionTimeout encodes the internal connection timeout for the publisher. Note that increasing this
// value may increase the total connection time significantly, due to internal retry logic in the gocql library.
var CassandraConnectionTimeout = NewDurationFlag("cassandra_connection_timeout", "Timout for communication with Cassandra cluster.", 0)

// CassandraSslEnabled determines whether the cassandra publisher should connect to the cluster over an SSL encrypted connection.
var CassandraSslEnabled = NewBoolFlag("cassandra_ssl", "Determines whether the cassandra publisher should connect to the cluster over an SSL encrypted connection. Flags CassandraSslHostValidation, CassandraSslCAPath, CassandraSslCertPath and CassandraSslKeyPath should be set accordingly.", false)

// CassandraSslHostValidation determines whether the publisher will attempt to validate the cluster at CassandraAddress.
var CassandraSslHostValidation = NewBoolFlag("cassandra_ssl_host_validation", "Determines whether the publisher will attempt to validate the host. Note that self-signed certificates and hostname mismatch, will cause the connection to fail if not set up correctly. The recommended setting is to enable this flag.", false)

// CassandraSslCAPath enables self-signed certificates by setting a certificate authority directly.
var CassandraSslCAPath = NewStringFlag("cassandra_ssl_ca_path", "Enables self-signed certificates by setting a certificate authority directly. This is not recommended in production settings.", "")

// CassandraSslCertPath sets the client certificate, in case the cluster requires client verification.
var CassandraSslCertPath = NewStringFlag("cassandra_ssl_cert_path", "Sets the client certificate, in case the cluster requires client verification.", "")

// CassandraSslKeyPath sets the client private key, in case the cluster requires client verification.
var CassandraSslKeyPath = NewStringFlag("cassandra_ssl_key_path", "Sets the client private key, in case the cluster requires client verification.", "")

// CassandraPort sets port of Cassandra DB endpoint.
var CassandraPort = NewIntFlag("cassandra_port", "Port of Cassandra DB endpoint.", 9042)

// CassandraTimeout sets query timeout for the publisher in seconds.
var CassandraTimeout = NewDurationFlag("cassandra_timeout", "Query timeout for the publisher in seconds.", 0)

// CassandraInitialHostLookup sets that driver will attempt to get host info from Cassandra.
var CassandraInitialHostLookup = NewBoolFlag("cassandra_initial_host_loopup", "If false, driver won't attempt to get host info from Cassandra.", true)

// CassandraIgnorePeerAddr sets that driver won't resolve internal nodes addresses.
var CassandraIgnorePeerAddr = NewBoolFlag("cassandra_ignore_peer_addr", "If true, doesn't resolve internal nodes addresses.", false)

// CassandraKeyspaceName sets keyspace name used by driver.
var CassandraKeyspaceName = NewStringFlag("cassandra_keyspace_name", "Keyspace used to store metadata.", "swan")

// CassandraCreateKeyspace sets that publisher will attempt to create keyspace
var CassandraCreateKeyspace = NewBoolFlag("cassandra_create_keyspace", "Attempt to create keyspace.", false)

// CassandraTagIndex allows to pass comma-separated list of tags that will be used to insert metrics into tags table to improve SELECT performance.
var CassandraTagIndex = NewStringFlag("cassandra_tag_index", fmt.Sprintf("Allows to pass comma-separated list of tags that will be used to insert metrics into %s.tags table to improve SELECT performance", CassandraKeyspaceName.Value()), "swan_experiment")

// InfluxDBAddress represents influxdb address flag.
var InfluxDBAddress = NewStringFlag("influxdb_address", "Address of InfluxDB DB endpoint for Metadata and Snap Publishers.", "127.0.0.1")

// InfluxDBUsername holds the user name which will be presented when connecting to the cluster at InfluxDBAddress.
var InfluxDBUsername = NewStringFlag("influxdb_username", "The user name which will be presented when connecting to the cluster at 'influxdb_address'.", "swan")

// InfluxDBPassword holds the password which will be presented when connecting to the cluster at InfluxDBAddress.
var InfluxDBPassword = NewStringFlag("influxdb_password", "The password which will be presented when connecting to the cluster at 'influxdb_address'.", "swan")

// InfluxDBPort sets port of InfluxDB DB endpoint.
var InfluxDBPort = NewIntFlag("influxdb_port", "Port of InfluxDB DB endpoint.", 8086)

// InfluxDBCreateDatabase controls whether driver shall try to create database in influx
var InfluxDBCreateDatabase = NewBoolFlag("influxdb_create_database", "Attempt to create database (if it does not exist)", true)

// InfluxDBInsecureSkipVerify controls the certificate verification step
var InfluxDBInsecureSkipVerify = NewBoolFlag("influxdb_insecure_skip_verify", "If set skip certificate validation", true)

// InfluxDBMetaName sets database name used by driver.
var InfluxDBMetaName = NewStringFlag("influxdb_metadata_db_name", "Database's name used to store metadata.", "swan_metadata")

// InfluxDBMetricsName sets database name used by driver.
var InfluxDBMetricsName = NewStringFlag("influxdb_metrics_db_name", "Database's name used to store metrics.", "swan_metrics")

// DefaultSnapPublisher  sets default publisher used by swan
var DefaultSnapPublisher = NewStringFlag("default_snap_publisher", "Publisher to use. Name shall be used from snap-plugin-publisher-<name>. Supported: cassandra, influxdb", "cassandra")

// DefaultMetadataDB sets default database for metadata
var DefaultMetadataDB = NewStringFlag("default_metadata_db", "Database to which metadata will be stored. Suported: cassandra, influxdb", "cassandra")
