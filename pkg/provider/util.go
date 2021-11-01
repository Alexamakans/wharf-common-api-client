package provider

func stripProject(p WharfProject) WharfProject {
	copy := p
	if copy.Provider != nil {
		copy.Provider.Token = nil
	}
	copy.Token = nil
	return copy
}
