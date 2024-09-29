package commander

// Declare CyberGhost service types // Declaration order matters
type ServiceType int

const (
	Traffic ServiceType = iota + 1
	Torrent
	Streaming
)

func (st ServiceType) String() string {
	return [...]string{"Traffic", "Torrent", "Streaming"}[st-1]
}
func (st ServiceType) CommandArg() string {
	return [...]string{"traffic", "torrent", "streaming"}[st-1]
}
func (st ServiceType) EnumIndex() int {
	return int(st)
}

// Declare VPNProtocols // Declaration order matters
type VpnProtocol int

const (
	OpenVpn VpnProtocol = iota + 1
	WireGuard
)

func (vp VpnProtocol) String() string {
	return [...]string{"OpenVPN", "WireGuard"}[vp-1]
}
func (vp VpnProtocol) CommandArg() string {
	return [...]string{"openvpn", "wireguard"}[vp-1]
}
func (vp VpnProtocol) EnumIndex() int {
	return int(vp)
}

// Declare TransmissionProtocols // Declaration order matters
type TransmissionProtocol int

const (
	Tcp TransmissionProtocol = iota + 1
	Udp
)

func (tp TransmissionProtocol) String() string {
	return [...]string{"TCP", "UDP"}[tp-1]
}
func (tp TransmissionProtocol) CommandArg() string {
	return [...]string{"tcp", "udp"}[tp-1]
}
func (tp TransmissionProtocol) EnumIndex() int {
	return int(tp)
}
