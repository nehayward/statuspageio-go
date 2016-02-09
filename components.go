package statuspageio

import "fmt"

// curl https://api.statuspage.io/v1/pages/17jhzflc8qjp/components.json \
//  -H "Authorization: OAuth 00b870eb-95b9-4e39-96db-5dc60bac1462"
func (a StatusPageIOApi) ComponentList() (r []ComponentListResponse, err error) {
	err = a.apiRequest("GET", "/components.json", nil, &r)
	// if r.Status != "no" {
	// 	err = errors.New(r.Status.Message)
	// }
	return r, err
}

/*
curl https://api.statuspage.io/v1/pages/17jhzflc8qjp/components/237p399shf5q.json \
	-H "Authorization: OAuth 00b870eb-95b9-4e39-96db-5dc60bac1462" \
	-X PATCH \
	-d "component[status]=major_outage"
*/
func (a StatusPageIOApi) ComponentUpdate(componentID string, status string) (r ComponentUpdateResponse, err error) {
	err = a.apiRequest("PATCH", fmt.Sprintf("/components/%s.json", componentID), status, &r)
	// if r.Status.Error != "no" {
	// 	err = errors.New(r.Status.Message)
	// }
	return r, err
}

// curl https://api.statuspage.io/v1/pages/17jhzflc8qjp/components.json \
// 	 -H "Authorization: OAuth 00b870eb-95b9-4e39-96db-5dc60bac1462" \
// 	 -X POST \
// 	 -d "component[name]=Drone Login Test (Stage)" \
//
//
//
// 	 curl https://api.statuspage.io/v1/pages/17jhzflc8qjp/page_access_groups.json \
// 	 -H "Authorization: OAuth 00b870eb-95b9-4e39-96db-5dc60bac1462" \
