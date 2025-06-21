- `Flyweight` is a space optimization technique that lets us use less memory by storing externally the data associated with similar objects.
- This is basically used to save memory.

- Let's take an example of storing formatted text.
- The unoptimized approach is to save the plain text along with the a boolean array of where the capital letters could be.
- This leads to the fact that for every text, there's a boolean array of the same length which is mostly empty.
- Instead what we could do is to have a struct which contains the start and end of the formatting along with boolean flags of what to do.

```go
type TextRange struct {
 Start, End               int
 Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
 return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
 plainText  string
 formatting []*TextRange
}
func (b *BetterFormattedText) Range(start, end int) *TextRange {
 r := &TextRange{
  Start: start,
  End:   end,
 }
 b.formatting = append(b.formatting, r)
 return r
}
func (b *BetterFormattedText) String() string {
 sb := strings.Builder{}

 for i := 0; i < len(b.plainText); i++ {
  c := b.plainText[i]
  for _, r := range b.formatting {
   if r.Covers(i) {
    c = byte(unicode.ToUpper(rune(c)))
   }
  }
  sb.WriteRune(rune(c))
 }

 return sb.String()
}

bft := NewBetterFormattedText("This is a brave new world")
bft.Range(10, 15).Capitalize = true
```

- The last line will have the same effect as making `brave` capital instead of saving the whole boolean array.

## Usernames

- Other very common usecase of flyweight pattern is saving usernames
- If we save usernames of the user as full strings, it will lead to a lot of memory usage as there are many repeating firstnames or lastnames or even fullnames.
- So an optimized approach is to maintain a list of names ( both first name and last name ) and store the fullname as an array of the indices of the different parts of name like this -

```go
var allNames []string

type User2 struct {
 names []uint8
}

func (u *User2) FullName() string {
 var parts []string
 for _, id := range u.names {
  parts = append(parts, allNames[id])
 }

 return strings.Join(parts, " ")
}

func NewUser2(fullName string) *User2 {
 getOrAdd := func(s string) uint8 {
  for i := range allNames {
   if allNames[i] == s {
    return uint8(i)
   }
  }
  allNames = append(allNames, s)
  return uint8(len(allNames) - 1)
 }

 result := User2{}
 parts := strings.Split(fullName, " ")
 for _, p := range parts {
  result.names = append(result.names, getOrAdd(p))
 }

 return &result
}
```

- Do keep in mind that in this usecase, the limit of the number of names in 256 as we are using uint8.
- But you can see that in case of very large number of users, there will be many repeated first names and last names and we wont have any trouble with it as we will not be duplicately storing the user names.
