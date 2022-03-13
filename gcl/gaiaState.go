package gaia

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// how are terraform inputs passed to providers and returned? I honestly just need to do some proof of concept work
// on this one before I can proceed anymore with the implementation of gaia

// we need to figure out the right way to store this shit
type gaiaState struct {
	data []schema.Resource
}

// this function will aggregate the states of the implemented archetypal resources into an instance of gaiaState
func (gs *gaiaState) multiplexState() error {

	return nil
}

// this return is to be used to configure implemented resource archetypes
func (gs *gaiaState) demultiplexState() ([]schema.Resource, error) {

	var demuxxedResources []schema.Resource

	demuxxedResources = nil
	// iterate over
	/*for i := 0 ; i < len(gc.state.data) ; i ++ {
		gc.state.data[i]
	}*/

	return demuxxedResources, nil
}
