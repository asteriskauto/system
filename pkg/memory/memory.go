
// Package memory defines Memory resource.
 //
 // This collector reports on the following meminfo metrics:
 //
		func (m *Memory) Report() {
                        m.client.Gauge("swap.percent", swapPercent(stat))
 
                        if m.Extended {
                               m.client.Gauge("total", bytes(stat.MemTotal))
                                m.client.Gauge("used", bytes(used(stat)))

                               m.client.Gauge("free", bytes(stat.MemFree))
                               m.client.Gauge("active", bytes(stat.Active))
                               m.client.Gauge("swap.total", bytes(stat.SwapTotal))
                               m.client.Gauge("swap.free", bytes(stat.SwapFree))
                        }
 
                case <-m.exit:
			func (m *Memory) Stop() error {
 }
 
 // calculate swap percentage.

func swapPercent(s *linux.MemInfo) int {
       total := s.SwapTotal
       used := total - s.SwapFree
	p := float64(used) / float64(total) * 100
 
        if math.IsNaN(p) {
		func swapPercent(s linux.MemInfo) int {
 }
 
 // calculate percentage.

func percent(s *linux.MemInfo) int {
       total := s.MemTotal
        p := float64(used(s)) / float64(total) * 100
 
        if math.IsNaN(p) {
		func percent(s linux.MemInfo) int {
 }
 
 // used memory.

func used(s *linux.MemInfo) uint64 {
       return s.MemTotal - s.MemFree - s.Buffers - s.Cached
 }
 
 // convert to bytes.
