package concatenation

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func join(in ...string) string {
	return strings.Join(in, " ")
}
func buff(in ...string) string {
	var b bytes.Buffer
	for _, s := range in {
		b.WriteString(s)
		b.WriteString(" ")
	}
	return b.String()
}

func sfmt(format string, in ...string) string {
	return fmt.Sprintf(format, in)
}

func ifmt(format string, in ...int) string {
	return fmt.Sprintf(format, in)
}
func ffmt(format string, in ...float64) string {
	return fmt.Sprintf(format, in)
}

func cnt100(in ...string) string {
	return in[0] + " " + in[1] + " " + in[2] + " " + in[3] + " " + in[4] + " " + in[5] + " " + in[6] + " " + in[7] + " " + in[8] + " " + in[9] + " " +
		in[10] + " " + in[11] + " " + in[12] + " " + in[13] + " " + in[14] + " " + in[15] + " " + in[16] + " " + in[17] + " " + in[18] + " " + in[19] + " " +
		in[20] + " " + in[21] + " " + in[22] + " " + in[23] + " " + in[24] + " " + in[25] + " " + in[26] + " " + in[27] + " " + in[28] + " " + in[29] + " " +
		in[30] + " " + in[31] + " " + in[32] + " " + in[33] + " " + in[34] + " " + in[35] + " " + in[36] + " " + in[37] + " " + in[38] + " " + in[39] + " " +
		in[40] + " " + in[41] + " " + in[42] + " " + in[43] + " " + in[44] + " " + in[45] + " " + in[46] + " " + in[47] + " " + in[48] + " " + in[49] + " " +
		in[50] + " " + in[51] + " " + in[52] + " " + in[53] + " " + in[54] + " " + in[55] + " " + in[56] + " " + in[57] + " " + in[58] + " " + in[59] + " " +
		in[60] + " " + in[61] + " " + in[62] + " " + in[63] + " " + in[64] + " " + in[65] + " " + in[66] + " " + in[67] + " " + in[68] + " " + in[69] + " " +
		in[70] + " " + in[71] + " " + in[72] + " " + in[73] + " " + in[74] + " " + in[75] + " " + in[76] + " " + in[77] + " " + in[78] + " " + in[79] + " " +
		in[80] + " " + in[81] + " " + in[82] + " " + in[83] + " " + in[84] + " " + in[85] + " " + in[86] + " " + in[87] + " " + in[88] + " " + in[89] + " " +
		in[90] + " " + in[91] + " " + in[92] + " " + in[93] + " " + in[94] + " " + in[95] + " " + in[96] + " " + in[97] + " " + in[98] + " " + in[99] + " "
}

func cnt100Int(in ...int) string {
	return strconv.Itoa(in[0]) + " " + strconv.Itoa(in[1]) + " " + strconv.Itoa(in[2]) + " " + strconv.Itoa(in[3]) + " " + strconv.Itoa(in[4]) + " " + strconv.Itoa(in[5]) + " " + strconv.Itoa(in[6]) + " " + strconv.Itoa(in[7]) + " " + strconv.Itoa(in[8]) + " " + strconv.Itoa(in[9]) + " " +
		strconv.Itoa(in[10]) + " " + strconv.Itoa(in[11]) + " " + strconv.Itoa(in[12]) + " " + strconv.Itoa(in[13]) + " " + strconv.Itoa(in[14]) + " " + strconv.Itoa(in[15]) + " " + strconv.Itoa(in[16]) + " " + strconv.Itoa(in[17]) + " " + strconv.Itoa(in[18]) + " " + strconv.Itoa(in[19]) + " " +
		strconv.Itoa(in[20]) + " " + strconv.Itoa(in[21]) + " " + strconv.Itoa(in[22]) + " " + strconv.Itoa(in[23]) + " " + strconv.Itoa(in[24]) + " " + strconv.Itoa(in[25]) + " " + strconv.Itoa(in[26]) + " " + strconv.Itoa(in[27]) + " " + strconv.Itoa(in[28]) + " " + strconv.Itoa(in[29]) + " " +
		strconv.Itoa(in[30]) + " " + strconv.Itoa(in[31]) + " " + strconv.Itoa(in[32]) + " " + strconv.Itoa(in[33]) + " " + strconv.Itoa(in[34]) + " " + strconv.Itoa(in[35]) + " " + strconv.Itoa(in[36]) + " " + strconv.Itoa(in[37]) + " " + strconv.Itoa(in[38]) + " " + strconv.Itoa(in[39]) + " " +
		strconv.Itoa(in[40]) + " " + strconv.Itoa(in[41]) + " " + strconv.Itoa(in[42]) + " " + strconv.Itoa(in[43]) + " " + strconv.Itoa(in[44]) + " " + strconv.Itoa(in[45]) + " " + strconv.Itoa(in[46]) + " " + strconv.Itoa(in[47]) + " " + strconv.Itoa(in[48]) + " " + strconv.Itoa(in[49]) + " " +
		strconv.Itoa(in[50]) + " " + strconv.Itoa(in[51]) + " " + strconv.Itoa(in[52]) + " " + strconv.Itoa(in[53]) + " " + strconv.Itoa(in[54]) + " " + strconv.Itoa(in[55]) + " " + strconv.Itoa(in[56]) + " " + strconv.Itoa(in[57]) + " " + strconv.Itoa(in[58]) + " " + strconv.Itoa(in[59]) + " " +
		strconv.Itoa(in[60]) + " " + strconv.Itoa(in[61]) + " " + strconv.Itoa(in[62]) + " " + strconv.Itoa(in[63]) + " " + strconv.Itoa(in[64]) + " " + strconv.Itoa(in[65]) + " " + strconv.Itoa(in[66]) + " " + strconv.Itoa(in[67]) + " " + strconv.Itoa(in[68]) + " " + strconv.Itoa(in[69]) + " " +
		strconv.Itoa(in[70]) + " " + strconv.Itoa(in[71]) + " " + strconv.Itoa(in[72]) + " " + strconv.Itoa(in[73]) + " " + strconv.Itoa(in[74]) + " " + strconv.Itoa(in[75]) + " " + strconv.Itoa(in[76]) + " " + strconv.Itoa(in[77]) + " " + strconv.Itoa(in[78]) + " " + strconv.Itoa(in[79]) + " " +
		strconv.Itoa(in[80]) + " " + strconv.Itoa(in[81]) + " " + strconv.Itoa(in[82]) + " " + strconv.Itoa(in[83]) + " " + strconv.Itoa(in[84]) + " " + strconv.Itoa(in[85]) + " " + strconv.Itoa(in[86]) + " " + strconv.Itoa(in[87]) + " " + strconv.Itoa(in[88]) + " " + strconv.Itoa(in[89]) + " " +
		strconv.Itoa(in[90]) + " " + strconv.Itoa(in[91]) + " " + strconv.Itoa(in[92]) + " " + strconv.Itoa(in[93]) + " " + strconv.Itoa(in[94]) + " " + strconv.Itoa(in[95]) + " " + strconv.Itoa(in[96]) + " " + strconv.Itoa(in[97]) + " " + strconv.Itoa(in[98]) + " " + strconv.Itoa(in[99]) + " "
}

func cnt100Float(in ...float64) string {
	return strconv.FormatFloat(in[0], 'g', -1, 64) + " " + strconv.FormatFloat(in[1], 'g', -1, 64) + " " + strconv.FormatFloat(in[2], 'g', -1, 64) + " " + strconv.FormatFloat(in[3], 'g', -1, 64) + " " + strconv.FormatFloat(in[4], 'g', -1, 64) + " " + strconv.FormatFloat(in[5], 'g', -1, 64) + " " + strconv.FormatFloat(in[6], 'g', -1, 64) + " " + strconv.FormatFloat(in[7], 'g', -1, 64) + " " + strconv.FormatFloat(in[8], 'g', -1, 64) + " " + strconv.FormatFloat(in[9], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[10], 'g', -1, 64) + " " + strconv.FormatFloat(in[11], 'g', -1, 64) + " " + strconv.FormatFloat(in[12], 'g', -1, 64) + " " + strconv.FormatFloat(in[13], 'g', -1, 64) + " " + strconv.FormatFloat(in[14], 'g', -1, 64) + " " + strconv.FormatFloat(in[15], 'g', -1, 64) + " " + strconv.FormatFloat(in[16], 'g', -1, 64) + " " + strconv.FormatFloat(in[17], 'g', -1, 64) + " " + strconv.FormatFloat(in[18], 'g', -1, 64) + " " + strconv.FormatFloat(in[19], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[20], 'g', -1, 64) + " " + strconv.FormatFloat(in[21], 'g', -1, 64) + " " + strconv.FormatFloat(in[22], 'g', -1, 64) + " " + strconv.FormatFloat(in[23], 'g', -1, 64) + " " + strconv.FormatFloat(in[24], 'g', -1, 64) + " " + strconv.FormatFloat(in[25], 'g', -1, 64) + " " + strconv.FormatFloat(in[26], 'g', -1, 64) + " " + strconv.FormatFloat(in[27], 'g', -1, 64) + " " + strconv.FormatFloat(in[28], 'g', -1, 64) + " " + strconv.FormatFloat(in[29], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[30], 'g', -1, 64) + " " + strconv.FormatFloat(in[31], 'g', -1, 64) + " " + strconv.FormatFloat(in[32], 'g', -1, 64) + " " + strconv.FormatFloat(in[33], 'g', -1, 64) + " " + strconv.FormatFloat(in[34], 'g', -1, 64) + " " + strconv.FormatFloat(in[35], 'g', -1, 64) + " " + strconv.FormatFloat(in[36], 'g', -1, 64) + " " + strconv.FormatFloat(in[37], 'g', -1, 64) + " " + strconv.FormatFloat(in[38], 'g', -1, 64) + " " + strconv.FormatFloat(in[39], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[40], 'g', -1, 64) + " " + strconv.FormatFloat(in[41], 'g', -1, 64) + " " + strconv.FormatFloat(in[42], 'g', -1, 64) + " " + strconv.FormatFloat(in[43], 'g', -1, 64) + " " + strconv.FormatFloat(in[44], 'g', -1, 64) + " " + strconv.FormatFloat(in[45], 'g', -1, 64) + " " + strconv.FormatFloat(in[46], 'g', -1, 64) + " " + strconv.FormatFloat(in[47], 'g', -1, 64) + " " + strconv.FormatFloat(in[48], 'g', -1, 64) + " " + strconv.FormatFloat(in[49], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[50], 'g', -1, 64) + " " + strconv.FormatFloat(in[51], 'g', -1, 64) + " " + strconv.FormatFloat(in[52], 'g', -1, 64) + " " + strconv.FormatFloat(in[53], 'g', -1, 64) + " " + strconv.FormatFloat(in[54], 'g', -1, 64) + " " + strconv.FormatFloat(in[55], 'g', -1, 64) + " " + strconv.FormatFloat(in[56], 'g', -1, 64) + " " + strconv.FormatFloat(in[57], 'g', -1, 64) + " " + strconv.FormatFloat(in[58], 'g', -1, 64) + " " + strconv.FormatFloat(in[59], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[60], 'g', -1, 64) + " " + strconv.FormatFloat(in[61], 'g', -1, 64) + " " + strconv.FormatFloat(in[62], 'g', -1, 64) + " " + strconv.FormatFloat(in[63], 'g', -1, 64) + " " + strconv.FormatFloat(in[64], 'g', -1, 64) + " " + strconv.FormatFloat(in[65], 'g', -1, 64) + " " + strconv.FormatFloat(in[66], 'g', -1, 64) + " " + strconv.FormatFloat(in[67], 'g', -1, 64) + " " + strconv.FormatFloat(in[68], 'g', -1, 64) + " " + strconv.FormatFloat(in[69], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[70], 'g', -1, 64) + " " + strconv.FormatFloat(in[71], 'g', -1, 64) + " " + strconv.FormatFloat(in[72], 'g', -1, 64) + " " + strconv.FormatFloat(in[73], 'g', -1, 64) + " " + strconv.FormatFloat(in[74], 'g', -1, 64) + " " + strconv.FormatFloat(in[75], 'g', -1, 64) + " " + strconv.FormatFloat(in[76], 'g', -1, 64) + " " + strconv.FormatFloat(in[77], 'g', -1, 64) + " " + strconv.FormatFloat(in[78], 'g', -1, 64) + " " + strconv.FormatFloat(in[79], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[80], 'g', -1, 64) + " " + strconv.FormatFloat(in[81], 'g', -1, 64) + " " + strconv.FormatFloat(in[82], 'g', -1, 64) + " " + strconv.FormatFloat(in[83], 'g', -1, 64) + " " + strconv.FormatFloat(in[84], 'g', -1, 64) + " " + strconv.FormatFloat(in[85], 'g', -1, 64) + " " + strconv.FormatFloat(in[86], 'g', -1, 64) + " " + strconv.FormatFloat(in[87], 'g', -1, 64) + " " + strconv.FormatFloat(in[88], 'g', -1, 64) + " " + strconv.FormatFloat(in[89], 'g', -1, 64) + " " +
		strconv.FormatFloat(in[90], 'g', -1, 64) + " " + strconv.FormatFloat(in[91], 'g', -1, 64) + " " + strconv.FormatFloat(in[92], 'g', -1, 64) + " " + strconv.FormatFloat(in[93], 'g', -1, 64) + " " + strconv.FormatFloat(in[94], 'g', -1, 64) + " " + strconv.FormatFloat(in[95], 'g', -1, 64) + " " + strconv.FormatFloat(in[96], 'g', -1, 64) + " " + strconv.FormatFloat(in[97], 'g', -1, 64) + " " + strconv.FormatFloat(in[98], 'g', -1, 64) + " " + strconv.FormatFloat(in[99], 'g', -1, 64) + " "
}
