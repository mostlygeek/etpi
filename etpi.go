package etpi

type ArmMode int

const (
	ArmAway = iota + 1
	ArmStay
	ArmNoEntryDelay
)

type PanelStatus struct {
	Zone      []ZoneStatus
	Partition []PartitionStatus
	Keypad    KeypadStatus
}

type ZoneStatus int

const UnknownStatus = 0

const (
	ZoneStatusAlarm = iota + 1
	ZoneStatusTamper
	ZoneStatusFault
	ZoneStatusOpen
	ZoneStatusRestored
)

func (z ZoneStatus) String() string {
	switch z {
	case ZoneStatusAlarm:
		return "ALARM"
	case ZoneStatusTamper:
		return "TAMPER"
	case ZoneStatusFault:
		return "FAULT"
	case ZoneStatusOpen:
		return "OPEN"
	case ZoneStatusRestored:
		return "RESTORED"
	default:
		return "UNKNOWN"
	}
}

type PartitionStatus int

const (
	PartitionStatusReady = iota + 1
	PartitionStatusNotReady
	PartitionStatusArmedAway
	PartitionStatusArmedStay
	PartitionStatusArmedZeroEntryAway
	PartitionStatusArmedZeroEntryStay
	PartitionStatusAlarm
	PartitionStatusDisarmed
	PartitionStatusExitDelay
	PartitionStatusEntryDelay
	PartitionStatusFailedToArm
	PartitionStatusBusy
)

func (p PartitionStatus) String() string {
	switch p {
	case PartitionStatusReady:
		return "DISARMED_READY"
	case PartitionStatusNotReady:
		return "DISARMED_NOT_READY"
	case PartitionStatusArmedAway:
		return "ARMED_AWAY"
	case PartitionStatusArmedStay:
		return "ARMED_STAY"
	case PartitionStatusArmedZeroEntryAway:
		return "ARMED_ZERO_ENTRY_AWAY"
	case PartitionStatusArmedZeroEntryStay:
		return "ARMED_ZERO_ENTRY_STAY"
	case PartitionStatusAlarm:
		return "ALARM"
	case PartitionStatusDisarmed:
		return "DISARMED"
	case PartitionStatusExitDelay:
		return "EXIT_DELAY"
	case PartitionStatusEntryDelay:
		return "ENTRY_DELAY"
	case PartitionStatusFailedToArm:
		return "FAILED_TO_ARM"
	case PartitionStatusBusy:
		return "BUSY"
	default:
		return "UNKNOWN"
	}
}

type KeypadStatus struct {
	Backlight bool
	Fire      bool
	Program   bool
	Trouble   bool
	Bypass    bool
	Memory    bool
	Armed     bool
	Ready     bool
}
