package mutilatesession

import (
	"time"

	"github.com/intelsdi-x/snap/mgmt/rest/client"
	"github.com/intelsdi-x/snap/scheduler/wmap"
	"github.com/intelsdi-x/swan/pkg/executor"
	"github.com/intelsdi-x/swan/pkg/snap"
	"github.com/intelsdi-x/swan/pkg/snap/sessions"
)

// DefaultConfig returns default configuration for Mutilate Collector session.
func DefaultConfig() Config {
	publisher := wmap.NewPublishNode("cassandra", snap.PluginAnyVersion)
	sessions.ApplyCassandraConfiguration(publisher)

	return Config{
		SnapteldAddress: snap.SnapteldHTTPEndpoint.Value(),
		Interval:        1 * time.Second,
		Publisher:       publisher,
	}
}

// Config contains configuration for Mutilate Collector session.
type Config struct {
	SnapteldAddress string
	Publisher       *wmap.PublishWorkflowMapNode
	Interval        time.Duration
}

// SessionLauncher configures & launches snap workflow for gathering
// SLIs from Mutilate.
type SessionLauncher struct {
	session    *snap.Session
	snapClient *client.Client
}

// NewSessionLauncher constructs MutilateSnapSessionLauncher.
func NewSessionLauncher(config Config) (*SessionLauncher, error) {
	snapClient, err := client.New(config.SnapteldAddress, "v1", true)
	if err != nil {
		return nil, err
	}

	loaderConfig := snap.DefaultPluginLoaderConfig()
	loaderConfig.SnapteldAddress = config.SnapteldAddress
	loader, err := snap.NewPluginLoader(loaderConfig)
	if err != nil {
		return nil, err
	}

	err = loader.Load(snap.MutilateCollector, snap.CassandraPublisher)
	if err != nil {
		return nil, err
	}

	return &SessionLauncher{
		session: snap.NewSession(
			"swan-mutilate-session",
			[]string{
				"/intel/swan/mutilate/*/avg",
				"/intel/swan/mutilate/*/std",
				"/intel/swan/mutilate/*/min",
				"/intel/swan/mutilate/*/percentile/5th",
				"/intel/swan/mutilate/*/percentile/10th",
				"/intel/swan/mutilate/*/percentile/90th",
				"/intel/swan/mutilate/*/percentile/95th",
				"/intel/swan/mutilate/*/percentile/99th",
				"/intel/swan/mutilate/*/qps",
				//TODO: Fetch the 99_999th value from MUTILATE task itself!
				//It shall be redesigned ASAP
				"/intel/swan/mutilate/*/percentile/*/custom",
			},
			config.Interval,
			snapClient,
			config.Publisher,
		),
		snapClient: snapClient,
	}, nil
}

// LaunchSession starts Snap Collection session and returns handle to that session.
func (s *SessionLauncher) LaunchSession(
	task executor.TaskInfo,
	tags string) (snap.SessionHandle, error) {

	// Obtain Mutilate output file.
	stdoutFile, err := task.StdoutFile()
	if err != nil {
		return nil, err
	}

	// Configuring Mutilate collector.
	s.session.CollectNodeConfigItems = []snap.CollectNodeConfigItem{
		snap.CollectNodeConfigItem{
			Ns:    "/intel/swan/mutilate",
			Key:   "stdout_file",
			Value: stdoutFile.Name(),
		},
	}

	// Start session.
	err = s.session.Start(tags)
	if err != nil {
		return nil, err
	}

	return s.session, nil
}
