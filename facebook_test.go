// A facebook graph api client in go.
// https://github.com/huandu/facebook/
//
// Copyright 2012, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/facebook/blob/master/LICENSE

package facebook

// Constants for test cases.
const (
	FB_TEST_APP_ID     = "169186383097898"
	FB_TEST_APP_SECRET = "b2e4262c306caa3c7f5215d2d099b319"

	// remeber to change it to a valid token to run test
	FB_TEST_VALID_ACCESS_TOKEN = ""

	// remember to change it to a valid signed request to run test
	//FB_TEST_VALID_SIGNED_REQUEST = "ZAxP-ILRQBOwKKxCBMNlGmVraiowV7WFNg761OYBNGc.eyJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsImV4cGlyZXMiOjEzNDM0OTg0MDAsImlzc3VlZF9hdCI6MTM0MzQ5MzI2NSwib2F1dGhfdG9rZW4iOiJBQUFDWkEzOFpBRDhDb0JBRFpCcmZ5TFpDanBNUVczdThVTWZmRldSWkNpZGw5Tkx4a1BsY2tTcXZaQnpzTW9OWkF2bVk2RUd2NG1hUUFaQ0t2VlpBWkJ5VXA5a0FCU2x6THFJejlvZTdOdHBzdzhyQVpEWkQiLCJ1c2VyIjp7ImNvdW50cnkiOiJ1cyIsImxvY2FsZSI6ImVuX1VTIiwiYWdlIjp7Im1pbiI6MjF9fSwidXNlcl9pZCI6IjUzODc0NDQ2OCJ9"
	FB_TEST_VALID_SIGNED_REQUEST = ""

	// test binary file base64 value
	FB_TEST_BINARY_JPG_FILE = "/9j/4AAQSkZJRgABAQEASABIAAD/4gv4SUNDX1BST0ZJTEUAAQEAAAvoAAAAAAIAAABtbnRy" +
		"UkdCIFhZWiAH2QADABsAFQAkAB9hY3NwAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAA" +
		"9tYAAQAAAADTLQAAAAAp+D3er/JVrnhC+uTKgzkNAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAABBkZXNjAAABRAAAAHliWFlaAAABwAAAABRiVFJDAAAB1AAACAxkbWRkAAAJ4AAA" +
		"AIhnWFlaAAAKaAAAABRnVFJDAAAB1AAACAxsdW1pAAAKfAAAABRtZWFzAAAKkAAAACRia3B0" +
		"AAAKtAAAABRyWFlaAAAKyAAAABRyVFJDAAAB1AAACAx0ZWNoAAAK3AAAAAx2dWVkAAAK6AAA" +
		"AId3dHB0AAALcAAAABRjcHJ0AAALhAAAADdjaGFkAAALvAAAACxkZXNjAAAAAAAAAB9zUkdC" +
		"IElFQzYxOTY2LTItMSBibGFjayBzY2FsZWQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"WFlaIAAAAAAAACSgAAAPhAAAts9jdXJ2AAAAAAAABAAAAAAFAAoADwAUABkAHgAjACgALQAy" +
		"ADcAOwBAAEUASgBPAFQAWQBeAGMAaABtAHIAdwB8AIEAhgCLAJAAlQCaAJ8ApACpAK4AsgC3" +
		"ALwAwQDGAMsA0ADVANsA4ADlAOsA8AD2APsBAQEHAQ0BEwEZAR8BJQErATIBOAE+AUUBTAFS" +
		"AVkBYAFnAW4BdQF8AYMBiwGSAZoBoQGpAbEBuQHBAckB0QHZAeEB6QHyAfoCAwIMAhQCHQIm" +
		"Ai8COAJBAksCVAJdAmcCcQJ6AoQCjgKYAqICrAK2AsECywLVAuAC6wL1AwADCwMWAyEDLQM4" +
		"A0MDTwNaA2YDcgN+A4oDlgOiA64DugPHA9MD4APsA/kEBgQTBCAELQQ7BEgEVQRjBHEEfgSM" +
		"BJoEqAS2BMQE0wThBPAE/gUNBRwFKwU6BUkFWAVnBXcFhgWWBaYFtQXFBdUF5QX2BgYGFgYn" +
		"BjcGSAZZBmoGewaMBp0GrwbABtEG4wb1BwcHGQcrBz0HTwdhB3QHhgeZB6wHvwfSB+UH+AgL" +
		"CB8IMghGCFoIbgiCCJYIqgi+CNII5wj7CRAJJQk6CU8JZAl5CY8JpAm6Cc8J5Qn7ChEKJwo9" +
		"ClQKagqBCpgKrgrFCtwK8wsLCyILOQtRC2kLgAuYC7ALyAvhC/kMEgwqDEMMXAx1DI4MpwzA" +
		"DNkM8w0NDSYNQA1aDXQNjg2pDcMN3g34DhMOLg5JDmQOfw6bDrYO0g7uDwkPJQ9BD14Peg+W" +
		"D7MPzw/sEAkQJhBDEGEQfhCbELkQ1xD1ERMRMRFPEW0RjBGqEckR6BIHEiYSRRJkEoQSoxLD" +
		"EuMTAxMjE0MTYxODE6QTxRPlFAYUJxRJFGoUixStFM4U8BUSFTQVVhV4FZsVvRXgFgMWJhZJ" +
		"FmwWjxayFtYW+hcdF0EXZReJF64X0hf3GBsYQBhlGIoYrxjVGPoZIBlFGWsZkRm3Gd0aBBoq" +
		"GlEadxqeGsUa7BsUGzsbYxuKG7Ib2hwCHCocUhx7HKMczBz1HR4dRx1wHZkdwx3sHhYeQB5q" +
		"HpQevh7pHxMfPh9pH5Qfvx/qIBUgQSBsIJggxCDwIRwhSCF1IaEhziH7IiciVSKCIq8i3SMK" +
		"IzgjZiOUI8Ij8CQfJE0kfCSrJNolCSU4JWgllyXHJfcmJyZXJocmtyboJxgnSSd6J6sn3CgN" +
		"KD8ocSiiKNQpBik4KWspnSnQKgIqNSpoKpsqzysCKzYraSudK9EsBSw5LG4soizXLQwtQS12" +
		"Last4S4WLkwugi63Lu4vJC9aL5Evxy/+MDUwbDCkMNsxEjFKMYIxujHyMioyYzKbMtQzDTNG" +
		"M38zuDPxNCs0ZTSeNNg1EzVNNYc1wjX9Njc2cjauNuk3JDdgN5w31zgUOFA4jDjIOQU5Qjl/" +
		"Obw5+To2OnQ6sjrvOy07azuqO+g8JzxlPKQ84z0iPWE9oT3gPiA+YD6gPuA/IT9hP6I/4kAj" +
		"QGRApkDnQSlBakGsQe5CMEJyQrVC90M6Q31DwEQDREdEikTORRJFVUWaRd5GIkZnRqtG8Ec1" +
		"R3tHwEgFSEtIkUjXSR1JY0mpSfBKN0p9SsRLDEtTS5pL4kwqTHJMuk0CTUpNk03cTiVObk63" +
		"TwBPSU+TT91QJ1BxULtRBlFQUZtR5lIxUnxSx1MTU19TqlP2VEJUj1TbVShVdVXCVg9WXFap" +
		"VvdXRFeSV+BYL1h9WMtZGllpWbhaB1pWWqZa9VtFW5Vb5Vw1XIZc1l0nXXhdyV4aXmxevV8P" +
		"X2Ffs2AFYFdgqmD8YU9homH1YklinGLwY0Njl2PrZEBklGTpZT1lkmXnZj1mkmboZz1nk2fp" +
		"aD9olmjsaUNpmmnxakhqn2r3a09rp2v/bFdsr20IbWBtuW4SbmtuxG8eb3hv0XArcIZw4HE6" +
		"cZVx8HJLcqZzAXNdc7h0FHRwdMx1KHWFdeF2Pnabdvh3VnezeBF4bnjMeSp5iXnnekZ6pXsE" +
		"e2N7wnwhfIF84X1BfaF+AX5ifsJ/I3+Ef+WAR4CogQqBa4HNgjCCkoL0g1eDuoQdhICE44VH" +
		"hauGDoZyhteHO4efiASIaYjOiTOJmYn+imSKyoswi5aL/IxjjMqNMY2Yjf+OZo7OjzaPnpAG" +
		"kG6Q1pE/kaiSEZJ6kuOTTZO2lCCUipT0lV+VyZY0lp+XCpd1l+CYTJi4mSSZkJn8mmia1ZtC" +
		"m6+cHJyJnPedZJ3SnkCerp8dn4uf+qBpoNihR6G2oiailqMGo3aj5qRWpMelOKWpphqmi6b9" +
		"p26n4KhSqMSpN6mpqhyqj6sCq3Wr6axcrNCtRK24ri2uoa8Wr4uwALB1sOqxYLHWskuywrM4" +
		"s660JbSctRO1irYBtnm28Ldot+C4WbjRuUq5wro7urW7LrunvCG8m70VvY++Cr6Evv+/er/1" +
		"wHDA7MFnwePCX8Lbw1jD1MRRxM7FS8XIxkbGw8dBx7/IPci8yTrJuco4yrfLNsu2zDXMtc01" +
		"zbXONs62zzfPuNA50LrRPNG+0j/SwdNE08bUSdTL1U7V0dZV1tjXXNfg2GTY6Nls2fHadtr7" +
		"24DcBdyK3RDdlt4c3qLfKd+v4DbgveFE4cziU+Lb42Pj6+Rz5PzlhOYN5pbnH+ep6DLovOlG" +
		"6dDqW+rl63Dr++yG7RHtnO4o7rTvQO/M8Fjw5fFy8f/yjPMZ86f0NPTC9VD13vZt9vv3ivgZ" +
		"+Kj5OPnH+lf65/t3/Af8mP0p/br+S/7c/23//2Rlc2MAAAAAAAAALklFQyA2MTk2Ni0yLTEg" +
		"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"

	FB_LATEST_VERSION = "v3.1"
)

var (
	testGlobalApp *App
)

func init() {
	testGlobalApp = New(FB_TEST_APP_ID, FB_TEST_APP_SECRET)
	testGlobalApp.EnableAppsecretProof = true
}
// cara que massaaaaaaaaaaaaaaaaaaa