package discovery

import (
	"net"

	"github.com/hashicorp/serf/serf"
	"go.uber.org/zap"
)

type Membership struct {
	Config
	handler Handler
	serf    *serf.Serf
	events  chan serf.Event
	logger  *zap.Logger
}

func New(handler Handler, config Config) (*Membership, error) {
	c := &Membership{
		Config:  config,
		handler: handler,
		logger:  zap.L().Named("membership"),
	}

	if err := c.setupSerf(); err != nil {
		return nil, err
	}

	return c, nil
}

func (m Membership) setupSerf() (err error) {
	addr, err := net.ResolveTCPAddr(
		"tcp", m.BindAddr)
	if err != nil {
		return err
	}

	config := serf.DefaultConfig()
	config.Init()
	config.MemberlistConfig.BindAddr = addr.IP.String()
	config.MemberlistConfig.BindPort = addr.Port
	m.events = make(chan serf.Event)
	config.EventCh = m.events
	config.Tags = m.Tags
	config.NodeName = m.NodeName
	m.serf, err = serf.Create(config)

	if err != nil {
		return err
	}

	go m.eventHandler()
	if m.StartJoinAddrs != nil {
		_, err = m.serf.Join(
			m.StartJoinAddrs,
			true,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Membership) eventHandler() {
	for e := range m.events {
		switch e.EventType() {

		case serf.EventMemberJoin:
			for _, member := range e.(serf.MemberEvent).Members {
				if m.isLocal(member) {
					continue
				}
				m.handleJoin(member)
			}
		case serf.EventMemberLeave, serf.EventMemberFailed:
			for _, member := range e.(serf.MemberEvent).Members {
				if m.isLocal(member) {
					continue
				}
				m.handleLeave(member)
			}
		}
	}
}

func (m *Membership) handleJoin(member serf.Member) {
	//TODO: complete
}
func (m *Membership) handleLeave(member serf.Member) {

	//TODO: complete
}

func (m *Membership) isLocal(member serf.Member) bool {
	return m.serf.LocalMember().Name == member.Name
}

type Config struct {
	NodeName       string
	BindAddr       string
	Tags           map[string]string
	StartJoinAddrs []string
}
type Handler struct{}
