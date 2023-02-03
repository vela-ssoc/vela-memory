package memory

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/vela-ssoc/vela-kit/kind"
	"time"
)

type snapshot struct {
	time int64
	err  error
}

type summary struct {
	Total        uint64  `json:"total"`
	Free         uint64  `json:"free"`
	Used         uint64  `json:"used"`
	Available    uint64  `json:"available"`
	UsedPct      float64 `json:"used_pct"`
	SwapTotal    uint64  `json:"swap_total"`
	SwapFree     uint64  `json:"swap_free"`
	SwapUsedPct  float64 `json:"swap_used_pct"`
	SwapInPages  uint64  `json:"swap_in_pages"`
	SwapOutPages uint64  `json:"swap_out_pages"`
	snap         snapshot
}

func New() *summary {
	return &summary{}
}

func (sum *summary) Memory() {
	st, err := mem.VirtualMemory()
	if err != nil {
		xEnv.Errorf("get memory stats error: %v", err)
		sum.snap.err = err
		return
	}

	sum.Total = st.Total
	sum.Free = st.Free
	sum.Used = st.Used
	sum.UsedPct = st.UsedPercent
	sum.Available = st.Available
}

func (sum *summary) Swap() {
	swap, err := mem.SwapMemory()
	if err != nil {
		xEnv.Errorf("got swap memory fail %v", err)
		return
	}

	sum.SwapTotal = swap.Total
	sum.SwapInPages = swap.PgIn
	sum.SwapOutPages = swap.PgOut
	sum.SwapFree = swap.Free
	sum.SwapUsedPct = float64(swap.Total-swap.Free) / float64(swap.Total)
}

func (sum *summary) Byte() []byte {
	buf := kind.NewJsonEncoder()
	buf.Tab("")
	buf.KL("total", int64(sum.Total))
	buf.KL("free", int64(sum.Free))
	buf.KL("used_pct", int64(sum.UsedPct))
	buf.KL("swap_total", int64(sum.SwapTotal))
	buf.KL("swap_free", int64(sum.SwapFree))
	buf.KL("swap_used_pct", int64(sum.SwapUsedPct))
	buf.KL("swap_in_pages", int64(sum.SwapInPages))
	buf.KL("swap_out_pages", int64(sum.SwapOutPages))
	buf.End("}")
	return buf.Bytes()
}

func (sum *summary) Update() {
	now := time.Now().Unix()
	if now-sum.snap.time > 5 {
		sum.snap.time = now
		sum.Memory()
		sum.Swap()
	}
}
