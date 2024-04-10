	var tbs cryptobyte.String
	wtf := func(s *cryptobyte.String) {
		switch {
		case s.PeekASN1Tag(cryptobyte_asn1.BIT_STRING):
			fmt.Println("bit string")
		case s.PeekASN1Tag(cryptobyte_asn1.BOOLEAN):
			fmt.Println("boolean")
		case s.PeekASN1Tag(cryptobyte_asn1.ENUM):
			fmt.Println("enum")
		case s.PeekASN1Tag(cryptobyte_asn1.GeneralString):
			fmt.Println("general string")
		case s.PeekASN1Tag(cryptobyte_asn1.GeneralizedTime):
			fmt.Println("gen time")
		case s.PeekASN1Tag(cryptobyte_asn1.IA5String):
			fmt.Println("ia5string")
		case s.PeekASN1Tag(cryptobyte_asn1.INTEGER):
			fmt.Println("int")
		case s.PeekASN1Tag(cryptobyte_asn1.OBJECT_IDENTIFIER):
			fmt.Println("OID")
		case s.PeekASN1Tag(cryptobyte_asn1.NULL):
			fmt.Println("null")
		case s.PeekASN1Tag(cryptobyte_asn1.OCTET_STRING):
			fmt.Println("octet string")
		case s.PeekASN1Tag(cryptobyte_asn1.PrintableString):
			fmt.Println("printable string")
		case s.PeekASN1Tag(cryptobyte_asn1.SEQUENCE):
			fmt.Println("sequins")
		case s.PeekASN1Tag(cryptobyte_asn1.SET):
			fmt.Println("set")
		case s.PeekASN1Tag(cryptobyte_asn1.T61String):
			fmt.Println("t61 string")
		case s.PeekASN1Tag(cryptobyte_asn1.UTCTime):
			fmt.Println("utc time")
		case s.PeekASN1Tag(cryptobyte_asn1.UTF8String):
			fmt.Println("utf8 string")
		default:
			fmt.Println("no idear")
		}
	}
