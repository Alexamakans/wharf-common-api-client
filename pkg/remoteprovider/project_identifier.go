package remoteprovider

// ProjectIdentifier holds the values necessary to uniquely identify a project
// when communicating with the remote provider.
//
// Instances of this object should be obtained through the
// Client.WharfProjectToIdentifier method.
type ProjectIdentifier struct {
	Values []string
}

// func (p ProjectIdentifier) ToPathEscapedString() string {
// 	sb := strings.Builder{}
// 	valuesLen := len(p.Values)
// 	for i := 0; i < valuesLen-1; i++ {
// 		sb.WriteString(p.Values[i])
// 		sb.WriteRune(';')
// 	}
// 	if valuesLen != 0 {
// 		sb.WriteString(p.Values[valuesLen-1])
// 	}
// 	return url.PathEscape(sb.String())
// }

// func projectIdentifierFromPathEscapedString(s string) (ProjectIdentifier, error) {
// 	unescapedString, err := url.PathUnescape(s)
// 	if err != nil {
// 		return ProjectIdentifier{}, err
// 	}
// 	return ProjectIdentifier{
// 		Values: strings.Split(unescapedString, ";"),
// 	}, nil
// }
