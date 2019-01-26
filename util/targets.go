package util

type Language uint

//  this should cover most of the languages
const (
	Unknown Language = iota
	Go
	CSharp
	JavaScript
	TypeScript
	Java
	Scala
	Kotlin
	Swift
	Ruby
	Python
)

//go:generate stringer -type=Language

func (l Language) String() string {
	var lang string

	switch l {
	case Go:
		lang = "Go"
	case CSharp:
		lang = "CSharp"
	case JavaScript:
		lang = "JavaScript"
	case TypeScript:
		lang = "TypeScript"
	case Java:
		lang = "Java"
	case Scala:
		lang = "Scala"
	case Kotlin:
		lang = "Kotlin"
	case Swift:
		lang = "Swift"
	case Ruby:
		lang = "Ruby"
	case Python:
		lang = "Python"
	default:
		lang = "Unknown"
	}

	return lang
}

func ExtensionToLanguage(ext string) (Language, error) {
	return Unknown, nil
}
