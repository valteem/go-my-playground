// Mighty fucking great attempt to tackle 'String Matching'. Not very successful, actually.

package bytesm // one folder - one package

func Count(s []byte, sep byte) int {

	if len(s) == 0 {
		return -1 // nowhere to look for
	}

	if sep == 0 {
		return -99 // nothing to look for
	}

/* 	if Index(s, sep) == -1 {
//  Need to first implement Index, IndexByte, IndexRabinKarp and some other stuff
	}
 */

 return 0 // that will do for the moment
}
