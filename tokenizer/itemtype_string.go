// Code generated by "stringer -type=ItemType"; DO NOT EDIT.

package tokenizer

import "strconv"

const _ItemType_name = "ItemErrorT_ABSTRACTT_AND_EQUALT_ARRAYT_ARRAY_CASTT_AST_BAD_CHARACTERT_BOOLEAN_ANDT_BOOLEAN_ORT_BOOL_CASTT_BREAKT_CALLABLET_CASET_CATCHT_CLASST_CLASS_CT_CLONET_CLOSE_TAGT_COALESCET_COMMENTT_CONCAT_EQUALT_CONSTT_CONSTANT_ENCAPSED_STRINGT_CONTINUET_CURLY_OPENT_DECT_DECLARET_DEFAULTT_DIRT_DIV_EQUALT_DOC_COMMENTT_DOT_DOLLAR_OPEN_CURLY_BRACEST_DOUBLE_ARROWT_DOUBLE_CASTT_PAAMAYIM_NEKUDOTAYIMT_ECHOT_ELLIPSIST_ELSET_ELSEIFT_EMPTYT_ENCAPSED_AND_WHITESPACET_ENDDECLARET_ENDFORT_ENDFOREACHT_ENDIFT_ENDSWITCHT_ENDWHILET_END_HEREDOCT_EVALT_EXITT_EXTENDST_FILET_FINALT_FINALLYT_FORT_FOREACHT_FUNCTIONT_FUNC_CT_GLOBALT_GOTOT_HALT_COMPILERT_IFT_IMPLEMENTST_INCT_INCLUDET_INCLUDE_ONCET_INLINE_HTMLT_INSTANCEOFT_INSTEADOFT_INT_CASTT_INTERFACET_ISSETT_IS_EQUALT_IS_GREATER_OR_EQUALT_IS_IDENTICALT_IS_NOT_EQUALT_IS_NOT_IDENTICALT_IS_SMALLER_OR_EQUALT_SPACESHIPT_LINET_LISTT_LOGICAL_ANDT_LOGICAL_ORT_LOGICAL_XORT_METHOD_CT_MINUS_EQUALT_MOD_EQUALT_MUL_EQUALT_NAMESPACET_NS_CT_NS_SEPARATORT_NEWT_NUM_STRINGT_OBJECT_CASTT_OBJECT_OPERATORT_OPEN_TAGT_OPEN_TAG_WITH_ECHOT_OR_EQUALT_PLUS_EQUALT_POWT_POW_EQUALT_PRINTT_PRIVATET_PUBLICT_PROTECTEDT_REQUIRET_REQUIRE_ONCET_RETURNT_SLT_SL_EQUALT_SRT_SR_EQUALT_START_HEREDOCT_STATICT_STRINGT_STRING_CASTT_STRING_VARNAMET_SWITCHT_THROWT_TRAITT_TRAIT_CT_TRYT_UNSETT_UNSET_CASTT_USET_VART_VARIABLET_WHILET_WHITESPACET_XOR_EQUALT_YIELDT_YIELD_FROMT_DNUMBERT_LNUMBERT_EOF"

var _ItemType_index = [...]uint16{0, 9, 19, 30, 37, 49, 53, 68, 81, 93, 104, 111, 121, 127, 134, 141, 150, 157, 168, 178, 187, 201, 208, 234, 244, 256, 261, 270, 279, 284, 295, 308, 312, 338, 352, 365, 387, 393, 403, 409, 417, 424, 449, 461, 469, 481, 488, 499, 509, 522, 528, 534, 543, 549, 556, 565, 570, 579, 589, 597, 605, 611, 626, 630, 642, 647, 656, 670, 683, 695, 706, 716, 727, 734, 744, 765, 779, 793, 811, 832, 843, 849, 855, 868, 880, 893, 903, 916, 927, 938, 949, 955, 969, 974, 986, 999, 1016, 1026, 1046, 1056, 1068, 1073, 1084, 1091, 1100, 1108, 1119, 1128, 1142, 1150, 1154, 1164, 1168, 1178, 1193, 1201, 1209, 1222, 1238, 1246, 1253, 1260, 1269, 1274, 1281, 1293, 1298, 1303, 1313, 1320, 1332, 1343, 1350, 1362, 1371, 1380, 1385}

func (i ItemType) String() string {
	if i < 0 || i >= ItemType(len(_ItemType_index)-1) {
		return "ItemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ItemType_name[_ItemType_index[i]:_ItemType_index[i+1]]
}
