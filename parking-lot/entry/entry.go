package entry

type EntryGate struct {
	GateNo    int
	Available bool
}

func EntryGates(totalGate int) []*EntryGate {
	gates := []*EntryGate{}
	for i := 1; i <= totalGate; i++ {
		gate := &EntryGate{GateNo: i, Available: true}
		gates = append(gates, gate)
	}
	return gates
}

// func (entryGate *EntryGate) GiveEntry() bool {

// }
