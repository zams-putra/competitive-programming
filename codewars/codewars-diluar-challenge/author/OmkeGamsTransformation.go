package author

// Name: OmkeGams Transformation

// Level: 8 kyu

// Discipline: algorithm

// Tags: crypthography

// Desc:
// Implement a function omkeGas Transformation that transforms a given word based on the following rules:
// -> Every consonant from the set 'p', 'b', and 'f' is replaced with the letter 'v'.
// -> If a vowel ('a', 'i', 'u', 'e', 'o') appears at an odd index
// (starting from 0), insert the letter 'm' after the vowel.
// -> Vowels at even indices remain unchanged.
// -> Return the transformed word.
// Example :
// oke gas -> omke gams
// rispek -> rimsvek
// apple -> amvlle

// Complete Solution :
func OmkeGamsTransform(word string) (r string) {
	d := true
	vowel := map[string]bool{}
	v := "aiueo"
	vlan := map[string]bool{}
	lan := "pbf"
	for _, l := range lan {
		vlan[string(l)] = true
	}
	for _, k := range v {
		vowel[string(k)] = true
	}
	for _, w := range word {
		strW := string(w)
		if vowel[strW] && d {
			d = false
			r += strW
			r += "m"
			continue
		} else if vowel[strW] {
			d = true
			r += strW
			continue
		} else if vlan[strW] {
			r += "v"
			continue
		}
		r += strW
	}
	return
}

// Initial solution :
// func omkeGamsTransform(word string) string {
// }

// Test Case :
// import (
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	. "codewarrior/kata"
//   )
//   var _ = Describe("omkeGamsTransform", func() {
// 	It("should return 'godvlan' for input 'godplan'", func() {
// 	  Expect(omkeGamsTransform("godplan")).To(Equal("godvlan"))
// 	})
// 	It("should return 'rimsvek' for input 'rispek'", func() {
// 	  Expect(omkeGamsTransform("rispek")).To(Equal("rimsvek"))
// 	})
// 	It("should return 'omke gams' for input 'oke gas'", func() {
// 	  Expect(omkeGamsTransform("oke gas")).To(Equal("omke gams"))
// 	})
//   })






// JS :
// function omkeGamsTransform(word) {
// 	let d = true;
// 	const vowels = new Set(['a', 'i', 'u', 'e', 'o']);
// 	const vlan = new Set(['p', 'b', 'f']);
// 	let result = '';
// 	for (let i = 0; i < word.length; i++) {
// 	  const char = word[i];
// 	  if (vowels.has(char) && d) {
// 		d = false;
// 		result += char + 'm';
// 		continue;
// 	  } else if (vowels.has(char)) {
// 		d = true;
// 		result += char;
// 		continue;
// 	  } else if (vlan.has(char)) {
// 		result += 'v';
// 		continue;
// 	  }
// 	  result += char;
// 	}
// 	return result;
//   }
  

// const chai = require("chai");
// const assert = chai.assert;
// chai.config.truncateThreshold = 0;
// describe("omkeGamsTransform", function() {
//   it("should return 'amvvle' for input 'apple'", function() {
//     assert.strictEqual(omkeGamsTransform("apple"), "amvvle");
//   });
//   it("should return 'rimsvek' for input 'rispek'", function() {
//     assert.strictEqual(omkeGamsTransform("rispek"), "rimsvek");
//   });
//   it("should return 'omke gams' for input 'oke gas'", function() {
//     assert.strictEqual(omkeGamsTransform("oke gas"), "omke gams");
//   });
// });
