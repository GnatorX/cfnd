// Code generated by go-enum
// DO NOT EDIT!

package cloudformation

import (
	"fmt"
)

const (
	// StackStatusCREATEINPROGRESS is a StackStatus of type CREATE_IN_PROGRESS
	StackStatusCREATEINPROGRESS StackStatus = iota
	// StackStatusCREATEFAILED is a StackStatus of type CREATE_FAILED
	StackStatusCREATEFAILED
	// StackStatusCREATECOMPLETE is a StackStatus of type CREATE_COMPLETE
	StackStatusCREATECOMPLETE
	// StackStatusROLLBACKINPROGRESS is a StackStatus of type ROLLBACK_IN_PROGRESS
	StackStatusROLLBACKINPROGRESS
	// StackStatusROLLBACKFAILED is a StackStatus of type ROLLBACK_FAILED
	StackStatusROLLBACKFAILED
	// StackStatusROLLBACKCOMPLETE is a StackStatus of type ROLLBACK_COMPLETE
	StackStatusROLLBACKCOMPLETE
	// StackStatusDELETEINPROGRESS is a StackStatus of type DELETE_IN_PROGRESS
	StackStatusDELETEINPROGRESS
	// StackStatusDELETEFAILED is a StackStatus of type DELETE_FAILED
	StackStatusDELETEFAILED
	// StackStatusDELETECOMPLETE is a StackStatus of type DELETE_COMPLETE
	StackStatusDELETECOMPLETE
	// StackStatusUPDATEINPROGRESS is a StackStatus of type UPDATE_IN_PROGRESS
	StackStatusUPDATEINPROGRESS
	// StackStatusUPDATECOMPLETECLEANUPINPROGRESS is a StackStatus of type UPDATE_COMPLETE_CLEANUP_IN_PROGRESS
	StackStatusUPDATECOMPLETECLEANUPINPROGRESS
	// StackStatusUPDATECOMPLETE is a StackStatus of type UPDATE_COMPLETE
	StackStatusUPDATECOMPLETE
	// StackStatusUPDATEROLLBACKINPROGRESS is a StackStatus of type UPDATE_ROLLBACK_IN_PROGRESS
	StackStatusUPDATEROLLBACKINPROGRESS
	// StackStatusUPDATEROLLBACKFAILED is a StackStatus of type UPDATE_ROLLBACK_FAILED
	StackStatusUPDATEROLLBACKFAILED
	// StackStatusUPDATEROLLBACKCOMPLETECLEANUPINPROGRESS is a StackStatus of type UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS
	StackStatusUPDATEROLLBACKCOMPLETECLEANUPINPROGRESS
	// StackStatusUPDATEROLLBACKCOMPLETE is a StackStatus of type UPDATE_ROLLBACK_COMPLETE
	StackStatusUPDATEROLLBACKCOMPLETE
	// StackStatusREVIEWINPROGRESS is a StackStatus of type REVIEW_IN_PROGRESS
	StackStatusREVIEWINPROGRESS
	// StackStatusIMPORTINPROGRESS is a StackStatus of type IMPORT_IN_PROGRESS
	StackStatusIMPORTINPROGRESS
	// StackStatusIMPORTCOMPLETE is a StackStatus of type IMPORT_COMPLETE
	StackStatusIMPORTCOMPLETE
	// StackStatusIMPORTROLLBACKINPROGRESS is a StackStatus of type IMPORT_ROLLBACK_IN_PROGRESS
	StackStatusIMPORTROLLBACKINPROGRESS
	// StackStatusIMPORTROLLBACKFAILED is a StackStatus of type IMPORT_ROLLBACK_FAILED
	StackStatusIMPORTROLLBACKFAILED
	// StackStatusIMPORTROLLBACKCOMPLETE is a StackStatus of type IMPORT_ROLLBACK_COMPLETE
	StackStatusIMPORTROLLBACKCOMPLETE
)

const _StackStatusName = "CREATE_IN_PROGRESSCREATE_FAILEDCREATE_COMPLETEROLLBACK_IN_PROGRESSROLLBACK_FAILEDROLLBACK_COMPLETEDELETE_IN_PROGRESSDELETE_FAILEDDELETE_COMPLETEUPDATE_IN_PROGRESSUPDATE_COMPLETE_CLEANUP_IN_PROGRESSUPDATE_COMPLETEUPDATE_ROLLBACK_IN_PROGRESSUPDATE_ROLLBACK_FAILEDUPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESSUPDATE_ROLLBACK_COMPLETEREVIEW_IN_PROGRESSIMPORT_IN_PROGRESSIMPORT_COMPLETEIMPORT_ROLLBACK_IN_PROGRESSIMPORT_ROLLBACK_FAILEDIMPORT_ROLLBACK_COMPLETE"

var _StackStatusMap = map[StackStatus]string{
	0:  _StackStatusName[0:18],
	1:  _StackStatusName[18:31],
	2:  _StackStatusName[31:46],
	3:  _StackStatusName[46:66],
	4:  _StackStatusName[66:81],
	5:  _StackStatusName[81:98],
	6:  _StackStatusName[98:116],
	7:  _StackStatusName[116:129],
	8:  _StackStatusName[129:144],
	9:  _StackStatusName[144:162],
	10: _StackStatusName[162:197],
	11: _StackStatusName[197:212],
	12: _StackStatusName[212:239],
	13: _StackStatusName[239:261],
	14: _StackStatusName[261:305],
	15: _StackStatusName[305:329],
	16: _StackStatusName[329:347],
	17: _StackStatusName[347:365],
	18: _StackStatusName[365:380],
	19: _StackStatusName[380:407],
	20: _StackStatusName[407:429],
	21: _StackStatusName[429:453],
}

// String implements the Stringer interface.
func (x StackStatus) String() string {
	if str, ok := _StackStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("StackStatus(%d)", x)
}

var _StackStatusValue = map[string]StackStatus{
	_StackStatusName[0:18]:    0,
	_StackStatusName[18:31]:   1,
	_StackStatusName[31:46]:   2,
	_StackStatusName[46:66]:   3,
	_StackStatusName[66:81]:   4,
	_StackStatusName[81:98]:   5,
	_StackStatusName[98:116]:  6,
	_StackStatusName[116:129]: 7,
	_StackStatusName[129:144]: 8,
	_StackStatusName[144:162]: 9,
	_StackStatusName[162:197]: 10,
	_StackStatusName[197:212]: 11,
	_StackStatusName[212:239]: 12,
	_StackStatusName[239:261]: 13,
	_StackStatusName[261:305]: 14,
	_StackStatusName[305:329]: 15,
	_StackStatusName[329:347]: 16,
	_StackStatusName[347:365]: 17,
	_StackStatusName[365:380]: 18,
	_StackStatusName[380:407]: 19,
	_StackStatusName[407:429]: 20,
	_StackStatusName[429:453]: 21,
}

// ParseStackStatus attempts to convert a string to a StackStatus
func ParseStackStatus(name string) (StackStatus, error) {
	if x, ok := _StackStatusValue[name]; ok {
		return x, nil
	}
	return StackStatus(0), fmt.Errorf("%s is not a valid StackStatus", name)
}
