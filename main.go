//// 2023/10/21 // 21:23

//// Maps

// Maps are similar to JavaScript objects, Python dictioanries, and Ruby hashes. Maps
// are a data structure that provides key->value mapping.

// The zero value of a map is nil.

// // We can create a map by using a literal or by usign the make() function:
// ages := make(map[string]int)
// ages["John"] = 37
// ages["Mary"] = 24
// ages["Mary"] = 21 // overwrites 24

// ages = map[string]int{
// 	"John": 37,
// 	"Mary": 21,
// }

// // The len() function works on a map, it returns the total number of key/value pairs.
// ages = map[string]int{
// 	"John": 37,
// 	"Mary": 21,
// }
// fmt.Println(len(ages)) // 2

//// Assignment

// We can speed up our contact-info loopups by usign a map! Looking up a value in a
// map by its key is much faster than searching through a slice.

// Complete the getUserMap function. It takes a slice of names and a slice of phone
// numbers, and returns a map of name -> user structs and potentially an error. A
// user struct just contains a user's name and phoen number.

// If the length of names and phoneNumbers is not equal, return an error with the string
// "invalid sizes".

// // The first name in the names slice mathces the first phone number, and so on.

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
// 	users := make(map[string]user)
// 	if len(names) != len(phoneNumbers) {
// 		return nil, errors.New("invalid sizes")
// 	}
// 	for i := range names {
// 		name := names[i]
// 		users[name] = user{
// 			name:        name,
// 			phoneNumber: phoneNumbers[i],
// 		}
// 	}
// 	return users, nil
// }

// // don't touch below this line

// type user struct {
// 	name        string
// 	phoneNumber int
// }

// func test(names []string, phoneNumbers []int) {
// 	fmt.Println("Creating map...")
// 	defer fmt.Println("====================================")
// 	users, err := getUserMap(names, phoneNumbers)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, name := range names {
// 		fmt.Printf("key: %v, value:\n", name)
// 		fmt.Println(" - name:", users[name].name)
// 		fmt.Println(" - number:", users[name].phoneNumber)
// 	}
// }

// func main() {
// 	test(
// 		[]string{"John", "Bob", "Jill"},
// 		[]int{14355550987, 98765550987, 18265554567},
// 	)
// 	test(
// 		[]string{"John", "Bob"},
// 		[]int{14355550987, 98765550987, 18265554567},
// 	)
// 	test(
// 		[]string{"George", "Sally", "Rich", "Sue"},
// 		[]int{20955559812, 38385550982, 48265554567, 16045559873},
// 	)
// }

//// 21:36

//// Mutations

// // Insert an element
// m[key] = elem

// // Get an element
// elem = m[key]

// // Delete an element
// delete(m, key)

// // Check if a key exists
// elem, ok := m[key]

// If key is in n, then ok is true. If not, ok is false.
// If key is not in the map, then elem is the zero value for the map's element type.

// Assignment

// It's important to keep up with privacy regulations and to respect our user's data. We
// need a function that will delete user records.

// Complete the deleteIfNecessary function.
// * If the user doesn't exist in the map, return the error not found.
// * If they exist but aren't scheduled for deletion, return deleted as false with no
// errors.
// * If they exist and are scheduled for deletion, return deleted as true with no
// errors and delete their record from the map.

// Mote on passing maps
// Like spices, maps are also passed by reference into functions. This means that when a
// map is passed into a function we can write, we can make changes to the original, we
// don't have a copy.

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"sort"
// )

// func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
// 	user, ok := users[name]
// 	if !ok {
// 		return false, errors.New("not found")
// 	}
// 	if !user.scheduledForDeletion {
// 		return false, nil
// 	}
// 	delete(users, name)
// 	return true, nil
// }

// type user struct {
// 	name                 string
// 	number               int
// 	scheduledForDeletion bool
// }

// func test(users map[string]user, name string) {
// 	fmt.Printf("Attempting to delete %s...\n", name)
// 	defer fmt.Println("======================================")
// 	deleted, err := deleteIfNecessary(users, name)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if deleted {
// 		fmt.Println("Deleted:", name)
// 		return
// 	}
// 	fmt.Println("Did not delete:", name)
// }

// func main() {
// 	users := map[string]user{
// 		"john": {
// 			name:                 "john",
// 			number:               18965554631,
// 			scheduledForDeletion: true,
// 		},
// 		"elon": {
// 			name:                 "elon",
// 			number:               19875556452,
// 			scheduledForDeletion: true,
// 		},
// 		"breanna": {
// 			name:                 "breanna",
// 			number:               98575554231,
// 			scheduledForDeletion: false,
// 		},
// 		"kade": {
// 			name:                 "kade",
// 			number:               10765557221,
// 			scheduledForDeletion: false,
// 		},
// 	}
// 	test(users, "john")
// 	test(users, "musk")
// 	test(users, "santa")
// 	test(users, "kade")

// 	keys := []string{}
// 	for name := range users {
// 		keys = append(keys, name)
// 	}
// 	sort.Strings(keys)

// 	fmt.Println("Final map keys:")
// 	for _, name := range keys {
// 		fmt.Println(" - ", name)
// 	}
// }

//// 21:50

//// Key Types

// // Any type can be used as the value in a map, but keys are more restrictive.

// // As mentioned earlier, map keys may be of any type that is comparable. The
// // language spec defines this precisely, but in short, comparable types are boolean,
// // numeric, string, pointer, channel, and interface types, and structs or arrays that
// // contain only those types. Notably absent from the list are slices, maps, and functions;
// // these types cannot be compared using ==, and may not be used as map keys.

// // It's abvious that strings, ints, and other basic types should be available as map keys,
// // but perhaps unexpected are struct keys. Struct can be used to key data by multiple
// // dimensions. For example, this map of maps could be used to tally web page hits by
// // country:

// hits := make(map[string]map[string]int)

// // This is map of string to (map of string to int). Each key of the outer map is the path to
// // a web page with its own inner map. Each inner map key is a two-letter country code.
// // This expression retrieves the number of times an Australian has loaded the
// // documentation page:
// n := hits["/doc/"]["au"]

// // Unfortunately, this approach becomes unwieldy whe naddign data, as for any given
// // outer key you must check if the inner map exists, and create it if needed:
// func add(m map[string]map[string]int, path, country string) {
// 	mm, ok := m[path]
// 	if !ok {
// 		mm = make(map[string]int)
// 		m[path] == mm
// 	}
// 	mm[country]++
// }
// add(hits, "/doc/", "au")

// // On the other had, a design that uses a single map with a struct key does away with
// // all that complexity:
// type Key struct {
// 	Path, Country string
// }
// hits := make(map[Key]int)

// // When a Vietnamese person visits the home page, incrementing (and possible
// // creating) the appropriate counter is a one-liner:
// hits[Key{"/", "vn"}]++

// // And it's similaryly straightforward to see how many Swiss people have read the spec:
// n := hits[Key{"/ref/spec/", "ch"}]

//// 2023/10/22 // 16:18 //

//// 4-maps_count
//// Count Instances

// // Remember that you can check if a key is already present in a map by using the
// // second return value from the index operation.

// names := map[string]int{}

// if _, ok := names["elon"]; !ok {
// 	// if the key doesn't exist yet,
// 	// initialize its value to 0
// 	names["elon"] = 0
// }

//// Assignment

// // We have a slice of user ids, and each instance of an id in the slice indicates that a
// // message was sent to that user. We need to count up how many times each user's id
// // appears in the slice to track how many messages they received.

// // Implement the getCounts function. It should return a map of string -> int so that
// // each int is a count of how many times each string was found in the slice.

// package main

// import (
// 	"crypto/md5"
// 	"fmt"
// 	"io"
// )

// func getCounts(userIDs []string) map[string]int {
// 	m := map[string]int{}
// 	for _, id := range userIDs {
// 		if _, ok := m[id]; !ok {
// 			m[id] = 0
// 		}
// 		m[id]++
// 	}
// 	return m
// }

// func test(userIDs []string, ids []string) {
// 	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

// 	counts := getCounts(userIDs)
// 	fmt.Println("Counts from select IDs:")
// 	for _, k := range ids {
// 		v := counts[k]
// 		fmt.Printf(" - %s: %d\n", k, v)
// 	}
// 	fmt.Println("============================================")
// }

// func main() {
// 	userIDs := []string{}
// 	for i := 0; i < 10000; i++ {
// 		h := md5.New()
// 		io.WriteString(h, fmt.Sprint(i))
// 		key := fmt.Sprintf("%x", h.Sum(nil))
// 		userIDs = append(userIDs, key[:2])
// 	}

// 	test(userIDs, []string{"00", "ff", "dd"})
// 	test(userIDs, []string{"aa", "12", "32"})
// 	test(userIDs, []string{"bb", "33"})
// }

//// 16:29

//// Nested

// // Maps can contain maps, creating a nested structure. For example:
// map[string]map[string]int
// map[rune]map[string]int
// map[int]map[string]map[string]int

// //// Assignment

// // Because Textio is a glorified customer database, we have a lot of internal logic for
// // sorting and dealing with customer names.

// // Complete the getNameCounts function. It takes a slice of string (names) and returns
// // a nested map where the first key is all the unique first characters of the names, the
// // second key is all the names themselves, and the value is the count of each name.

// // For example:
// // billy
// // billy
// // bob
// // joe

// // Creates the folowing nested map:
// b: {
// 	billy: 2,
// 	bob: 1
// },
// j: {
// 	joe: 1
// }

// Note that the test suite is not printing the map you're returning directly, but instead
// checking some specific keys.

package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		firstLetter := rune(0)
		if len(name) != 0 {
			firstLetter = rune(name[0])
		}

		if _, ok := nameCounts[firstLetter]; !ok {
			nameCounts[firstLetter] = make(map[string]int)
		}
		if _, ok := nameCounts[firstLetter][name]; !ok {
			nameCounts[firstLetter][name] = 0
		}
		nameCounts[firstLetter][name]++
	}
	return nameCounts
}

func test(names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %v names...\n", len(names))

	nameCounts := getNameCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
	fmt.Println("=============================================")
}

func main() {
	test(getNames(50), 'M', "Matthew")
	test(getNames(100), 'G', "George")
	test(getNames(150), 'D', "Drew")
	test(getNames(200), 'P', "Phillip")
	test(getNames(250), 'B', "Bryant")
	test(getNames(300), 'M', "Matthew")
}

func getNames(length int) []string {
	names := []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew", "Peter", "Peter", "Henry", "Parker", "Parker", "Parker", "Collin", "Hayden", "George", "Bradley", "Mitchell", "Devon", "Ricardo", "Shawn", "Taylor", "Nicolas", "Gregory", "Francisco", "Liam", "Kaleb", "Preston", "Erik", "Alexis", "Owen", "Omar", "Diego", "Dustin", "Corey", "Fernando", "Clayton", "Carter", "Ivan", "Jaden", "Javier", "Alec", "Johnathan", "Scott", "Manuel", "Cristian", "Alan", "Raymond", "Brett", "Max", "Drew", "Andres", "Gage", "Mario", "Dawson", "Dillon", "Cesar", "Wesley", "Levi", "Jakob", "Chandler", "Martin", "Malik", "Edgar", "Sergio", "Trenton", "Josiah", "Nolan", "Marco", "Drew", "Peyton", "Harrison", "Drew", "Hector", "Micah", "Roberto", "Drew", "Brady", "Erick", "Conner", "Jonah", "Casey", "Jayden", "Edwin", "Emmanuel", "Andre", "Phillip", "Brayden", "Landon", "Giovanni", "Bailey", "Ronald", "Braden", "Damian", "Donovan", "Ruben", "Frank", "Gerardo", "Pedro", "Andy", "Chance", "Abraham", "Calvin", "Trey", "Cade", "Donald", "Derrick", "Payton", "Darius", "Enrique", "Keith", "Raul", "Jaylen", "Troy", "Jonathon", "Cory", "Marc", "Eli", "Skyler", "Rafael", "Trent", "Griffin", "Colby", "Johnny", "Chad", "Armando", "Kobe", "Caden", "Marcos", "Cooper", "Elias", "Brenden", "Israel", "Avery", "Zane", "Zane", "Zane", "Zane", "Dante", "Josue", "Zackary", "Allen", "Philip", "Mathew", "Dennis", "Leonardo", "Ashton", "Philip", "Philip", "Philip", "Julio", "Miles", "Damien", "Ty", "Gustavo", "Drake", "Jaime", "Simon", "Jerry", "Curtis", "Kameron", "Lance", "Brock", "Bryson", "Alberto", "Dominick", "Jimmy", "Kaden", "Douglas", "Gary", "Brennan", "Zachery", "Randy", "Louis", "Larry", "Nickolas", "Albert", "Tony", "Fabian", "Keegan", "Saul", "Danny", "Tucker", "Myles", "Damon", "Arturo", "Corbin", "Deandre", "Ricky", "Kristopher", "Lane", "Pablo", "Darren", "Jarrett", "Zion", "Alfredo", "Micheal", "Angelo", "Carl", "Oliver", "Kyler", "Tommy", "Walter", "Dallas", "Jace", "Quinn", "Theodore", "Grayson", "Lorenzo", "Joe", "Arthur", "Bryant", "Roman", "Brent", "Russell", "Ramon", "Lawrence", "Moises", "Aiden", "Quentin", "Jay", "Tyrese", "Tristen", "Emanuel", "Salvador", "Terry", "Morgan", "Jeffery", "Esteban", "Tyson", "Braxton", "Branden", "Marvin", "Brody", "Craig", "Ismael", "Rodney", "Isiah", "Marshall", "Maurice", "Ernesto", "Emilio", "Brendon", "Kody", "Eddie", "Malachi", "Abel", "Keaton", "Jon", "Shaun", "Skylar", "Ezekiel", "Nikolas", "Santiago", "Kendall", "Axel", "Camden", "Trevon", "Bobby", "Conor", "Jamal", "Lukas", "Malcolm", "Zackery", "Jayson", "Javon", "Roger", "Reginald", "Zachariah", "Desmond", "Felix", "Johnathon", "Dean", "Quinton", "Ali", "Davis", "Gerald", "Rodrigo", "Demetrius", "Billy", "Rene", "Reece", "Kelvin", "Leo", "Justice", "Chris", "Guillermo", "Matthew", "Matthew", "Matthew", "Kevon", "Steve", "Frederick", "Clay", "Weston", "Dorian", "Hugo", "Roy", "Orlando", "Terrance", "Kai", "Khalil", "Khalil", "Khalil", "Graham", "Noel", "Willie", "Nathanael", "Terrell", "Tyrone",
	}
	if length > len(names) {
		length = len(names)
	}
	return names[:length]
}
