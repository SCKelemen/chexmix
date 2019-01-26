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

func (l Language) String() (lang string) {
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

func ExtensionToLanguage(ext string) Language {
	switch ext {
	case ".go":
		return Go
	case ".cs":
	case ".csproj":
		return CSharp
	case ".js":
	case ".jsx":
		return JavaScript
	case ".ts":
	case ".tsx":
		return TypeScript
	case ".java":
		return Java
	case ".scala":
		return Scala
	case ".kotlin":
	case ".kt":
	case ".ktm":
	case ".kts":
		return Kotlin
	case ".swift":
		return Swift
	case ".ruby":
	case ".rb":
		return Ruby
	case ".py":
		return Python
	default:
		return Unknown
	}
	return Unknown
}
